package repo

import (
	"database/sql"
	"fmt"
)

func UserLoginRepo(d sql.DB) string {

	// fetch data from db
	rows, err := d.Query("select 1 from dual")
	if err != nil {
		panic(err)
	}

	// iterate through the result set
	for rows.Next() {
		var msg string
		err := rows.Scan(&msg)
		if err != nil {
			panic(err)
		}
		fmt.Println(msg)
	}

	return "OK"

}
