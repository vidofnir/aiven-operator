// Copyright (c) 2024 Aiven, Helsinki, Finland. https://aiven.io/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ClickhouseUserSpec defines the desired state of ClickhouseUser
// +kubebuilder:validation:XValidation:rule="has(oldSelf.connInfoSecretTargetDisabled) == has(self.connInfoSecretTargetDisabled)",message="connInfoSecretTargetDisabled can only be set during resource creation."
type ClickhouseUserSpec struct {
	// +kubebuilder:validation:MaxLength=63
	// +kubebuilder:validation:Format="^[a-zA-Z0-9_-]*$"
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Value is immutable"
	// Project to link the user to
	Project string `json:"project"`

	// +kubebuilder:validation:MaxLength=63
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Value is immutable"
	// Service to link the user to
	ServiceName string `json:"serviceName"`

	// Information regarding secret creation.
	// Exposed keys: `CLICKHOUSEUSER_HOST`, `CLICKHOUSEUSER_PORT`, `CLICKHOUSEUSER_USER`, `CLICKHOUSEUSER_PASSWORD`
	ConnInfoSecretTarget ConnInfoSecretTarget `json:"connInfoSecretTarget,omitempty"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="connInfoSecretTargetDisabled is immutable."
	// When true, the secret containing connection information will not be created, defaults to false. This field cannot be changed after resource creation.
	ConnInfoSecretTargetDisabled *bool `json:"connInfoSecretTargetDisabled,omitempty"`

	// Authentication reference to Aiven token in a secret
	AuthSecretRef *AuthSecretReference `json:"authSecretRef,omitempty"`
}

// ClickhouseUserStatus defines the observed state of ClickhouseUser
type ClickhouseUserStatus struct {
	// Clickhouse user UUID
	UUID string `json:"uuid"`

	// Conditions represent the latest available observations of an ClickhouseUser state
	// +kubebuilder:validation:type=array
	Conditions []metav1.Condition `json:"conditions"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ClickhouseUser is the Schema for the clickhouseusers API
// +kubebuilder:printcolumn:name="Service Name",type="string",JSONPath=".spec.serviceName"
// +kubebuilder:printcolumn:name="Project",type="string",JSONPath=".spec.project"
// +kubebuilder:printcolumn:name="Connection Information Secret",type="string",JSONPath=".spec.connInfoSecretTarget.name"
type ClickhouseUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClickhouseUserSpec   `json:"spec,omitempty"`
	Status ClickhouseUserStatus `json:"status,omitempty"`
}

var _ AivenManagedObject = &ClickhouseUser{}

func (in *ClickhouseUser) NoSecret() bool {
	return in.Spec.ConnInfoSecretTargetDisabled != nil && *in.Spec.ConnInfoSecretTargetDisabled
}

func (in *ClickhouseUser) AuthSecretRef() *AuthSecretReference {
	return in.Spec.AuthSecretRef
}

func (in *ClickhouseUser) Conditions() *[]metav1.Condition {
	return &in.Status.Conditions
}

func (in *ClickhouseUser) GetConnInfoSecretTarget() ConnInfoSecretTarget {
	return in.Spec.ConnInfoSecretTarget
}

//+kubebuilder:object:root=true

// ClickhouseUserList contains a list of ClickhouseUser
type ClickhouseUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClickhouseUser `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClickhouseUser{}, &ClickhouseUserList{})
}
