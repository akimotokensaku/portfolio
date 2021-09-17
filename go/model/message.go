package model

import (
	"github.com/jinzhu/gorm"
)

type Message struct {
	gorm.Model
	Name       string `validate:"required,unique,max=50"`
	Title      string `validate:"required,unique,max=50"`
	Body       string `validate:"required,unique,max=255"`
	FromUserId int    `validate:"required"`
	ToUserId   int    `validate:"required"`
}