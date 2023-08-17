package main

import (
	"fmt"
	"log"

	telegramapi "github.com/IrvingWash/ac6-when-got/internal/telegram-api"
	"github.com/IrvingWash/ac6-when-got/internal/utils"
)

func main() {
	variable, err := utils.ExtractEnvironmentVariable("TEST")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(variable)

	telegramapi.GetMe()
}
