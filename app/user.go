package app

import (
	"../internal/constant"
	"../internal/token"
	"../models/users"
	"errors"
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

func (app *ProjectApp) SignIn(phone, password string) (string, error) {
	user, err := app.users.FindByPhone(phone)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	if user.Status != constant.APPROVED {
		return "", errors.New("unauthorized. not approved yet")
	}

	accessToken := token.GenerateUserAccessToken(user.UUID, user.Name, user.Phone)

	return accessToken, nil
}

func (app *ProjectApp) GetUserInfo(userUUID string) (*users.User, error) {
	user, err := app.users.Find(userUUID)
	if err != nil {
		return &users.User{}, err
	}

	return user, nil
}