package handlers

import (
	db "haste/infra/driven/database/sqlc"
)

type AccountPort interface {
	GetAllAccounts() ([]AccountResponse, error)
	GetAccountById() (AccountResponse, error)
	CreateAccount(accountForm AccountForm) (AccountResponse, error)
	TransferAmount(transerForm TransferForm) (TransferResponse, error)
}

type AccountForm struct {
	Owner    string `json:"owner"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency"`
}

type AccountResponse struct {
	ID       int64  `json:"id"`
	Owner    string `json:"owner"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency"`
}

type TransferForm struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferResponse struct {
	FromAccount db.Account
	ToAccount   db.Account
	Amount      int64
}
