package main

import (
	"github.com/mlin-aviatrix/pulumi-aviatrix/sdk/go/aviatrix"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	accountName := pulumi.String("mlin-aviatrix")
	cloudType := pulumi.Int(1)
	region := pulumi.String("us-west-2")
	gwName := "mlin-gateway-test"
	gwSize := pulumi.String("t2.nano")

	pulumi.Providers()

	pulumi.Run(func(ctx *pulumi.Context) error {
		vpc, err := aviatrix.NewVpc(ctx, "mlin-vpc", &aviatrix.VpcArgs{
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

		gateway, err := aviatrix.NewGateway(ctx, gwName, &aviatrix.GatewayArgs{
			CloudType: vpc.CloudType,
			AccountName: vpc.AccountName,
			GwName: pulumi.String(gwName),
			VpcId: vpc.VpcId,
			VpcReg: vpc.Region.Elem(),
			GwSize: gwSize,
			Subnet: vpc.PublicSubnets.Index(pulumi.Int(0)).Cidr().Elem(),
			IdleTimeout: pulumi.Int(600),
			RenegotiationInterval: pulumi.Int(600),
			Tags: pulumi.StringMap{
				"k1": pulumi.String("v1"),
			},
		})

		ctx.Export("gw_name", gateway.GwName)
		return nil
	})
}
