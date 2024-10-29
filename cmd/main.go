package main

import (
	"log"
	"discord-chatgpt-bot/internal/bot"
)

func main() {
	if err := bot.Start(); err != nil {
		log.Fatalf("Error starting bot: %v", err)
	}

	select {}
}
