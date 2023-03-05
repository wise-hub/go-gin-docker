package repository

import (
	"database/sql"

	"ginws/model"
)

func GetAccountsRepo(d *sql.DB, customer_id string) ([]model.Account, error) {

	stmt, err := d.Prepare("select iban, balance from accounts where customer_id = :1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(customer_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accounts []model.Account
	for rows.Next() {
		var account model.Account
		err := rows.Scan(&account.IBAN, &account.Balance)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return accounts, nil
}
