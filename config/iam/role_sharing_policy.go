package iam

import (
	"github.com/upbound/upjet/pkg/config"
)

// RoleSharingPolicyConfigure configures individual resources by adding custom ResourceConfigurators.
func RoleSharingPolicyConfigure(p *config.Provider) {
	p.AddResourceConfigurator("hsdp_iam_role_sharing_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.References["role_id"] = config.Reference{
			Type:         "Role",
			RefFieldName: "RoleRef",
		}
		r.References["target_organization_id"] = config.Reference{
			Type:         "Organization",
			RefFieldName: "OrganizationRef",
		}
	})
}
