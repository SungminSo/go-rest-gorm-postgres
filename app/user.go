package app

import (
	"../internal/constant"
	"../models/users"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

func (app *ProjectApp) UserRegister(name, phone, email, password string) (string, error) {
	_, err := app.users.FindByPhone(phone)
	if err == nil {
		return "", nil
	} else if !strings.Contains(err.Error(), constant.NotFoundStr) {
		return "", err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	hashedPasswordStr := string(hashedPassword[:])

	user := &users.User{
		UUID:  uuid.New().String(),
		Name:  name,
		Phone: phone,
		Email: email,
		Password: hashedPasswordStr,
		Status: constant.APPLIED,
	}

	err = app.users.Save(user)
	if err != nil {
		return "", err
	}

	return user.UUID, nil
}