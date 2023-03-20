package repository

import (
	"database/sql"
	"ginws/model"
)

func CustomerRepo(d *sql.DB, customer_id string) (model.Customer, error) {

	stmt, err := d.Prepare("SELECT customer_id, name, egn, address from customers where customer_id = :1")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	var customer model.Customer
	err = stmt.QueryRow(customer_id).Scan(&customer.CustomerID, &customer.Name, &customer.EGN, &customer.Address)
	if err != nil {
		if err == sql.ErrNoRows {
			return customer, err
		}
		panic(err.Error())
	}

	return customer, nil

}
