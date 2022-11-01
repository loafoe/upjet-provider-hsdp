package iam

import (
	"github.com/upbound/upjet/pkg/config"
)

// EmailTemplateConfigure configures individual resources by adding custom ResourceConfigurators.
func EmailTemplateConfigure(p *config.Provider) {
	p.AddResourceConfigurator("hsdp_iam_email_template", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.References["managing_organization"] = config.Reference{
			Type:         "Organization",
			RefFieldName: "OrganizationRef",
		}
	})
}
