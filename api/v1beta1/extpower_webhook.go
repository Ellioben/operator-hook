/*
Copyright 2023.

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

package v1beta1

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var extpowerlog = logf.Log.WithName("extpower-resource")

func (r *ExtPower) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-plugin-github-com-ellioben-v1beta1-extpower,mutating=true,failurePolicy=fail,sideEffects=None,groups=plugin.github.com.ellioben,resources=extpowers,verbs=create;update,versions=v1beta1,name=mextpower.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &ExtPower{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
// This webhook will populate the default values for the ExtPower object
func (r *ExtPower) Default() {
	extpowerlog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
	r.Spec.EnableIngress = !r.Spec.EnableIngress
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-plugin-github-com-ellioben-v1beta1-extpower,mutating=false,failurePolicy=fail,sideEffects=None,groups=plugin.github.com.ellioben,resources=extpowers,verbs=create;update,versions=v1beta1,name=vextpower.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &ExtPower{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
//
//	This webhook will validate the ExtPower object
func (r *ExtPower) ValidateCreate() error {
	extpowerlog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *ExtPower) ValidateUpdate(old runtime.Object) error {
	extpowerlog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *ExtPower) ValidateDelete() error {
	extpowerlog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
