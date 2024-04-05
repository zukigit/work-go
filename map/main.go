package main

import "fmt"

type student struct {
	name string
	age  int
}

func unregister(dataStore map[string]student, username string) {
	delete(dataStore, username)
}

func register(dataStore map[string]student, student student) {
	dataStore[student.name] = student
}

func main() {
	fmt.Println("maps test")
	data_store := make(map[string]student)
	register(data_store, student{
		name: "zuki",
		age:  23,
	})
	fmt.Println(data_store)

	unregister(data_store, "zuki")
	fmt.Println(data_store)
}
