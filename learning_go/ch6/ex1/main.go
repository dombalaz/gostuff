package main

type Person struct {
	firstName string
	lastName  string
	age       int
}

func MakePerson(firstName string, lastName string, age int) Person {
	return Person{firstName, lastName, age}
}

func MakePersonPointer(firstName string, lastName string, age int) *Person {
	return &Person{firstName, lastName, age}
}

func main() {
	MakePerson("peter", "spalnicek", 10)
	MakePerson("anna", "karenina", 20)
}
