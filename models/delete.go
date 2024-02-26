package models

import "github.com/vbmelo/todo-golang-api/db"

func Delete(id int64) (int64, error) {
	conn, err := db.OpenConnection() // Stablishes the Connection
	if err != nil {                  // If there's an error, it returns 0 and the error
		return 0, err
	}
	defer conn.Close() // Defer the closing of the connection

	res, err := conn.Exec(`DELETE FROM todos WHERE id=$1`, id)
	if err != nil {
		return 0, err
	}

	// We only expect one row to be affected
	return res.RowsAffected()

}
