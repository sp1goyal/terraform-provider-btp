package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/SAP/terraform-provider-btp/internal/btpcli/types/cis"
)

const (
	DirectoryFeatureDefault        = "DEFAULT"
	DirectoryFeatureAuthorizations = "ENTITLEMENTS"
	DirectoryFeatureEntitlements   = "AUTHORIZATIONS"
)

type directoryType struct {
	ID           types.String `tfsdk:"id"`
	CreatedBy    types.String `tfsdk:"created_by"`
	CreatedDate  types.String `tfsdk:"created_date"`
	Description  types.String `tfsdk:"description"`
	Features     types.Set    `tfsdk:"features"`
	Labels       types.Map    `tfsdk:"labels"`
	LastModified types.String `tfsdk:"last_modified"`
	Name         types.String `tfsdk:"name"`
	ParentID     types.String `tfsdk:"parent_id"`
	State        types.String `tfsdk:"state"`
	Subdomain    types.String `tfsdk:"subdomain"`
}

func directoryValueFrom(ctx context.Context, value cis.DirectoryResponseObject) (directoryType, diag.Diagnostics) {
	directory := directoryType{
		ID:           types.StringValue(value.Guid),
		CreatedBy:    types.StringValue(value.CreatedBy),
		CreatedDate:  timeToValue(value.CreatedDate.Time()),
		Description:  types.StringValue(value.Description),
		LastModified: timeToValue(value.ModifiedDate.Time()),
		Name:         types.StringValue(value.DisplayName),
		ParentID:     types.StringValue(value.ParentGUID),
		State:        types.StringValue(value.EntityState),
		Subdomain:    types.StringValue(value.Subdomain),
	}

	var summary, diags diag.Diagnostics

	directory.Features, diags = types.SetValueFrom(ctx, types.StringType, value.DirectoryFeatures)
	summary.Append(diags...)

	directory.Labels, diags = types.MapValueFrom(ctx, types.SetType{ElemType: types.StringType}, value.Labels)
	summary.Append(diags...)

	return directory, summary
}

func getAllDirectories(ctx context.Context, resp *datasource.ReadResponse, dirResponses []cis.DirectoryResponseObject) []directoryType {
	dirs := []directoryType{}
	return recursivelyMapDirectories(ctx, resp, dirs, dirResponses)
}

func recursivelyMapDirectories(ctx context.Context, resp *datasource.ReadResponse, dirs []directoryType, dirResponses []cis.DirectoryResponseObject) []directoryType {
	for _, dirRes := range dirResponses {
		dir, diags := directoryValueFrom(ctx, dirRes)
		resp.Diagnostics.Append(diags...)
		dirs = append(dirs, dir)
		if len(dirRes.Children) > 0 {
			dirs = recursivelyMapDirectories(ctx, resp, dirs, dirRes.Children)
		}
	}
	return dirs
}
