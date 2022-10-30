package iam

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func OrgConfigure(p *config.Provider) {
    p.AddResourceConfigurator("hsdp_iam_org", func(r *config.Resource) {
        r.ShortGroup = shortGroup
    })
}