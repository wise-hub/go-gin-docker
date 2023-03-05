package repository

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"fmt"
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

func ValidateTokenAtDb(d *sql.DB, token string) bool {
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

func InsertNewToken(d *sql.DB, username string, ipAddr string) (string, error) {

	tokenBytes := make([]byte, 32)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		panic(err)
	}
	token := hex.EncodeToString(tokenBytes)

	// set token, expiry and last login
	/*
		table structure
		id, username, role, user_status, token, session_expiry_dt, last_login_dt
	*/
	updStmt, err := d.Prepare(`update users 
	set last_login_dt = sysdate, 
	session_expiry_dt = sysdate+1/24, 
	last_login_ip = :1,
	token = :2
	where username = :3 `)

	if err != nil {
		panic(err)
	}
	defer func() {
		_ = updStmt.Close()
	}()

	_, err = updStmt.Exec(ipAddr, token, username)
	if err != nil {
		panic(err)
	}
	fmt.Println(token)

	return token, nil
}
