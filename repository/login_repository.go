package repository

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
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

func InsertNewToken(d *sql.DB, username string, token string, ipAddr string) (bool, error) {

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

	return true, nil
}

func UpdateTokenExpiry(d *sql.DB, username string) error {

	updStmt, err := d.Prepare(`update users 
	set session_expiry_dt = sysdate+1/24
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

func InsertNewTokenCreateToken(d *sql.DB, username string, role string, ipAddr string) (string, error) {

	// tokenBytes := make([]byte, 32)
	// _, err := rand.Read(tokenBytes)
	// if err != nil {
	// 	panic(err)
	// }
	// token := hex.EncodeToString(tokenBytes)

	b := make([]byte, 48) // 48 bytes = 64 chars
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	token := base64.URLEncoding.EncodeToString(b)[:64] // strip padding

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
	fmt.Println(len(token))

	return token, nil
}
