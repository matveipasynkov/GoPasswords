package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (acc Account) Output() {
	fmt.Println(acc.Login)
	fmt.Println(acc.Password)
	fmt.Println(acc.Url)
}

func (acc *Account) generatePassword(n int) {
	password := make([]rune, n)
	for i := range password {
		password[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.Password = string(password)
}

func NewAccount(login, password, urlString string) (*Account, error) {
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	acc := &Account{
		Login:     login,
		Password:  password,
		Url:       urlString,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}

	if password == "" {
		acc.generatePassword(12)
	}

	return acc, nil
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%&*0123456789")
