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

type ClientObservation struct {
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ClientParameters struct {

	// +kubebuilder:validation:Optional
	AccessTokenLifetime *float64 `json:"accessTokenLifetime,omitempty" tf:"access_token_lifetime,omitempty"`

	// +crossplane:generate:reference:type=Application
	// +crossplane:generate:reference:refFieldName=ApplicationRef
	// +kubebuilder:validation:Optional
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// Selector for a Application to populate applicationId.
	// +kubebuilder:validation:Optional
	ApplicationIDSelector *v1.Selector `json:"applicationIdSelector,omitempty" tf:"-"`

	// Reference to a Application to populate applicationId.
	// +kubebuilder:validation:Optional
	ApplicationRef *v1.Reference `json:"applicationRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	ClientID *string `json:"clientId" tf:"client_id,omitempty"`

	// +kubebuilder:validation:Optional
	ConsentImplied *bool `json:"consentImplied,omitempty" tf:"consent_implied,omitempty"`

	// +kubebuilder:validation:Required
	DefaultScopes []*string `json:"defaultScopes" tf:"default_scopes,omitempty"`

	// +kubebuilder:validation:Required
	Description *string `json:"description" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	GlobalReferenceID *string `json:"globalReferenceId" tf:"global_reference_id,omitempty"`

	// +kubebuilder:validation:Optional
	IDTokenLifetime *float64 `json:"idTokenLifetime,omitempty" tf:"id_token_lifetime,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	RedirectionUris []*string `json:"redirectionUris" tf:"redirection_uris,omitempty"`

	// +kubebuilder:validation:Optional
	RefreshTokenLifetime *float64 `json:"refreshTokenLifetime,omitempty" tf:"refresh_token_lifetime,omitempty"`

	// +kubebuilder:validation:Required
	ResponseTypes []*string `json:"responseTypes" tf:"response_types,omitempty"`

	// +kubebuilder:validation:Required
	Scopes []*string `json:"scopes" tf:"scopes,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// ClientSpec defines the desired state of Client
type ClientSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClientParameters `json:"forProvider"`
}

// ClientStatus defines the observed state of Client.
type ClientStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClientObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Client is the Schema for the Clients API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,hsdpjet}
type Client struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClientSpec   `json:"spec"`
	Status            ClientStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClientList contains a list of Clients
type ClientList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Client `json:"items"`
}

// Repository type metadata.
var (
	Client_Kind             = "Client"
	Client_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Client_Kind}.String()
	Client_KindAPIVersion   = Client_Kind + "." + CRDGroupVersion.String()
	Client_GroupVersionKind = CRDGroupVersion.WithKind(Client_Kind)
)

func init() {
	SchemeBuilder.Register(&Client{}, &ClientList{})
}
