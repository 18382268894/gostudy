package main

import (
	"fmt"

	"day6/user"
)

func main() {
	name := "cxz"
	var age uint8 = 23
	photoUrl := "www.cxz.com/jpg.index"
	var user = user.NewUser(name, age, photoUrl)
	fmt.Printf("%#v", user)
}
