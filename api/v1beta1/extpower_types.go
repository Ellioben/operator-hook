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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ExtPowerSpec defines the desired state of ExtPower
type ExtPowerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of ExtPower. Edit extpower_types.go to remove/update
	EnableIngress bool   `json:"enable_ingress,omitempty"`
	EnableService bool   `json:"enable_service"`
	Replicas      int32  `json:"replicas"`
	Image         string `json:"image"`
}

// ExtPowerStatus defines the observed state of ExtPower
type ExtPowerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ExtPower is the Schema for the extpowers API
type ExtPower struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ExtPowerSpec   `json:"spec,omitempty"`
	Status ExtPowerStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ExtPowerList contains a list of ExtPower
type ExtPowerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ExtPower `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ExtPower{}, &ExtPowerList{})
}
