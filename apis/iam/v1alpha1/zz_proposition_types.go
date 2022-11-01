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

type PropositionObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type PropositionParameters struct {

	// +kubebuilder:validation:Required
	Description *string `json:"description" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	GlobalReferenceID *string `json:"globalReferenceId,omitempty" tf:"global_reference_id,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +crossplane:generate:reference:type=Organization
	// +crossplane:generate:reference:refFieldName=OrganizationRef
	// +kubebuilder:validation:Optional
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// Selector for a Organization to populate organizationId.
	// +kubebuilder:validation:Optional
	OrganizationIDSelector *v1.Selector `json:"organizationIdSelector,omitempty" tf:"-"`

	// Reference to a Organization to populate organizationId.
	// +kubebuilder:validation:Optional
	OrganizationRef *v1.Reference `json:"organizationRef,omitempty" tf:"-"`
}

// PropositionSpec defines the desired state of Proposition
type PropositionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PropositionParameters `json:"forProvider"`
}

// PropositionStatus defines the observed state of Proposition.
type PropositionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PropositionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Proposition is the Schema for the Propositions API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,hsdpjet}
type Proposition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PropositionSpec   `json:"spec"`
	Status            PropositionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PropositionList contains a list of Propositions
type PropositionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Proposition `json:"items"`
}

// Repository type metadata.
var (
	Proposition_Kind             = "Proposition"
	Proposition_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Proposition_Kind}.String()
	Proposition_KindAPIVersion   = Proposition_Kind + "." + CRDGroupVersion.String()
	Proposition_GroupVersionKind = CRDGroupVersion.WithKind(Proposition_Kind)
)

func init() {
	SchemeBuilder.Register(&Proposition{}, &PropositionList{})
}
