package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

type account struct {
	login    string
	password string
	url      string
}

func (acc account) OutputPassword() {
	fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *account) generatePassword(n int) {
	password := make([]rune, n)
	for i := range password {
		password[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(password)
}

func NewAccount(login, password, urlString string) (*account, error) {
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		fmt.Println("Неверный формат URL.")
		return nil, errors.New("INVALID_URL")
	}

	acc := &account{
		login:    login,
		password: password,
		url:      urlString,
	}

	if login == "" {
		fmt.Println("Неверный формат логина.")
		return nil, errors.New("INVALID_LOGIN")
	}

	if password == "" {
		acc.generatePassword(12)
	}

	return acc, nil
}

type accountWithTimeStamp struct {
	createdAt time.Time
	updatedAt time.Time
	account
}

func NewAccountWithTimeStamp(login, password, urlString string) (*accountWithTimeStamp, error) {
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		fmt.Println("Неверный формат URL.")
		return nil, errors.New("INVALID_URL")
	}

	acc := &accountWithTimeStamp{
		account: account{
			login: login,
			password: password,
			url: urlString,
		},
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}

	if login == "" {
		fmt.Println("Неверный формат логина.")
		return nil, errors.New("INVALID_LOGIN")
	}

	if password == "" {
		acc.generatePassword(12)
	}

	return acc, nil
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%&*0123456789")