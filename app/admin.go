package app

import (
	"../internal/constant"
	"../models/admins"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

func (app *ProjectApp) Register(adminID, password string) (string, error) {
	_, err := app.admins.FindByID(adminID)
	if err == nil {
		return "", nil
	} else if !strings.Contains(err.Error(), constant.NotFouncStr) {
		return "", err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	hashedPasswordStr := string(hashedPassword[:])

	admin := &admins.Admin{
		UUID: uuid.New().String(),
		AdminID: adminID,
		Password: hashedPasswordStr,
	}

	err = app.admins.Save(admin)
	if err != nil {
		return "", err
	}

	return admin.UUID, nil
}