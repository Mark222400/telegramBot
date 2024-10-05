package main

var whArr = []string{
	"Пустышка",
	"Коледино",
	"Подольск",
	"Тула",
	"Электросталь",
	"Казань",
	"Краснодар (Тихорецкая)",
	"Невинномысск",
	"Екатеринбург - Испытателей 14г",
	"Санкт-Петербург (Уткина Заводь)",
	"Новосибирск",
	"Алматы Атакент",
	"Астана",
	"Астана 2",
	"Белые Столбы",
	"Внуково СГТ",
	"Волгоград",
	"Голицыно СГТ",
	"Екатеринбург - Перспективный 12/2",
	"Котовск",
	"Минск",
	"Новосемейкино",
	"Обухово",
	"Обухово СГТ",
	"Подольск 3",
	"Подольск 4",
	"Радумля СГТ",
	"Рязань (Тюшевское)",
	"СЦ Абакан",
	"СЦ Абакан 2",
	"СЦ Актобе",
	"СЦ Артем",
	"СЦ Архангельск (ул Ленина)",
	"СЦ Астрахань",
	"СЦ Астрахань (Солянка)",
	"СЦ Байсерке",
	"СЦ Барнаул",
	"СЦ Белая Дача",
	"СЦ Белогорск",
	"СЦ Бишкек",
	"СЦ Брест",
	"СЦ Видное",
	"СЦ Владикавказ",
	"СЦ Владимир",
	"СЦ Внуково",
	"СЦ Вологда",
	"СЦ Вологда 2",
	"СЦ Воронеж",
	"СЦ Вёшки",
	"СЦ Иваново",
	"СЦ Ижевск",
	"СЦ Иркутск",
	"СЦ Калуга",
	"СЦ Кемерово",
	"СЦ Киров",
	"СЦ Крыловская",
	"СЦ Кузнецк",
	"СЦ Курск",
	"СЦ Липецк",
	"СЦ Махачкала",
	"СЦ Мурманск",
	"СЦ Набережные Челны",
	"СЦ Нижний Новгород",
	"СЦ Нижний Тагил",
	"СЦ Новокузнецк",
	"СЦ Новосибирск Пасечная",
	"СЦ Ноябрьск",
	"СЦ Обухово 2",
	"СЦ Омск",
	"СЦ Оренбург Центральная",
	"СЦ Ош",
	"СЦ Пермь 2",
	"СЦ Псков",
	"СЦ Пушкино",
	"СЦ Пятигорск",
	"СЦ Пятигорск (Этока)",
	"СЦ Ростов-на-Дону",
	"СЦ Самара",
	"СЦ Семей",
	"СЦ Серов",
	"СЦ Симферополь",
	"СЦ Смоленск 3",
	"СЦ Сургут",
	"СЦ Сыктывкар",
	"СЦ Тамбов",
	"СЦ Тверь",
	"СЦ Томск",
	"СЦ Тюмень",
	"СЦ Ульяновск",
	"СЦ Уральск",
	"СЦ Уфа",
	"СЦ Хабаровск",
	"СЦ Чебоксары 2",
	"СЦ Челябинск 2",
	"СЦ Череповец",
	"СЦ Чита 2",
	"СЦ Шушары",
	"СЦ Шымкент",
	"СЦ Ярославль",
	"Сц Брянск 2",
	"Чашниково",
	"Чехов 1, Новоселки вл 11 стр 2",
	"Чехов 2, Новоселки вл 11 стр 7",
	"Шушары СГТ",
}

