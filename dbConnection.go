package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
)

var host = os.Getenv("DBHOST")
var port = os.Getenv("DBPORT")
var user = os.Getenv("DBUSER")
var password = os.Getenv("DBPASS")
var dbname = os.Getenv("DBNAME")

const insertIntoCustomers = `INSERT INTO "Customers"."Customers"("firstName","lastName","userName","email","password") VALUES ($1,$2,$3,$4,$5) RETURNING id;`
const selectCustomer = `SELECT * FROM "Customers"."Customers" WHERE id=$1;`

var db *sql.DB

func connectToDatabase() {
	i_port, _ := strconv.Atoi(port)
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		user, password, host, i_port, dbname)

	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
}

func PostCustomerDB(requestedCustomer *Customer) (customerId int64) {

	connectToDatabase()
	err := db.QueryRow(insertIntoCustomers, &requestedCustomer.FirstName, &requestedCustomer.LastName, &requestedCustomer.Username, &requestedCustomer.Email, &requestedCustomer.Password).Scan(&customerId)
	if err != nil {
		log.Fatal("Failed to execute query: ", err)
	}

	defer db.Close()

	return
}

func GetCustomerDB(customerId int) (requestedCustomer *Customer) {

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
	requestedCustomer = &Customer{
		Id:        id,
		FirstName: firstName,
		LastName:  lastName,
		Username:  username,
		Email:     email,
	}

	return
}
