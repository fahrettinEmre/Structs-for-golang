package main

import (
	"fmt"
	"strconv"
)

type Human struct {
	firstName string
	lastName  string
	age       int
}

func NemHuman() *Human { // yeni nesne tanımlama
	z := new(Human)
	return z
}

func NemHumanwithParams(firstname, lastname string, age int) *Human { //parametrelerle nesne tanımlama
	h := new(Human)
	h.firstName = firstname
	h.lastName = lastname
	h.age = age
	return h
}

func main() {
	//versiyon 1
	x := Human{firstName: "Fahrettin"}
	fmt.Println(x.firstName)

	//versiyon 2
	y := new(Human)
	y.firstName = "Fahrettin"
	y.lastName = "Emre"
	y.age = 29
	fmt.Println(y)
	fmt.Println(y.age)

	//versiyon 3
	h1 := NemHuman()
	h1.age = 29
	h1.firstName = "Emre"
	fmt.Println(h1, h1.firstName)

	//versiyon
	h := NemHumanwithParams("Fahrettin", "Emre", 29)
	fmt.Println(h)
	fmt.Println(h.age, h.firstName, h.lastName)

	data := h.firstName + h.lastName + strconv.Itoa(h.age)

	fmt.Println(data)
}
