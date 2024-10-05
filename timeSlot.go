package main

import (
	"strconv"
	"strings"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func timeChose(bot *tg.BotAPI, update tg.Update, small string) {
	wh := strings.Split(small, "|")[0]

	buttonClose := tg.NewInlineKeyboardButtonData("Закрыть", "close_1")
	rowOne := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("Сегодня", "coef_"+small+"|"+"0"))
	rowTwo := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("Завтра", "coef_"+small+"|"+"1"))
	rowThree := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("Месяц", "coef_"+small+"|"+"0"+":"+"27"))
	rowFour := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("Выбрать промежуток", "timeOne_"+small))
	rowHome := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("🏠Главная страница", "start_1"))
	rowBack := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("🡸Тип поставки", "boxtype_"+wh), buttonClose)
	text := texterforTIMEone(small)
	keyboard := tg.NewInlineKeyboardMarkup(rowOne, rowTwo, rowThree, rowFour, rowHome, rowBack)

	msg := tg.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, text)
	msg.ReplyMarkup = &keyboard
	bot.Send(msg)
}

func texterforTIMEone(small string) (result string) {
	wh := strings.Split(small, "|")[0]
	bt := strings.Split(small, "|")[1]

	result = `
	✪Склады:`
	if wh == "" {
		result += ""
		return
	}
	wh = strings.Trim(wh, ":")
	arr := strings.Split(wh, ":")
	for _, v := range arr {
		index, _ := strconv.Atoi(v)

		result += "\n- " + whArr[index]
		
	}

	result += "\n✪Тип поставки: " + boxtypeMap[bt]
	result += "\n☆Выберите даты"
	return
}

func timeChoseOne(bot *tg.BotAPI, update tg.Update, small string) {
	buttonClose := tg.NewInlineKeyboardButtonData("Закрыть", "close_1")
	wh := strings.Split(small, "|")[0]
	bt := strings.Split(small, "|")[1]
	n:=[][]tg.InlineKeyboardButton{}
	
	arr:=date()
	for i,v:=range arr{
		if i==27{
			row:=tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData(v,"coef_"+small+"|"+strconv.Itoa(i)))
			n=append(n, row)
		} else{
		row:=tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData(v,"timeTwo_"+small+"|"+strconv.Itoa(i)))
		n=append(n, row)}
	}
	rowHome := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("🏠Главная страница", "start_1"))
	rowBack := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("🡸Назад", "time_"+wh+"|"+bt), buttonClose)
	n=append(n, rowHome,rowBack)
	keyboard:=tg.InlineKeyboardMarkup{
		InlineKeyboard: n,
	}
	text:=textForTimeChoseOne(small)

	msg:=tg.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, text)
	msg.ReplyMarkup=&keyboard
	bot.Send(msg)
}

func textForTimeChoseOne(small string)(result string) {
	wh := strings.Split(small, "|")[0]
	bt := strings.Split(small, "|")[1]

	result = `
	✪Cклады:`
	if wh == "" {
		result += ""
		return
	}
	wh = strings.Trim(wh, ":")
	arr := strings.Split(wh, ":")
	for _, v := range arr {
		index, _ := strconv.Atoi(v)

		result += "\n- " + whArr[index]

	}

	result += "\n✪Тип поставки: " + boxtypeMap[bt]
	result += "\n☆Выберите начальную дату"
	return
	
}
func textForTimeChoseTwo(small string)(result string) {
	wh := strings.Split(small, "|")[0]
	bt := strings.Split(small, "|")[1]
	index1:=strings.Split(small, "|")[2]
	ind1,_:=strconv.Atoi(index1)
	result = `
	✪Склады:`
	if wh == "" {
		result += ""
		return
	}
	wh = strings.Trim(wh, ":")
	arr := strings.Split(wh, ":")
	for _, v := range arr {
		index, _ := strconv.Atoi(v)

		result += "\n- " + whArr[index]

	}

	result += "\n✪Тип поставки: " + boxtypeMap[bt]
	result += "\n☆Даты: с " + date()[ind1] + " по ..."

	result += "\n☆Выберите конечную дату"
	return
	
}


func timeChoseTwo(bot *tg.BotAPI, update tg.Update, small string){
	buttonClose := tg.NewInlineKeyboardButtonData("Закрыть", "close_1")
	wh := strings.Split(small, "|")[0]
	bt := strings.Split(small, "|")[1]
	index1:=strings.Split(small, "|")[2]
	ind1,_:=strconv.Atoi(index1)
	n:=[][]tg.InlineKeyboardButton{}
	
	arr:=date()
	for i:=ind1;i<len(arr);i++{
		if i==ind1 {
			v:=arr[i]
			row:=tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData(v,"coef_"+wh+"|"+bt+"|"+index1))
			n=append(n, row)
		} else{
		v:=arr[i]
		row:=tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData(v,"coef_"+wh+"|"+bt+"|"+index1+":"+strconv.Itoa(i)))
		n=append(n, row)}
	}
	rowHome := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("🏠Главная страница", "start_1"))
	rowBack := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("🡸Начальная дата", "timeOne_"+wh+"|"+bt), buttonClose)
	n=append(n, rowHome,rowBack)
	keyboard:=tg.InlineKeyboardMarkup{
		InlineKeyboard: n,
	}
	text:=textForTimeChoseTwo(small)

	msg:=tg.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, text)
	msg.ReplyMarkup=&keyboard
	bot.Send(msg)
}
