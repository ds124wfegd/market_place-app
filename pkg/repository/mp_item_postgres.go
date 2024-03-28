package repository

import (
	"fmt"
	"strings"

	"github.com/ds124wfegd/mp-app"
	"github.com/jmoiron/sqlx"
)

type MpItemPostgres struct {
	db *sqlx.DB
}

func NewMpItemPostgres(db *sqlx.DB) *MpItemPostgres {
	return &MpItemPostgres{db: db}
}

func (r *MpItemPostgres) Create(listId int, item mp.MpItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description) values ($1, $2) RETURNING id", mpItemsTable)

	row := tx.QueryRow(createItemQuery, item.Title, item.Description, item.Filepath, item.Price)
	err = row.Scan(&itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createListItemsQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) values ($1, $2)", listsItemsTable)
	_, err = tx.Exec(createListItemsQuery, listId, itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}

func (r *MpItemPostgres) GetAll(userId, listId int) ([]mp.MpItem, error) {
	var items []mp.MpItem
	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, tl.filepath, tl.price, ti.done FROM %s ti INNER JOIN %s li on li.item_id = ti.id
									INNER JOIN %s ul on ul.list_id = li.list_id WHERE li.list_id = $1 AND ul.user_id = $2`,
		mpItemsTable, listsItemsTable, usersListsTable)
	if err := r.db.Select(&items, query, listId, userId); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *MpItemPostgres) GetById(userId, itemId int) (mp.MpItem, error) {
	var item mp.MpItem
	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, tl.filepath, tl.price, ti.done FROM %s ti INNER JOIN %s li on li.item_id = ti.id
									INNER JOIN %s ul on ul.list_id = li.list_id WHERE ti.id = $1 AND ul.user_id = $2`,
		mpItemsTable, listsItemsTable, usersListsTable)
	if err := r.db.Get(&item, query, itemId, userId); err != nil {
		return item, err
	}

	return item, nil
}

func (r *MpItemPostgres) Delete(userId, itemId int) error {
	query := fmt.Sprintf(`DELETE FROM %s ti USING %s li, %s ul 
									WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $1 AND ti.id = $2`,
		mpItemsTable, listsItemsTable, usersListsTable)
	_, err := r.db.Exec(query, userId, itemId)
	return err
}

func (r *MpItemPostgres) Update(userId, itemId int, input mp.UpdateItemInput) error {
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

	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, *input.Done)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s ti SET %s FROM %s li, %s ul
									WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $%d AND ti.id = $%d`,
		mpItemsTable, setQuery, listsItemsTable, usersListsTable, argId, argId+1)
	args = append(args, userId, itemId)

	_, err := r.db.Exec(query, args...)
	return err
}
