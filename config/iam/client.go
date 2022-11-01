package iam

import (
	"github.com/upbound/upjet/pkg/config"
)

// ClientConfigure configures individual resources by adding custom ResourceConfigurators.
func ClientConfigure(p *config.Provider) {
	p.AddResourceConfigurator("hsdp_iam_client", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.References["application_id"] = config.Reference{
			Type:         "Application",
			RefFieldName: "ApplicationRef",
		}
	})
}
