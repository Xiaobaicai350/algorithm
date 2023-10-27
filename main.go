package main

import "fmt"

type User struct {
	Name string
}

func main() {

	var user1 User
	user2 := new(User)
	fmt.Println(user1)
	fmt.Println(user2)

}
