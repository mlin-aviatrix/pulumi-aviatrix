// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aviatrix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GatewaySNAT struct {
	pulumi.CustomResourceState

	// Name of the gateway.
	GwName pulumi.StringOutput `pulumi:"gwName"`
	// Nat mode. Currently only supports 'customized_snat'.
	SnatMode pulumi.StringPtrOutput `pulumi:"snatMode"`
	// Policy rules applied for 'snat_mode'' of 'customized_snat'.'
	SnatPolicies GatewaySNATSnatPolicyArrayOutput `pulumi:"snatPolicies"`
	// Whether to sync the policies to the HA gateway.
	SyncToHa pulumi.BoolPtrOutput `pulumi:"syncToHa"`
}

// NewGatewaySNAT registers a new resource with the given unique name, arguments, and options.
func NewGatewaySNAT(ctx *pulumi.Context,
	name string, args *GatewaySNATArgs, opts ...pulumi.ResourceOption) (*GatewaySNAT, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GwName == nil {
		return nil, errors.New("invalid value for required argument 'GwName'")
	}
	var resource GatewaySNAT
	err := ctx.RegisterResource("aviatrix:index/gatewaySNAT:GatewaySNAT", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGatewaySNAT gets an existing GatewaySNAT resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGatewaySNAT(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewaySNATState, opts ...pulumi.ResourceOption) (*GatewaySNAT, error) {
	var resource GatewaySNAT
	err := ctx.ReadResource("aviatrix:index/gatewaySNAT:GatewaySNAT", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GatewaySNAT resources.
type gatewaySNATState struct {
	// Name of the gateway.
	GwName *string `pulumi:"gwName"`
	// Nat mode. Currently only supports 'customized_snat'.
	SnatMode *string `pulumi:"snatMode"`
	// Policy rules applied for 'snat_mode'' of 'customized_snat'.'
	SnatPolicies []GatewaySNATSnatPolicy `pulumi:"snatPolicies"`
	// Whether to sync the policies to the HA gateway.
	SyncToHa *bool `pulumi:"syncToHa"`
}

type GatewaySNATState struct {
	// Name of the gateway.
	GwName pulumi.StringPtrInput
	// Nat mode. Currently only supports 'customized_snat'.
	SnatMode pulumi.StringPtrInput
	// Policy rules applied for 'snat_mode'' of 'customized_snat'.'
	SnatPolicies GatewaySNATSnatPolicyArrayInput
	// Whether to sync the policies to the HA gateway.
	SyncToHa pulumi.BoolPtrInput
}

func (GatewaySNATState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewaySNATState)(nil)).Elem()
}

type gatewaySNATArgs struct {
	// Name of the gateway.
	GwName string `pulumi:"gwName"`
	// Nat mode. Currently only supports 'customized_snat'.
	SnatMode *string `pulumi:"snatMode"`
	// Policy rules applied for 'snat_mode'' of 'customized_snat'.'
	SnatPolicies []GatewaySNATSnatPolicy `pulumi:"snatPolicies"`
	// Whether to sync the policies to the HA gateway.
	SyncToHa *bool `pulumi:"syncToHa"`
}

// The set of arguments for constructing a GatewaySNAT resource.
type GatewaySNATArgs struct {
	// Name of the gateway.
	GwName pulumi.StringInput
	// Nat mode. Currently only supports 'customized_snat'.
	SnatMode pulumi.StringPtrInput
	// Policy rules applied for 'snat_mode'' of 'customized_snat'.'
	SnatPolicies GatewaySNATSnatPolicyArrayInput
	// Whether to sync the policies to the HA gateway.
	SyncToHa pulumi.BoolPtrInput
}

func (GatewaySNATArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewaySNATArgs)(nil)).Elem()
}

type GatewaySNATInput interface {
	pulumi.Input

	ToGatewaySNATOutput() GatewaySNATOutput
	ToGatewaySNATOutputWithContext(ctx context.Context) GatewaySNATOutput
}

func (*GatewaySNAT) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewaySNAT)(nil)).Elem()
}

func (i *GatewaySNAT) ToGatewaySNATOutput() GatewaySNATOutput {
	return i.ToGatewaySNATOutputWithContext(context.Background())
}

func (i *GatewaySNAT) ToGatewaySNATOutputWithContext(ctx context.Context) GatewaySNATOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewaySNATOutput)
}

