package transaction

import (
	wallet_cli "github.com/hyperyuri/wallet-cli/pkg/wallet"
)

type TransactionService struct {
	
}

func NewTransactionService() *TransactionService {
	return &TransactionService{}
}

func (s TransactionService) SendTransaction(t *wallet_cli.Transaction) error {
	panic("implement me")
}

func (s TransactionService) GetFee(t *wallet_cli.Transaction) (string, error) {
	panic("implement me")
}

