/*
Copyright 2019 Tyler French.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// VcdClusterProviderSpecSpec defines the desired state of VcdClusterProviderSpec
type VcdClusterProviderSpecSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// VcdClusterProviderSpecStatus defines the observed state of VcdClusterProviderSpec
type VcdClusterProviderSpecStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VcdClusterProviderSpec is the Schema for the vcdclusterproviderspecs API
// +k8s:openapi-gen=true
type VcdClusterProviderSpec struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VcdClusterProviderSpecSpec   `json:"spec,omitempty"`
	Status VcdClusterProviderSpecStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VcdClusterProviderSpecList contains a list of VcdClusterProviderSpec
type VcdClusterProviderSpecList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VcdClusterProviderSpec `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VcdClusterProviderSpec{}, &VcdClusterProviderSpecList{})
}
