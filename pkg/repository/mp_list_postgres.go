package repository

import (
	"fmt"
	"strings"

	"github.com/ds124wfegd/mp-app"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type MpListPostgres struct {
	db *sqlx.DB
}

func NewMpListPostgres(db *sqlx.DB) *MpListPostgres {
	return &MpListPostgres{db: db}
}

func (r *MpListPostgres) Create(userId int, list mp.MpList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", mpListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description, list.Filepath, list.Price)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTable)
	_, err = tx.Exec(createUsersListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *MpListPostgres) GetAll(userId int) ([]mp.MpList, error) {
	var lists []mp.MpList

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description, tl.filepath, tl.price FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1",
		mpListsTable, usersListsTable)
	err := r.db.Select(&lists, query, userId)

	return lists, err
}

func (r *MpListPostgres) GetById(userId, listId int) (mp.MpList, error) {
	var list mp.MpList

	query := fmt.Sprintf(`SELECT tl.id, tl.title, tl.description, tl.filepath, tl.price FROM %s tl
								INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2`,
		mpListsTable, usersListsTable)
	err := r.db.Get(&list, query, userId, listId)

	return list, err
}

func (r *MpListPostgres) Delete(userId, listId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id=$1 AND ul.list_id=$2",
		mpListsTable, usersListsTable)
	_, err := r.db.Exec(query, userId, listId)

	return err
}

func (r *MpListPostgres) Update(userId, listId int, input mp.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	if input.Filepath != nil {
		setValues = append(setValues, fmt.Sprintf("filepath=$%d", argId))
		args = append(args, *input.Filepath)
		argId++
	}

	if input.Price != nil {
		setValues = append(setValues, fmt.Sprintf("price=$%d", argId))
		args = append(args, *input.Price)
		argId++
	}

	// title=$1, description=$2...
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d",
		mpListsTable, setQuery, usersListsTable, argId, argId+1)
	args = append(args, listId, userId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}
