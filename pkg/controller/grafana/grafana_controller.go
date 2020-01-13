package grafana

import (
	"context"
	"fmt"

	i8ly "github.com/integr8ly/grafana-operator/pkg/apis/integreatly/v1alpha1"
	"github.com/integr8ly/grafana-operator/pkg/controller/common"
	core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_grafana")

const (
	PhaseConfigFiles int = iota
	PhaseInstallGrafana
	PhaseDone
	PhaseReconcile
)

const OpenShiftOAuthRedirect = "serviceaccounts.openshift.io/oauth-redirectreference.primary"

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new Grafana Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileGrafana{
		client:  mgr.GetClient(),
		scheme:  mgr.GetScheme(),
		helper:  common.NewKubeHelper(),
		plugins: newPluginsHelper(),
		config:  common.GetControllerConfig(),
	}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("grafana-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource Grafana
	return c.Watch(&source.Kind{Type: &i8ly.Grafana{}}, &handler.EnqueueRequestForObject{})
}

var _ reconcile.Reconciler = &ReconcileGrafana{}

// ReconcileGrafana reconciles a Grafana object
type ReconcileGrafana struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client  client.Client
	scheme  *runtime.Scheme
	helper  *common.KubeHelperImpl
	plugins *PluginsHelperImpl
	config  *common.ControllerConfig
}

// Reconcile reads that state of the cluster for a Grafana object and makes changes based on the state read
// and what is in the Grafana.Spec
func (r *ReconcileGrafana) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	instance := &i8ly.Grafana{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Stop the dashboard controller from reconciling when grafana is not installed
			r.config.RemoveConfigItem(common.ConfigDashboardLabelSelector)
			return reconcile.Result{}, nil
		}
		return reconcile.Result{}, err
	}

	cr := instance.DeepCopy()

	switch cr.Status.Phase {
	case PhaseConfigFiles:
		return r.createConfigFiles(cr)
	case PhaseInstallGrafana:
		return r.installGrafana(cr)
	case PhaseDone:
		log.Info("Grafana installation complete")
		return r.updatePhase(cr, PhaseReconcile)
	case PhaseReconcile:
		return r.ReconcileGrafana(cr)
	}

	return reconcile.Result{}, nil
}

// ReconcileGrafana is constantly reconcile the grafana config and plugins
func (r *ReconcileGrafana) ReconcileGrafana(cr *i8ly.Grafana) (reconcile.Result, error) {
	// Update the label selector and make it available to the dashboard controller
	r.config.AddConfigItem(common.ConfigDashboardLabelSelector, cr.Spec.DashboardLabelSelector)

	// Config updated?
	newConfig := NewIniConfig(cr)
	if err := newConfig.Build(); err != nil {
		return reconcile.Result{}, err
	}

	if newConfig.DiffersFrom(cr.Status.LastConfig) {
		if err := r.helper.UpdateGrafanaConfig(newConfig.Contents, cr); err != nil {
			return reconcile.Result{}, err
		}

		// Store the new config hash
		cr.Status.LastConfig = newConfig.Hash
		if err := r.client.Update(context.TODO(), cr); err != nil {
			return reconcile.Result{}, err
		}

		// Grafana needs to be restarted after a config change
		if err := r.helper.UpdateGrafanaDeployment(cr.Status.LastConfig); err != nil {
			return reconcile.Result{}, err
		}
		log.Info("grafana updated configuration hash due to config change")

		// Skip plugins reconciliation while grafana is restarting
		return reconcile.Result{RequeueAfter: common.RequeueDelay}, nil
	}

	// Plugins updated?
	if err := r.ReconcileDashboardPlugins(cr); err != nil {
		return reconcile.Result{}, err
	}

	return reconcile.Result{RequeueAfter: common.RequeueDelay}, nil
}

