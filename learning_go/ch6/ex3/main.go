package main

type Person struct {
	firstName string
	lastName  string
	age       int
}

func MakePerson(firstName string, lastName string, age int) Person {
	return Person{firstName, lastName, age}
}

func main() {
	var p []Person
	for range 10_000_000 {
		p = append(p, MakePerson("first", "last", 10))
	}
}
