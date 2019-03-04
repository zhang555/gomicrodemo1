package handler

import (
	"common/db"
	"common/log"
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
}
