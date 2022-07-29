package main

import (
	"fmt"

	user "github.com/tse-lao/ether-user/wallet"
)

func (a *App) GetAccountExist() (bool, error) {
	status, result := user.DirExists()

	//check if the dir exists.

	if status {
		return status, result
	}

	return status, result

}

func (a *App) CreateAccount(password string) user.Notification {
	//create a account by writing

	result := user.CreateNewAccount(password)

	fmt.Println(result)

	return result
}
