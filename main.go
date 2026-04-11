package main

import (
	"fmt"

	"example.com/note/user"
)

func main() {
	userDetails, err := user.GetUserDetails()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("User Details: %+v\n", userDetails)
	// user.GetAdminDetails()
}
