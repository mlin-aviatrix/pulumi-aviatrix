// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func GetControllerIp(ctx *pulumi.Context) string {
	return config.Get(ctx, "aviatrix:controllerIp")
}
func GetPassword(ctx *pulumi.Context) string {
	return config.Get(ctx, "aviatrix:password")
}
func GetPathToCaCertificate(ctx *pulumi.Context) string {
	return config.Get(ctx, "aviatrix:pathToCaCertificate")
}
func GetSkipVersionValidation(ctx *pulumi.Context) bool {
	v, err := config.TryBool(ctx, "aviatrix:skipVersionValidation")
	if err == nil {
		return v
	}
	return true
}
func GetUsername(ctx *pulumi.Context) string {
	return config.Get(ctx, "aviatrix:username")
}
func GetVerifySslCertificate(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "aviatrix:verifySslCertificate")
}