var whMap = map[string]string{
	"0":   "",
	"1":   "Коледино",
	"2":   "Подольск",
	"3":   "Тула",
	"4":   "Электросталь",
	"5":   "Казань",
	"6":   "Краснодар (Тихорецкая)",
	"7":   "Невинномысск",
	"8":   "Екатеринбург - Испытателей 14г",
	"9":   "Санкт-Петербург (Уткина Заводь)",
	"10":  "Новосибирск",
	"11":  "Алматы Атакент",
	"12":  "Астана",
	"13":  "Астана 2",
	"14":  "Белые Столбы",
	"15":  "Внуково СГТ",
	"16":  "Волгоград",
	"17":  "Голицыно СГТ",
	"18":  "Екатеринбург - Перспективный 12/2",
	"19":  "Котовск",
	"20":  "Минск",
	"21":  "Новосемейкино",
	"22":  "Обухово",
	"23":  "Обухово СГТ",
	"24":  "Подольск 3",
	"25":  "Подольск 4",
	"26":  "Радумля СГТ",
	"27":  "Рязань (Тюшевское)",
	"28":  "СЦ Абакан",
	"29":  "СЦ Абакан 2",
	"30":  "СЦ Актобе",
	"31":  "СЦ Артем",
	"32":  "СЦ Архангельск (ул Ленина)",
	"33":  "СЦ Астрахань",
	"34":  "СЦ Астрахань (Солянка)",
	"35":  "СЦ Байсерке",
	"36":  "СЦ Барнаул",
	"37":  "СЦ Белая Дача",
	"38":  "СЦ Белогорск",
	"39":  "СЦ Бишкек",
	"40":  "СЦ Брест",
	"41":  "СЦ Видное",
	"42":  "СЦ Владикавказ",
	"43":  "СЦ Владимир",
	"44":  "СЦ Внуково",
	"45":  "СЦ Вологда",
	"46":  "СЦ Вологда 2",
	"47":  "СЦ Воронеж",
	"48":  "СЦ Вёшки",
	"49":  "СЦ Иваново",
	"50":  "СЦ Ижевск",
	"51":  "СЦ Иркутск",
	"52":  "СЦ Калуга",
	"53":  "СЦ Кемерово",
	"54":  "СЦ Киров",
	"55":  "СЦ Крыловская",
	"56":  "СЦ Кузнецк",
	"57":  "СЦ Курск",
	"58":  "СЦ Липецк",
	"59":  "СЦ Махачкала",
	"60":  "СЦ Мурманск",
	"61":  "СЦ Набережные Челны",
	"62":  "СЦ Нижний Новгород",
	"63":  "СЦ Нижний Тагил",
	"64":  "СЦ Новокузнецк",
	"65":  "СЦ Новосибирск Пасечная",
	"66":  "СЦ Ноябрьск",
	"67":  "СЦ Обухово 2",
	"68":  "СЦ Омск",
	"69":  "СЦ Оренбург Центральная",
	"70":  "СЦ Ош",
	"71":  "СЦ Пермь 2",
	"72":  "СЦ Псков",
	"73":  "СЦ Пушкино",
	"74":  "СЦ Пятигорск",
	"75":  "СЦ Пятигорск (Этока)",
	"76":  "СЦ Ростов-на-Дону",
	"77":  "СЦ Самара",
	"78":  "СЦ Семей",
	"79":  "СЦ Серов",
	"80":  "СЦ Симферополь",
	"81":  "СЦ Смоленск 3",
	"82":  "СЦ Сургут",
	"83":  "СЦ Сыктывкар",
	"84":  "СЦ Тамбов",
	"85":  "СЦ Тверь",
	"86":  "СЦ Томск",
	"87":  "СЦ Тюмень",
	"88":  "СЦ Ульяновск",
	"89":  "СЦ Уральск",
	"90":  "СЦ Уфа",
	"91":  "СЦ Хабаровск",
	"92":  "СЦ Чебоксары 2",
	"93":  "СЦ Челябинск 2",
	"94":  "СЦ Череповец",
	"95":  "СЦ Чита 2",
	"96":  "СЦ Шушары",
	"97":  "СЦ Шымкент",
	"98":  "СЦ Ярославль",
	"99":  "Сц Брянск 2",
	"100": "Чашниково",
	"101": "Чехов 1, Новоселки вл 11 стр 2",
	"102": "Чехов 2, Новоселки вл 11 стр 7",
	"103": "Шушары СГТ",

}


var boxtypeMap = map[string]string{
"1":"Короба",
"2":"Монопаллеты",
"3":"Суперсейф",
"4":"QR-поставка с коробами",
}


