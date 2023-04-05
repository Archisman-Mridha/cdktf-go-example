package main

import (
	"log"
	"os"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v12/instance"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v12/provider"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/joho/godotenv"
)

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	if err := godotenv.Load("aws.credentials.env"); err != nil {
		log.Fatalf("error loading aws.credentials.env : %v", err)}

	// The code that defines your stack goes here

	provider.NewAwsProvider(stack, jsii.String("aws"), &provider.AwsProviderConfig{

		AccessKey: jsii.String(os.Getenv("AWS_ACCESS_KEY_ID")),
		SecretKey: jsii.String(os.Getenv("AWS_SECRET_ACCESS_KEY")),

		Region: cdktf.NewTerraformVariable(stack, jsii.String("region"), &cdktf.TerraformVariableConfig{
			Type: jsii.String("string"),
			Description: jsii.String("AWS region where infrastructure will be deployed"),

			Default: "us-east-1",
		}).StringValue( ),

	})

	instance.NewInstance(stack, jsii.String("vm"), &instance.InstanceConfig{

		InstanceType: jsii.String("t2.micro"),
		Ami: jsii.String("ami-00c39f71452c08778"),

	})

	return stack
}

func main( ) {
	terraformApp := cdktf.NewApp(nil)

	NewMyStack(terraformApp, "cdktf")

	terraformApp.Synth( )
}
