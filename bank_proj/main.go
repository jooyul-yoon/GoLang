package main

import (
	"fmt"

	"github.com/jooyul-yoon/learngo/bank_proj/accounts"
)

func main() {
	account := accounts.NewAccount("juyeol")
	fmt.Println(account.Balance())
	account.Deposit(1000)
	fmt.Println(account.Balance())
	err := account.Withdraw(2000)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account)
	account.ChangeOwner("Jungkee")
	fmt.Println(account)
}
