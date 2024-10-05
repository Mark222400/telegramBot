package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/mattn/go-sqlite3"
)

type query1 struct {
	ID     string
	ChatID string
	WH     string
	BT     string
	Date   string
	CF     string
}

func uvedomlator(bot *tg.BotAPI) {
	t:=time.NewTicker(500*time.Millisecond)
	for range t.C {	n := selectorQuery()
	for _, v := range n {
		db, err := sql.Open("sqlite3", "eldar.db")
		if err != nil {
			fmt.Println(err)
		}
		defer db.Close()
		rows, err := db.Query("SELECT Дата, Коэффициент, Название_Склада, Тип_Поставки FROM AWU WHERE (Коэффициент>-1 and (Коэффициент<? or Коэффициент=?)) and Название_Склада=? and  Тип_Поставки=? and Дата=?", v.CF, v.CF, v.WH, v.BT, v.Date)
		if err != nil {
			fmt.Println(err)
		}
		for rows.Next() {
			n := query1{}
			rows.Scan(&n.Date, &n.CF, &n.WH, &n.BT)
			if n.Date != "" {
				chatID, err := strconv.Atoi(v.ChatID)
				if err != nil {
					fmt.Println(err)
				}
				text := fmt.Sprintf(`⭐Лот по вашей заявке
Дата: %s
Склад: %s
Тип Поставки: %s
Коэффициент: %s`, n.Date, n.WH, n.BT, n.CF)
				msg := tg.NewMessage(int64(chatID), text)
				bot.Send(msg)
				deleteQuery(v.ID)
				goto restart
			}
			
		}
			
		}
		restart:
	}

}

func selectorQuery() []query1 {
	db, err := sql.Open("sqlite3", "query.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM HHM;")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	arr := []query1{}
	for rows.Next() {
		g := query1{}
		rows.Scan(&g.ID, &g.ChatID, &g.WH, &g.Date, &g.CF, &g.BT)
		g.Date = strings.Trim(g.Date, "|nil")
		arrTemp := dateParser(g)
		arr = append(arr, arrTemp...)

	}
	return arr
}

func dateParser(g query1) (s []query1) {
	if len(strings.Split(g.Date, "|")) == 1 {
		dateTime, err := time.Parse("02.01.2006", g.Date)
		if err != nil {
			fmt.Println(err)
		}
		g.Date = dateTime.Format("2006-01-02") + "T00:00:00Z"
		return []query1{
			{ID: g.ID,
				ChatID: g.ChatID,
				WH:     g.WH,
				Date:   g.Date,
				CF:     g.CF,
				BT:     g.BT},
		}
	} else {
		dateOne := strings.Split(g.Date, "|")[0]
		dateTwo := strings.Split(g.Date, "|")[1]
		dateTime, err := time.Parse("02.01.2006", dateTwo)
		if err != nil {
			fmt.Println(err)
		}
		dateTwo = dateTime.Format("2006-01-02") + "T00:00:00Z"
		dateTime, err = time.Parse("02.01.2006", dateOne)
		if err != nil {
			fmt.Println(err)
		}
		dateOne = dateTime.Format("2006-01-02") + "T00:00:00Z"
		arr := []string{dateOne}
		var date string
		for date != dateTwo {
			dateTime, err := time.Parse("2006-01-02T00:00:00Z", arr[len(arr)-1])
			if err != nil {
				fmt.Println(err)
			}
			dateTime = dateTime.Add(24 * time.Hour)
			date = dateTime.Format("2006-01-02") + "T00:00:00Z"
			arr = append(arr, date)
		}
		for _, v := range arr {
			ni := query1{ID: g.ID,
				ChatID: g.ChatID,
				WH:     g.WH,
				Date:   v,
				CF:     g.CF,
				BT:     g.BT}

			s = append(s, ni)
		}
		return s
	}
}

func deleteQuery(ID string) {
	db,err:=sql.Open("sqlite3","query.db")
	if err!=nil{
		fmt.Println(err)
	}
	_,err=db.Exec("DELETE FROM HHM WHERE ID=?",ID)
	if err!=nil{fmt.Println(err)}
	fmt.Println("Заявку удалил!")
	db.Close()
}