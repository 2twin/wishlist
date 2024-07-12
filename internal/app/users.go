package app

import (
	"errors"
	"log"
)

type User struct {
	Username   string
}

var (
	ErrUserAlreadyExists = errors.New("user with the same username is already exists")
	ErrUserNotFound      = errors.New("username not found")
)

func (a *App) AddUser(username string) error {
	for _, user := range a.Users {
		if user.Username == username {
			return ErrUserAlreadyExists
		}
	}

	a.Users = append(a.Users, User{Username: username})
	log.Printf("Added user with username: %s", username)

	return nil
}

func (a *App) RemoveUser(username string) error {
	i := 0
	for i < len(a.Users) {
		if a.Users[i].Username == username {
			a.Users = append(a.Users[:i], a.Users[i+1:]...)
			log.Printf("Removed user with username: %s", username)
			return nil
		}
		i++
	}

	log.Printf("<Error> User with username %s not found", username)
	return ErrUserNotFound
}

func (a *App) GetUsers() []User {
	return a.Users
}