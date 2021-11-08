package balance

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math"
	"math/big"
	"github.com/hyperyuri/wallet-cli/pkg/wallet"
)

type BalanceService struct {
}

func NewBalanceService() *BalanceService {
	return &BalanceService{}
}

func (b *BalanceService) GetBalance(a *wallet.Address) (*wallet.Balance, error) {
	client, err := ethclient.Dial("https://bsc-dataseed.binance.org/")
	if err != nil {
		return nil, err
	}

	_ = client // we'll use this in the upcoming sections

	account := common.HexToAddress(a.Code())

	balance, err := client.BalanceAt(context.Background(), account, nil)

	if err != nil {
		return nil, err
	}

	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)

	if err != nil {
		return nil, err
	}


	return wallet.NewBalance(ethValue.String(), pendingBalance.String()), nil
}