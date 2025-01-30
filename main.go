package main

import (
	"GoPasswords/app/account"
	"GoPasswords/app/files"
	"GoPasswords/app/output"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("___Менеджер паролей___")
	Menu()
}

func createAccount(vault *account.VaultWithDb) {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintError("Неверный формат URL или Логин.")
		return
	}
	vault.AddAccount(*myAccount)
}

func findAccount(vault *account.VaultWithDb) {
	url := promptData("Введите url")
	accounts, err := vault.FindURL(url)
	if err != nil {
		output.PrintError("Аккаунт не найден.")
		return
	}
	fmt.Println("~~~~~~~~~~~~~~~~~~~")
	for index, account := range *accounts {
		fmt.Println("Аккаунт", index+1)
		account.Output()
		fmt.Println("~~~~~~~~~~~~~~~~~~~")
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	url := promptData("Введите url")
	success := vault.DeleteURL(url)
	if success {
		color.Green("Удалены нужные элементы.")
	} else {
		output.PrintError("Такие элементы не найдены.")
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
	vault := account.NewVault(files.NewJsonDb("data.json"))
MenuLoop:
	for {
		printMenu()
		cmd := getCommandFromUser()
		switch cmd {
		case 1:
			createAccount(vault)
		case 2:
			findAccount(vault)
		case 3:
			deleteAccount(vault)
		case 4:
			fmt.Println("Программа завершена.")
			break MenuLoop
		}
		fmt.Println("-------------------")
	}
}
