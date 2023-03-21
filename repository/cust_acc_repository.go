package repository

import (
	"database/sql"
	"ginws/model"
)

func CustAccRepo(d *sql.DB, customer_id string) (model.CustAccount, error) {

	stmt, err := d.Prepare("SELECT customer_id, name, egn, address from customers where customer_id = :1")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	var custAccount model.CustAccount

	custAccount.CustomerData, err = CustomerRepo(d, customer_id)

	if err != nil {
		if err == sql.ErrNoRows {
			return custAccount, err
		}
		panic(err.Error())
	}

	custAccount.AccountData, err = AccountsRepo(d, customer_id)

	if err != nil {
		if err == sql.ErrNoRows {
			return custAccount, err
		}
		panic(err.Error())
	}

	return custAccount, nil

}
