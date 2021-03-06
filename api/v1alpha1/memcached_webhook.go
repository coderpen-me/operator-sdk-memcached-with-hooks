/*
Copyright 2021.

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
	"encoding/json"
	"fmt"

	schema "github.com/Kangaroux/go-map-schema"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var memcachedlog = logf.Log.WithName("memcached-resource")

func (r *Memcached) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-cache-example-com-v1alpha1-memcached,mutating=true,failurePolicy=fail,sideEffects=None,groups=cache.example.com,resources=memcacheds,verbs=create;update,versions=v1alpha1,name=mmemcached.kb.io,admissionReviewVersions={v1,v1beta1}

var _ webhook.Defaulter = &Memcached{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *Memcached) Default() {
	memcachedlog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-cache-example-com-v1alpha1-memcached,mutating=false,failurePolicy=fail,sideEffects=None,groups=cache.example.com,resources=memcacheds,verbs=create;update,versions=v1alpha1,name=vmemcached.kb.io,admissionReviewVersions={v1,v1beta1}

var _ webhook.Validator = &Memcached{}

func TypeConversion(from schema.CompareResults) ([]MissingField, []MismatchedField) {
	var toMissing []MissingField
	var toMismatched []MismatchedField

	for _, b := range from.MissingFields {
		newMissingField := MissingField{Field: b.Field, Path: b.Path} // pulled out for clarity
		toMissing = append(toMissing, newMissingField)
	}

	for _, b := range from.MismatchedFields {
		newMismatchedField := MismatchedField{Actual: b.Actual, Expected: b.Expected, Field: b.Field, Path: b.Path} // pulled out for clarity
		toMismatched = append(toMismatched, newMismatchedField)
	}

	return toMissing, toMismatched
}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Memcached) ValidateCreate() error {
	memcachedlog.Info("validate create", "name", r.Name)

	resp := ValidatorResponse{}
	resp.Status = true
	resp.PluginError = make(map[string]string)

	for key, value := range r.Spec.Config {
		m := make(map[string]interface{})
		switch key {
		case "plugin_1_name":
			{
				p := Plugin1{}
				// Load the JSON into a map.
				if err := json.Unmarshal(value.Raw, &m); err != nil {
					resp.PluginError[key] = "Plugin format is invalid"
					resp.Status = false
				}

				// Check the types.
				rs, err := schema.CompareMapToStruct(&p, m, nil)

				if err != nil {
					resp.PluginError[key] = "Plugin format is invalid"
					resp.Status = false
				}

				missingFields, mismatchedFields := TypeConversion(*rs)

				resp.Data.MissingFields = append(resp.Data.MissingFields, missingFields...)
				resp.Data.MismatchedFields = append(resp.Data.MismatchedFields, mismatchedFields...)
			}
		case "plugin_2_name":
			{
				p := Plugin2{}
				// Load the JSON into a map.
				if err := json.Unmarshal(value.Raw, &m); err != nil {
					resp.PluginError[key] = "Plugin format is invalid"
					resp.Status = false
				}

				// Check the types.
				rs, err := schema.CompareMapToStruct(&p, m, nil)

				if err != nil {
					resp.PluginError[key] = "Plugin format is invalid"
					resp.Status = false
				}

				missingFields, mismatchedFields := TypeConversion(*rs)

				resp.Data.MissingFields = append(resp.Data.MissingFields, missingFields...)
				resp.Data.MismatchedFields = append(resp.Data.MismatchedFields, mismatchedFields...)
			}
		default:
			{
				resp.PluginError[key] = "Plugin format is invalid"
				resp.Status = false
			}
		}
	}

	if !resp.Status {
		memcachedlog.Info("validate create error", "mismatch", (resp.Data.MismatchedFields), "missing", resp.Data.MissingFields, "plugin_error", resp.PluginError)
		return fmt.Errorf("config format invalid")
	}

	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Memcached) ValidateUpdate(old runtime.Object) error {
	memcachedlog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Memcached) ValidateDelete() error {
	memcachedlog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
