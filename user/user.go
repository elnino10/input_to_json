package user

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"time"
)

type str string

func (s str) customStringStorage() string {
	strVal := fmt.Sprintf("%s_rhonn", s)

	return strVal
}

type user struct {
	firstName string
	lastName  string
	email     string
	role      str
	createdAt time.Time
}

type adminUser struct {
	adminId  string
	password string
	user     user
}

func (u *user) New() user {
	return user{
		firstName: u.firstName,
		lastName:  u.lastName,
		email:     u.email,
		role:      str(str(u.role).customStringStorage()),
		createdAt: time.Now(),
	}
}

func GetUserDetails() (user, error) {

	firstName, err := getUserInput("Enter firstName: ")
	lastName, err2 := getUserInput("Enter lastName: ")
	email, err3 := getUserInput("Enter email: ")
	role, err4 := getUserInput("Enter role: ")

	if err != nil || err2 != nil || err3 != nil || err4 != nil {
		fmt.Println("Error:", err)
		return user{}, errors.New("firstName, lastName, email are required")
	}

	return user{
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		role:      str(str(role).customStringStorage()),
		createdAt: time.Now(),
	}, nil
}

func GetAdminDetails() (adminUser, error) {
	adminId := fmt.Sprintf("admin_%d", rand.IntN(1000))
	password := fmt.Sprintf("password_%d", rand.IntN(1000))
	user, err := GetUserDetails()
	if err != nil {
		return adminUser{}, err
	}

	return adminUser{
		adminId:  adminId,
		password: password,
		user:     user,
	}, nil
}

func getUserInput(prompt string) (string, error) {
	fmt.Print(prompt)
	var value string
	fmt.Scanln(&value)

	return value, nil
}
