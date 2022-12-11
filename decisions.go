package main

import "fmt"

func main() {
	var age int = 17
	if age >= 20 {
		fmt.Println("You have the required age limit")
	} else if age < 14 {
		fmt.Println("You are still a small kid")
	} else {
		fmt.Println("You are a teenager")
	}
}
