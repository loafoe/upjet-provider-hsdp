package iam

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func GroupConfigure(p *config.Provider) {
    p.AddResourceConfigurator("hsdp_iam_group", func(r *config.Resource) {
        r.ShortGroup = shortGroup
    })
}
