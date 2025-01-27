package main

import (
	"GoPasswords/app/account"
	"GoPasswords/app/files"
	"fmt"
)

func main() {
	files.ReadFile("dasdasd.txt")
	files.WriteFile("dasdasdasds.txt")
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	myAccount, err := account.NewAccountWithTimeStamp(login, password, url)
	if err != nil {
		return
	}
	// myAccount.generatePassword(12)
	myAccount.OutputPassword()
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}
