package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/greymd/ojichat/generator"
)

type Payload struct {
	Name string `json:"name"`
}

type Response struct {
	Message string `json:"message"`
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var payload Payload
	err := json.Unmarshal([]byte(request.Body), &payload)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: 500,
		}, err
	}

	config := generator.Config{
		TargetName:       payload.Name,
		EmojiNum:         4,
		PunctuationLevel: 0,
	}

	ojiResult, err := generator.Start(config)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: 500,
		}, err
	}

	res := Response{
		Message: ojiResult,
	}

	jsonResult, _ := json.Marshal(res)

	return events.APIGatewayProxyResponse{
		Body:       string(jsonResult),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
