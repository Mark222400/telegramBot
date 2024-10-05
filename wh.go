package main

import (
	"strconv"
	"strings"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func wh_1(bot *tg.BotAPI, update tg.Update, small string) {
	rowPages:=tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("#1","whOne_"+small),tg.NewInlineKeyboardButtonData("2","whTwo_"+small),tg.NewInlineKeyboardButtonData("3","whThree_"+small),tg.NewInlineKeyboardButtonData("4","whFour_"+small),tg.NewInlineKeyboardButtonData("5","whFive_"+small),tg.NewInlineKeyboardButtonData("6","whSix_"+small),tg.NewInlineKeyboardButtonData("7","whSeven_"+small))
	var rows [][]tg.InlineKeyboardButton
	for i := 1; i < 17; i += 2 {
		name, callback := arrHandlerFromWH(small, i)
		bl := tg.NewInlineKeyboardButtonData(name, "whOne_"+callback)
		name, callback = arrHandlerFromWH(small, i+1)
		br := tg.NewInlineKeyboardButtonData(name, "whOne_"+callback)
		row := tg.NewInlineKeyboardRow(bl, br)
		rows = append(rows, row)
	}
	small = strings.Trim(small, ":")
	buttonClose := tg.NewInlineKeyboardButtonData("❌Закрыть", "close_1")
	rowFunc := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData(">>", "whTwo_"+small))
	rowHome := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("🏠Главная страница", "start_1"), buttonClose)
	if small != "" {
		rowAdmire := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("🟢Далее➠", "boxtype_"+small))
		rows = append(rows,rowPages, rowFunc, rowAdmire, rowHome)
	} else {
		rows = append(rows, rowPages,rowFunc, rowHome)
	}
	text := textWH(small)
	if small == ""{
		text="☆Выберите склады(максимально 5)"
	}
	msg := tg.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, text)
	keyboard := tg.InlineKeyboardMarkup{
		InlineKeyboard: rows,
	}
	msg.ReplyMarkup = &keyboard
	bot.Send(msg)
}

func wh_2(bot *tg.BotAPI, update tg.Update, small string) {
	rowPages:=tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("1","whOne_"+small),tg.NewInlineKeyboardButtonData("#2","whTwo_"+small),tg.NewInlineKeyboardButtonData("3","whThree_"+small),tg.NewInlineKeyboardButtonData("4","whFour_"+small),tg.NewInlineKeyboardButtonData("5","whFive_"+small),tg.NewInlineKeyboardButtonData("6","whSix_"+small),tg.NewInlineKeyboardButtonData("7","whSeven_"+small))
	var rows [][]tg.InlineKeyboardButton
	for i := 17; i < 33; i += 2 {
		name, callback := arrHandlerFromWH(small, i)
		bl := tg.NewInlineKeyboardButtonData(name, "whTwo_"+callback)
		name, callback = arrHandlerFromWH(small, i+1)
		br := tg.NewInlineKeyboardButtonData(name, "whTwo_"+callback)
		row := tg.NewInlineKeyboardRow(bl, br)
		rows = append(rows, row)
	}
	small = strings.Trim(small, ":")
	buttonClose := tg.NewInlineKeyboardButtonData("❌Закрыть", "close_1")
	rowFunc := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("<<", "whOne_"+small), tg.NewInlineKeyboardButtonData(">>", "whThree_"+small))

	rowHome := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("🏠Главная страница", "start_1"), buttonClose)
	if small != "" {
		rowAdmire := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("🟢Далее➠", "boxtype_"+small))
		rows = append(rows,rowPages, rowFunc, rowAdmire, rowHome)
	} else {
		rows = append(rows,rowPages, rowFunc, rowHome)
	}
	text := textWH(small)
	if small == ""{
		text="☆Выберите склады(максимально 5)"
	}
	msg := tg.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, text)
	keyboard := tg.InlineKeyboardMarkup{
		InlineKeyboard: rows,
	}
	msg.ReplyMarkup = &keyboard
	bot.Send(msg)
}

func wh_3(bot *tg.BotAPI, update tg.Update, small string) {
	rowPages:=tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("1","whOne_"+small),tg.NewInlineKeyboardButtonData("2","whTwo_"+small),tg.NewInlineKeyboardButtonData("#3","whThree_"+small),tg.NewInlineKeyboardButtonData("4","whFour_"+small),tg.NewInlineKeyboardButtonData("5","whFive_"+small),tg.NewInlineKeyboardButtonData("6","whSix_"+small),tg.NewInlineKeyboardButtonData("7","whSeven_"+small))
	var rows [][]tg.InlineKeyboardButton
	for i := 33; i < 49; i += 2 {
		name, callback := arrHandlerFromWH(small, i)
		bl := tg.NewInlineKeyboardButtonData(name, "whThree_"+callback)
		name, callback = arrHandlerFromWH(small, i+1)
		br := tg.NewInlineKeyboardButtonData(name, "whThree_"+callback)
		row := tg.NewInlineKeyboardRow(bl, br)
		rows = append(rows, row)
	}
	small = strings.Trim(small, ":")
	buttonClose := tg.NewInlineKeyboardButtonData("❌Закрыть", "close_1")
	rowFunc := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("<<", "whTwo_"+small), tg.NewInlineKeyboardButtonData(">>", "whFour_"+small))

	rowHome := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("🏠Главная страница", "start_1"), buttonClose)
	if small != "" {
		rowAdmire := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("🟢Далее➠", "boxtype_"+small))
		rows = append(rows, rowPages,rowFunc, rowAdmire, rowHome)
	} else {
		rows = append(rows,rowPages, rowFunc, rowHome)
	}
	text := textWH(small)
	if small == ""{
		text="☆Выберите склады(максимально 5)"
	}
	msg := tg.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, text)
	keyboard := tg.InlineKeyboardMarkup{
		InlineKeyboard: rows,
	}
	msg.ReplyMarkup = &keyboard
	bot.Send(msg)
}

