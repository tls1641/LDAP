package users

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email    string
	Password string
}

type authUser struct {
	email        string
	passwordHash string
}

var DefaultUserService userService

type userService struct {
}

func (userService) VerifyUser(user User) bool {
	authUser, ok := authUserDB[user.Email]
	if !ok {
		return false
	}
	err := bcrypt.CompareHashAndPassword(
		[]byte(authUser.passwordHash),
		[]byte(user.Password))
	return err == nil

}

var authUserDB = map[string]authUser{}

func (userService) CreateUser(newUser User) error {
	_, ok := authUserDB[newUser.Email]
	if ok {
		fmt.Println("user already exists")
		return errors.New("user already exists")

	}
	passwordHash, err := getPasswordHash(newUser.Password)
	if err != nil {
		fmt.Println("getPasswordHash")
		return err
	}
	newAuthUser := authUser{
		email:        newUser.Email,
		passwordHash: passwordHash,
	}
	authUserDB[newAuthUser.email] = newAuthUser
	return nil
}

func getPasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	return string(hash), err
}
