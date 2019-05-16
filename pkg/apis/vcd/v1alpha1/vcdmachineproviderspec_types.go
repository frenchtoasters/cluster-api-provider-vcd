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
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/cluster-api/pkg/apis/cluster/common"
	"sigs.k8s.io/cluster-api/pkg/apis/cluster/v1alpha1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// VcdMachineProviderSpecSpec defines the desired state of VcdMachineProviderSpec
type VcdMachineProviderSpecSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Taints []v1.Taint `json:"taints,omitempty"`

	ProviderSpec v1alpha1.ProviderSpec `json:"providerSpec"`

	Versions v1alpha1.MachineVersionInfo `json:"versions,omitempty"`

	ConfigSource *v1.NodeConfigSource `json:"configSource,omitempty"`
}

// VcdMachineProviderSpecStatus defines the observed state of VcdMachineProviderSpec
type VcdMachineProviderSpecStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	NodeRef *v1.ObjectReference `json:"nodeRef,omitempty"`

	LastUpdated *metav1.Time `json:"lastUpdated,omitempty"`

	Versions *v1alpha1.MachineVersionInfo `json:"versions,omitempty"`

	ErrorReason *common.MachineStatusError `json:"errorReason,omitempty"`

	ErrorMessage *string `json:"errorMessage,omitempty"`

	ProviderStatus *runtime.RawExtension `json:"providerStatus,omitempty"`

	Addresses []v1.NodeAddress `json:"addresses,omitempty"`

	Conditions []v1.NodeCondition `json:"conditions,omitempty"`

	LastOperation *v1alpha1.LastOperation `json:"lastOperation,omitempty"`

	Phase *string `json:"phase,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VcdMachineProviderSpec is the Schema for the vcdmachineproviderspecs API
// +k8s:openapi-gen=true
type VcdMachineProviderSpec struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VcdMachineProviderSpecSpec   `json:"spec,omitempty"`
	Status VcdMachineProviderSpecStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VcdMachineProviderSpecList contains a list of VcdMachineProviderSpec
type VcdMachineProviderSpecList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VcdMachineProviderSpec `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VcdMachineProviderSpec{}, &VcdMachineProviderSpecList{})
}
