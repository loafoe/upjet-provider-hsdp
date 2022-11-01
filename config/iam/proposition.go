package iam

import (
	"github.com/upbound/upjet/pkg/config"
)

// PropositionConfigure configures individual resources by adding custom ResourceConfigurators.
func PropositionConfigure(p *config.Provider) {
	p.AddResourceConfigurator("hsdp_iam_proposition", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.References["organization_id"] = config.Reference{
			Type:         "Organization",
			RefFieldName: "OrganizationRef",
		}
	})
}
