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

type FolderPermissionObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type FolderPermissionParameters struct {

	// Reference to a Folder in oss to populate folderUid.
	// +kubebuilder:validation:Optional
	FolderRef *v1.Reference `json:"folderRef,omitempty" tf:"-"`

	// Selector for a Folder in oss to populate folderUid.
	// +kubebuilder:validation:Optional
	FolderSelector *v1.Selector `json:"folderSelector,omitempty" tf:"-"`

	// The UID of the folder.
	// +crossplane:generate:reference:type=github.com/grafana/crossplane-provider-grafana/apis/oss/v1alpha1.Folder
	// +crossplane:generate:reference:extractor=github.com/grafana/crossplane-provider-grafana/config/grafana.UIDExtractor()
	// +crossplane:generate:reference:refFieldName=FolderRef
	// +crossplane:generate:reference:selectorFieldName=FolderSelector
	// +kubebuilder:validation:Optional
	FolderUID *string `json:"folderUid,omitempty" tf:"folder_uid,omitempty"`

	// The permission items to add/update. Items that are omitted from the list will be removed.
	// +kubebuilder:validation:Required
	Permissions []FolderPermissionPermissionsParameters `json:"permissions" tf:"permissions,omitempty"`
}

type FolderPermissionPermissionsObservation struct {
}

type FolderPermissionPermissionsParameters struct {

	// Permission to associate with item. Must be one of `View`, `Edit`, or `Admin`.
	// +kubebuilder:validation:Required
	Permission *string `json:"permission" tf:"permission,omitempty"`

	// Manage permissions for `Viewer` or `Editor` roles.
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// ID of the team to manage permissions for. Defaults to `0`.
	// +kubebuilder:validation:Optional
	TeamID *float64 `json:"teamId,omitempty" tf:"team_id,omitempty"`

	// ID of the user to manage permissions for. Defaults to `0`.
	// +kubebuilder:validation:Optional
	UserID *float64 `json:"userId,omitempty" tf:"user_id,omitempty"`
}

// FolderPermissionSpec defines the desired state of FolderPermission
type FolderPermissionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FolderPermissionParameters `json:"forProvider"`
}

// FolderPermissionStatus defines the observed state of FolderPermission.
type FolderPermissionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FolderPermissionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FolderPermission is the Schema for the FolderPermissions API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,grafana}
type FolderPermission struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FolderPermissionSpec   `json:"spec"`
	Status            FolderPermissionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FolderPermissionList contains a list of FolderPermissions
type FolderPermissionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FolderPermission `json:"items"`
}

// Repository type metadata.
var (
	FolderPermission_Kind             = "FolderPermission"
	FolderPermission_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FolderPermission_Kind}.String()
	FolderPermission_KindAPIVersion   = FolderPermission_Kind + "." + CRDGroupVersion.String()
	FolderPermission_GroupVersionKind = CRDGroupVersion.WithKind(FolderPermission_Kind)
)

func init() {
	SchemeBuilder.Register(&FolderPermission{}, &FolderPermissionList{})
}
