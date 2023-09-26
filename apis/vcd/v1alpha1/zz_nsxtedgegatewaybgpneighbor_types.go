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

type NsxtEdgeGatewayBgpNeighborObservation struct {

	// A flag indicating whether BGP neighbors can receive routes with same Autonomous System (AS) (default 'false')
	AllowAsIn *bool `json:"allowAsIn,omitempty" tf:"allow_as_in,omitempty"`

	// Number of times a heartbeat packet is missed before BFD declares that the neighbor is down
	BfdDeadMultiple *float64 `json:"bfdDeadMultiple,omitempty" tf:"bfd_dead_multiple,omitempty"`

	// BFD configuration for failure detection
	BfdEnabled *bool `json:"bfdEnabled,omitempty" tf:"bfd_enabled,omitempty"`

	// Time interval (in milliseconds) between heartbeat packets
	BfdInterval *float64 `json:"bfdInterval,omitempty" tf:"bfd_interval,omitempty"`

	// Edge gateway ID for BGP Neighbor Configuration
	EdgeGatewayID *string `json:"edgeGatewayId,omitempty" tf:"edge_gateway_id,omitempty"`

	// One of 'DISABLE', 'HELPER_ONLY', 'GRACEFUL_AND_HELPER'
	GracefulRestartMode *string `json:"gracefulRestartMode,omitempty" tf:"graceful_restart_mode,omitempty"`

	// Time interval (in seconds) before declaring a peer dead
	HoldDownTimer *float64 `json:"holdDownTimer,omitempty" tf:"hold_down_timer,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// BGP Neighbor IP address (IPv4 or IPv6)
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// An optional IP Prefix List ID for filtering 'IN' direction.
	InFilterIPPrefixListID *string `json:"inFilterIpPrefixListId,omitempty" tf:"in_filter_ip_prefix_list_id,omitempty"`

	// Time interval (in seconds) between sending keep alive messages to a peer
	KeepAliveTimer *float64 `json:"keepAliveTimer,omitempty" tf:"keep_alive_timer,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// An optional IP Prefix List ID for filtering 'OUT' direction.
	OutFilterIPPrefixListID *string `json:"outFilterIpPrefixListId,omitempty" tf:"out_filter_ip_prefix_list_id,omitempty"`

	// Remote Autonomous System (AS) number
	RemoteAsNumber *string `json:"remoteAsNumber,omitempty" tf:"remote_as_number,omitempty"`

	// One of 'DISABLED', 'IPV4', 'IPV6'
	RouteFiltering *string `json:"routeFiltering,omitempty" tf:"route_filtering,omitempty"`
}

