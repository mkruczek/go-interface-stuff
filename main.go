package main

import (
	"fmt"
	"interface-stuff/user"
	"interface-stuff/user/repository"
)

func main() {

	ud := repository.UserDao
	u := user.User{Login: "joe"}
	fmt.Printf("1) u = %v\n", u)

	fmt.Printf("db : %d\n", ud.Amount())
	nu, _ := ud.Save(&u)
	fmt.Printf("2) u = %v\n", nu)
	fmt.Printf("db : %d\n", ud.Amount())

	fmt.Printf("3) u = %v\n", u)
	nnu, _ := ud.User(nu.ID)
	fmt.Printf("3) u = %v\n", nnu)

}
