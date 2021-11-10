package transaction

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
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
	client, err := ethclient.Dial(t.Currency().Uri())
	if err != nil {
		return "", err
	}

	gwei, err := client.SuggestGasPrice(context.Background())

	if err != nil {
		return "", err
	}

	val := decimal.NewFromBigInt(gwei, 0).Div(decimal.NewFromInt32(1e9))
	fee := val.Mul(decimal.NewFromFloat32(0.000000001))
	gasLimit := fee.Mul(decimal.NewFromInt32(21000)).Truncate(8)

	return gasLimit.String(), nil
}

