package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	ac6countdown "github.com/IrvingWash/ac6-when-got/internal/ac6-count-down"
	telegramapi "github.com/IrvingWash/ac6-when-got/internal/telegram-api"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var update telegramapi.TelegramUpdate

	err := json.Unmarshal([]byte(request.Body), &update)

	if err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Failed to read request body",
		}, nil
	}

	remainingTime := ac6countdown.GetRemainingTime()

	telegramapi.SendMessage(&telegramapi.TelegramSendMessagePayload{
		ChatID: update.Message.Chat.ID,
		Text:   remainingTime,
	})

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
