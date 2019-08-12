package grafanadashboard

import (
	"context"
	"crypto/md5"
	defaultErrors "errors"
	"fmt"
	i8ly "github.com/integr8ly/grafana-operator/pkg/apis/integreatly/v1alpha1"
	"github.com/integr8ly/grafana-operator/pkg/controller/common"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/json"
	"time"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_grafanadashboard")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new GrafanaDashboard Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileGrafanaDashboard{
		client: mgr.GetClient(),
		scheme: mgr.GetScheme(),
		config: common.GetControllerConfig(),
		helper: common.NewKubeHelper(),
	}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("grafanadashboard-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource GrafanaDashboard
	err = c.Watch(&source.Kind{Type: &i8ly.GrafanaDashboard{}}, &handler.EnqueueRequestForObject{})
	if err == nil {
		log.Info("Starting dashboard controller")
	}

	return err
}

var _ reconcile.Reconciler = &ReconcileGrafanaDashboard{}

// ReconcileGrafanaDashboard reconciles a GrafanaDashboard object
type ReconcileGrafanaDashboard struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
	config *common.ControllerConfig
	helper *common.KubeHelperImpl
}

func (r *ReconcileGrafanaDashboard) matchesSelector(d *i8ly.GrafanaDashboard, s *v1.LabelSelector) (bool, error) {
	selector, err := v1.LabelSelectorAsSelector(s)
	if err != nil {
		return false, err
	}

	return selector.Empty() || selector.Matches(labels.Set(d.Labels)), nil
}

func (r *ReconcileGrafanaDashboard) matchesSelectors(d *i8ly.GrafanaDashboard, s []*v1.LabelSelector) (bool, error) {
	result := false

	for _, selector := range s {
		match, err := r.matchesSelector(d, selector)
		if err != nil {
			return false, err
		}

		result = result || match
	}

	return result, nil
}

// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileGrafanaDashboard) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	dashboardLabelSelectors := r.config.GetConfigItem(common.ConfigDashboardLabelSelector, nil)
	if dashboardLabelSelectors == nil {
		return reconcile.Result{RequeueAfter: time.Second * 10}, nil
	}

	// Fetch the GrafanaDashboard instance
	instance := &i8ly.GrafanaDashboard{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	cr := instance.DeepCopy()
	if match, err := r.matchesSelectors(cr, dashboardLabelSelectors.([]*v1.LabelSelector)); err != nil {
		return reconcile.Result{}, err
	} else if !match {
		log.Info(fmt.Sprintf("found dashboard '%s/%s' but labels do not match", cr.Namespace, cr.Name))
		return reconcile.Result{}, nil
	}

	// Resource deleted?
	if cr.DeletionTimestamp != nil {
		return r.deleteDashboard(cr)
	}

	switch cr.Status.Phase {
	case common.StatusResourceUninitialized:
		// New resource
		return r.updatePhase(cr, common.StatusResourceSetFinalizer)
	case common.StatusResourceSetFinalizer:
		// Set finalizer first
		if len(cr.Finalizers) > 0 {
			return r.updatePhase(cr, common.StatusResourceCreated)
		} else {
			return r.setFinalizer(cr)
		}
	case common.StatusResourceCreated:
		// Import / update dashboard
		return r.importDashboard(cr)
	default:
		return reconcile.Result{}, nil
	}
}

func (r *ReconcileGrafanaDashboard) checkPrerequisites(d *i8ly.GrafanaDashboard) bool {
	changed, hash := r.hasDashboardChanged(d)
	if !changed {
		log.Info("dashboard reconciled but no changes")
		return false
	}
	d.Status.LastConfig = hash

	valid, err := r.isJsonValid(d)
	if valid {
		return true
	}

	// Don't append the same error twice
	msg := fmt.Sprintf("invalid JSON, error: %s", err)
	for _, statusMessage := range d.Status.Messages {
		if statusMessage.Message == msg {
			log.Info("dashboard reconciled but json still invalid")
			return false
		}
	}

	common.AppendMessage(msg, d)
	err = r.client.Update(context.TODO(), d)
	if err != nil {
		log.Error(err, "update dashboard messages failed")
	}

	log.Info("dashboard reconciled but json invalid")
	return false
}

func (r *ReconcileGrafanaDashboard) importDashboard(d *i8ly.GrafanaDashboard) (reconcile.Result, error) {
	operatorNamespace := r.config.GetConfigString(common.ConfigOperatorNamespace, "")
	if operatorNamespace == "" {
		return reconcile.Result{}, defaultErrors.New("operator namespace not yet known")
	}

	valid := r.checkPrerequisites(d)
	if valid == false {
		return reconcile.Result{Requeue: false}, nil
	}

	updated, err := r.helper.UpdateDashboard(operatorNamespace, d)
	if err != nil {
		return reconcile.Result{}, err
	}

	if !updated {
		return reconcile.Result{RequeueAfter: time.Second * common.RequeueDelaySeconds}, err
	}

	// Reconcile dashboard plugins
	r.config.SetPluginsFor(d)

	msg := fmt.Sprintf("dashboard '%s/%s' imported", d.Namespace, d.Spec.Name)
	common.AppendMessage(msg, d)

	// Update the dashboard to persist the new hash and the satus message
	err = r.client.Update(context.TODO(), d)
	if err == nil {
		log.Info(msg)
	}

	return reconcile.Result{}, err
}

func (r *ReconcileGrafanaDashboard) deleteDashboard(d *i8ly.GrafanaDashboard) (reconcile.Result, error) {
	operatorNamespace := r.config.GetConfigString(common.ConfigOperatorNamespace, "")
	if operatorNamespace == "" {
		return reconcile.Result{}, defaultErrors.New("no monitoring namespace set")
	}

	err := r.helper.DeleteDashboard(operatorNamespace, d.Namespace, d)
	if err == nil {
		log.Info(fmt.Sprintf("dashboard '%s/%s' deleted", d.Namespace, d.Spec.Name))
	}

	r.config.RemovePluginsFor(d)
	return r.removeFinalizer(d)
}

func (r *ReconcileGrafanaDashboard) removeFinalizer(cr *i8ly.GrafanaDashboard) (reconcile.Result, error) {
	cr.Finalizers = nil
	err := r.client.Update(context.TODO(), cr)
	return reconcile.Result{}, err
}

func (r *ReconcileGrafanaDashboard) setFinalizer(cr *i8ly.GrafanaDashboard) (reconcile.Result, error) {
	if len(cr.Finalizers) == 0 {
		cr.Finalizers = append(cr.Finalizers, common.ResourceFinalizerName)
	}
	err := r.client.Update(context.TODO(), cr)
	return reconcile.Result{}, err
}

func (r *ReconcileGrafanaDashboard) updatePhase(cr *i8ly.GrafanaDashboard, phase int) (reconcile.Result, error) {
	cr.Status.Phase = phase
	err := r.client.Update(context.TODO(), cr)
	return reconcile.Result{}, err
}

func (r *ReconcileGrafanaDashboard) isJsonValid(cr *i8ly.GrafanaDashboard) (bool, error) {
	var js map[string]interface{}
	err := json.Unmarshal([]byte(cr.Spec.Json), &js)
	return err == nil, err

}

func (r *ReconcileGrafanaDashboard) hasDashboardChanged(cr *i8ly.GrafanaDashboard) (bool, string) {
	hash := fmt.Sprintf("%x", md5.Sum([]byte(cr.Spec.Json)))
	return hash != cr.Status.LastConfig, hash
}