func wh_4(bot *tg.BotAPI, update tg.Update, small string) {
	rowPages:=tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("1","whOne_"+small),tg.NewInlineKeyboardButtonData("2","whTwo_"+small),tg.NewInlineKeyboardButtonData("3","whThree_"+small),tg.NewInlineKeyboardButtonData("#4","whFour_"+small),tg.NewInlineKeyboardButtonData("5","whFive_"+small),tg.NewInlineKeyboardButtonData("6","whSix_"+small),tg.NewInlineKeyboardButtonData("7","whSeven_"+small))
	var rows [][]tg.InlineKeyboardButton
	for i := 49; i < 65; i += 2 {
		name, callback := arrHandlerFromWH(small, i)
		bl := tg.NewInlineKeyboardButtonData(name, "whFour_"+callback)
		name, callback = arrHandlerFromWH(small, i+1)
		br := tg.NewInlineKeyboardButtonData(name, "whFour_"+callback)
		row := tg.NewInlineKeyboardRow(bl, br)
		rows = append(rows, row)
	}
	small = strings.Trim(small, ":")
	buttonClose := tg.NewInlineKeyboardButtonData("❌Закрыть", "close_1")
	rowFunc := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("<<", "whThree_"+small), tg.NewInlineKeyboardButtonData(">>", "whFive_"+small))
	rowHome := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("🏠Главная страница", "start_1"), buttonClose)
	if small != "" {
		rowAdmire := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("🟢Далее➠", "boxtype_"+small))
		rows = append(rows,rowPages, rowFunc, rowAdmire, rowHome)
	} else {
		rows = append(rows,rowPages, rowFunc, rowHome)
	}
	text := textWH(small)
	if small == ""{
		text="☆Выберите склады(максимально 5)"
	}
	msg := tg.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, text)
	keyboard := tg.InlineKeyboardMarkup{
		InlineKeyboard: rows,
	}
	msg.ReplyMarkup = &keyboard
	bot.Send(msg)
}

func wh_5(bot *tg.BotAPI, update tg.Update, small string) {
	rowPages:=tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("1","whOne_"+small),tg.NewInlineKeyboardButtonData("2","whTwo_"+small),tg.NewInlineKeyboardButtonData("3","whThree_"+small),tg.NewInlineKeyboardButtonData("4","whFour_"+small),tg.NewInlineKeyboardButtonData("#5","whFive_"+small),tg.NewInlineKeyboardButtonData("6","whSix_"+small),tg.NewInlineKeyboardButtonData("7","whSeven_"+small))
	var rows [][]tg.InlineKeyboardButton
	for i := 65; i < 81; i += 2 {
		name, callback := arrHandlerFromWH(small, i)
		bl := tg.NewInlineKeyboardButtonData(name, "whFive_"+callback)
		name, callback = arrHandlerFromWH(small, i+1)
		br := tg.NewInlineKeyboardButtonData(name, "whFive_"+callback)
		row := tg.NewInlineKeyboardRow(bl, br)
		rows = append(rows, row)
	}
	small = strings.Trim(small, ":")
	buttonClose := tg.NewInlineKeyboardButtonData("❌Закрыть", "close_1")
	rowFunc := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("<<", "whFour_"+small), tg.NewInlineKeyboardButtonData(">>", "whSix_"+small))
	rowHome := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("🏠Главная страница", "start_1"), buttonClose)
	if small != "" {
		rowAdmire := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("🟢Далее➠", "boxtype_"+small))
		rows = append(rows,rowPages, rowFunc, rowAdmire, rowHome)
	} else {
		rows = append(rows,rowPages, rowFunc, rowHome)
	}
	text := textWH(small)
	if small == ""{
		text="Выберите склады(максимально 5)"
	}
	msg := tg.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, text)
	keyboard := tg.InlineKeyboardMarkup{
		InlineKeyboard: rows,
	}
	msg.ReplyMarkup = &keyboard
	bot.Send(msg)
}

