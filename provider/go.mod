module github.com/mlin-aviatrix/pulumi-avxmlin/provider

go 1.17

replace (
	github.com/hashicorp/go-getter v1.5.0 => github.com/hashicorp/go-getter v1.4.0
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20211230170131-3a7c83bfab87
)

require (
	github.com/AviatrixSystems/terraform-provider-aviatrix/v2 v2.22.0
	github.com/hashicorp/terraform-plugin-sdk v1.9.1 // indirect
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.21.0
	github.com/pulumi/pulumi/sdk/v3 v3.30.0
)
