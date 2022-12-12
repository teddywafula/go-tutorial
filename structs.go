package main

import (
	"fmt"
)

type person struct {
	name    string
	age     int
	address string
}

func main() {
	p := person{name: "John", age: 23, address: "Nairobi"}
	fmt.Printf("Person details: \nName: %s\nAge: %d \nAddress: %s\n", p.name, p.age, p.address)
	fmt.Println(p.age)
	fmt.Println(&p)
	s := &p
	fmt.Println(s)
	fmt.Println(s.address)
	s.address = "Nairobi, Kenya"
	fmt.Println(s)
	fmt.Println(p.address)

}
