package main

import (
	"context"
	"github.com/achaquisse/skulla-api/handler"
	"github.com/achaquisse/skulla-api/helper"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gofiber/fiber/v2"
	"log"

	"github.com/aws/aws-lambda-go/events"
	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
)

var fiberLambda *fiberadapter.FiberLambda

// init the Fiber Server
func init() {
	log.Printf("Fiber cold start")
	var app *fiber.App
	app = fiber.New()

	app.Get("/", handler.HealthCheck)
	app.Get("/users", handler.ReturnUsers)

	if helper.IsLambda() {
		fiberLambda = fiberadapter.New(app)
	} else {
		err := app.Listen(":3000")
		if err != nil {
			panic(err)
		}
	}
}

// Handler will deal with Fiber working with Lambda
func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return fiberLambda.ProxyWithContextV2(ctx, req)
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	if helper.IsLambda() {
		lambda.Start(Handler)
	}
}