var whMapTwo = map[string]string{
	"":                    "0",
	"Коледино":            "1",
	"Подольск":             "2",
	"Тула":                "3",
	"Электросталь":        "4",
	"Казань":              "5",
	"Краснодар (Тихорецкая)": "6",
	"Невинномысск":        "7",
	"Екатеринбург - Испытателей 14г": "8",
	"Санкт-Петербург (Уткина Заводь)": "9",
	"Новосибирск":         "10",
	"Алматы Атакент":      "11",
	"Астана":              "12",
	"Астана 2":            "13",
	"Белые Столбы":        "14",
	"Внуково СГТ":         "15",
	"Волгоград":           "16",
	"Голицыно СГТ":        "17",
	"Екатеринбург - Перспективный 12/2": "18",
	"Котовск":             "19",
	"Минск":               "20",
	"Новосемейкино":       "21",
	"Обухово":             "22",
	"Обухово СГТ":         "23",
	"Подольск 3":          "24",
	"Подольск 4":          "25",
	"Радумля СГТ":         "26",
	"Рязань (Тюшевское)":  "27",
	"СЦ Абакан":           "28",
	"СЦ Абакан 2":         "29",
	"СЦ Актобе":           "30",
	"СЦ Артем":            "31",
	"СЦ Архангельск (ул Ленина)": "32",
	"СЦ Астрахань":        "33",
	"СЦ Астрахань (Солянка)": "34",
	"СЦ Байсерке":         "35",
	"СЦ Барнаул":          "36",
	"СЦ Белая Дача":       "37",
	"СЦ Белогорск":        "38",
	"СЦ Бишкек":           "39",
	"СЦ Брест":            "40",
	"СЦ Видное":           "41",
	"СЦ Владикавказ":      "42",
	"СЦ Владимир":         "43",
	"СЦ Внуково":          "44",
	"СЦ Вологда":          "45",
	"СЦ Вологда 2":        "46",
	"СЦ Воронеж":          "47",
	"СЦ Вёшки":            "48",
	"СЦ Иваново":          "49",
	"СЦ Ижевск":           "50",
	"СЦ Иркутск":          "51",
	"СЦ Калуга":           "52",
	"СЦ Кемерово":         "53",
	"СЦ Киров":            "54",
	"СЦ Крыловская":       "55",
	"СЦ Кузнецк":          "56",
	"СЦ Курск":            "57",
	"СЦ Липецк":           "58",
	"СЦ Махачкала":        "59",
	"СЦ Мурманск":         "60",
	"СЦ Набережные Челны": "61",
	"СЦ Нижний Новгород":  "62",
	"СЦ Нижний Тагил":     "63",
	"СЦ Новокузнецк":      "64",
	"СЦ Новосибирск Пасечная": "65",
	"СЦ Ноябрьск":         "66",
	"СЦ Обухово 2":        "67",
	"СЦ Омск":             "68",
	"СЦ Оренбург Центральная": "69",
	"СЦ Ош":               "70",
	"СЦ Пермь 2":          "71",
	"СЦ Псков":            "72",
	"СЦ Пушкино":          "73",
	"СЦ Пятигорск":        "74",
	"СЦ Пятигорск (Этока)": "75",
	"СЦ Ростов-на-Дону":   "76",
	"СЦ Самара":           "77",
	"СЦ Семей":            "78",
	"СЦ Серов":            "79",
	"СЦ Симферополь":      "80",
	"СЦ Смоленск 3":       "81",
	"СЦ Сургут":           "82",
	"СЦ Сыктывкар":        "83",
	"СЦ Тамбов":           "84",
	"СЦ Тверь":            "85",
	"СЦ Томск":            "86",
	"СЦ Тюмень":           "87",
	"СЦ Ульяновск":        "88",
	"СЦ Уральск":          "89",
	"СЦ Уфа":              "90",
	"СЦ Хабаровск":        "91",
	"СЦ Чебоксары 2":      "92",
	"СЦ Челябинск 2":      "93",
	"СЦ Череповец":        "94",
	"СЦ Чита 2":           "95",
	"СЦ Шушары":           "96",
	"СЦ Шымкент":          "97",
	"СЦ Ярославль":        "98",
	"Сц Брянск 2":         "99",
	"Чашниково":           "100",
	"Чехов 1, Новоселки вл 11 стр 2": "101",
	"Чехов 2, Новоселки вл 11 стр 7": "102",
	"Шушары СГТ":         "103",
}


var boxtypeMapTwo = map[string]string{
	"Короба": "1",
	"Монопаллеты":"2",
	"Суперсейф":"3",
	"QR-поставка с коробами":"4",
	}

