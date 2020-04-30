package db

import (
	"database/sql"
	"fmt"
)

func ReadSelect(rows *sql.Rows) []interface{} {
	users := make([]interface{}, 0)

	for rows.Next() {
		var user interface{}
		if err := rows.Scan(&user); err != nil {
			fmt.Println("ERROR", err)
			return nil
		}
		users = append(users, user)
	}

	return users
}
