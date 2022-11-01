package iam

import "github.com/upbound/upjet/pkg/config"

// GroupConfigure configures individual resources by adding custom ResourceConfigurators.
func GroupConfigure(p *config.Provider) {
	p.AddResourceConfigurator("hsdp_iam_group", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.References["managing_organization"] = config.Reference{
			Type:         "Organization",
			RefFieldName: "OrganizationRef",
		}
		r.References["users"] = config.Reference{
			Type:         "User",
			RefFieldName: "UserRef",
		}
		r.References["services"] = config.Reference{
			Type:         "Service",
			RefFieldName: "ServiceRef",
		}
		r.References["roles"] = config.Reference{
			Type:         "Role",
			RefFieldName: "RoleRef",
		}
	})
}
