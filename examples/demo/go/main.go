package main

import (
	"github.com/mlin-aviatrix/pulumi-aviatrix/sdk/go/aviatrix"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	accountName := pulumi.String("mlin-aviatrix")
	cloudType := pulumi.Int(1)
	region := pulumi.String("us-west-2")
	gwName := "mlin-gateway-test-2"
	gwSize := pulumi.String("t2.micro")

	pulumi.Providers()

	pulumi.Run(func(ctx *pulumi.Context) error {
		name := "snat"

		vpc, err := aviatrix.NewVpc(ctx, name, &aviatrix.VpcArgs{
			AccountName: accountName,
			CloudType: cloudType,
			Region: region,
			Name: pulumi.String("mlin-vpc"),
			Cidr: pulumi.String("10.1.0.0/16"),
		})

		if err != nil {
			return err
		}

		ctx.Export("vpc_id", vpc.VpcId)

		//gateway, err := aviatrix.NewGateway(ctx, gwName, &aviatrix.GatewayArgs{
		//	CloudType: vpc.CloudType,
		//	AccountName: vpc.AccountName,
		//	GwName: pulumi.String(gwName),
		//	VpcId: vpc.VpcId,
		//	VpcReg: vpc.Region.Elem(),
		//	GwSize: gwSize,
		//	Subnet: vpc.PublicSubnets.Index(pulumi.Int(0)).Cidr().Elem(),
		//	IdleTimeout: pulumi.Int(600),
		//	RenegotiationInterval: pulumi.Int(600),
		//	Tags: pulumi.StringMap{
		//		"k1": pulumi.String("v1"),
		//	},
		//})

		gateway, err := aviatrix.NewGateway(ctx, name, &aviatrix.GatewayArgs{
			CloudType: vpc.CloudType,
			AccountName: vpc.AccountName,
			GwName: pulumi.String(gwName),
			VpcId: vpc.VpcId,
			VpcReg: vpc.Region.Elem(),
			GwSize: gwSize,
			IdleTimeout: pulumi.Int(600),
			RenegotiationInterval: pulumi.Int(600),
			Subnet: vpc.PublicSubnets.Index(pulumi.Int(0)).Cidr().Elem(),
			Tags: pulumi.StringMap{
				"k1": pulumi.String("v1"),
			},
		})

		//snat, err := aviatrix.New
		policies := []aviatrix.GatewaySNATSnatPolicyInput{
			aviatrix.GatewaySNATSnatPolicyInput(aviatrix.GatewaySNATSnatPolicyArgs{
				DstCidr:         pulumi.String("10.2.0.0/24"),
				SnatIps:         pulumi.String("10.5.0.1"),
				SrcCidr:         pulumi.String("10.3.0.21/32"),
				Interface: pulumi.String("eth0"),
			}),
			aviatrix.GatewaySNATSnatPolicyInput(aviatrix.GatewaySNATSnatPolicyArgs{
				DstCidr:         pulumi.String("10.2.0.0/24"),
				SnatIps:         pulumi.String("10.5.0.1"),
				SrcCidr:         pulumi.String("10.3.0.21/32"),
				Interface: pulumi.String("eth0"),
			}),
			aviatrix.GatewaySNATSnatPolicyInput(aviatrix.GatewaySNATSnatPolicyArgs{
				DstCidr:         pulumi.String("10.2.0.0/24"),
				SnatIps:         pulumi.String("10.5.0.1"),
				SrcCidr:         pulumi.String("10.3.0.21/32"),
				Interface: pulumi.String("eth0"),
			}),
		}

		aviatrix.NewGatewaySNAT(ctx, name, &aviatrix.GatewaySNATArgs{
			GwName:       gateway.GwName,
			SnatMode:     pulumi.String("customized_snat"),
			SnatPolicies: aviatrix.GatewaySNATSnatPolicyArray(policies),
			SyncToHa:     pulumi.Bool(false),
		})

		ctx.Export("gw_name", gateway.GwName)
		return nil
	})
}
