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

package v1alpha1

import (
	"errors"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var cakelog = logf.Log.WithName("cake-resource")

func (r *Cake) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-food-github-com-reetas-v1alpha1-cake,mutating=true,failurePolicy=fail,sideEffects=None,groups=food.github.com.reetas,resources=cakes,verbs=create;update,versions=v1alpha1,name=mcake.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &Cake{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *Cake) Default() {
	cakelog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-food-github-com-reetas-v1alpha1-cake,mutating=false,failurePolicy=fail,sideEffects=None,groups=food.github.com.reetas,resources=cakes,verbs=create;update,versions=v1alpha1,name=vcake.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &Cake{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Cake) ValidateCreate() error {
	cakelog.Info("validate create", "name", r.Name)

	if r.Spec.Delivery.DeliveryType == Pickup && len(r.Spec.Delivery.DeliveryAddress) > 0 {
		return errors.New("DeliverAddress should be empty for Pickup")
	}
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Cake) ValidateUpdate(old runtime.Object) error {
	cakelog.Info("validate update", "name", r.Name)

	if r.Spec.Delivery.DeliveryType == Pickup && len(r.Spec.Delivery.DeliveryAddress) > 0 {
		return errors.New("DeliverAddress should be empty  for Pickup")
	}
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Cake) ValidateDelete() error {
	cakelog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
