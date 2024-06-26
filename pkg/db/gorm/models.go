package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

var db *gorm.DB


type Model struct {
	ID          int   `gorm:"primary_key" json:"id"`
	CreatedOn   int   `json:"created_on"`
	ModifiedOn  int   `json:"modified-on"`
}


func init() {
	var (
		err error
		dbType, dbName, user, password, host, tablePrefix string
	)
	sec, err := Cfg.GetSection("database")
	if err != nil {
		log.Fatalf("Fail to get section database: %v", err)
	}
	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("Name").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8" +
		"&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))
	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDb() {
	defer db.Close()
}