type NsxtEdgeGatewayBgpNeighborParameters struct {

	// A flag indicating whether BGP neighbors can receive routes with same Autonomous System (AS) (default 'false')
	// +kubebuilder:validation:Optional
	AllowAsIn *bool `json:"allowAsIn,omitempty" tf:"allow_as_in,omitempty"`

	// Number of times a heartbeat packet is missed before BFD declares that the neighbor is down
	// +kubebuilder:validation:Optional
	BfdDeadMultiple *float64 `json:"bfdDeadMultiple,omitempty" tf:"bfd_dead_multiple,omitempty"`

	// BFD configuration for failure detection
	// +kubebuilder:validation:Optional
	BfdEnabled *bool `json:"bfdEnabled,omitempty" tf:"bfd_enabled,omitempty"`

	// Time interval (in milliseconds) between heartbeat packets
	// +kubebuilder:validation:Optional
	BfdInterval *float64 `json:"bfdInterval,omitempty" tf:"bfd_interval,omitempty"`

	// Edge gateway ID for BGP Neighbor Configuration
	// +kubebuilder:validation:Optional
	EdgeGatewayID *string `json:"edgeGatewayId,omitempty" tf:"edge_gateway_id,omitempty"`

	// One of 'DISABLE', 'HELPER_ONLY', 'GRACEFUL_AND_HELPER'
	// +kubebuilder:validation:Optional
	GracefulRestartMode *string `json:"gracefulRestartMode,omitempty" tf:"graceful_restart_mode,omitempty"`

	// Time interval (in seconds) before declaring a peer dead
	// +kubebuilder:validation:Optional
	HoldDownTimer *float64 `json:"holdDownTimer,omitempty" tf:"hold_down_timer,omitempty"`

	// BGP Neighbor IP address (IPv4 or IPv6)
	// +kubebuilder:validation:Optional
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// An optional IP Prefix List ID for filtering 'IN' direction.
	// +kubebuilder:validation:Optional
	InFilterIPPrefixListID *string `json:"inFilterIpPrefixListId,omitempty" tf:"in_filter_ip_prefix_list_id,omitempty"`

	// Time interval (in seconds) between sending keep alive messages to a peer
	// +kubebuilder:validation:Optional
	KeepAliveTimer *float64 `json:"keepAliveTimer,omitempty" tf:"keep_alive_timer,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	// +kubebuilder:validation:Optional
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// An optional IP Prefix List ID for filtering 'OUT' direction.
	// +kubebuilder:validation:Optional
	OutFilterIPPrefixListID *string `json:"outFilterIpPrefixListId,omitempty" tf:"out_filter_ip_prefix_list_id,omitempty"`

	// Neighbor password
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// Remote Autonomous System (AS) number
	// +kubebuilder:validation:Optional
	RemoteAsNumber *string `json:"remoteAsNumber,omitempty" tf:"remote_as_number,omitempty"`

	// One of 'DISABLED', 'IPV4', 'IPV6'
	// +kubebuilder:validation:Optional
	RouteFiltering *string `json:"routeFiltering,omitempty" tf:"route_filtering,omitempty"`
}

// NsxtEdgeGatewayBgpNeighborSpec defines the desired state of NsxtEdgeGatewayBgpNeighbor
type NsxtEdgeGatewayBgpNeighborSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NsxtEdgeGatewayBgpNeighborParameters `json:"forProvider"`
}

// NsxtEdgeGatewayBgpNeighborStatus defines the observed state of NsxtEdgeGatewayBgpNeighbor.
type NsxtEdgeGatewayBgpNeighborStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NsxtEdgeGatewayBgpNeighborObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NsxtEdgeGatewayBgpNeighbor is the Schema for the NsxtEdgeGatewayBgpNeighbors API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vcd}
type NsxtEdgeGatewayBgpNeighbor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.edgeGatewayId)",message="edgeGatewayId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.ipAddress)",message="ipAddress is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.remoteAsNumber)",message="remoteAsNumber is a required parameter"
	Spec   NsxtEdgeGatewayBgpNeighborSpec   `json:"spec"`
	Status NsxtEdgeGatewayBgpNeighborStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NsxtEdgeGatewayBgpNeighborList contains a list of NsxtEdgeGatewayBgpNeighbors
type NsxtEdgeGatewayBgpNeighborList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NsxtEdgeGatewayBgpNeighbor `json:"items"`
}

// Repository type metadata.
var (
	NsxtEdgeGatewayBgpNeighbor_Kind             = "NsxtEdgeGatewayBgpNeighbor"
	NsxtEdgeGatewayBgpNeighbor_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NsxtEdgeGatewayBgpNeighbor_Kind}.String()
	NsxtEdgeGatewayBgpNeighbor_KindAPIVersion   = NsxtEdgeGatewayBgpNeighbor_Kind + "." + CRDGroupVersion.String()
	NsxtEdgeGatewayBgpNeighbor_GroupVersionKind = CRDGroupVersion.WithKind(NsxtEdgeGatewayBgpNeighbor_Kind)
)

func init() {
	SchemeBuilder.Register(&NsxtEdgeGatewayBgpNeighbor{}, &NsxtEdgeGatewayBgpNeighborList{})
}