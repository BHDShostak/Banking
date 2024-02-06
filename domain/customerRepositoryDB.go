package domain

import (
	"banking/errs"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDB struct {
	client *sql.DB
}

func (d CustomerRepositoryDB) FindAll() ([]Customer, error) {
	findAllsql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
	rows, err := d.client.Query(findAllsql)
	if err != nil {
		log.Println("Error while queryng customer table " + err.Error())
		return nil, err
	}
	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
		if err != nil {
			log.Println("Error while scunning customers" + err.Error())
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func (d CustomerRepositoryDB) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id =?"
	row := d.client.QueryRow(customerSql, id)
	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			log.Println("Error while scunning customer" + err.Error())
			return nil, errs.NewUnexpectedError("unexpected db error")
		}
	}
	return &c, nil
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	client, err := sql.Open("mysql", "root:my-secret-pw@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDB{client}
}
