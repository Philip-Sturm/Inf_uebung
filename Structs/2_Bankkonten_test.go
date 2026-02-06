package structs

import "fmt"

func ExampleBankAccount() {
	account := BankAccount{
		Owner:   "Max Mustermann",
		Number:  "DE1234567890",
		Balance: 100.0,
	}

	account.Deposit(50.0)
	fmt.Println(account.Balance) // 150.0

	success := account.Withdraw(30.0)
	fmt.Println(success)         // true
	fmt.Println(account.Balance) // 120.0

	success = account.Withdraw(200.0)
	fmt.Println(success)         // false
	fmt.Println(account.Balance) // 120.0

	// Output:
	// 150
	// true
	// 120
	// false
	// 120
}
