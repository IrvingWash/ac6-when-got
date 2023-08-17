package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	telegramapi "github.com/IrvingWash/ac6-when-got/internal/telegram-api"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	body := request.Body

	fmt.Println(body)

	var update telegramapi.TelegramUpdate

	err := json.Unmarshal([]byte(body), update)

	if err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Failed to read request body",
		}, nil
	}

	telegramapi.SendMessage(&telegramapi.TelegramSendMessagePayload{
		ChatID: update.Message.Chat.ID,
		Text:   "Hello, World!",
	})

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
