package transaction

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	wallet_cli "github.com/hyperyuri/wallet-cli/pkg/wallet"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/shopspring/decimal"
	"log"
)

type TransactionService struct {

}

func NewTransactionService() *TransactionService {
	return &TransactionService{}
}

func (s TransactionService) SendTransaction(mnemonic string, t *wallet_cli.Transaction) (*wallet_cli.Transaction, error)  {
	client, err := ethclient.Dial(t.Currency().Uri())
	if err != nil {
		log.Fatal(err)
	}

	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return nil, err
	}

	gasLimit := uint64(21000)

	amount, err := decimal.NewFromString(t.Amount())
	if err != nil {
		return nil, err
	}

	gas, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	gasPrice, err := decimal.NewFromString(gas.String())
	if err != nil {
		return nil, err
	}

	nonce, err := client.PendingNonceAt(context.Background(), common.HexToAddress(t.Address().Code()))
	if err != nil {
		return nil, err
	}

	toAddress := common.HexToAddress(t.ToAddress())

	weiAmount := amount.Mul(decimal.NewFromInt(1e18))

	var data []byte
	tx := types.NewTransaction(nonce, toAddress, weiAmount.BigInt(), gasLimit, gasPrice.BigInt(), data)

	derivationPath := fmt.Sprintf("m/44'/60'/0'/0/%s", t.Address().Derivation())
	path := hdwallet.MustParseDerivationPath(derivationPath)
	account, err := wallet.Derive(path, true)
	if err != nil {
		return nil, fmt.Errorf("account %s", err.Error())
	}

	sign, err := wallet.SignTx(account, tx, nil)
	if err != nil {
		return nil, err
	}

	err = client.SendTransaction(context.Background(), sign)
	if err != nil {
		return nil, err
	}

	t.SetTxid(sign.Hash().Hex())

	return t, nil
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