// GatewaySNATArrayInput is an input type that accepts GatewaySNATArray and GatewaySNATArrayOutput values.
// You can construct a concrete instance of `GatewaySNATArrayInput` via:
//
//          GatewaySNATArray{ GatewaySNATArgs{...} }
type GatewaySNATArrayInput interface {
	pulumi.Input

	ToGatewaySNATArrayOutput() GatewaySNATArrayOutput
	ToGatewaySNATArrayOutputWithContext(context.Context) GatewaySNATArrayOutput
}

type GatewaySNATArray []GatewaySNATInput

func (GatewaySNATArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewaySNAT)(nil)).Elem()
}

func (i GatewaySNATArray) ToGatewaySNATArrayOutput() GatewaySNATArrayOutput {
	return i.ToGatewaySNATArrayOutputWithContext(context.Background())
}

func (i GatewaySNATArray) ToGatewaySNATArrayOutputWithContext(ctx context.Context) GatewaySNATArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewaySNATArrayOutput)
}

// GatewaySNATMapInput is an input type that accepts GatewaySNATMap and GatewaySNATMapOutput values.
// You can construct a concrete instance of `GatewaySNATMapInput` via:
//
//          GatewaySNATMap{ "key": GatewaySNATArgs{...} }
type GatewaySNATMapInput interface {
	pulumi.Input

	ToGatewaySNATMapOutput() GatewaySNATMapOutput
	ToGatewaySNATMapOutputWithContext(context.Context) GatewaySNATMapOutput
}

type GatewaySNATMap map[string]GatewaySNATInput

func (GatewaySNATMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewaySNAT)(nil)).Elem()
}

func (i GatewaySNATMap) ToGatewaySNATMapOutput() GatewaySNATMapOutput {
	return i.ToGatewaySNATMapOutputWithContext(context.Background())
}

func (i GatewaySNATMap) ToGatewaySNATMapOutputWithContext(ctx context.Context) GatewaySNATMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewaySNATMapOutput)
}

type GatewaySNATOutput struct{ *pulumi.OutputState }

func (GatewaySNATOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewaySNAT)(nil)).Elem()
}

func (o GatewaySNATOutput) ToGatewaySNATOutput() GatewaySNATOutput {
	return o
}

func (o GatewaySNATOutput) ToGatewaySNATOutputWithContext(ctx context.Context) GatewaySNATOutput {
	return o
}

type GatewaySNATArrayOutput struct{ *pulumi.OutputState }

func (GatewaySNATArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewaySNAT)(nil)).Elem()
}

func (o GatewaySNATArrayOutput) ToGatewaySNATArrayOutput() GatewaySNATArrayOutput {
	return o
}

func (o GatewaySNATArrayOutput) ToGatewaySNATArrayOutputWithContext(ctx context.Context) GatewaySNATArrayOutput {
	return o
}

func (o GatewaySNATArrayOutput) Index(i pulumi.IntInput) GatewaySNATOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GatewaySNAT {
		return vs[0].([]*GatewaySNAT)[vs[1].(int)]
	}).(GatewaySNATOutput)
}

type GatewaySNATMapOutput struct{ *pulumi.OutputState }

func (GatewaySNATMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewaySNAT)(nil)).Elem()
}

func (o GatewaySNATMapOutput) ToGatewaySNATMapOutput() GatewaySNATMapOutput {
	return o
}

func (o GatewaySNATMapOutput) ToGatewaySNATMapOutputWithContext(ctx context.Context) GatewaySNATMapOutput {
	return o
}

func (o GatewaySNATMapOutput) MapIndex(k pulumi.StringInput) GatewaySNATOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GatewaySNAT {
		return vs[0].(map[string]*GatewaySNAT)[vs[1].(string)]
	}).(GatewaySNATOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GatewaySNATInput)(nil)).Elem(), &GatewaySNAT{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewaySNATArrayInput)(nil)).Elem(), GatewaySNATArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewaySNATMapInput)(nil)).Elem(), GatewaySNATMap{})
	pulumi.RegisterOutputType(GatewaySNATOutput{})
	pulumi.RegisterOutputType(GatewaySNATArrayOutput{})
	pulumi.RegisterOutputType(GatewaySNATMapOutput{})
}
