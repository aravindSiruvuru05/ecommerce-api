package account

import (
	"haste/core/components"
	"haste/core/ports/handlers"
)

type AccountComponent struct {
	components.BaseComponent
}

func (c *AccountComponent) GetAllAccounts() ([]handlers.AccountResponse, error) {
	var accounts []handlers.AccountResponse
	for i := 0; i < 5; i++ {
		accounts = append(accounts, handlers.AccountResponse{})

	}
	return accounts, nil
}

func (c *AccountComponent) GetAccountById() (handlers.AccountResponse, error) {
	var accounts handlers.AccountResponse

	return accounts, nil
}

func (c *AccountComponent) CreateAccount(accountForm handlers.AccountForm) (handlers.AccountResponse, error) {
	return handlers.AccountResponse{}, nil
}

func (c *AccountComponent) TransferAmount(tansferForm handlers.TransferForm) (handlers.TransferResponse, error) {
	return handlers.TransferResponse{}, nil
}

func init() {
	components.ComponentMap["Account"] = func(bc *components.BaseComponent) interface{} {
		c := &AccountComponent{
			BaseComponent: *bc,
		}
		return handlers.AccountPort(c)
	}
}
