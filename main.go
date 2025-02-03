package main

import (
	"GoPasswords/app/account"
	"GoPasswords/app/files"
	"GoPasswords/app/output"
	"fmt"
	"strings"

	"github.com/fatih/color"
)

var menuVariants = []string{
	"Меню.",
	"1. Создать аккаунт",
	"2. Найти аккаунт по URL",
	"3. Найти аккаунт по логину",
	"4. Удалить аккаунт",
	"5. Выход",
	"Введите команду",
}

var menu = map[int]func(*account.VaultWithDb){
	1: createAccount,
	2: findAccountByURL,
	3: findAccountByLogin,
	4: deleteAccount,
}

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

func findAccountByLogin(vault *account.VaultWithDb) {
	login := promptData("Введите логин")
	accounts, err := vault.FindAccounts(login, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Login, login)
	})
	findResults(accounts, err)
}

func findAccountByURL(vault *account.VaultWithDb) {
	url := promptData("Введите url")
	accounts, err := vault.FindAccounts(url, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Url, str)
	})
	findResults(accounts, err)
}

func findResults(accounts *[]account.Account, err error) {
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

func promptData(prompt ...string) string {
	for index, value := range prompt {
		if index != len(prompt)-1 {
			fmt.Println(value)
		} else {
			fmt.Printf("%v: ", value)
		}
	}
	var res string
	fmt.Scanln(&res)
	return res

}

func getCommandFromUser() int {
	for {
		answer := promptData(menuVariants...)
		if answer == "1" || answer == "2" || answer == "3" || answer == "4" || answer == "5" {
			return int(answer[0] - '1' + 1)
		}
		fmt.Println("Команда неверная. Попробуйте ещё раз.")
		fmt.Println("-------------------")
	}
}

func Menu() {
	vault := account.NewVault(files.NewJsonDb("data.json"))
MenuLoop:
	for {
		cmd := getCommandFromUser()
		menuFunc := menu[cmd]
		if menuFunc == nil {
			fmt.Println("Программа завершена.")
			break MenuLoop
		}
		menuFunc(vault)
		fmt.Println("-------------------")
	}
}
