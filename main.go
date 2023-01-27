package main

import (
	"fmt"
	"something/pojo"
)

func main() {
	user := pojo.NewUser(pojo.UserWithName("John"), pojo.UserWithAge(18), pojo.UserWithSex("male"))
	user2 := pojo.NewUser(pojo.UserWithName("Mike"), pojo.UserWithAge(16))
	fmt.Printf("%v\n", user)
	fmt.Printf("%v", user2)
}
