package models

import "gorm.io/gorm"

type Test struct {
	gorm.Model
	Name string
	Sex int
	Password string
}
