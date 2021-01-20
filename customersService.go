package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/nvellon/hal"
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

type Review struct {
}

func (c Customer) GetMap() hal.Entry {
	return hal.Entry{
		"firstName": c.FirstName,
		"lastName":  c.LastName,
		"email":     c.Email,
		"username":  c.Username,
	}
}

func (r Review) GetMap() hal.Entry {
	return hal.Entry{}
}

func GetCustomerService(id int, req *http.Request) (hypermedia []byte) {
	customer := GetCustomerDB(id)

	r := hal.NewResource(customer, "https://"+req.Host+req.URL.Path)

	review := hal.NewResource(Review{}, "http://reviews-psidi.herokuapp.com/reviews?customerId="+strconv.Itoa(id))

	r.Embed("review", review)
	var err error
	hypermedia, err = json.MarshalIndent(r, "", "  ")
	if err != nil {
		fmt.Printf("%s", err)
	}

	fmt.Printf("%s", hypermedia)

	fmt.Println("Fetched Customer is:", customer)
	return hypermedia
}
