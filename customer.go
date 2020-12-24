package main

type Customer struct {
	Id int64 `json:"id"`

	Username string `json:"username"`

	Password string `json:"password"`

	FirstName string `json:"firstName"`

	LastName string `json:"lastName"`

	Email string `json:"email"`

	//	Reviews string`json:"reviews,omitempty"`
}
