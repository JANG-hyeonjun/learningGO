package main

import "fmt"

type person struct {
	name    string
	age     int
	favfood []string
}

func main() {
	favfood := []string{"kimchi", "ramen"}
	nico := person{name:"nico",age:18,favfood:favfood}
	fmt.Println(nico)

}