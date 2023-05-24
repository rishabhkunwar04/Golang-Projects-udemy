package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {

	findAllsql := "select customer_id,name,city,zipcode,date_of_birth,status from customers"

	rows, err := d.client.Query(findAllsql)
	if err != nil {
		log.Println("error while querying" + err.Error())
		return nil, err
	}
	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
		if err != nil {
			log.Println("error while Scanning" + err.Error())
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, nil
}
func (d CustomerRepositoryDb) ById(id string) (*Customer, error) {
	customerSql := "select customer_id,name,city,zipcode,date_of_birth,status from customers where customer_id=?"
	row := d.client.QueryRow(customerSql, id)
	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
	if err != nil {
		log.Println("error while Scanning customer" + err.Error())
		return nil, err
	}
	return &c, nil
}
func NewCustomerRepositioryDb() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "root:root@tcp/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client}
}
