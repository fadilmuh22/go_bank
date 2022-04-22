package db

import (
	"context"
)

type Store struct {
}

type Account struct{}

func (store *Store) GetAccount(ctx context.Context, id int64) (Account, error) {
	return Account{}, nil
}
