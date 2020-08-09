package users

import (
	"../../internal/constant"
	"errors"
	"fmt"
)

func (schema *UserDatabase) Save(user *User) error {
	err := schema.table.Save(user).Error
	if err != nil {
		return errors.New(fmt.Sprintf(constant.SavingError, "user", user))
	}

	return nil
}