package main

import (
	"strconv"
	"strings"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func boxType (bot *tg.BotAPI, update tg.Update, small string) {
	
	text:=textBT(small)
	msg:=tg.NewEditMessageText(update.CallbackQuery.Message.Chat.ID,update.CallbackQuery.Message.MessageID,text)
	rowCor:=tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("Короба","timeOne_"+"|"+small+"|"+":1"))
	rowMon:=tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("Монопаллеты","timeOne_"+"|"+small+"|"+":2"))
	rowSup:=tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("Сейфы","timeOne_"+"|"+small+"|"+":3"))
	rowQR:=tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("QR-постаки","timeOne_"+"|"+small+"|"+":4"))
	keyboard:=tg.NewInlineKeyboardMarkup(rowCor,rowMon,rowSup,rowQR)
	msg.ReplyMarkup=&keyboard
	bot.Send(msg)
}
func textBT(small string) (result string) {

	result = `Выберите тип поставки
	Склады:`
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

	return
}