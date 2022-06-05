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
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	foodv1alpha1 "github.com/reetas/project/api/v1alpha1"
)

// CakeReconciler reconciles a Cake object
type CakeReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=food.github.com.reetas,resources=cakes,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=food.github.com.reetas,resources=cakes/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=food.github.com.reetas,resources=cakes/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Cake object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.2/pkg/reconcile
func (r *CakeReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	l := log.FromContext(ctx)

	cake := &foodv1alpha1.Cake{}
	err := r.Get(ctx, req.NamespacedName, cake)
	if err != nil {
		return ctrl.Result{}, fmt.Errorf("unable to get cake %s", req.NamespacedName)
	}
	new_cake := cake.DeepCopy()
	l.Info(cake.GenerateName)
	requeTime := time.Duration(10 * time.Minute)
	//TODO figure why all the state is getting updated together
	if !new_cake.Status.OrderReceived {
		new_cake.Status.OrderReceived = true
	} else if !new_cake.Status.InProgress {
		new_cake.Status.InProgress = true
	} else if !new_cake.Status.Ready {
		new_cake.Status.Ready = true
	} else if new_cake.Spec.Delivery.DeliveryType == foodv1alpha1.Deliver && !new_cake.Status.Delivered {
		new_cake.Status.Delivered = true
	} else if new_cake.Spec.Delivery.DeliveryType == foodv1alpha1.Pickup {
		new_cake.Status.PickedUp = true
	}

	r.Status().Update(ctx, new_cake)
	return ctrl.Result{RequeueAfter: requeTime}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *CakeReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&foodv1alpha1.Cake{}).
		Complete(r)
}
