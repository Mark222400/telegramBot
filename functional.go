package main

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func close(bot *tg.BotAPI, update tg.Update) {
	msg := tg.NewDeleteMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID)
	bot.Send(msg)
}
