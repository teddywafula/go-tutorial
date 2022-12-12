package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	fmt.Println(m)
	v1 := m["a"]
	fmt.Println(v1)
	delete(m, "a")
	fmt.Println(m)
	a, prs := m["b"]
	fmt.Println("prs:", prs)
	fmt.Println("prs:", a)

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		fmt.Println(i, num)
	}
	for i, c := range "abcdefghijkl" {
		fmt.Println(i, c)
	}
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for _, k := range kvs {
		fmt.Println("key:", k)
	}

}
