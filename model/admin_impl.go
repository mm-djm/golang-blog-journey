package model

import (
	"errors"
	"golang-blog-journey/util/log"
)

func CheckUserNameExist(email, requestId string) (bool, error) {
	data, err := getSimpleAdminData(email, requestId)
	if err != nil {
		log.Errorw("CheckUserNameExist", "RequestID", requestId, "err", err)
		return false, err
	}
	log.Infow("CheckUserNameExist", "RequestID", requestId, "data", data)
	if len(data) > 0 {
		return true, nil
	}
	return false, nil
}

func GetSimpleAdminData(email, requestId string) (*Admin, error) {
	data, err := getSimpleAdminData(email, requestId)
	if err != nil {
		log.Errorw("GetSimpleAdminData", "RequestID", requestId, "err", err)
		return nil, err
	}
	if len(data) == 0 {
		log.Errorw("GetSimpleAdminData", "RequestID", requestId, "err", "no admin data")
		return nil, errors.New("no admin data")
	}
	return data[0], nil
}

func GetAdminData(requestId string) ([]*Admin, error) {
	data, err := getAdminData(requestId)
	if err != nil {
		log.Errorw("GetAdminData", "RequestID", requestId, "err", err)
		return nil, err
	}
	if len(data) == 0 {
		log.Errorw("GetAdminData", "RequestID", requestId, "err", "no admin data")
		return nil, errors.New("no admin data")
	}
	return data, nil
}

func AddAdmin(name, password, email, requestId string) error {
	role := "guest" //initial
	err := addAdminData(name, password, role, email, requestId)
	if err != nil {
		log.Errorw("AddAdmin", "RequestID", requestId, "err", err)
		return err
	}
	return nil
}

func RemoveAdmin(email, requestId string) error {
	err := updateAdminData("", "", "", email, requestId, 1)
	if err != nil {
		log.Errorw("RemoveAdmin", "RequestID", requestId, "err", err)
		return err
	}
	return nil
}

func UpdateAdmin(name, password, role, email, requestId string) error {
	err := updateAdminData(name, password, role, email, requestId, 0)
	if err != nil {
		log.Errorw("UpdateAdmin", "RequestID", requestId, "err", err)
		return err
	}
	return nil
}
