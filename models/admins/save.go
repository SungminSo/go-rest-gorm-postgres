package admins

import (
	"../internal/constant"
	"errors"
	"fmt"
)

func (schema *AdminDatabase) Save(admin *Admin) error {
	err := schema.table.Save(admin).Error
	if err != nil {
		return errors.New(fmt.Sprintf(constant.SavingError, "admin", admin))
	}

	return nil
}