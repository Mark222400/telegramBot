package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/mattn/go-sqlite3"
)

func accept(bot *tg.BotAPI, update tg.Update, small string) {

	text := ""
	smallWH := strings.Split(small, "|")[0]
	smallBT := strings.Split(small, "|")[1]
	indexes := strings.Split(small, "|")[2]
	coef := strings.Split(small, "|")[3]
	if len(strings.Split(indexes, ":")) > 1 {
		index1 := strings.Split(indexes, ":")[0]
		index2 := strings.Split(indexes, ":")[1]
		text = textForResultWithTwoAcc(smallWH, smallBT, index1, index2, coef)
	}
	if len(strings.Split(indexes, ":")) == 1 {
		index1 := strings.Split(strings.Split(small, "|")[2], ":")[0]
		text = textForResultWithOneAcc(smallWH, smallBT, index1, coef)
	}
	rowHome := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("🏠Главная страница", "start_1"))
	buttonClose := tg.NewInlineKeyboardButtonData("❌Закрыть", "close_1")

	rowFunc := tg.NewInlineKeyboardRow(buttonClose)

	keyboard := tg.NewInlineKeyboardMarkup(rowHome, rowFunc)

	msg := tg.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, text)
	msg.ReplyMarkup = &keyboard
	bot.Send(msg)
	chatID := update.CallbackQuery.Message.Chat.ID
	for _, v := range strings.Split(smallWH, ":") {
		sendQuery(int(chatID), coef, indexes, whMap[v], boxtypeMap[smallBT])
	}

}

func textForResultWithOneAcc(wh string, bt string, index1 string, coef string) (text string) {
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
	text += "\n✪Коэффициент: до " + coef
	text += "\n✪Заявка подтвеждена"
	return text
}

func textForResultWithTwoAcc(wh string, bt string, index1 string, index2 string, coef string) (text string) {
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
	text += "\n✪Коэффициент: до " + coef
	text += "\n✪Заявка подтвеждена☑"
	return text
}

func sendQuery(chatID int, coef string, data string, warehouse string, boxType string) {
	var dateOne, dateTwo string
	arr := strings.Split(data, ":")
	if len(arr) > 1 {
		ind1, _ := strconv.Atoi(arr[0])
		dateOne = date()[ind1]
		ind2, _ := strconv.Atoi(arr[1])
		dateTwo = date()[ind2]
	} else {
		ind1, _ := strconv.Atoi(arr[0])
		dateOne = date()[ind1]
		dateTwo = "nil"
	}
	db, err := sql.Open("sqlite3", "query.db")

	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO HHM (ChatID,WH,DATE,Coef,BoxType) VALUES(?,?,?,?,?)", chatID, warehouse, dateOne +"|"+ dateTwo, coef, boxType)
	if err != nil {
		fmt.Println(err)
	}

}


func delete(bot *tg.BotAPI, update tg.Update, ID string){
db, err := sql.Open("sqlite3", "query.db")
if err != nil {
	fmt.Println(err)
}
defer db.Close()
_, err = db.Exec("DELETE FROM HHM WHERE ID=?",ID)
if err!=nil{
	fmt.Println()
}
list(bot,update, strconv.Itoa(int(update.CallbackQuery.Message.Chat.ID)))
}