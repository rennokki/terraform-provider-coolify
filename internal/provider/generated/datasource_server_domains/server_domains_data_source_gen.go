// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package datasource_server_domains

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func ServerDomainsDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"server_domains": schema.SetNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"domains": schema.ListAttribute{
							ElementType: types.StringType,
							Computed:    true,
						},
						"ip": schema.StringAttribute{
							Computed: true,
						},
					},
					CustomType: ServerDomainsType{
						ObjectType: types.ObjectType{
							AttrTypes: ServerDomainsValue{}.AttributeTypes(ctx),
						},
					},
				},
				Computed: true,
			},
			"uuid": schema.StringAttribute{
				Required:            true,
				Description:         "Server's UUID",
				MarkdownDescription: "Server's UUID",
			},
		},
	}
}

type ServerDomainsModel struct {
	ServerDomains types.Set    `tfsdk:"server_domains"`
	Uuid          types.String `tfsdk:"uuid"`
}

var _ basetypes.ObjectTypable = ServerDomainsType{}

type ServerDomainsType struct {
	basetypes.ObjectType
}

func (t ServerDomainsType) Equal(o attr.Type) bool {
	other, ok := o.(ServerDomainsType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t ServerDomainsType) String() string {
	return "ServerDomainsType"
}

func (t ServerDomainsType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	domainsAttribute, ok := attributes["domains"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`domains is missing from object`)

		return nil, diags
	}

	domainsVal, ok := domainsAttribute.(basetypes.ListValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`domains expected to be basetypes.ListValue, was: %T`, domainsAttribute))
	}

	ipAttribute, ok := attributes["ip"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`ip is missing from object`)

		return nil, diags
	}

	ipVal, ok := ipAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`ip expected to be basetypes.StringValue, was: %T`, ipAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return ServerDomainsValue{
		Domains: domainsVal,
		Ip:      ipVal,
		state:   attr.ValueStateKnown,
	}, diags
}

func NewServerDomainsValueNull() ServerDomainsValue {
	return ServerDomainsValue{
		state: attr.ValueStateNull,
	}
}

func NewServerDomainsValueUnknown() ServerDomainsValue {
	return ServerDomainsValue{
		state: attr.ValueStateUnknown,
	}
}

func NewServerDomainsValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (ServerDomainsValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing ServerDomainsValue Attribute Value",
				"While creating a ServerDomainsValue value, a missing attribute value was detected. "+
					"A ServerDomainsValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("ServerDomainsValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid ServerDomainsValue Attribute Type",
				"While creating a ServerDomainsValue value, an invalid attribute value was detected. "+
					"A ServerDomainsValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("ServerDomainsValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("ServerDomainsValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra ServerDomainsValue Attribute Value",
				"While creating a ServerDomainsValue value, an extra attribute value was detected. "+
					"A ServerDomainsValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra ServerDomainsValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewServerDomainsValueUnknown(), diags
	}

	domainsAttribute, ok := attributes["domains"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`domains is missing from object`)

		return NewServerDomainsValueUnknown(), diags
	}

	domainsVal, ok := domainsAttribute.(basetypes.ListValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`domains expected to be basetypes.ListValue, was: %T`, domainsAttribute))
	}

	ipAttribute, ok := attributes["ip"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`ip is missing from object`)

		return NewServerDomainsValueUnknown(), diags
	}

	ipVal, ok := ipAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`ip expected to be basetypes.StringValue, was: %T`, ipAttribute))
	}

	if diags.HasError() {
		return NewServerDomainsValueUnknown(), diags
	}

	return ServerDomainsValue{
		Domains: domainsVal,
		Ip:      ipVal,
		state:   attr.ValueStateKnown,
	}, diags
}

func NewServerDomainsValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) ServerDomainsValue {
	object, diags := NewServerDomainsValue(attributeTypes, attributes)

	if diags.HasError() {
		// This could potentially be added to the diag package.
		diagsStrings := make([]string, 0, len(diags))

		for _, diagnostic := range diags {
			diagsStrings = append(diagsStrings, fmt.Sprintf(
				"%s | %s | %s",
				diagnostic.Severity(),
				diagnostic.Summary(),
				diagnostic.Detail()))
		}

		panic("NewServerDomainsValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t ServerDomainsType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewServerDomainsValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewServerDomainsValueUnknown(), nil
	}

	if in.IsNull() {
		return NewServerDomainsValueNull(), nil
	}

	attributes := map[string]attr.Value{}

	val := map[string]tftypes.Value{}

	err := in.As(&val)

	if err != nil {
		return nil, err
	}

	for k, v := range val {
		a, err := t.AttrTypes[k].ValueFromTerraform(ctx, v)

		if err != nil {
			return nil, err
		}

		attributes[k] = a
	}

	return NewServerDomainsValueMust(ServerDomainsValue{}.AttributeTypes(ctx), attributes), nil
}

func (t ServerDomainsType) ValueType(ctx context.Context) attr.Value {
	return ServerDomainsValue{}
}

var _ basetypes.ObjectValuable = ServerDomainsValue{}

type ServerDomainsValue struct {
	Domains basetypes.ListValue   `tfsdk:"domains"`
	Ip      basetypes.StringValue `tfsdk:"ip"`
	state   attr.ValueState
}

func (v ServerDomainsValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 2)

	var val tftypes.Value
	var err error

	attrTypes["domains"] = basetypes.ListType{
		ElemType: types.StringType,
	}.TerraformType(ctx)
	attrTypes["ip"] = basetypes.StringType{}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 2)

		val, err = v.Domains.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["domains"] = val

		val, err = v.Ip.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["ip"] = val

		if err := tftypes.ValidateValue(objectType, vals); err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		return tftypes.NewValue(objectType, vals), nil
	case attr.ValueStateNull:
		return tftypes.NewValue(objectType, nil), nil
	case attr.ValueStateUnknown:
		return tftypes.NewValue(objectType, tftypes.UnknownValue), nil
	default:
		panic(fmt.Sprintf("unhandled Object state in ToTerraformValue: %s", v.state))
	}
}

func (v ServerDomainsValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v ServerDomainsValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v ServerDomainsValue) String() string {
	return "ServerDomainsValue"
}

func (v ServerDomainsValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	var domainsVal basetypes.ListValue
	switch {
	case v.Domains.IsUnknown():
		domainsVal = types.ListUnknown(types.StringType)
	case v.Domains.IsNull():
		domainsVal = types.ListNull(types.StringType)
	default:
		var d diag.Diagnostics
		domainsVal, d = types.ListValue(types.StringType, v.Domains.Elements())
		diags.Append(d...)
	}

	if diags.HasError() {
		return types.ObjectUnknown(map[string]attr.Type{
			"domains": basetypes.ListType{
				ElemType: types.StringType,
			},
			"ip": basetypes.StringType{},
		}), diags
	}

	attributeTypes := map[string]attr.Type{
		"domains": basetypes.ListType{
			ElemType: types.StringType,
		},
		"ip": basetypes.StringType{},
	}

	if v.IsNull() {
		return types.ObjectNull(attributeTypes), diags
	}

	if v.IsUnknown() {
		return types.ObjectUnknown(attributeTypes), diags
	}

	objVal, diags := types.ObjectValue(
		attributeTypes,
		map[string]attr.Value{
			"domains": domainsVal,
			"ip":      v.Ip,
		})

	return objVal, diags
}

func (v ServerDomainsValue) Equal(o attr.Value) bool {
	other, ok := o.(ServerDomainsValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.Domains.Equal(other.Domains) {
		return false
	}

	if !v.Ip.Equal(other.Ip) {
		return false
	}

	return true
}

func (v ServerDomainsValue) Type(ctx context.Context) attr.Type {
	return ServerDomainsType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v ServerDomainsValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"domains": basetypes.ListType{
			ElemType: types.StringType,
		},
		"ip": basetypes.StringType{},
	}
}
