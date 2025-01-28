package main

import (
	"GoPasswords/app/account"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("___Менеджер паролей___")
	Menu()
}

func createAccount() {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println(err)
		return
	}
	var vault = account.NewVault()
	vault.AddAccount(*myAccount)
}

func findAccount() {
	url := promptData("Введите url")
	vault := account.NewVault()
	accounts, err := vault.FindURL(url)
	if err != nil {
		color.Red("Аккаунт не найден.")
		return
	}
	fmt.Println("~~~~~~~~~~~~~~~~~~~")
	for index, account := range *accounts {
		fmt.Println("Аккаунт", index+1)
		account.Output()
		fmt.Println("~~~~~~~~~~~~~~~~~~~")
	}
}

func deleteAccount() {
	url := promptData("Введите url")
	vault := account.NewVault()
	success := vault.DeleteURL(url)
	if success {
		color.Green("Удалены нужные элементы.")
	} else {
		color.Red("Такие элементы не найдены.")
	}
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}

func printMenu() {
	fmt.Println("Меню.")
	fmt.Println("1. Создать аккаунт")
	fmt.Println("2. Найти аккаунт")
	fmt.Println("3. Удалить аккаунт")
	fmt.Println("4. Выход")
}

func getCommandFromUser() int {
	for {
		answer := promptData("Введите команду")
		if answer == "1" || answer == "2" || answer == "3" || answer == "4" {
			return int(answer[0] - '1' + 1)
		}
		fmt.Println("Команда неверная. Попробуйте ещё раз.")
	}
}

func Menu() {
MenuLoop:
	for {
		printMenu()
		cmd := getCommandFromUser()
		switch cmd {
		case 1:
			createAccount()
		case 2:
			findAccount()
		case 3:
			deleteAccount()
		case 4:
			fmt.Println("Программа завершена.")
			break MenuLoop
		}
		fmt.Println("-------------------")
	}
}
