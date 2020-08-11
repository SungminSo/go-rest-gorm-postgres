package admins

import (
	"../../internal/constant"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

func (schema *AdminDatabase) Find(uuid string) (*Admin, error) {
	admin := &Admin{}
	err := schema.table.First(admin, "uuid=?", uuid).Error
	if err !=nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, errors.New(fmt.Sprintf(constant.NotFoundError, "admins", uuid))
		}
		return nil, errors.New(fmt.Sprintf(constant.FindingError, "admins", uuid))
	}

	return admin, nil
}