/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type RouteMsteamsObservation struct {
}

type RouteMsteamsParameters struct {

	// Enable notification in MS teams. Defaults to `true`.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// MS teams channel id. Alerts will be directed to this channel in Microsoft teams.
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RouteObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RouteParameters struct {

	// The ID of the escalation chain.
	// +kubebuilder:validation:Required
	EscalationChainID *string `json:"escalationChainId" tf:"escalation_chain_id,omitempty"`

	// The ID of the integration.
	// +kubebuilder:validation:Required
	IntegrationID *string `json:"integrationId" tf:"integration_id,omitempty"`

	// MS teams-specific settings for a route.
	// +kubebuilder:validation:Optional
	Msteams []RouteMsteamsParameters `json:"msteams,omitempty" tf:"msteams,omitempty"`

	// The position of the route (starts from 0).
	// +kubebuilder:validation:Required
	Position *float64 `json:"position" tf:"position,omitempty"`

	// Python Regex query. Route is chosen for an alert if there is a match inside the alert payload.
	// +kubebuilder:validation:Required
	RoutingRegex *string `json:"routingRegex" tf:"routing_regex,omitempty"`

	// Slack-specific settings for a route.
	// +kubebuilder:validation:Optional
	Slack []RouteSlackParameters `json:"slack,omitempty" tf:"slack,omitempty"`

	// Telegram-specific settings for a route.
	// +kubebuilder:validation:Optional
	Telegram []RouteTelegramParameters `json:"telegram,omitempty" tf:"telegram,omitempty"`
}

type RouteSlackObservation struct {
}

type RouteSlackParameters struct {

	// Slack channel id. Alerts will be directed to this channel in Slack.
	// +kubebuilder:validation:Optional
	ChannelID *string `json:"channelId,omitempty" tf:"channel_id,omitempty"`

	// Enable notification in Slack. Defaults to `true`.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type RouteTelegramObservation struct {
}

type RouteTelegramParameters struct {

	// Enable notification in Telegram. Defaults to `true`.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Telegram channel id. Alerts will be directed to this channel in Telegram.
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

// RouteSpec defines the desired state of Route
type RouteSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RouteParameters `json:"forProvider"`
}

// RouteStatus defines the observed state of Route.
type RouteStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RouteObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Route is the Schema for the Routes API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,grafana}
type Route struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RouteSpec   `json:"spec"`
	Status            RouteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RouteList contains a list of Routes
type RouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Route `json:"items"`
}

// Repository type metadata.
var (
	Route_Kind             = "Route"
	Route_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Route_Kind}.String()
	Route_KindAPIVersion   = Route_Kind + "." + CRDGroupVersion.String()
	Route_GroupVersionKind = CRDGroupVersion.WithKind(Route_Kind)
)

func init() {
	SchemeBuilder.Register(&Route{}, &RouteList{})
}
