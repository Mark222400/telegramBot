package main

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func startPage(bot *tg.BotAPI, update tg.Update, editStatus int) {
	buttonWH := tg.NewInlineKeyboardButtonData("Создать заявку", "whOne_")
	rowWH := tg.NewInlineKeyboardRow(buttonWH)
	buttonClose := tg.NewInlineKeyboardButtonData("Закрыть", "close_1")
	rowFunc := tg.NewInlineKeyboardRow(buttonClose)
	keyboard := tg.NewInlineKeyboardMarkup(rowWH, rowFunc)
	text := `Стартовая страница`
	if editStatus == 1 {
		msg := tg.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, text)
		msg.ReplyMarkup = &keyboard
		bot.Send(msg)
	} else {
		msg := tg.NewMessage(update.Message.Chat.ID, text)
		msg.ReplyMarkup = &keyboard
		bot.Send(msg)
	}

}