// ReconcileDashboardPlugins is responsible for reconciling the grafana dashboards
func (r *ReconcileGrafana) ReconcileDashboardPlugins(cr *i8ly.Grafana) error {
	// Waited long enough for dashboards to be ready?
	if r.plugins.CanUpdatePlugins() == false {
		return nil
	}

	// Fetch all plugins of all dashboards
	var requestedPlugins i8ly.PluginList
	for _, v := range common.GetControllerConfig().Plugins {
		requestedPlugins = append(requestedPlugins, v...)
	}

	// Consolidate plugins and create the new list of plugin requirements
	// If 'updated' is false then no changes have to be applied
	filteredPlugins, updated := r.plugins.FilterPlugins(cr, requestedPlugins)

	if updated {
		r.ReconcilePlugins(cr, filteredPlugins)

		// Update the dashboards that had their plugins modified
		// to let the owners know about the status
		err := r.updateDashboardMessages(filteredPlugins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *ReconcileGrafana) ReconcilePlugins(cr *i8ly.Grafana, plugins []i8ly.GrafanaPlugin) error {
	var validPlugins []i8ly.GrafanaPlugin
	var failedPlugins []i8ly.GrafanaPlugin

	for _, plugin := range plugins {
		if r.plugins.PluginExists(plugin) == false {
			log.Info(fmt.Sprintf("invalid plugin: %s@%s", plugin.Name, plugin.Version))
			failedPlugins = append(failedPlugins, plugin)
			continue
		}

		log.Info(fmt.Sprintf("installing plugin: %s@%s", plugin.Name, plugin.Version))
		validPlugins = append(validPlugins, plugin)
	}

	cr.Status.InstalledPlugins = validPlugins
	cr.Status.FailedPlugins = failedPlugins

	err := r.client.Update(context.TODO(), cr)
	if err != nil {
		return err
	}

	newEnv := r.plugins.BuildEnv(cr)
	err = r.helper.UpdateGrafanaInitContainersDeployment(newEnv)
	return err
}

func (r *ReconcileGrafana) updateDashboardMessages(plugins i8ly.PluginList) error {
	for _, plugin := range plugins {
		err := r.client.Update(context.TODO(), plugin.Origin)
		if err != nil {
			return err
		}
	}
	return nil
}

// Initially create the config map that contains grafana.ini
func (r *ReconcileGrafana) createGrafanaConfig(cr *i8ly.Grafana) error {
	grafanaIni := NewIniConfig(cr)
	err := grafanaIni.Build()
	if err != nil {
		return err
	}

	configMap := core.ConfigMap{}
	configMap.ObjectMeta = v1.ObjectMeta{
		Name:      common.GrafanaConfigMapName,
		Namespace: cr.Namespace,
	}
	configMap.Data = map[string]string{}
	configMap.Data[common.GrafanaConfigFileName] = grafanaIni.Contents
	err = controllerutil.SetControllerReference(cr, &configMap, r.scheme)
	if err != nil {
		return err
	}

	cr.Status.LastConfig = grafanaIni.Hash
	err = r.client.Create(context.TODO(), &configMap)
	if err != nil {
		// This might be called multiple times if creating one of the other config
		// resources fails
		if errors.IsAlreadyExists(err) {
			return nil
		}
	}
	return err
}

func (r *ReconcileGrafana) createConfigFiles(cr *i8ly.Grafana) (reconcile.Result, error) {
	log.Info("Phase: Create Config Files")

	ingressType := common.GrafanaIngressName
	if common.GetControllerConfig().GetConfigBool(common.ConfigOpenshift, false) == true {
		ingressType = common.GrafanaRouteName
	}

	err := r.createServiceAccount(cr, common.GrafanaServiceAccountName)
	if err != nil {
		return reconcile.Result{}, err
	}

	err = r.createGrafanaConfig(cr)
	if err != nil {
		return reconcile.Result{}, err
	}

	resources := []string{
		common.GrafanaDashboardsConfigMapName,
		common.GrafanaProvidersConfigMapName,
		common.GrafanaDatasourcesConfigMapName,
		common.GrafanaServiceName,
		common.GrafanaAuthProxyConfigMapName,
		common.GrafanaAuthProxyIngressName,
		common.GrafanaAuthProxyServiceName,
	}

	// Ingress / Route has to be enabled
	if cr.Spec.Ingress.Enabled {
		resources = append(resources, ingressType)
	}

	for _, resourceName := range resources {
		log.Info(fmt.Sprintf("Creating the %s resource", resourceName))
		if err := r.createResource(cr, resourceName); err != nil {
			log.Info(fmt.Sprintf("Error in CreateConfigFiles, resourceName=%s : err=%s", resourceName, err))
			// Requeue so it can be attempted again
			return reconcile.Result{}, err
		}
	}
	log.Info("Config files created")

	return r.updatePhase(cr, PhaseInstallGrafana)
}

func (r *ReconcileGrafana) installGrafana(cr *i8ly.Grafana) (reconcile.Result, error) {
	log.Info("Phase: Install Grafana")

	if err := r.createDeployment(cr, common.GrafanaDeploymentName); err != nil {
		return reconcile.Result{}, err
	}
	if cr.Spec.AuthProxy.Enabled {
		if err := r.createDeployment(cr, common.GrafanaAuthProxyDeploymentName); err != nil {
			return reconcile.Result{}, err
		}
	}

	return r.updatePhase(cr, PhaseDone)
}

// Creates the deployment from the template and injects any specified extra containers before
// submitting it
func (r *ReconcileGrafana) createDeployment(cr *i8ly.Grafana, resourceName string) error {
	resourceHelper := newResourceHelper(cr)
	resource, err := resourceHelper.createResource(resourceName)
	if err != nil {
		return err
	}

	rawResource := newUnstructuredResourceMap(resource.(*unstructured.Unstructured))
	var extraVolumeMounts []interface{}

	// Extra secrets to be added as volumes?
	if len(cr.Spec.Secrets) > 0 {
		volumes := rawResource.access("spec").access("template").access("spec").get("volumes").([]interface{})

		for _, secret := range cr.Spec.Secrets {
			volumeName := fmt.Sprintf("secret-%s", secret)
			log.Info(fmt.Sprintf("adding volume for secret '%s' as '%s'", secret, volumeName))
			volumes = append(volumes, core.Volume{
				Name: volumeName,
				VolumeSource: core.VolumeSource{
					Secret: &core.SecretVolumeSource{
						SecretName: secret,
					},
				},
			})
			extraVolumeMounts = append(extraVolumeMounts, map[string]interface{}{
				"name":      volumeName,
				"readOnly":  true,
				"mountPath": common.SecretsMountDir + secret,
			})
		}

		rawResource.access("spec").access("template").access("spec").set("volumes", volumes)
	}

	// Extra config maps to be added as volumes?
	if len(cr.Spec.ConfigMaps) > 0 {
		volumes := rawResource.access("spec").access("template").access("spec").get("volumes").([]interface{})

		for _, configmap := range cr.Spec.ConfigMaps {
			volumeName := fmt.Sprintf("configmap-%s", configmap)
			log.Info(fmt.Sprintf("adding volume for configmap '%s' as '%s'", configmap, volumeName))
			volumes = append(volumes, core.Volume{
				Name: volumeName,
				VolumeSource: core.VolumeSource{
					ConfigMap: &core.ConfigMapVolumeSource{
						LocalObjectReference: core.LocalObjectReference{
							Name: configmap,
						},
					},
				},
			})
			extraVolumeMounts = append(extraVolumeMounts, map[string]interface{}{
				"name":      volumeName,
				"readOnly":  true,
				"mountPath": common.ConfigMapsMountDir + configmap,
			})
		}

		rawResource.access("spec").access("template").access("spec").set("volumes", volumes)
	}

	// Extra containers to add to the deployment?
	if len(cr.Spec.Containers) > 0 {
		// Otherwise append extra containers before submitting the resource
		containers := rawResource.access("spec").access("template").access("spec").get("containers").([]interface{})

		for _, container := range cr.Spec.Containers {
			containers = append(containers, container)
			log.Info(fmt.Sprintf("adding extra container '%v' to '%v'", container.Name, common.GrafanaDeploymentName))
		}

		rawResource.access("spec").access("template").access("spec").set("containers", containers)
	}

	// Append extra volume mounts to all containers
	if len(extraVolumeMounts) > 0 {
		containers := rawResource.access("spec").access("template").access("spec").get("containers").([]interface{})

		for _, container := range containers {
			volumeMounts := container.(map[string]interface{})["volumeMounts"].([]interface{})
			volumeMounts = append(volumeMounts, extraVolumeMounts...)
			container.(map[string]interface{})["volumeMounts"] = volumeMounts
		}
	}

	return r.deployResource(cr, resource, resourceName)

}

func (r *ReconcileGrafana) createServiceAccount(cr *i8ly.Grafana, resourceName string) error {
	resourceHelper := newResourceHelper(cr)
	resource, err := resourceHelper.createResource(resourceName)

	if err != nil {
		return err
	}

	// Deploy the unmodified resource if not on OpenShift
	if common.GetControllerConfig().GetConfigBool(common.ConfigOpenshift, false) == false {
		return r.deployResource(cr, resource, resourceName)
	}

	// Otherwise add an annotation that allows using the OAuthProxy (and will have no
	// effect otherwise)
	annotations := make(map[string]string)
	annotations[OpenShiftOAuthRedirect] = fmt.Sprintf(`{"kind":"OAuthRedirectReference","apiVersion":"v1","reference":{"kind":"Route","name":"%s"}}`, common.GrafanaRouteName)

	rawResource := newUnstructuredResourceMap(resource.(*unstructured.Unstructured))
	rawResource.access("metadata").set("annotations", annotations)

	return r.deployResource(cr, resource, resourceName)
}

// Creates a generic kubernetes resource from a template
func (r *ReconcileGrafana) createResource(cr *i8ly.Grafana, resourceName string) error {
	resourceHelper := newResourceHelper(cr)
	resource, err := resourceHelper.createResource(resourceName)

	if err != nil {
		return err
	}

	return r.deployResource(cr, resource, resourceName)
}

// Deploys a resource given by a runtime object
func (r *ReconcileGrafana) deployResource(cr *i8ly.Grafana, resource runtime.Object, resourceName string) error {
	// Try to find the resource, it may already exist
	selector := types.NamespacedName{
		Namespace: cr.Namespace,
		Name:      resourceName,
	}
	err := r.client.Get(context.TODO(), selector, resource)

	// The resource exists, do nothing
	if err == nil {
		return nil
	}

	// Resource does not exist or something went wrong
	if errors.IsNotFound(err) {
		log.Info(fmt.Sprintf("Resource %s does not exist, creating now", resourceName))
	} else {
		return err
	}

	// Set the CR as the owner of this resource so that when
	// the CR is deleted this resource also gets removed
	err = controllerutil.SetControllerReference(cr, resource.(v1.Object), r.scheme)
	if err != nil {
		return err
	}

	return r.client.Create(context.TODO(), resource)
}

func (r *ReconcileGrafana) updatePhase(cr *i8ly.Grafana, phase int) (reconcile.Result, error) {
	cr.Status.Phase = phase
	err := r.client.Update(context.TODO(), cr)
	return reconcile.Result{}, err
}
