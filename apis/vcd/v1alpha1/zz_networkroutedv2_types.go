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

type NetworkRoutedV2MetadataEntryObservation struct {

	// Domain for this metadata entry. true if it belongs to SYSTEM, false if it belongs to GENERAL
	IsSystem *bool `json:"isSystem,omitempty" tf:"is_system,omitempty"`

	// Key of this metadata entry. Required if the metadata entry is not empty
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Type of this metadata entry. One of: 'MetadataStringValue', 'MetadataNumberValue', 'MetadataBooleanValue', 'MetadataDateTimeValue'
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// User access level for this metadata entry. One of: 'READWRITE', 'READONLY', 'PRIVATE'
	UserAccess *string `json:"userAccess,omitempty" tf:"user_access,omitempty"`

	// Value of this metadata entry. Required if the metadata entry is not empty
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type NetworkRoutedV2MetadataEntryParameters struct {

	// Domain for this metadata entry. true if it belongs to SYSTEM, false if it belongs to GENERAL
	// +kubebuilder:validation:Optional
	IsSystem *bool `json:"isSystem,omitempty" tf:"is_system,omitempty"`

	// Key of this metadata entry. Required if the metadata entry is not empty
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Type of this metadata entry. One of: 'MetadataStringValue', 'MetadataNumberValue', 'MetadataBooleanValue', 'MetadataDateTimeValue'
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// User access level for this metadata entry. One of: 'READWRITE', 'READONLY', 'PRIVATE'
	// +kubebuilder:validation:Optional
	UserAccess *string `json:"userAccess,omitempty" tf:"user_access,omitempty"`

	// Value of this metadata entry. Required if the metadata entry is not empty
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type NetworkRoutedV2Observation struct {

	// DNS suffix
	DNSSuffix *string `json:"dnsSuffix,omitempty" tf:"dns_suffix,omitempty"`

	// Network description
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// DNS server 1
	Dns1 *string `json:"dns1,omitempty" tf:"dns1,omitempty"`

	// DNS server 1
	Dns2 *string `json:"dns2,omitempty" tf:"dns2,omitempty"`

	// Boolean value if Dual-Stack mode should be enabled (default `false`)
	DualStackEnabled *bool `json:"dualStackEnabled,omitempty" tf:"dual_stack_enabled,omitempty"`

	// Edge gateway ID in which Routed network should be located
	EdgeGatewayID *string `json:"edgeGatewayId,omitempty" tf:"edge_gateway_id,omitempty"`

	// Gateway IP address
	Gateway *string `json:"gateway,omitempty" tf:"gateway,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Optional interface type (only for NSX-V networks). One of 'INTERNAL' (default), 'DISTRIBUTED', 'SUBINTERFACE'
	InterfaceType *string `json:"interfaceType,omitempty" tf:"interface_type,omitempty"`

	// Key value map of metadata to assign to this network. Key and value can be any string
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// Metadata entries for the given Network
	MetadataEntry []NetworkRoutedV2MetadataEntryObservation `json:"metadataEntry,omitempty" tf:"metadata_entry,omitempty"`

	// Network name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// ID of VDC or VDC Group
	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	// Network prefix
	PrefixLength *float64 `json:"prefixLength,omitempty" tf:"prefix_length,omitempty"`

	// Secondary gateway (can only be IPv6 and requires enabled Dual Stack mode)
	SecondaryGateway *string `json:"secondaryGateway,omitempty" tf:"secondary_gateway,omitempty"`

	// Secondary prefix (can only be IPv6 and requires enabled Dual Stack mode)
	SecondaryPrefixLength *string `json:"secondaryPrefixLength,omitempty" tf:"secondary_prefix_length,omitempty"`

	// Secondary IP ranges used for static pool allocation in the network
	SecondaryStaticIPPool []NetworkRoutedV2SecondaryStaticIPPoolObservation `json:"secondaryStaticIpPool,omitempty" tf:"secondary_static_ip_pool,omitempty"`

	// IP ranges used for static pool allocation in the network
	StaticIPPool []NetworkRoutedV2StaticIPPoolObservation `json:"staticIpPool,omitempty" tf:"static_ip_pool,omitempty"`

	// The name of VDC to use, optional if defined at provider level
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

type NetworkRoutedV2Parameters struct {

	// DNS suffix
	// +kubebuilder:validation:Optional
	DNSSuffix *string `json:"dnsSuffix,omitempty" tf:"dns_suffix,omitempty"`

	// Network description
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// DNS server 1
	// +kubebuilder:validation:Optional
	Dns1 *string `json:"dns1,omitempty" tf:"dns1,omitempty"`

	// DNS server 1
	// +kubebuilder:validation:Optional
	Dns2 *string `json:"dns2,omitempty" tf:"dns2,omitempty"`

	// Boolean value if Dual-Stack mode should be enabled (default `false`)
	// +kubebuilder:validation:Optional
	DualStackEnabled *bool `json:"dualStackEnabled,omitempty" tf:"dual_stack_enabled,omitempty"`

	// Edge gateway ID in which Routed network should be located
	// +kubebuilder:validation:Optional
	EdgeGatewayID *string `json:"edgeGatewayId,omitempty" tf:"edge_gateway_id,omitempty"`

	// Gateway IP address
	// +kubebuilder:validation:Optional
	Gateway *string `json:"gateway,omitempty" tf:"gateway,omitempty"`

	// Optional interface type (only for NSX-V networks). One of 'INTERNAL' (default), 'DISTRIBUTED', 'SUBINTERFACE'
	// +kubebuilder:validation:Optional
	InterfaceType *string `json:"interfaceType,omitempty" tf:"interface_type,omitempty"`

	// Key value map of metadata to assign to this network. Key and value can be any string
	// +kubebuilder:validation:Optional
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// Metadata entries for the given Network
	// +kubebuilder:validation:Optional
	MetadataEntry []NetworkRoutedV2MetadataEntryParameters `json:"metadataEntry,omitempty" tf:"metadata_entry,omitempty"`

	// Network name
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	// +kubebuilder:validation:Optional
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// Network prefix
	// +kubebuilder:validation:Optional
	PrefixLength *float64 `json:"prefixLength,omitempty" tf:"prefix_length,omitempty"`

	// Secondary gateway (can only be IPv6 and requires enabled Dual Stack mode)
	// +kubebuilder:validation:Optional
	SecondaryGateway *string `json:"secondaryGateway,omitempty" tf:"secondary_gateway,omitempty"`

	// Secondary prefix (can only be IPv6 and requires enabled Dual Stack mode)
	// +kubebuilder:validation:Optional
	SecondaryPrefixLength *string `json:"secondaryPrefixLength,omitempty" tf:"secondary_prefix_length,omitempty"`

	// Secondary IP ranges used for static pool allocation in the network
	// +kubebuilder:validation:Optional
	SecondaryStaticIPPool []NetworkRoutedV2SecondaryStaticIPPoolParameters `json:"secondaryStaticIpPool,omitempty" tf:"secondary_static_ip_pool,omitempty"`

	// IP ranges used for static pool allocation in the network
	// +kubebuilder:validation:Optional
	StaticIPPool []NetworkRoutedV2StaticIPPoolParameters `json:"staticIpPool,omitempty" tf:"static_ip_pool,omitempty"`

	// The name of VDC to use, optional if defined at provider level
	// +kubebuilder:validation:Optional
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

type NetworkRoutedV2SecondaryStaticIPPoolObservation struct {

	// End address of the IP range
	EndAddress *string `json:"endAddress,omitempty" tf:"end_address,omitempty"`

	// Start address of the IP range
	StartAddress *string `json:"startAddress,omitempty" tf:"start_address,omitempty"`
}

type NetworkRoutedV2SecondaryStaticIPPoolParameters struct {

	// End address of the IP range
	// +kubebuilder:validation:Required
	EndAddress *string `json:"endAddress" tf:"end_address,omitempty"`

	// Start address of the IP range
	// +kubebuilder:validation:Required
	StartAddress *string `json:"startAddress" tf:"start_address,omitempty"`
}

type NetworkRoutedV2StaticIPPoolObservation struct {

	// End address of the IP range
	EndAddress *string `json:"endAddress,omitempty" tf:"end_address,omitempty"`

	// Start address of the IP range
	StartAddress *string `json:"startAddress,omitempty" tf:"start_address,omitempty"`
}

type NetworkRoutedV2StaticIPPoolParameters struct {

	// End address of the IP range
	// +kubebuilder:validation:Required
	EndAddress *string `json:"endAddress" tf:"end_address,omitempty"`

	// Start address of the IP range
	// +kubebuilder:validation:Required
	StartAddress *string `json:"startAddress" tf:"start_address,omitempty"`
}

// NetworkRoutedV2Spec defines the desired state of NetworkRoutedV2
type NetworkRoutedV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NetworkRoutedV2Parameters `json:"forProvider"`
}

// NetworkRoutedV2Status defines the observed state of NetworkRoutedV2.
type NetworkRoutedV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NetworkRoutedV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkRoutedV2 is the Schema for the NetworkRoutedV2s API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vcd}
type NetworkRoutedV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.edgeGatewayId)",message="edgeGatewayId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.gateway)",message="gateway is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.prefixLength)",message="prefixLength is a required parameter"
	Spec   NetworkRoutedV2Spec   `json:"spec"`
	Status NetworkRoutedV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkRoutedV2List contains a list of NetworkRoutedV2s
type NetworkRoutedV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkRoutedV2 `json:"items"`
}

// Repository type metadata.
var (
	NetworkRoutedV2_Kind             = "NetworkRoutedV2"
	NetworkRoutedV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NetworkRoutedV2_Kind}.String()
	NetworkRoutedV2_KindAPIVersion   = NetworkRoutedV2_Kind + "." + CRDGroupVersion.String()
	NetworkRoutedV2_GroupVersionKind = CRDGroupVersion.WithKind(NetworkRoutedV2_Kind)
)

func init() {
	SchemeBuilder.Register(&NetworkRoutedV2{}, &NetworkRoutedV2List{})
}
