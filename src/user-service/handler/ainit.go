package handler

import (
	"gomicrodemo1/common/db"
	"gomicrodemo1/common/log"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

var (
	DB  *gorm.DB
	Log *logrus.Logger
)

func init() {
	DB = db.DB
	Log = log.Log

	DB.LogMode(true)
}
