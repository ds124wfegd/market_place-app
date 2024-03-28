package repository

import (
	"github.com/ds124wfegd/mp-app"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user mp.User) (int, error)
	GetUser(username, password string) (mp.User, error)
}

type MpList interface {
	Create(userId int, list mp.MpList) (int, error)
	GetAll(userId int) ([]mp.MpList, error)
	GetById(userId, listId int) (mp.MpList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input mp.UpdateListInput) error
}

type MpItem interface {
	Create(listId int, item mp.MpItem) (int, error)
	GetAll(userId, listId int) ([]mp.MpItem, error)
	GetById(userId, itemId int) (mp.MpItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input mp.UpdateItemInput) error
}

type Repository struct {
	Authorization
	MpList
	MpItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		MpList:        NewMpListPostgres(db),
		MpItem:        NewMpItemPostgres(db),
	}
}
