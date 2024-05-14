/*
Copyright(c) 2023-present Accton. All rights reserved. www.accton.com.tw
*/

package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type EvenSpec struct {
	DeploymentName string `json:"deploymentname"`
	// +kubebuilder:validation:Maximum=5
	// +kubebuilder:validation:Minimum=1
	Replicas int `json:"replicas,omitempty"`
}

type EvenStatus struct {
	Replicas int `json:"replicas,omitempty"`
}

// +kubebuilder:subresource:status
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Even is the Schema for the Even API
type Even struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EvenSpec   `json:"spec,omitempty"`
	Status            EvenStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// EvenList is the Schema for the Even API
type EvenList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Even `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Even{}, &EvenList{})
}
