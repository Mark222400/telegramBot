package main

import (
	"strings"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func callbackHandler(bot *tg.BotAPI, update tg.Update) {
	data := update.CallbackQuery.Data
	big:=strings.Split(data,"_")[0]
	small:=strings.Split(data,"_")[1]
	switch big {
	case "start":
		startPage(bot, update, 1)


	case "whOne":
			wh_1(bot, update, small)

	case "whTwo":
		wh_2(bot, update, small)
		
	case "whThree":
		wh_3(bot,update, small)
	
	case "whFour":
		wh_4(bot,update, small)
	
	case "whFive":
		wh_5(bot,update, small)
	case "whSix":
		wh_6(bot,update,small)
	case "boxtype":
		boxType(bot, update, small)
	}

}
