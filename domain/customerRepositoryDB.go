package domain

import (
	"banking/errs"
	"banking/logger"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDB struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDB) FindAll(status string) ([]Customer, *errs.AppError) {

	customers := make([]Customer, 0)
	var err error

	if status == "" {
		findAllsql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = d.client.Select(&customers, findAllsql)

	} else {
		findAllsql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		err = d.client.Select(&customers, findAllsql, status)
	}

	if err != nil {
		logger.Error("Error while queryng customer table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
  
	return customers, nil
}

func (d CustomerRepositoryDB) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id =?"

	var c Customer
	err := d.client.Get(&c, customerSql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			logger.Error("Error while scunning customer" + err.Error())
			return nil, errs.NewUnexpectedError("unexpected db error")
		}
	}
	return &c, nil
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	client, err := sqlx.Open("mysql", "root:my-secret-pw@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDB{client}
}
