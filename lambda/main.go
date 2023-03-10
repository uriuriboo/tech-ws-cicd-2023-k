package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, event events.LambdaFunctionURLRequest) (string, error) {
	return fmt.Sprintf("Hello %s!", event.QueryStringParameters["name"]), nil
}

func main() {
	lambda.Start(HandleRequest)
}
