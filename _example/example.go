package main

import (
	"fmt"
	"govalidator"
)

// User contains user information
type User struct {
	FirstName      string
	LastName       string
	Age            uint8
	Email          string
	FavouriteColor string
	Myage			int		   `min:"1" max:"10"`
	Address			Address
}

// Address houses a users address information
type Address struct {
	Street string 	`format:"email" type:"text"`
	City   string	`format:"alphanumeric"`
	Planet string 	`startswith:"u" endswith:"s"`
	Phone  string 	`min:"1" max:"10" type:"number"`
}
type User1 struct {
	FirstName      string		`value_of:"jhonny,john" field:"required"`
	LastName       int
	Age            float64		`min:"4" max:"100" field:"required"`
	Email          string
	FavouriteColor []string		`field:"required"`
	Myage			int			`min:"1" max:"10"`
	Address			[]Address
	Testing 		float64		`field:"required"`
}

func main() {
	fmt.Println("RETURN VALUE OF SIMPLE STRUCT")
	simpleStruct()
	fmt.Println("RETURN VALUE OF SIMPLE NESTEDSTRUCT")
	nestedSturct()
	fmt.Println("RETURN VALUE OF ARRAY STRUCT")
	arrayStruct()
	fmt.Println("RETURN VALUE OF NESTED ARRAY STRUCT")
	nestedArrayStruct()
}

func simpleStruct(){
	d := Address{City:"BANDUNG TEST 123",Phone:"9663",Street:"aas@gmail.com",Planet:"uranus"}
	err := validator.Validate(d)
	fmt.Println(err)
}

func nestedSturct(){
	d := Address{City:"BANDUNG",Phone:"88766524928347384",Street:"asd@gmail.com"}
	a := User{FirstName:"udin",Age:0,Myage:3,Address:d}
	err := validator.Validate(a)
	fmt.Println(err)
}

func arrayStruct(){
	var d []Address
	d = append(d,Address{City:"ASD",Phone:"asd",Planet:"uranus",Street:"asd@gmail.com"})
	d = append(d,Address{City:"ASD",Phone:"asd",Planet:"mars",Street:"asd@agmail.com"})
	err := validator.Validate(d)
	fmt.Println(err)
}
func nestedArrayStruct(){
	var d []Address
	d = append(d,Address{City:"ASD",Phone:"asd",Planet:"uranus",Street:"asd@gmail.com"})
	d = append(d,Address{City:"ASD",Phone:"asd",Planet:"uranus",Street:"asd@agmail.com"})
	a := User1{FirstName:"john",Age:100,Myage:8,FavouriteColor:[]string{"asd"},Address:d}
	err := validator.Validate(a)
	fmt.Println(err)
}
