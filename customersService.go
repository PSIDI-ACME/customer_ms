package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"strings"
)

func PostCustomer(customer *Customer) (id int64, code int, err error) {

	input := strings.NewReader(customer.Password)

	hash := sha256.New()
	if _, err = io.Copy(hash, input); err != nil {
		code = 500
		return
	}
	sum := hash.Sum(nil)
	pass := fmt.Sprintf("%x", sum)

	log.Println(customer.Password)
	fmt.Printf("%s\n", pass)

	customer.Password = pass

	id = PostCustomerDB(customer)
	fmt.Println("New record ID is:", id)
	return
}

/*
func PutCustomer(customerId int, review string) (code int, err error) {
	customer := GetCustomerDB(customerId)
	reviews := customer.Reviews
	reviews = reviews[:len(reviews)-1]
	if reviews == "[" {
		reviews = reviews +review+"]"
	}else{
	reviews = reviews + ","+review+"]"
	}
	fmt.Println("reviews", reviews)
	UpdateCustomerReviews(customerId, reviews)
	return
}
*/
func GetCustomerService(id int) (customer *Customer) {
	customer = GetCustomerDB(id)
	fmt.Println("Fetched Customer is:", customer)
	return customer
}
