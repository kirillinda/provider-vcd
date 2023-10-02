/*
Copyright 2023 ANKASOFT
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SecurityTagObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Security tag name to be created
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// List of VM IDs that the security tags is going to be tied to
	VMIds []*string `json:"vmIds,omitempty" tf:"vm_ids,omitempty"`
}

type SecurityTagParameters struct {

	// Security tag name to be created
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	// +kubebuilder:validation:Optional
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// List of VM IDs that the security tags is going to be tied to
	// +kubebuilder:validation:Optional
	VMIds []*string `json:"vmIds,omitempty" tf:"vm_ids,omitempty"`
}

// SecurityTagSpec defines the desired state of SecurityTag
type SecurityTagSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecurityTagParameters `json:"forProvider"`
}

// SecurityTagStatus defines the observed state of SecurityTag.
type SecurityTagStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecurityTagObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityTag is the Schema for the SecurityTags API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vcd}
type SecurityTag struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.vmIds)",message="vmIds is a required parameter"
	Spec   SecurityTagSpec   `json:"spec"`
	Status SecurityTagStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityTagList contains a list of SecurityTags
type SecurityTagList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityTag `json:"items"`
}

// Repository type metadata.
var (
	SecurityTag_Kind             = "SecurityTag"
	SecurityTag_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecurityTag_Kind}.String()
	SecurityTag_KindAPIVersion   = SecurityTag_Kind + "." + CRDGroupVersion.String()
	SecurityTag_GroupVersionKind = CRDGroupVersion.WithKind(SecurityTag_Kind)
)

func init() {
	SchemeBuilder.Register(&SecurityTag{}, &SecurityTagList{})
}