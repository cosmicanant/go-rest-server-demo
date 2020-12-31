package models

import (
	"errors"
	"fmt"
)

// User Struct Type
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

// GetUsers function returns the users
func GetUsers() []*User {
	return users
}

// AddUser function add a new user
func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("can't add a user with existing ID")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

// UpdateUser API
func UpdateUser(u User) (User, error) {
	for i, v := range users {
		if v.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID %d is not found", u.ID)
}

// GetUserByID API
func GetUserByID(ID int) (User, error) {
	for _, v := range users {
		if v.ID == ID {
			return *v, nil
		}
	}
	return User{}, fmt.Errorf("User with ID %d is not found", ID)
}

// RemoveUserByID API
func RemoveUserByID(ID int) error {
	for i, v := range users {
		if v.ID == ID {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User with ID %d is not found", ID)
}
