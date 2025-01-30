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
	login := promptData([]string{"Введите логин"})
	password := promptData([]string{"Введите пароль"})
	url := promptData([]string{"Введите URL"})

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintError("Неверный формат URL или Логин.")
		return
	}
	vault.AddAccount(*myAccount)
}

func findAccount(vault *account.VaultWithDb) {
	url := promptData([]string{"Введите url"})
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
	url := promptData([]string{"Введите url"})
	success := vault.DeleteURL(url)
	if success {
		color.Green("Удалены нужные элементы.")
	} else {
		output.PrintError("Такие элементы не найдены.")
	}
}

func promptData[T any](prompt []T) string {
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

func printMenu() {
	fmt.Println("Меню.")
	fmt.Println("1. Создать аккаунт")
	fmt.Println("2. Найти аккаунт")
	fmt.Println("3. Удалить аккаунт")
	fmt.Println("4. Выход")
}

func getCommandFromUser() int {
	for {
		answer := promptData([]string{
			"Меню.",
			"1. Создать аккаунт",
			"2. Найти аккаунт",
			"3. Удалить аккаунт",
			"4. Выход",
			"Введите команду"})
		if answer == "1" || answer == "2" || answer == "3" || answer == "4" {
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