func wh_6(bot *tg.BotAPI, update tg.Update, small string) {
	rowPages:=tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("1","whOne_"+small),tg.NewInlineKeyboardButtonData("2","whTwo_"+small),tg.NewInlineKeyboardButtonData("3","whThree_"+small),tg.NewInlineKeyboardButtonData("4","whFour_"+small),tg.NewInlineKeyboardButtonData("5","whFive_"+small),tg.NewInlineKeyboardButtonData("#6","whSix_"+small),tg.NewInlineKeyboardButtonData("7","whSeven_"+small))
	var rows [][]tg.InlineKeyboardButton
	for i := 81; i < 97; i += 2 {
		name, callback := arrHandlerFromWH(small, i)
		bl := tg.NewInlineKeyboardButtonData(name, "whSix_"+callback)
		name, callback = arrHandlerFromWH(small, i+1)
		br := tg.NewInlineKeyboardButtonData(name, "whSix_"+callback)
		row := tg.NewInlineKeyboardRow(bl, br)
		rows = append(rows, row)
	}
	small = strings.Trim(small, ":")
	buttonClose := tg.NewInlineKeyboardButtonData("❌Закрыть", "close_1")
	rowFunc := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("<<", "whFive_"+small), tg.NewInlineKeyboardButtonData(">>", "whSeven_"+small))
	rowHome := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("🏠Главная страница", "start_1"), buttonClose)
	if small != "" {
		rowAdmire := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("🟢Далее➠", "boxtype_"+small))
		rows = append(rows,rowPages, rowFunc, rowAdmire, rowHome)
	} else {
		rows = append(rows,rowPages, rowFunc, rowHome)
	}
	text := textWH(small)
	if small == ""{
		text="☆Выберите склады(максимально 5)"
	}
	msg := tg.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, text)
	keyboard := tg.InlineKeyboardMarkup{
		InlineKeyboard: rows,
	}
	msg.ReplyMarkup = &keyboard
	bot.Send(msg)
}


func wh_7(bot *tg.BotAPI, update tg.Update, small string) {
	rowPages:=tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("1","whOne_"+small),tg.NewInlineKeyboardButtonData("2","whTwo_"+small),tg.NewInlineKeyboardButtonData("3","whThree_"+small),tg.NewInlineKeyboardButtonData("4","whFour_"+small),tg.NewInlineKeyboardButtonData("5","whFive_"+small),tg.NewInlineKeyboardButtonData("6","whSix_"+small),tg.NewInlineKeyboardButtonData("#7","whSeven_"+small))

	var rows [][]tg.InlineKeyboardButton
	for i := 97; i < 103; i += 2 {
		name, callback := arrHandlerFromWH(small, i)
		bl := tg.NewInlineKeyboardButtonData(name, "whSeven_"+callback)
		name, callback = arrHandlerFromWH(small, i+1)
		br := tg.NewInlineKeyboardButtonData(name, "whSeven_"+callback)
		row := tg.NewInlineKeyboardRow(bl, br)
		rows = append(rows, row)
	}
	name,callback:= arrHandlerFromWH(small, 103)
	br := tg.NewInlineKeyboardButtonData(name, "whSeven_"+callback)
	row := tg.NewInlineKeyboardRow(br)
	rows = append(rows, row)
	small = strings.Trim(small, ":")
	buttonClose := tg.NewInlineKeyboardButtonData("❌Закрыть", "close_1")
	rowFunc := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("<<", "whSix_"+small))
	rowHome := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("🏠Главная страница", "start_1"), buttonClose)
	if small != "" {
		rowAdmire := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("🟢Далее➠", "boxtype_"+small))
		rows = append(rows,rowPages, rowFunc, rowAdmire, rowHome)
	} else {
		rows = append(rows,rowPages, rowFunc, rowHome)
	}
	text := textWH(small)
	if small == ""{
		text="☆Выберите склады(максимально 5)"
	}
	msg := tg.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, text)
	keyboard := tg.InlineKeyboardMarkup{
		InlineKeyboard: rows,
	}
	msg.ReplyMarkup = &keyboard
	bot.Send(msg)
}

func arrHandlerFromWH(small string, i int) (name string, callback string) {
	small = strings.Trim(small, ":")
	arr := strings.Split(small, ":")
	if len(arr) >= 5 {
		contain := false
		for _, v := range arr {
			if v == strconv.Itoa(i) {
				contain = true
			}
		}
		if contain {
			name = "✅" + whArr[i]
			callback = strings.Trim(strings.ReplaceAll(":"+small+":", ":"+strconv.Itoa(i)+":", ":"), ":")
			return
		} else {
			name = whArr[i]
			callback = small
			return
		}
	} else {
		contain := false
		for _, v := range arr {
			if v == strconv.Itoa(i) {
				contain = true
			}
		}
		if contain {
			name = "✅" + whArr[i]
			callback = strings.Trim(strings.ReplaceAll(":"+small+":", ":"+strconv.Itoa(i)+":", ":"), ":")
			return
		}

		name = whArr[i]
		callback = small + ":" + strconv.Itoa(i)
		return
	}
}

func textWH(small string) (result string) {

	result = `☆Выбранные склады:`
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
	if len(arr) >= 5 {
		result += "\n\n!⚠️Выбрано максимальное количество складов(5)⚠️!\nНажмите повторно на выбранный склад, для исключения его из списка."
	}
	return
}
