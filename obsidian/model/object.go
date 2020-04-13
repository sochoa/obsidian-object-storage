package model

import (
	"github.com/jinzhu/gorm"
)

type Object struct {
	gorm.Model
	size     int64
	checkSum string
}
