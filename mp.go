package mp

import "errors"

type MpList struct {
	Id          int     `json:"id" db:"id"`
	Title       string  `json:"title" db:"title" binding:"required"`
	Description string  `json:"description" db:"description"`
	Filepath    string  `json:"filepath" db:"filepath"`
	Price       float32 `json:"price" db:"price"`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type MpItem struct {
	Id          int     `json:"id" db:"id"`
	Title       string  `json:"title" db:"title" binding:"required"`
	Description string  `json:"description" db:"description"`
	Filepath    string  `json:"filepath" db:"filepath"`
	Price       float32 `json:"price" db:"price"`
	Done        bool    `json:"done" db:"done"`
}

type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}

type UpdateListInput struct {
	Title       *string  `json:"title"`
	Description *string  `json:"description"`
	Filepath    *string  `json:"filepath"`
	Price       *float32 `json:"price"`
}

func (i UpdateListInput) Validate() error {
	if i.Title == nil && i.Description == nil && i.Filepath == nil && i.Price == nil {
		return errors.New("update structure has no values")
	}

	return nil
}

type UpdateItemInput struct {
	Title       *string  `json:"title"`
	Description *string  `json:"description"`
	Filepath    *string  `json:"filepath"`
	Price       *float32 `json:"price"`
	Done        *bool    `json:"done"`
}

func (i UpdateItemInput) Validate() error {
	if i.Title == nil && i.Description == nil && i.Done == nil && i.Filepath == nil && i.Price == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
