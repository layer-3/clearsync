/*
Copyright 2024.

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

package controller

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	appsv1beta1 "github.com/layer-3/clearsync/api/v1beta1"
)

// GrafanaScalerReconciler reconciles a GrafanaScaler object
type GrafanaScalerReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=apps.yellow.com,resources=grafanascalers,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=apps.yellow.com,resources=grafanascalers/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=apps.yellow.com,resources=grafanascalers/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the GrafanaScaler object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.16.3/pkg/reconcile
func (r *GrafanaScalerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    var scaler apps.GrafanaScaler
    if err := r.Get(ctx, req.NamespacedName, &scaler); err != nil {
        log.Error(err, "unable to fetch GrafanaScaler")
        return ctrl.Result{}, client.IgnoreNotFound(err)
    }

    // Scale logic here
    var deployment appsv1beta1.Deployment
    if err := r.Get(ctx, types.NamespacedName{Name: scaler.Spec.DeploymentName, Namespace: req.Namespace}, &deployment); err != nil {
        return ctrl.Result{}, err
    }

    if *deployment.Spec.Replicas != 0 {
        deployment.Spec.Replicas = new(int32) // scale down to 0
        if err := r.Update(ctx, &deployment); err != nil {
            return ctrl.Result{}, err
        }
        scaler.Status.Scaled = true
        if err := r.Status().Update(ctx, &scaler); err != nil {
            return ctrl.Result{}, err
        }
    }

    return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *GrafanaScalerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&appsv1beta1.GrafanaScaler{}).
		Complete(r)
}
