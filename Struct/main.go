package main

import "fmt"

type User struct {
	ID   int
	Name string
	Age  int
}

func main() {
	// Membuat struct secara langsung
	user1 := User{ID: 1, Name: "Muha", Age: 21}

	// Cara lainnya
	var user2 User
	user2.ID = 2
	user2.Name = "Muha"
	user2.Age = 21

	fmt.Println(user1)
	fmt.Println(user2)
}
