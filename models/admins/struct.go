package admins

import "github.com/jinzhu/gorm"

type Admin struct {
	gorm.Model
	UUID     string `json:"uuid" gorm:"type:uuid;unique_index"`
	AdminID  string `json:"adminID" gorm:"column:admin_id"`
	Password string `json:"password" gorm:"column:password"`
}
