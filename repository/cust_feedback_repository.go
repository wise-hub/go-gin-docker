package repository

import (
	"database/sql"
	"ginws/model"
	"ginws/model_in"
)

func InsertCustFeedback(d *sql.DB, cf *model_in.InCustomerFeedback, username string) error {

	insQuery, err := d.Prepare(`insert into cust_feedback 
	(feedback_id, ins_dt, customer_no, user_name, feedback, is_visible)
	values (SEQ_FEEDBACK.nextval, sysdate, :1, :2, :3, 'Y')`)

	if err != nil {
		return err
	}
	defer func() {
		_ = insQuery.Close()
	}()

	_, err = insQuery.Exec(cf.CustomerID, username, cf.Feedback)
	if err != nil {
		return err
	}

	return nil
}

func ReadCustFeedback(d *sql.DB, customer_id string) ([]model.CustomerFeedback, error) {

	stmt, err := d.Prepare(`SELECT ins_dt, customer_no, user_name, feedback 
	from cust_feedback where is_visible = 'Y' and customer_no = :1`)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(customer_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var feedbacks []model.CustomerFeedback
	for rows.Next() {
		var feedback model.CustomerFeedback
		err := rows.Scan(&feedback.InsertDate,
			&feedback.CustomerID, &feedback.UserName, &feedback.Feedback)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, err
			}
			panic(err.Error())
		}
		feedbacks = append(feedbacks, feedback)
	}

	return feedbacks, nil

}
