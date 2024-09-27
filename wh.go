package main

import (
	"strconv"
	"strings"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func wh_1(bot *tg.BotAPI, update tg.Update, small string) {
	var rows [][]tg.InlineKeyboardButton
	for i := 1; i < 21; i += 2 {
		name, callback := arrHandlerFromWH(small, i)
		bl := tg.NewInlineKeyboardButtonData(name, "whOne_"+callback)
		name, callback = arrHandlerFromWH(small, i+1)
		br := tg.NewInlineKeyboardButtonData(name, "whOne_"+callback)
		row := tg.NewInlineKeyboardRow(bl, br)
		rows = append(rows, row)
	}
	small = strings.Trim(small, ":")
	rowFunc := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("Далее", "whTwo_"+small))
	rowHome := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("🏠Главная страница", "start_1"))
	rowAdmire := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("Подтвердить", "boxtype_"+small))
	rows = append(rows, rowFunc, rowAdmire, rowHome)
	text := textWH(small)
	msg := tg.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, text)
	keyboard := tg.InlineKeyboardMarkup{
		InlineKeyboard: rows,
	}
	msg.ReplyMarkup = &keyboard
	bot.Send(msg)
}

func wh_2(bot *tg.BotAPI, update tg.Update, small string) {
	var rows [][]tg.InlineKeyboardButton
	for i := 20; i < 41; i += 2 {
		name, callback := arrHandlerFromWH(small, i)
		bl := tg.NewInlineKeyboardButtonData(name, "whTwo_"+callback)
		name, callback = arrHandlerFromWH(small, i+1)
		br := tg.NewInlineKeyboardButtonData(name, "whTwo_"+callback)
		row := tg.NewInlineKeyboardRow(bl, br)
		rows = append(rows, row)
	}
	small = strings.Trim(small, ":")
	rowFunc := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("Назад", "whOne_"+small), tg.NewInlineKeyboardButtonData("Далее", "whThree_"+small))
	rowHome := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("🏠Главная страница", "start_1"))
	rows = append(rows, rowFunc, rowHome)
	text := textWH(small)
	msg := tg.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, text)
	keyboard := tg.InlineKeyboardMarkup{
		InlineKeyboard: rows,
	}
	msg.ReplyMarkup = &keyboard
	bot.Send(msg)
}

func wh_3(bot *tg.BotAPI, update tg.Update, small string) {
	var rows [][]tg.InlineKeyboardButton
	for i := 41; i < 61; i += 2 {
		name, callback := arrHandlerFromWH(small, i)
		bl := tg.NewInlineKeyboardButtonData(name, "whThree_"+callback)
		name, callback = arrHandlerFromWH(small, i+1)
		br := tg.NewInlineKeyboardButtonData(name, "whThree_"+callback)
		row := tg.NewInlineKeyboardRow(bl, br)
		rows = append(rows, row)
	}
	small = strings.Trim(small, ":")
	rowFunc := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("Назад", "whTwo_"+small), tg.NewInlineKeyboardButtonData("Далее", "whFour_"+small))
	rowHome := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("🏠Главная страница", "start_1"))
	rows = append(rows, rowFunc, rowHome)
	text := textWH(small)
	msg := tg.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, text)
	keyboard := tg.InlineKeyboardMarkup{
		InlineKeyboard: rows,
	}
	msg.ReplyMarkup = &keyboard
	bot.Send(msg)
}

func wh_4(bot *tg.BotAPI, update tg.Update, small string) {
	var rows [][]tg.InlineKeyboardButton
	for i := 61; i < 81; i += 2 {
		name, callback := arrHandlerFromWH(small, i)
		bl := tg.NewInlineKeyboardButtonData(name, "whFour_"+callback)
		name, callback = arrHandlerFromWH(small, i+1)
		br := tg.NewInlineKeyboardButtonData(name, "whFour_"+callback)
		row := tg.NewInlineKeyboardRow(bl, br)
		rows = append(rows, row)
	}
	small = strings.Trim(small, ":")
	rowFunc := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("Назад", "whThree_"+small), tg.NewInlineKeyboardButtonData("Далее", "whFive_"+small))
	rowHome := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("🏠Главная страница", "start_1"))
	rows = append(rows, rowFunc, rowHome)
	text := textWH(small)
	msg := tg.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, text)
	keyboard := tg.InlineKeyboardMarkup{
		InlineKeyboard: rows,
	}
	msg.ReplyMarkup = &keyboard
	bot.Send(msg)
}

func wh_5(bot *tg.BotAPI, update tg.Update, small string) {
	var rows [][]tg.InlineKeyboardButton
	for i := 81; i < 101; i += 2 {
		name, callback := arrHandlerFromWH(small, i)
		bl := tg.NewInlineKeyboardButtonData(name, "whFive_"+callback)
		name, callback = arrHandlerFromWH(small, i+1)
		br := tg.NewInlineKeyboardButtonData(name, "whFive_"+callback)
		row := tg.NewInlineKeyboardRow(bl, br)
		rows = append(rows, row)
	}
	small = strings.Trim(small, ":")
	rowFunc := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("Назад", "whFour_"+small), tg.NewInlineKeyboardButtonData("Далее", "whSix_"+small))
	rowHome := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("🏠Главная страница", "start_1"))
	rows = append(rows, rowFunc, rowHome)
	text := textWH(small)
	msg := tg.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, text)
	keyboard := tg.InlineKeyboardMarkup{
		InlineKeyboard: rows,
	}
	msg.ReplyMarkup = &keyboard
	bot.Send(msg)
}

func wh_6(bot *tg.BotAPI, update tg.Update, small string) {
	var rows [][]tg.InlineKeyboardButton
	name, callback := arrHandlerFromWH(small, 101)
	bl := tg.NewInlineKeyboardButtonData(name, "whSix_"+callback)
	name, callback = arrHandlerFromWH(small, 102)
	br := tg.NewInlineKeyboardButtonData(name, "whSix_"+callback)
	name, callback = arrHandlerFromWH(small, 103)
	bdownRow := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData(name, "whSix_"+callback))
	row := tg.NewInlineKeyboardRow(bl, br)
	rows = append(rows, row, bdownRow)

	small = strings.Trim(small, ":")
	rowFunc := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("Назад", "whFive_"+small))
	rowHome := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("🏠Главная страница", "start_1"))
	rows = append(rows, rowFunc, rowHome)
	text := textWH(small)
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

	result = `Выбранные склады:`
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
		result += "\n-------------------------------------\nВыбрано максимально количество складов"
	}
	return
}
