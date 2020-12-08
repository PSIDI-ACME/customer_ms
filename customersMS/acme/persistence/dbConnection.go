package persistence

import (
	"database/sql"
	"fmt"
	"log"

	"../models"
)

const (
	host     = "ec2-99-81-238-134.eu-west-1.compute.amazonaws.com"
	port     = 5432
	user     = "uktrtdoklswema"
	password = "2c174b8522aa6c46fbfe4668014e383c1fbce0f706e7cd2d3e3b7dabac6d653e"
	dbname   = "d56vth083636vl"
)

const insertIntoCustomers = `INSERT INTO "Customers"."Customers"("firstName","lastName","userName","email","password") VALUES ($1,$2,$3,$4,$5) RETURNING id;`
const selectCustomer = `SELECT * FROM "Customers"."Customers" WHERE id=$1;`

var db *sql.DB

func connectToDatabase() {
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		user, password, host, port, dbname)

	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
}

func PostCustomer(requestedCustomer *models.Customer) (customerId int64) {

	connectToDatabase()
	err := db.QueryRow(insertIntoCustomers, &requestedCustomer.FirstName, &requestedCustomer.LastName, &requestedCustomer.Username, &requestedCustomer.Email, &requestedCustomer.Password).Scan(&customerId)
	if err != nil {
		log.Fatal("Failed to execute query: ", err)
	}

	defer db.Close()

	return
}

func GetCustomer(customerId int) (requestedCustomer *models.Customer) {

	connectToDatabase()

	var id int64
	var firstName string
	var lastName string
	var username string
	var email string
	var password string

	defer db.Close()

	row := db.QueryRow(selectCustomer, customerId)
	switch err := row.Scan(&id, &firstName, &lastName, &username, &email, &password); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(id, firstName)
	default:
		panic(err)
	}
	requestedCustomer = &models.Customer{
		Id:        id,
		FirstName: firstName,
		LastName:  lastName,
		Username:  username,
		Email:     email,
	}

	return
}
