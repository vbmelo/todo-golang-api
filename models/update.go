package models

import "github.com/vbmelo/todo-golang-api/db"

func Update(id int64, todo Todo) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`UPDATE todos SET title=$2, description=$3, done=$4 WHERE id=$1`, id, todo.Title, todo.Description, todo.Done)
	if err != nil {
		return 0, err
	}

	// Returns the number of rows affected
	// Because we've only used one id, it should return 1
	// affected row only if the id exists
	return res.RowsAffected()

}
