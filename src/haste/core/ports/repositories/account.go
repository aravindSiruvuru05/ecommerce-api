package repositories

import (
	"context"
	db "haste/infra/driven/database/sqlc"
)

type Account interface {
	createAccountInDB(ctx context.Context, arg db.CreateAccountParams) (db.Account, error)
	GetAccountInDB(ctx context.Context, arg db.GetAccountsParams) (db.Account, error)
	TransferAmount(ctx context.Context, arg db.CreateTransferParams) (db.TransferTxResult, error)
}
