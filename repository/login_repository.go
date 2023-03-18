package repository

import (
	"database/sql"
	"fmt"
	"ginws/config"
)

func ValidateUserAtDb(d *sql.DB, username string) bool {
	// check if customer exists and is active
	stmt, err := d.Prepare("SELECT count(1) cnt from users where username = :1 and user_status = 'E'")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer stmt.Close()

	var result int
	err = stmt.QueryRow(username).Scan(&result)

	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	if result == 0 {
		return false
	}

	return true
}

func GetRoleFromUser(d *sql.DB, username string) (string, error) {
	// check if customer exists and is active
	stmt, err := d.Prepare("SELECT user_role from users where username = :1 and user_status = 'E'")
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	defer stmt.Close()

	var result string
	err = stmt.QueryRow(username).Scan(&result)

	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	if len(result) == 0 {
		return "", err
	}

	return result, nil
}

func GetUserFromToken(d *sql.DB, token string) bool {
	// check if token exists and user owner is active
	stmt, err := d.Prepare("SELECT username cnt from users where token = :1 and user_status = 'E'")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	var result int
	err = stmt.QueryRow(token).Scan(&result)

	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	if result == 0 {
		return false
	}

	return true
}

func ValidateTokenOnline(d *sql.DB, token string) bool {
	// check if token exists and user owner is active
	stmt, err := d.Prepare("SELECT count(1) cnt from users where token = :1 and user_status = 'E'")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	var result int
	err = stmt.QueryRow(token).Scan(&result)

	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	if result == 0 {
		return false
	}

	return true
}

func InsertNewToken(d *config.Dependencies, username string, token string, ipAddr string) (bool, error) {

	updStmt, err := d.Db.Prepare(`update users 
	set last_login_dt = sysdate, 
	session_expiry_dt = sysdate + NUMTODSINTERVAL(:1, 'MINUTE'),
	last_login_ip = :2,
	token = :3
	where username = :4 `)

	if err != nil {
		panic(err)
	}
	defer func() {
		_ = updStmt.Close()
	}()

	_, err = updStmt.Exec(d.Cfg.SessionLifetimeMins, ipAddr, token, username)
	if err != nil {
		panic(err)
	}

	return true, nil
}

func UpdateTokenExpiry(d *config.Dependencies, username string) error {

	updStmt, err := d.Db.Prepare(`update users 
	set session_expiry_dt = sysdate + interval '1' hour
	where username = :1`)

	if err != nil {
		panic(err)
	}
	defer func() {
		_ = updStmt.Close()
	}()

	_, err = updStmt.Exec(username)
	if err != nil {
		panic(err)
	}

	return nil
}
