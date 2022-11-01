package iam

import "github.com/upbound/upjet/pkg/config"

// PasswordPolicyConfigure configures individual resources by adding custom ResourceConfigurators.
func PasswordPolicyConfigure(p *config.Provider) {
	p.AddResourceConfigurator("hsdp_iam_service", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.References["managing_organization"] = config.Reference{
			Type:         "Organization",
			RefFieldName: "OrganizationRef",
		}
	})
}
