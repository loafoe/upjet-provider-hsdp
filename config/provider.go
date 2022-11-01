/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/upbound/upjet/pkg/config"

	"github.com/philips-software/provider-hsdp/config/iam"
)

const (
	resourcePrefix = "hsdp"
	modulePath     = "github.com/philips-software/provider-hsdp"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		iam.GroupConfigure,
		iam.OrgConfigure,
		iam.ApplicationConfigure,
		iam.EmailTemplateConfigure,
		iam.ClientConfigure,
		iam.PropositionConfigure,
		iam.ServiceConfigure,
		iam.RoleConfigure,
		iam.PasswordPolicyConfigure,
		iam.RoleSharingPolicyConfigure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
