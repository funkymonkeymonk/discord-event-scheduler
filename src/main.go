// main.go
package main

// TODO: Get the iteration zero lambda deployed and running on a cron

import (
	"github.com/aws/aws-lambda-go/lambda"
)

func handleRequest() (string, error) {
	return "Hello λ!", nil
}

func main() {
	lambda.Start(handleRequest)
}
