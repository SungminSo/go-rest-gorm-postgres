package users

import (
	"../../internal/constant"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

func (schema *UserDatabase) Find(uuid string) (*User, error) {
	user := &User{}
	err := schema.table.First(user, "uuid=?", uuid).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, errors.New(fmt.Sprintf(constant.NotFoundError, "users", uuid))
		}
		return nil, errors.New(fmt.Sprintf(constant.FindingError, "users", uuid))
	}

	return user, nil
}

func (schema *UserDatabase) All() ([]*User, error) {
	var users []*User
	err := schema.table.Find(users).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return users, errors.New(fmt.Sprintf(constant.NotFoundError, "users", users))
		}
		return nil, errors.New(fmt.Sprintf(constant.FindingError, "users", users))
	}

	return users, nil
}
