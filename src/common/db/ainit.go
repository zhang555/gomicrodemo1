package db

import (
	"common/log"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	"time"
)

var (
	DB *gorm.DB
)

func init() {

	var (
		err error
	)

	host := os.Getenv("DB_HOST")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	if host == "" {
		host = "localhost"
	}
	if dbName == "" {
		dbName = "micro"
	}
	if port == "" {
		port = "3306"
	}

	path := fmt.Sprintf("root:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		password, host, port, dbName,
	)

	waitSecond := 2

	for {
		DB, err = gorm.Open("mysql", path)
		if err != nil {
			log.Log.
				//WithField("host", host).
				//WithField("dbName", dbName).
				//WithField("port", port).
				//WithField("password", password).
				Error("failed to connect database")

			time.Sleep(time.Duration(waitSecond) * time.Second)
			//waitSecond++
			continue
		}
		break
	}

	log.Log.Info("connect database success ")

	//DB.LogMode(true)
	DB.LogMode(false)
	DB.SingularTable(true)

}
