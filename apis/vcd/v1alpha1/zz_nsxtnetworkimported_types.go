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

type NsxtNetworkImportedObservation struct {

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

	// ID of used Distributed Virtual Port Group
	DvpgID *string `json:"dvpgId,omitempty" tf:"dvpg_id,omitempty"`

	// Name of existing Distributed Virtual Port Group
	DvpgName *string `json:"dvpgName,omitempty" tf:"dvpg_name,omitempty"`

	// Gateway IP address
	Gateway *string `json:"gateway,omitempty" tf:"gateway,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Network name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of used NSX-T Logical Switch
	NsxtLogicalSwitchID *string `json:"nsxtLogicalSwitchId,omitempty" tf:"nsxt_logical_switch_id,omitempty"`

	// Name of existing NSX-T Logical Switch
	NsxtLogicalSwitchName *string `json:"nsxtLogicalSwitchName,omitempty" tf:"nsxt_logical_switch_name,omitempty"`

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
	SecondaryStaticIPPool []NsxtNetworkImportedSecondaryStaticIPPoolObservation `json:"secondaryStaticIpPool,omitempty" tf:"secondary_static_ip_pool,omitempty"`

	// IP ranges used for static pool allocation in the network
	StaticIPPool []NsxtNetworkImportedStaticIPPoolObservation `json:"staticIpPool,omitempty" tf:"static_ip_pool,omitempty"`

	// The name of VDC to use, optional if defined at provider level
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

type NsxtNetworkImportedParameters struct {

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

	// Name of existing Distributed Virtual Port Group
	// +kubebuilder:validation:Optional
	DvpgName *string `json:"dvpgName,omitempty" tf:"dvpg_name,omitempty"`

	// Gateway IP address
	// +kubebuilder:validation:Optional
	Gateway *string `json:"gateway,omitempty" tf:"gateway,omitempty"`

	// Network name
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Name of existing NSX-T Logical Switch
	// +kubebuilder:validation:Optional
	NsxtLogicalSwitchName *string `json:"nsxtLogicalSwitchName,omitempty" tf:"nsxt_logical_switch_name,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	// +kubebuilder:validation:Optional
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// ID of VDC or VDC Group
	// +kubebuilder:validation:Optional
	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

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
	SecondaryStaticIPPool []NsxtNetworkImportedSecondaryStaticIPPoolParameters `json:"secondaryStaticIpPool,omitempty" tf:"secondary_static_ip_pool,omitempty"`

	// IP ranges used for static pool allocation in the network
	// +kubebuilder:validation:Optional
	StaticIPPool []NsxtNetworkImportedStaticIPPoolParameters `json:"staticIpPool,omitempty" tf:"static_ip_pool,omitempty"`

	// The name of VDC to use, optional if defined at provider level
	// +kubebuilder:validation:Optional
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

type NsxtNetworkImportedSecondaryStaticIPPoolObservation struct {

	// End address of the IP range
	EndAddress *string `json:"endAddress,omitempty" tf:"end_address,omitempty"`

	// Start address of the IP range
	StartAddress *string `json:"startAddress,omitempty" tf:"start_address,omitempty"`
}

type NsxtNetworkImportedSecondaryStaticIPPoolParameters struct {

	// End address of the IP range
	// +kubebuilder:validation:Required
	EndAddress *string `json:"endAddress" tf:"end_address,omitempty"`

	// Start address of the IP range
	// +kubebuilder:validation:Required
	StartAddress *string `json:"startAddress" tf:"start_address,omitempty"`
}

type NsxtNetworkImportedStaticIPPoolObservation struct {

	// End address of the IP range
	EndAddress *string `json:"endAddress,omitempty" tf:"end_address,omitempty"`

	// Start address of the IP range
	StartAddress *string `json:"startAddress,omitempty" tf:"start_address,omitempty"`
}

type NsxtNetworkImportedStaticIPPoolParameters struct {

	// End address of the IP range
	// +kubebuilder:validation:Required
	EndAddress *string `json:"endAddress" tf:"end_address,omitempty"`

	// Start address of the IP range
	// +kubebuilder:validation:Required
	StartAddress *string `json:"startAddress" tf:"start_address,omitempty"`
}

// NsxtNetworkImportedSpec defines the desired state of NsxtNetworkImported
type NsxtNetworkImportedSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NsxtNetworkImportedParameters `json:"forProvider"`
}

// NsxtNetworkImportedStatus defines the observed state of NsxtNetworkImported.
type NsxtNetworkImportedStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NsxtNetworkImportedObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NsxtNetworkImported is the Schema for the NsxtNetworkImporteds API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vcd}
type NsxtNetworkImported struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.gateway)",message="gateway is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.prefixLength)",message="prefixLength is a required parameter"
	Spec   NsxtNetworkImportedSpec   `json:"spec"`
	Status NsxtNetworkImportedStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NsxtNetworkImportedList contains a list of NsxtNetworkImporteds
type NsxtNetworkImportedList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NsxtNetworkImported `json:"items"`
}

// Repository type metadata.
var (
	NsxtNetworkImported_Kind             = "NsxtNetworkImported"
	NsxtNetworkImported_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NsxtNetworkImported_Kind}.String()
	NsxtNetworkImported_KindAPIVersion   = NsxtNetworkImported_Kind + "." + CRDGroupVersion.String()
	NsxtNetworkImported_GroupVersionKind = CRDGroupVersion.WithKind(NsxtNetworkImported_Kind)
)

func init() {
	SchemeBuilder.Register(&NsxtNetworkImported{}, &NsxtNetworkImportedList{})
}
