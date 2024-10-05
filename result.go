package main

import (
	"strconv"
	"strings"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func result(bot *tg.BotAPI, update tg.Update, small string) {
	text := ""
	smallWH := strings.Split(small, "|")[0]
	smallBT := strings.Split(small, "|")[1]
	indexes:=strings.Split(small, "|")[2]
	coef:= strings.Split(small, "|")[3]
	if len(strings.Split(indexes, ":")) > 1 {
		index1 := strings.Split(indexes, ":")[0]
		index2 := strings.Split(indexes, ":")[1]
		text = textForResultWithTwo(smallWH, smallBT, index1, index2,coef)
	}
	if  len(strings.Split(indexes, ":")) == 1 {
		index1 := strings.Split(strings.Split(small, "|")[2], ":")[0]
		text = textForResultWithOne(smallWH, smallBT, index1,coef)
	}
	rowHome := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("🏠Главная страница", "start_1"))
	buttonClose := tg.NewInlineKeyboardButtonData("❌Закрыть", "close_1")
	buttonback := tg.NewInlineKeyboardButtonData("🡸Коэффициенты", "time_"+smallWH+"|"+smallBT+"|"+indexes)
	buttonAccept := tg.NewInlineKeyboardButtonData("🟢Подтвердить", "accept_"+small)
	rowFunc := tg.NewInlineKeyboardRow(buttonback, buttonClose)
	rowAccept:=tg.NewInlineKeyboardRow(buttonAccept)
	keyboard := tg.NewInlineKeyboardMarkup(rowHome,rowAccept, rowFunc)

	msg := tg.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, text)
	msg.ReplyMarkup = &keyboard
	bot.Send(msg)

}

func textForResultWithOne(wh string, bt string, index1 string,coef string) (text string) {
	ind1, _ := strconv.Atoi(index1)
	text += `✪Склады:`
	wh = strings.Trim(wh, ":")
	arr := strings.Split(wh, ":")
	for _, v := range arr {
		index, _ := strconv.Atoi(v)

		text += "\n- " + whArr[index]

	}
	switch bt {
	case "1":
		text += "\n✪Тип поставки: короба"
	case "2":
		text += "\n✪Тип поставки: монопаллеты"
	case "3":
		text += "\n✪Тип поставки: суперсейф"
	case "4":
		text += "\n✪Тип поставки: QR - поставки"
	}
	text +=
		"\n✪Дата: " + date()[ind1]
		text+="\n✪Коэффициент: до "+coef
		text+="\n☆Подтвердите заявку ..."
	return text
}

func textForResultWithTwo(wh string, bt string, index1 string, index2 string,coef string) (text string) {
	ind1, _ := strconv.Atoi(index1)
	ind2, _ := strconv.Atoi(index2)
	text += `✪Склады:`
	wh = strings.Trim(wh, ":")
	arr := strings.Split(wh, ":")
	for _, v := range arr {
		index, _ := strconv.Atoi(v)

		text += "\n- " + whArr[index]

	}
	switch bt {
	case "1":
		text += "\n✪Тип поставки: короба"
	case "2":
		text += "\n✪Тип поставки: монопаллеты"
	case "3":
		text += "\n✪Тип поставки: суперсейф"
	case "4":
		text += "\n✪Тип поставки: QR - поставки"
	}

	text += "\n✪Даты: c " + date()[ind1] + ` по ` + date()[ind2]

	text+="\n✪Коэффициент: до "+coef
	text+="\n☆Подтвердите заявку ..."
	return text
}
