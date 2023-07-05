package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"

	awscdkapigw "github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha/v2"
	awsapigwintegrations "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha/v2"
	awscdklambdago "github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type InfraCdkStackProps struct {
	awscdk.StackProps
}

func NewInfraCdkStack(scope constructs.Construct, id string, props *InfraCdkStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	awslambda.NewLayerVersion(stack, jsii.String("skulla-assets"), &awslambda.LayerVersionProps{
		Code: awslambda.Code_FromAsset(jsii.String("../assets"), &awss3assets.AssetOptions{}),
	})

	skullaFunc := awscdklambdago.NewGoFunction(stack, jsii.String("SkullaFunc"), &awscdklambdago.GoFunctionProps{
		FunctionName: jsii.String("SkullaFunc"),
		Description:  jsii.String("an api-gw handler for the skulla-api"),
		Entry:        jsii.String("../backend-api/cmd/main.go"),
		Architecture: awslambda.Architecture_ARM_64(),
	})

	skullaApi := awscdkapigw.NewHttpApi(stack, jsii.String("SkullaApi"), nil)

	skullaApi.AddRoutes(&awscdkapigw.AddRoutesOptions{
		Path:        jsii.String("/{proxy+}"),
		Methods:     &[]awscdkapigw.HttpMethod{awscdkapigw.HttpMethod_ANY},
		Integration: awsapigwintegrations.NewHttpLambdaIntegration(jsii.String("HealthCheckIntegration"), skullaFunc, nil),
	})

	awscdk.NewCfnOutput(stack, jsii.String("SkullaApiURL"), &awscdk.CfnOutputProps{
		Value:       skullaApi.ApiEndpoint(),
		Description: jsii.String("the URL to the skulla-api"),
		ExportName:  jsii.String("SkullaApiUrl"),
	})

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewInfraCdkStack(app, "SkullaApiCdkStack", &InfraCdkStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}
