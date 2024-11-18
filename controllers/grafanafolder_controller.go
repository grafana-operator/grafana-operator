/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/go-logr/logr"
	genapi "github.com/grafana/grafana-openapi-client-go/client"
	"github.com/grafana/grafana-openapi-client-go/client/folders"
	"github.com/grafana/grafana-openapi-client-go/models"
	client2 "github.com/grafana/grafana-operator/v5/controllers/client"
	"github.com/grafana/grafana-operator/v5/controllers/metrics"
	kuberr "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	grafanav1beta1 "github.com/grafana/grafana-operator/v5/api/v1beta1"
)

const (
	conditionFolderSynchronized = "FolderSynchronized"
)

// GrafanaFolderReconciler reconciles a GrafanaFolder object
type GrafanaFolderReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=grafana.integreatly.org,resources=grafanafolders,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=grafana.integreatly.org,resources=grafanafolders/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=grafana.integreatly.org,resources=grafanafolders/finalizers,verbs=update

func (r *GrafanaFolderReconciler) syncFolders(ctx context.Context) (ctrl.Result, error) {
	syncLog := log.FromContext(ctx).WithName("GrafanaFolderReconciler")
	foldersSynced := 0

	// get all grafana instances
	grafanas := &grafanav1beta1.GrafanaList{}
	var opts []client.ListOption
	err := r.Client.List(ctx, grafanas, opts...)
	if err != nil {
		return ctrl.Result{
			Requeue: true,
		}, err
	}

	// no instances, no need to sync
	if len(grafanas.Items) == 0 {
		return ctrl.Result{Requeue: false}, nil
	}

	// get all folders
	allFolders := &grafanav1beta1.GrafanaFolderList{}
	err = r.Client.List(ctx, allFolders, opts...)
	if err != nil {
		return ctrl.Result{
			Requeue: true,
		}, err
	}

	// sync folders, delete folders from grafana that do no longer have a cr
	foldersToDelete := map[*grafanav1beta1.Grafana][]grafanav1beta1.NamespacedResource{}
	for _, grafana := range grafanas.Items {
		grafana := grafana
		for _, folder := range grafana.Status.Folders {
			if allFolders.Find(folder.Namespace(), folder.Name()) == nil {
				foldersToDelete[&grafana] = append(foldersToDelete[&grafana], folder)
			}
		}
	}

	// delete all folders that no longer have a cr
	for grafana, existingFolders := range foldersToDelete {
		grafanaClient, err := client2.NewGeneratedGrafanaClient(ctx, r.Client, grafana)
		if err != nil {
			return ctrl.Result{Requeue: true}, err
		}

		for _, folder := range existingFolders {
			// avoid bombarding the grafana instance with a large number of requests at once, limit
			// the sync to a certain number of folders per cycle. This means that it will take longer to sync
			// a large number of deleted dashboard crs, but that should be an edge case.
			if foldersSynced >= syncBatchSize {
				return ctrl.Result{Requeue: true}, nil
			}

			namespace, name, uid := folder.Split()

			reftrue := true
			params := folders.NewDeleteFolderParams().WithFolderUID(uid).WithForceDeleteRules(&reftrue)
			_, err = grafanaClient.Folders.DeleteFolder(params) //nolint
			if err != nil {
				var notFound *folders.DeleteFolderNotFound
				if errors.As(err, &notFound) {
					syncLog.Info("folder no longer exists", "namespace", namespace, "name", name)
				} else {
					return ctrl.Result{Requeue: false}, err
				}
			}

			grafana.Status.Folders = grafana.Status.Folders.Remove(namespace, name)
			foldersSynced += 1
		}

		// one update per grafana - this will trigger a reconcile of the grafana controller
		// so we should minimize those updates
		err = r.Client.Status().Update(ctx, grafana)
		if err != nil {
			return ctrl.Result{Requeue: false}, err
		}
	}

	if foldersSynced > 0 {
		syncLog.Info("successfully synced folders", "folders", foldersSynced)
	}
	return ctrl.Result{Requeue: false}, nil
}

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the GrafanaFolder object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.9.2/pkg/reconcile
func (r *GrafanaFolderReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	controllerLog := log.FromContext(ctx).WithName("GrafanaFolderReconciler")
	r.Log = controllerLog

	// periodic sync reconcile
	if req.Namespace == "" && req.Name == "" {
		start := time.Now()
		syncResult, err := r.syncFolders(ctx)
		elapsed := time.Since(start).Milliseconds()
		metrics.InitialFoldersSyncDuration.Set(float64(elapsed))
		return syncResult, err
	}

	folder := &grafanav1beta1.GrafanaFolder{}

	err := r.Client.Get(ctx, client.ObjectKey{
		Namespace: req.Namespace,
		Name:      req.Name,
	}, folder)
	if err != nil {
		if kuberr.IsNotFound(err) {
			if err := r.onFolderDeleted(ctx, req.Namespace, req.Name); err != nil {
				return ctrl.Result{RequeueAfter: RequeueDelay}, err
			}
			return ctrl.Result{}, nil
		}
		controllerLog.Error(err, "error getting grafana folder cr")
		return ctrl.Result{RequeueAfter: RequeueDelay}, err
	}
	defer func() {
		if err := r.UpdateStatus(ctx, folder); err != nil {
			r.Log.Error(err, "updating status")
		}
	}()

	if folder.Spec.ParentFolderUID == folder.CustomUIDOrUID() {
		setInvalidSpec(&folder.Status.Conditions, folder.Generation, "CyclicParent", "The value of parentFolderUID must not be the uid of the current folder")
		meta.RemoveStatusCondition(&folder.Status.Conditions, conditionFolderSynchronized)
		return ctrl.Result{}, fmt.Errorf("cyclic folder reference")
	}
	removeInvalidSpec(&folder.Status.Conditions)

	instances, err := r.GetMatchingFolderInstances(ctx, folder, r.Client)
	if err != nil {
		setNoMatchingInstance(&folder.Status.Conditions, folder.Generation, "ErrFetchingInstances", fmt.Sprintf("error occurred during fetching of instances: %s", err.Error()))
		meta.RemoveStatusCondition(&folder.Status.Conditions, conditionFolderSynchronized)
		r.Log.Error(err, "could not find matching instances")
		return ctrl.Result{}, err
	}
	if len(instances.Items) == 0 {
		setNoMatchingInstance(&folder.Status.Conditions, folder.Generation, "EmptyAPIReply", "Instances could not be fetched, reconciliation will be retried")
		meta.RemoveStatusCondition(&folder.Status.Conditions, conditionFolderSynchronized)
		return ctrl.Result{}, fmt.Errorf("no instances found")
	}
	removeNoMatchingInstance(&folder.Status.Conditions)
	controllerLog.Info("found matching Grafana instances for folder", "count", len(instances.Items))

	applyErrors := make(map[string]string)
	for _, grafana := range instances.Items {
		grafana := grafana

		// check if this is a cross namespace import
		if grafana.Namespace != folder.Namespace && !folder.IsAllowCrossNamespaceImport() {
			continue
		}

		if grafana.Status.Stage != grafanav1beta1.OperatorStageComplete || grafana.Status.StageStatus != grafanav1beta1.OperatorStageResultSuccess {
			controllerLog.Info("grafana instance not ready", "grafana", grafana.Name)
			continue
		}

		err = r.onFolderCreated(ctx, &grafana, folder)
		if err != nil {
			controllerLog.Error(err, "error reconciling folder", "folder", folder.Name, "grafana", grafana.Name)
			applyErrors[fmt.Sprintf("%s/%s", grafana.Namespace, grafana.Name)] = err.Error()
		}
	}
	condition := buildSynchronizedCondition("Folder", conditionFolderSynchronized, folder.Generation, applyErrors, len(instances.Items))
	meta.SetStatusCondition(&folder.Status.Conditions, condition)

	if len(applyErrors) != 0 {
		return ctrl.Result{RequeueAfter: RequeueDelay}, nil
	}

	if folder.ResyncPeriodHasElapsed() {
		folder.Status.LastResync = metav1.Time{Time: time.Now()}
	}
	return ctrl.Result{RequeueAfter: folder.GetResyncPeriod()}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *GrafanaFolderReconciler) SetupWithManager(mgr ctrl.Manager, ctx context.Context) error {
	err := ctrl.NewControllerManagedBy(mgr).
		For(&grafanav1beta1.GrafanaFolder{}).
		Complete(r)

	if err == nil {
		d, err := time.ParseDuration(initialSyncDelay)
		if err != nil {
			return err
		}

		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case <-time.After(d):
					result, err := r.Reconcile(ctx, ctrl.Request{})
					if err != nil {
						r.Log.Error(err, "error synchronizing folders")
						continue
					}
					if result.Requeue {
						r.Log.Info("more folders left to synchronize")
						continue
					}
					r.Log.Info("folder sync complete")
					return
				}
			}
		}()
	}

	return err
}

