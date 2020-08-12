package users

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UUID     string `json:"uuid" gorm:"type:uuid;unique_index"`
	Name     string `json:"name" gorm:"column:name"`
	Phone    string `json:"phone" gorm:"column:phone"`
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"password" gorm:"column:password"`
	Status   string `json:"status" gorm:"column:status"`
}
