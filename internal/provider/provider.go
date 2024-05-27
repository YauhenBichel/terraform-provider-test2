// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var (
	_ provider.Provider = &hashicupsProvider{}
)

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &hashicupsProvider{
			version: version,
		}
	}
}

// hashicupsProvider defines the provider implementation.
type hashicupsProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// to define the provider type name for inclusion in each data source
// and resource type name.
// for example, a resource type named "hashicups_order" would have a provider type name of "hashicups"
func (p *hashicupsProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "hashicups"
	resp.Version = p.version
}

// to define the schema for provider-level configurartion.
// we will update this method to accept a hashicups api token and endpoint
func (p *hashicupsProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{}
}

// to configure shared clients for data source and resource implementations
func (p *hashicupsProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

}

// to define the provider's resources
func (p *hashicupsProvider) Resources(ctx context.Context) []func() resource.Resource {
	return nil
}

// to define the provider's data sources
func (p *hashicupsProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return nil
}
