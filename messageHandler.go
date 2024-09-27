package main

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func messageHandler(bot *tg.BotAPI, update tg.Update) {
	message := update.Message.Text
	switch message {
	case "/start":
		startPage(bot, update, 0)
	}
}
