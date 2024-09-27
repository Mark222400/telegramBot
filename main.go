package main

import (
	"log"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tg.NewBotAPI("7130209675:AAESp2UBKfKjVr8YbgYFrvAK0u3t-ecbdGo")
	if err != nil {
		log.Panic("Ошибка взаимодействия с API бота: ", err)
	}
	g := tg.BotCommand{
		Command:     "/start",
		Description: "Запуск бота",
	}
	config := tg.NewSetMyCommands(g)
	bot.Request(config)
	u := tg.NewUpdate(0)
	u.Timeout = 60
	bot.Debug = true
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.CallbackQuery == nil && update.Message == nil {
			continue
		}
		if update.Message != nil {
			messageHandler(bot, update)
			continue
		}
		if update.CallbackQuery != nil {
			callbackHandler(bot, update)
			continue
		}
	}
}
