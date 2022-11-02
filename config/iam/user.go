package iam

import (
	"github.com/philips-software/provider-hsdp/config/common"
	"github.com/upbound/upjet/pkg/config"
)

// UserConfigure configures individual resources by adding custom ResourceConfigurators.
func UserConfigure(p *config.Provider) {
	p.AddResourceConfigurator("hsdp_iam_user", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.References["organization_id"] = config.Reference{
			Type:         "Organization",
			Extractor:    common.ExtractResourceIDFuncPath,
			RefFieldName: "OrganizationRef",
		}
	})
}
