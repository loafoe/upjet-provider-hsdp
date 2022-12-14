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

type EmailTemplateObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	MessageBase64 *string `json:"messageBase64,omitempty" tf:"message_base64,omitempty"`
}

type EmailTemplateParameters struct {

	// +kubebuilder:validation:Optional
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// +kubebuilder:validation:Optional
	From *string `json:"from,omitempty" tf:"from,omitempty"`

	// +kubebuilder:validation:Optional
	Link *string `json:"link,omitempty" tf:"link,omitempty"`

	// +kubebuilder:validation:Optional
	Locale *string `json:"locale,omitempty" tf:"locale,omitempty"`

	// +crossplane:generate:reference:type=Organization
	// +crossplane:generate:reference:refFieldName=OrganizationRef
	// +kubebuilder:validation:Optional
	ManagingOrganization *string `json:"managingOrganization,omitempty" tf:"managing_organization,omitempty"`

	// Selector for a Organization to populate managingOrganization.
	// +kubebuilder:validation:Optional
	ManagingOrganizationSelector *v1.Selector `json:"managingOrganizationSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Message *string `json:"message" tf:"message,omitempty"`

	// Reference to a Organization to populate managingOrganization.
	// +kubebuilder:validation:Optional
	OrganizationRef *v1.Reference `json:"organizationRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// EmailTemplateSpec defines the desired state of EmailTemplate
type EmailTemplateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EmailTemplateParameters `json:"forProvider"`
}

// EmailTemplateStatus defines the observed state of EmailTemplate.
type EmailTemplateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EmailTemplateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EmailTemplate is the Schema for the EmailTemplates API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,hsdpjet}
type EmailTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EmailTemplateSpec   `json:"spec"`
	Status            EmailTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EmailTemplateList contains a list of EmailTemplates
type EmailTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EmailTemplate `json:"items"`
}

// Repository type metadata.
var (
	EmailTemplate_Kind             = "EmailTemplate"
	EmailTemplate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EmailTemplate_Kind}.String()
	EmailTemplate_KindAPIVersion   = EmailTemplate_Kind + "." + CRDGroupVersion.String()
	EmailTemplate_GroupVersionKind = CRDGroupVersion.WithKind(EmailTemplate_Kind)
)

func init() {
	SchemeBuilder.Register(&EmailTemplate{}, &EmailTemplateList{})
}
