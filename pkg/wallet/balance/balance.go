package balance

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hyperyuri/wallet-cli/pkg/wallet"
	token "github.com/hyperyuri/wallet-cli/utils/token"
	"math"
	"math/big"
)

type BalanceService struct {
}

func NewBalanceService() *BalanceService {
	return &BalanceService{}
}

func (b *BalanceService) GetBalance(a *wallet.Address, c *wallet.Currency) (*wallet.Balance, error) {
	if c.TokenType() == "erc20" {
		return b.getContractBalance(a, c)
	}

	client, err := ethclient.Dial(c.Uri())
	if err != nil {
		return nil, err
	}

	account := common.HexToAddress(a.Code())

	wei, err := client.BalanceAt(context.Background(), account, nil)

	if err != nil {
		return nil, err
	}

	fbalance := new(big.Float)
	fbalance.SetString(wei.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)

	if err != nil {
		return nil, err
	}


	return wallet.NewBalance(ethValue.String(), pendingBalance.String()), nil
}

func (b *BalanceService) getContractBalance(a *wallet.Address, c *wallet.Currency) (*wallet.Balance, error) {
	client, err := ethclient.Dial(c.Uri())

	if err != nil {
		return nil, err
	}

	tokenAddress := common.HexToAddress(c.ContractAddress())
	instance, err := token.NewToken(tokenAddress, client)
	if err != nil {
		return nil, err
	}

	address := common.HexToAddress(a.Code())

	wei, err := instance.BalanceOf(&bind.CallOpts{}, address)

	if err != nil {
		return nil, err
	}

	fbalance := new(big.Float)
	fbalance.SetString(wei.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	return wallet.NewBalance(ethValue.String(), "0"), nil
}