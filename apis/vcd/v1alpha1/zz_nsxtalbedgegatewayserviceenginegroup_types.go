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

type NsxtAlbEdgegatewayServiceEngineGroupObservation struct {

	// Number of deployed virtual services for this Service Engine Group
	DeployedVirtualServices *float64 `json:"deployedVirtualServices,omitempty" tf:"deployed_virtual_services,omitempty"`

	// Edge Gateway ID in which ALB Service Engine Group should be located
	EdgeGatewayID *string `json:"edgeGatewayId,omitempty" tf:"edge_gateway_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Maximum number of virtual services to be used in this Service Engine Group
	MaxVirtualServices *float64 `json:"maxVirtualServices,omitempty" tf:"max_virtual_services,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// Number of reserved virtual services for this Service Engine Group
	ReservedVirtualServices *string `json:"reservedVirtualServices,omitempty" tf:"reserved_virtual_services,omitempty"`

	// Service Engine Group ID to attach to this NSX-T Edge Gateway
	ServiceEngineGroupID *string `json:"serviceEngineGroupId,omitempty" tf:"service_engine_group_id,omitempty"`

	// Service Engine Group Name which is attached to NSX-T Edge Gateway
	ServiceEngineGroupName *string `json:"serviceEngineGroupName,omitempty" tf:"service_engine_group_name,omitempty"`

	// The name of VDC to use, optional if defined at provider level
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

type NsxtAlbEdgegatewayServiceEngineGroupParameters struct {

	// Edge Gateway ID in which ALB Service Engine Group should be located
	// +kubebuilder:validation:Optional
	EdgeGatewayID *string `json:"edgeGatewayId,omitempty" tf:"edge_gateway_id,omitempty"`

	// Maximum number of virtual services to be used in this Service Engine Group
	// +kubebuilder:validation:Optional
	MaxVirtualServices *float64 `json:"maxVirtualServices,omitempty" tf:"max_virtual_services,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	// +kubebuilder:validation:Optional
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// Number of reserved virtual services for this Service Engine Group
	// +kubebuilder:validation:Optional
	ReservedVirtualServices *string `json:"reservedVirtualServices,omitempty" tf:"reserved_virtual_services,omitempty"`

	// Service Engine Group ID to attach to this NSX-T Edge Gateway
	// +kubebuilder:validation:Optional
	ServiceEngineGroupID *string `json:"serviceEngineGroupId,omitempty" tf:"service_engine_group_id,omitempty"`

	// The name of VDC to use, optional if defined at provider level
	// +kubebuilder:validation:Optional
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

// NsxtAlbEdgegatewayServiceEngineGroupSpec defines the desired state of NsxtAlbEdgegatewayServiceEngineGroup
type NsxtAlbEdgegatewayServiceEngineGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NsxtAlbEdgegatewayServiceEngineGroupParameters `json:"forProvider"`
}

// NsxtAlbEdgegatewayServiceEngineGroupStatus defines the observed state of NsxtAlbEdgegatewayServiceEngineGroup.
type NsxtAlbEdgegatewayServiceEngineGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NsxtAlbEdgegatewayServiceEngineGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NsxtAlbEdgegatewayServiceEngineGroup is the Schema for the NsxtAlbEdgegatewayServiceEngineGroups API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vcd}
type NsxtAlbEdgegatewayServiceEngineGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.edgeGatewayId)",message="edgeGatewayId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.serviceEngineGroupId)",message="serviceEngineGroupId is a required parameter"
	Spec   NsxtAlbEdgegatewayServiceEngineGroupSpec   `json:"spec"`
	Status NsxtAlbEdgegatewayServiceEngineGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NsxtAlbEdgegatewayServiceEngineGroupList contains a list of NsxtAlbEdgegatewayServiceEngineGroups
type NsxtAlbEdgegatewayServiceEngineGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NsxtAlbEdgegatewayServiceEngineGroup `json:"items"`
}

// Repository type metadata.
var (
	NsxtAlbEdgegatewayServiceEngineGroup_Kind             = "NsxtAlbEdgegatewayServiceEngineGroup"
	NsxtAlbEdgegatewayServiceEngineGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NsxtAlbEdgegatewayServiceEngineGroup_Kind}.String()
	NsxtAlbEdgegatewayServiceEngineGroup_KindAPIVersion   = NsxtAlbEdgegatewayServiceEngineGroup_Kind + "." + CRDGroupVersion.String()
	NsxtAlbEdgegatewayServiceEngineGroup_GroupVersionKind = CRDGroupVersion.WithKind(NsxtAlbEdgegatewayServiceEngineGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&NsxtAlbEdgegatewayServiceEngineGroup{}, &NsxtAlbEdgegatewayServiceEngineGroupList{})
}
