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

type ProbeObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The tenant ID of the probe.
	TenantID *float64 `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type ProbeParameters struct {

	// Custom labels to be included with collected metrics and logs.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Latitude coordinates.
	// +kubebuilder:validation:Required
	Latitude *float64 `json:"latitude" tf:"latitude,omitempty"`

	// Longitude coordinates.
	// +kubebuilder:validation:Required
	Longitude *float64 `json:"longitude" tf:"longitude,omitempty"`

	// Name of the probe.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Public probes are run by Grafana Labs and can be used by all users. Only Grafana Labs managed public probes will be set to `true`. Defaults to `false`.
	// +kubebuilder:validation:Optional
	Public *bool `json:"public,omitempty" tf:"public,omitempty"`

	// Region of the probe.
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`
}

// ProbeSpec defines the desired state of Probe
type ProbeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProbeParameters `json:"forProvider"`
}

// ProbeStatus defines the observed state of Probe.
type ProbeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProbeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Probe is the Schema for the Probes API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,grafana}
type Probe struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProbeSpec   `json:"spec"`
	Status            ProbeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProbeList contains a list of Probes
type ProbeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Probe `json:"items"`
}

// Repository type metadata.
var (
	Probe_Kind             = "Probe"
	Probe_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Probe_Kind}.String()
	Probe_KindAPIVersion   = Probe_Kind + "." + CRDGroupVersion.String()
	Probe_GroupVersionKind = CRDGroupVersion.WithKind(Probe_Kind)
)

func init() {
	SchemeBuilder.Register(&Probe{}, &ProbeList{})
}
