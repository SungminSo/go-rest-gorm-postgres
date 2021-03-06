package app

import (
	"../internal/constant"
	"../internal/token"
	"../models/admins"
	"../models/users"
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"strings"
)

func (app *ProjectApp) AdminRegister(adminID, password string) (string, error) {
	_, err := app.admins.FindByID(adminID)
	if err == nil {
		return "", errors.New("ID already exists")
	} else if !strings.Contains(err.Error(), constant.NotFoundStr) {
		return "", err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	hashedPasswordStr := string(hashedPassword[:])

	admin := &admins.Admin{
		UUID:     uuid.New().String(),
		AdminID:  adminID,
		Password: hashedPasswordStr,
	}

	err = app.admins.Save(admin)
	if err != nil {
		return "", err
	}

	return admin.UUID, nil
}

func (app *ProjectApp) Login(adminID, password string) (string, error) {
	admin, err := app.admins.FindByID(adminID)
	if err != nil {
		return "", errors.New("unauthorized." + err.Error())
	}

	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password))
	if err != nil {
		return "", errors.New("unauthorized." + err.Error())
	}

	accessToken := token.GenerateAdminAccessToken(admin.UUID, admin.AdminID)

	return accessToken, nil
}

func (app *ProjectApp) GetUserList(page string) (int, []*users.User, error) {
	userList, err := app.users.All()
	if err != nil {
		return 0, []*users.User{}, err
	}

	limits := 10
	pageNum, err := strconv.Atoi(page)
	if err != nil {
		return 0, []*users.User{}, err
	}

	fromIndex := limits * (pageNum - 1)
	toIndex := limits * pageNum

	if fromIndex > len(userList) {
		return 0, []*users.User{}, errors.New("not found any users in this page")
	}
	if toIndex > len(userList) {
		toIndex = len(userList)
	}

	return len(userList), userList[fromIndex:toIndex], nil
}

func (app *ProjectApp) ApproveRegistration(userUUID string) (string, string, error) {
	user, err := app.users.Find(userUUID)
	if err != nil {
		return "", "", err
	}

	if user.Status != constant.APPLIED {
		return "", user.Status, errors.New("invalid status change request")
	}

	user.Status = constant.APPROVED
	err = app.users.Save(user)
	if err != nil {
		return "","", err
	}

	return user.UUID, user.Status, nil
}

func (app *ProjectApp) RejectRegistration(userUUID string) (string, string, error) {
	user, err := app.users.Find(userUUID)
	if err != nil {
		return "", "", err
	}

	if user.Status != constant.APPLIED {
		return "", user.Status, errors.New("invalid status change request")
	}

	user.Status = constant.REJECTED
	err = app.users.Save(user)
	if err != nil {
		return "","", err
	}

	return user.UUID, user.Status, nil
}