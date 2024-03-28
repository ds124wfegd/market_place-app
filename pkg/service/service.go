package service

import (
	"github.com/ds124wfegd/mp-app"
	"github.com/ds124wfegd/mp-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user mp.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type MpList interface {
	Create(userId int, list mp.MpList) (int, error)
	GetAll(userId int) ([]mp.MpList, error)
	GetById(userId, listId int) (mp.MpList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input mp.UpdateListInput) error
}

type MpItem interface {
	Create(userId, listId int, item mp.MpItem) (int, error)
	GetAll(userId, listId int) ([]mp.MpItem, error)
	GetById(userId, itemId int) (mp.MpItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input mp.UpdateItemInput) error
}

type Service struct {
	Authorization
	MpList
	MpItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		MpList:        NewMpListService(repos.MpList),
		MpItem:        NewMpItemService(repos.MpItem, repos.MpList),
	}
}
