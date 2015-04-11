package ormapper

import (
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var PrefixLines = map[string][]string{
	"あ": []string{"a", "i", "u", "e", "o"},
	"か": []string{"ka", "ki", "ku", "ke", "ko", "ga", "gi", "gu", "ge", "go"},
	"さ": []string{"sa", "si", "su", "se", "so", "za", "zi", "zu", "ze", "zo"},
	"た": []string{"ta", "ti", "tu", "te", "to", "da", "di", "du", "de", "do"},
	"な": []string{"na", "ni", "nu", "ne", "no"},
	"は": []string{"ha", "hi", "hu", "he", "ho", "ba", "bi", "bu", "be", "bo", "pa", "pi", "pu", "pe", "po"},
	"ま": []string{"ma", "mi", "mu", "me", "mo"},
	"や": []string{"ya", "yu", "yo"},
	"ら": []string{"ra", "ri", "ru", "re", "ro"},
	"わ": []string{"wa"},
}

var BracupLines = map[string][]string{
	"C": []string{"AAA", "AA", "A", "B", "C"},
	"D": []string{"D"},
	"E": []string{"E"},
	"F": []string{"F"},
	"G": []string{"G"},
	"H": []string{"H"},
	"I": []string{"I", "J", "K", "L", "M", "N", "O", "P"},
}

var DB gorm.DB

func init() {
	var err error

	runmode := beego.AppConfig.String("runmode")
	datasource := beego.AppConfig.String("datasource")

	switch runmode {
	case "dev":
		fallthrough
	case "prod":
		DB, err = gorm.Open("postgres", datasource)
		if err != nil {
			panic(err)
		}

		DB.DB()
		DB.DB().Ping()
		DB.DB().SetMaxIdleConns(100)
		DB.DB().SetMaxOpenConns(100)

		if runmode == "dev" {
			DB.LogMode(true)
		}
	default:
		DB, err = gorm.Open("sqlite3", datasource)
		if err != nil {
			panic(err)
		}

		DB.DB()
	}

	DB.SingularTable(true)
}
