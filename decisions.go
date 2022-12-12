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
	var total int = 50
	switch total {
	case 90:
		fmt.Println("A")
	case 70:
		fmt.Println("B")
	case 50, 60:
		fmt.Println("C")
	case 40:
		fmt.Println("D")
	default:
		fmt.Println("E")

	}
}
