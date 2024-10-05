package main

import (
	"database/sql"
	"fmt"
	

	"strings"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/mattn/go-sqlite3"
)

type query struct{
	ID string
	DATA string
	WH string
	CF string
	Bt string
}

func list(bot *tg.BotAPI, update tg.Update,small string){
	arrQuery:=[]query{}
	text:=`Cписок ваших заявок:                                                                                                         гЫ `
	msg:=tg.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, text)
	db,err:=sql.Open("sqlite3","query.db")
	if err!=nil{
		fmt.Println(err)
	}
	defer db.Close()
	rows,err:=db.Query("SELECT ID, WH, DATE, Coef, BoxType FROM HHM WHERE chatID=?;",small)
	if err!=nil{
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next(){
		n:=query{}
		rows.Scan(&n.ID,&n.WH,&n.DATA,&n.CF,&n.Bt)
		arrQuery=append(arrQuery, n)
	}
	stroki:=[][]tg.InlineKeyboardButton{}
	for _,v:=range arrQuery{
		DATA:=strings.Trim(v.DATA,"|nil")
		DATA=strings.ReplaceAll(DATA,".2024","")
		DATA=strings.ReplaceAll(DATA,"|"," - ")
		
		rowOne:=tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("✏️"+v.WH+" "+string(v.Bt[:6])+" "+DATA+" до x"+v.CF,"timeEdit_"+whMapTwo[v.WH]+"|"+boxtypeMapTwo[v.Bt]))
		rowTwo:=tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("🔺🗑️Удалить заявку","delete_"+v.ID))
		stroki=append(stroki, rowOne,rowTwo)
	}
	
	buttonClose := tg.NewInlineKeyboardButtonData("❌Закрыть", "close_1")
rowHome := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("🏠Главная страница", "start_1"),buttonClose)

	stroki = append(stroki, rowHome)
	keyboard:=tg.InlineKeyboardMarkup{
		InlineKeyboard: stroki,
	}
	msg.ReplyMarkup=&keyboard
	bot.Send(msg)
}