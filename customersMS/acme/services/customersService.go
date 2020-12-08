package services

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"strings"

	"../models"
	"../persistence"
)

func PostCustomer(customer *models.Customer) (id int64) {

	input := strings.NewReader(customer.Password)

	hash := sha256.New()
	if _, err := io.Copy(hash, input); err != nil {
		log.Fatal(err)
	}
	sum := hash.Sum(nil)
	pass := fmt.Sprintf("%x", sum)

	log.Println(customer.Password)
	fmt.Printf("%s\n", pass)

	customer.Password = pass

	id = persistence.PostCustomer(customer)
	fmt.Println("New record ID is:", id)
	return id
}

func GetCustomer(id int) (customer *models.Customer) {
	customer = persistence.GetCustomer(id)
	fmt.Println("Fetched Customer is:", customer)
	return customer
}
