package main

import (
	"strconv"
	"strings"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)


func coefficient(bot *tg.BotAPI, update tg.Update, small string){
	text := ""
	smallWH := strings.Split(small, "|")[0]
	smallBT := strings.Split(small, "|")[1]
	bt:=smallBT
	indexes:=strings.Split(small, "|")[2]
	rowZero := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("🎁 Бесплатно", "result_"+smallWH+"|"+bt+"|"+indexes+"|"+"0"))
	rowOne := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("до X1", "result_"+smallWH+"|"+bt+"|"+indexes+"|"+"1"))
	rowTwo := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("до X2", "result_"+smallWH+"|"+bt+"|"+indexes+"|"+"2"))
	rowThree := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("до X3", "result_"+smallWH+"|"+bt+"|"+indexes+"|"+"3"))
	rowFour := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("до X4", "result_"+smallWH+"|"+bt+"|"+indexes+"|"+"4"))
	rowFive := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("до X5", "result_"+smallWH+"|"+bt+"|"+indexes+"|"+"5"))
	rowSix := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("до X7", "result_"+smallWH+"|"+bt+"|"+indexes+"|"+"7"))
	if len(strings.Split(indexes, ":")) > 1 {
		index1 := strings.Split(indexes, ":")[0]
		index2 := strings.Split(indexes, ":")[1]
		text = textForCoefWithTwo(smallWH, smallBT, index1, index2)
	}
	if  len(strings.Split(indexes, ":")) == 1 {
		index1 := strings.Split(strings.Split(small, "|")[2], ":")[0]
		text = textForCoefOne(smallWH, smallBT, index1)
	}
	rowHome := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("🏠Главная страница", "start_1"))
	buttonClose := tg.NewInlineKeyboardButtonData("❌Закрыть", "close_1")
	buttonback := tg.NewInlineKeyboardButtonData("🡸Даты", "time_"+smallWH+"|"+smallBT)

	rowFunc := tg.NewInlineKeyboardRow(buttonback, buttonClose)
	keyboard := tg.NewInlineKeyboardMarkup(rowZero, rowOne, rowTwo, rowThree, rowFour, rowFive, rowSix,rowHome, rowFunc)

	msg := tg.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, text)
	msg.ReplyMarkup = &keyboard
	bot.Send(msg)

}

func textForCoefOne(wh string, bt string, index1 string) (text string) {
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
		text+="\n☆Выберите масимальный коэффициент"
	return text
}

func textForCoefWithTwo(wh string, bt string, index1 string, index2 string) (text string) {
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
	text+="\n☆Выберите максимальный коэффициент"
	return text
}
