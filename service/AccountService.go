package service

import (
	"encoding/json"
	"github.com/globalsign/mgo/bson"
	"github.com/ndphu/csn-go-api/dao"
	"github.com/ndphu/csn-go-api/entity"
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

func (s *AccountService) FindAll() ([]entity.DriveAccount, error) {
	var list []entity.DriveAccount
	err := dao.Collection("drive_account").Find(bson.M{}).All(&list)
	return list, err
}

func (s *AccountService) FindAccount(id string) (*entity.DriveAccount, error) {
	var acc entity.DriveAccount
	err := dao.Collection("drive_account").FindId(bson.ObjectIdHex(id)).One(&acc)
	return &acc, err
}

func (s *AccountService) InitializeKey(acc*entity.DriveAccount, key []byte) (error) {
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

func (s*AccountService) UpdateKey(id string, key []byte) (error)  {
	var acc entity.DriveAccount
	err := dao.Collection("drive_account").FindId(bson.ObjectIdHex(id)).One(&acc)
	if err != nil {
		return err
	}
	err=s.InitializeKey(&acc, key)
	if err != nil {
		return err
	}
	return dao.Collection("drive_account").UpdateId(bson.ObjectIdHex(id), &acc)
}

var accountService *AccountService

func GetAccountService() *AccountService {
	if accountService == nil {
		accountService = &AccountService{}
	}

	return accountService
}
