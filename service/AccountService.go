package service

import (
	"encoding/json"
	"github.com/globalsign/mgo/bson"
	"github.com/ndphu/csn-go-api/dao"
	"github.com/ndphu/csn-go-api/entity"
	"github.com/ndphu/google-api-helper"
	"time"
)

type AccountService struct {
}

type KeyDetails struct {
	Type        string `json:"type"`
	ProjectId   string `json:"project_id"`
	ClientEmail string `json:"client_email"`
	ClientId    string `json:"client_id"`
}

func (s *AccountService) Save(account *entity.DriveAccount) error {
	account.Id = bson.NewObjectId();
	return dao.Collection("drive_account").Insert(account)
}

func (s *AccountService) FindAll() ([]*entity.DriveAccount, error) {
	var list []*entity.DriveAccount
	err := dao.Collection("drive_account").Find(bson.M{}).All(&list)
	return list, err
}

func (s *AccountService) FindAccount(id string) (*entity.DriveAccount, error) {
	var acc entity.DriveAccount
	err := dao.Collection("drive_account").FindId(bson.ObjectIdHex(id)).One(&acc)
	return &acc, err
}

func (s *AccountService) InitializeKey(acc *entity.DriveAccount, key []byte) (error) {
	var kd KeyDetails
	err := json.Unmarshal(key, &kd)
	if err != nil {
		return err
	}
	acc.Key = string(key)
	acc.ClientId = kd.ClientId
	acc.ClientEmail = kd.ClientEmail
	acc.ProjectId = kd.ProjectId
	acc.Type = kd.Type

	return nil
}

func (s *AccountService) UpdateKey(id string, key []byte) (error) {
	var acc entity.DriveAccount
	err := dao.Collection("drive_account").FindId(bson.ObjectIdHex(id)).One(&acc)
	if err != nil {
		return err
	}
	err = s.InitializeKey(&acc, key)
	if err != nil {
		return err
	}
	return dao.Collection("drive_account").UpdateId(bson.ObjectIdHex(id), &acc)
}
func (s *AccountService) UpdateCachedQuota(id string) error {
	driveAccount, err := s.FindAccount(id)
	if err != nil {
		return err
	}
	driveService, err := google_api_helper.GetDriveService([]byte(driveAccount.Key))
	if err != nil {
		return err
	}
	quota, err := driveService.GetQuotaUsage()
	if err != nil {
		return err
	}
	return dao.Collection("drive_account").Update(
		bson.M{"_id": bson.ObjectIdHex(id)},
		bson.M{
			"$set": bson.M{
				"usage":                quota.Usage,
				"limit":                quota.Limit,
				"quotaUpdateTimestamp": time.Now(),
			},
		})
}

type FileLookup struct {
	Id      bson.ObjectId         `json:"_id" bson:"_id"`
	DriveId string                `json:"driveId" bson:"driveId"`
	Name    string                `json:"name" bson:"name"`
	Account []entity.DriveAccount `json:"account" bson:"account"`
}

func (s *AccountService) GetDownloadLink(fileId string) (string, error) {
	fileLookup := FileLookup{}
	err := dao.Collection("file").Pipe([]bson.M{
		{
			"$match": bson.M{
				"_id": bson.ObjectIdHex(fileId),
			},
		},
		{
			"$lookup": bson.M{
				"from":         "drive_account",
				"localField":   "driveAccount",
				"foreignField": "_id",
				"as":           "account",
			},
		},
	}).One(&fileLookup)
	if err != nil {
		return "", err
	}
	driveService, err := google_api_helper.GetDriveService([]byte(fileLookup.Account[0].Key))
	if err != nil {
		return "", err
	}
	link, err := driveService.GetDownloadLink(fileLookup.DriveId)
	if err != nil {
		return "", err
	}
	return link, nil
}

var accountService *AccountService

func GetAccountService() *AccountService {
	if accountService == nil {
		accountService = &AccountService{}
	}

	return accountService
}
