package main

import (
	"strconv"
	"strings"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func boxType(bot *tg.BotAPI, update tg.Update, small string) {

	text := textBT(small)
	msg := tg.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, text)
	buttonClose := tg.NewInlineKeyboardButtonData("❌Закрыть", "close_1")
	rowCor := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("Короба", "time_"+small+"|"+"1"),tg.NewInlineKeyboardButtonData("Монопаллеты", "time_"+small+"|"+"2"))

	rowSup := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("Cуперсейфы", "time_"+small+"|"+"3"),tg.NewInlineKeyboardButtonData("QR-постаки", "time_"+small+"|"+"4"))

	rowHome := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("🏠Главная страница", "start_1"))
	rowBack := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("🡸Склады", "whOne_"+small),buttonClose)
	keyboard := tg.NewInlineKeyboardMarkup(rowCor, rowSup, rowHome, rowBack)
	msg.ReplyMarkup = &keyboard
	bot.Send(msg)
}
func textBT(small string) (result string) {

	result = `
	✪Склады:`
	if small == "" {
		result += ""
		return
	}
	small = strings.Trim(small, ":")
	arr := strings.Split(small, ":")
	for _, v := range arr {
		index, _ := strconv.Atoi(v)

		result += "\n- " + whArr[index]

	}
	result+="\n☆Выберите тип поставки"

	return
}
