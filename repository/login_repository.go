package repository

import (
	"database/sql"
)

func UserLoginRepo(d *sql.DB, username string, password string) (int, error) {

	stmt, err := d.Prepare("SELECT count(1) cnt from users where username = :1")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	var result int
	err = stmt.QueryRow(username).Scan(&result)

	if err != nil {
		if err == sql.ErrNoRows {
			return result, err
		}
		panic(err.Error())
	}

	return result, nil
}
