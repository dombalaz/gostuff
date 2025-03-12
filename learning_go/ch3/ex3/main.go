package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	id        int
}

func main() {
	e1 := Employee{"Jakub", "Novák", 10}
	e2 := Employee{firstName: "Peter", lastName: "Hrivnák", id: 20}
	var e3 Employee
	e3.firstName = "Martina"
	e3.lastName = "Kováčová"
	e3.id = 30

	fmt.Println(e1)
	fmt.Println(e2)
	fmt.Println(e3)
}