func (r *GrafanaFolderReconciler) onFolderDeleted(ctx context.Context, namespace string, name string) error {
	list := grafanav1beta1.GrafanaList{}
	var opts []client.ListOption
	err := r.Client.List(ctx, &list, opts...)
	if err != nil {
		return err
	}

	for _, grafana := range list.Items {
		grafana := grafana
		if found, uid := grafana.Status.Folders.Find(namespace, name); found {
			grafanaClient, err := client2.NewGeneratedGrafanaClient(ctx, r.Client, &grafana)
			if err != nil {
				return err
			}
			reftrue := true
			params := folders.NewDeleteFolderParams().WithFolderUID(*uid).WithForceDeleteRules(&reftrue)
			_, err = grafanaClient.Folders.DeleteFolder(params) //nolint
			if err != nil {
				var notFound *folders.DeleteFolderNotFound
				if !errors.As(err, &notFound) {
					return err
				}
			}

			grafana.Status.Folders = grafana.Status.Folders.Remove(namespace, name)
			err = r.Client.Status().Update(ctx, &grafana)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (r *GrafanaFolderReconciler) onFolderCreated(ctx context.Context, grafana *grafanav1beta1.Grafana, cr *grafanav1beta1.GrafanaFolder) error {
	title := cr.GetTitle()
	uid := cr.CustomUIDOrUID()

	grafanaClient, err := client2.NewGeneratedGrafanaClient(ctx, r.Client, grafana)
	if err != nil {
		return err
	}

	parentFolderUID, err := getFolderUID(ctx, r.Client, cr)
	if err != nil {
		return err
	}

	exists, remoteUID, remoteParent, err := r.Exists(grafanaClient, cr)
	if err != nil {
		return err
	}

	// always update after resync period has elapsed even if cr is unchanged.
	if exists && cr.Unchanged() && !cr.ResyncPeriodHasElapsed() && parentFolderUID == remoteParent {
		return nil
	}

	if exists {
		// make sure we use the correct UID
		uid = remoteUID
		// Add to status to cover cases:
		// - operator have previously failed to update status
		// - the folder was created outside of operator
		// - the folder was created through dashboard controller
		if found, _ := grafana.Status.Folders.Find(cr.Namespace, cr.Name); !found {
			grafana.Status.Folders = grafana.Status.Folders.Add(cr.Namespace, cr.Name, uid)
			err = r.Client.Status().Update(ctx, grafana)
			if err != nil {
				return err
			}
		}

		if !cr.Unchanged() {
			_, err = grafanaClient.Folders.UpdateFolder(remoteUID, &models.UpdateFolderCommand{ //nolint
				Overwrite: true,
				Title:     title,
			})
			if err != nil {
				return err
			}
		}

		if parentFolderUID != remoteParent {
			_, err = grafanaClient.Folders.MoveFolder(remoteUID, &models.MoveFolderCommand{ //nolint
				ParentUID: parentFolderUID,
			})
			if err != nil {
				return err
			}
		}
	} else {
		body := &models.CreateFolderCommand{
			Title:     title,
			UID:       uid,
			ParentUID: parentFolderUID,
		}

		folderResp, err := grafanaClient.Folders.CreateFolder(body)
		if err != nil {
			return err
		}

		grafana.Status.Folders = grafana.Status.Folders.Add(cr.Namespace, cr.Name, folderResp.Payload.UID)
		err = r.Client.Status().Update(ctx, grafana)
		if err != nil {
			return err
		}
	}

	// NOTE: it's up to a user to reset permissions with correct json
	if cr.Spec.Permissions != "" {
		permissions := models.UpdateDashboardACLCommand{}
		err = json.Unmarshal([]byte(cr.Spec.Permissions), &permissions)
		if err != nil {
			return fmt.Errorf("failed to unmarshal spec.permissions: %w", err)
		}

		_, err = grafanaClient.FolderPermissions.UpdateFolderPermissions(uid, &permissions) //nolint
		if err != nil {
			return fmt.Errorf("failed to update folder permissions: %w", err)
		}
	}

	return nil
}

func (r *GrafanaFolderReconciler) UpdateStatus(ctx context.Context, cr *grafanav1beta1.GrafanaFolder) error {
	cr.Status.Hash = cr.Hash()
	return r.Client.Status().Update(ctx, cr)
}

// Check if the folder exists. Matches UID first and fall back to title. Title matching only works for non-nested folders
func (r *GrafanaFolderReconciler) Exists(client *genapi.GrafanaHTTPAPI, cr *grafanav1beta1.GrafanaFolder) (bool, string, string, error) {
	title := cr.GetTitle()
	uid := cr.CustomUIDOrUID()

	uidResp, err := client.Folders.GetFolderByUID(uid)
	if err == nil {
		return true, uidResp.Payload.UID, uidResp.Payload.ParentUID, nil
	}

	page := int64(1)
	limit := int64(10000)
	for {
		params := folders.NewGetFoldersParams().WithPage(&page).WithLimit(&limit)

		foldersResp, err := client.Folders.GetFolders(params)
		if err != nil {
			return false, "", "", err
		}
		folders := foldersResp.GetPayload()

		for _, remoteFolder := range folders {
			if strings.EqualFold(remoteFolder.Title, title) {
				return true, remoteFolder.UID, remoteFolder.ParentUID, nil
			}
		}
		if len(folders) < int(limit) {
			return false, "", "", nil
		}
		page++
	}
}

func (r *GrafanaFolderReconciler) GetMatchingFolderInstances(ctx context.Context, folder *grafanav1beta1.GrafanaFolder, k8sClient client.Client) (grafanav1beta1.GrafanaList, error) {
	instances, err := GetMatchingInstances(ctx, k8sClient, folder.Spec.InstanceSelector)
	if err != nil || len(instances.Items) == 0 {
		folder.Status.NoMatchingInstances = true
		if err := r.Client.Status().Update(ctx, folder); err != nil {
			r.Log.Info("unable to update the status of %v, in %v", folder.Name, folder.Namespace)
		}
		return grafanav1beta1.GrafanaList{}, err
	}
	folder.Status.NoMatchingInstances = false
	if err := r.Client.Status().Update(ctx, folder); err != nil {
		r.Log.Info("unable to update the status of %v, in %v", folder.Name, folder.Namespace)
	}

	return instances, err
}
