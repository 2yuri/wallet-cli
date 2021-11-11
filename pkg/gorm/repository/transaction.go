package repository

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hyperyuri/wallet-cli/pkg/gorm"
	"github.com/hyperyuri/wallet-cli/pkg/gorm/models"
	wallet_cli "github.com/hyperyuri/wallet-cli/pkg/wallet"
	"strconv"
)

type GormTransaction struct {

}

func (g GormTransaction) CreateTransaction(t *wallet_cli.Transaction) error {
	return gorm.DB.Create(&models.Transaction{
		Txid:              t.Txid(),
		Amount:            t.Amount(),
		Fee:               t.Fee(),
		Status:            "pending",
		BlockConfirmatios: "0",
		ToAddress:         t.ToAddress(),
		CurrencyID:        t.Currency().Id(),
		AddressID:         t.Address().Id(),
	}).Error
}

func (g GormTransaction) UpdateTranscations() error {
	var query []models.Transaction
	err := gorm.DB.Find(&query, "status = ?", "pending").Error
	if err != nil {
		return err
	}

	for _, v := range query {
		var cur models.Currency
		err := gorm.DB.First(&cur, "id = ?", v.CurrencyID).Error
		if err != nil {
			return err
		}

		client, err := ethclient.Dial(cur.URI)
		if err != nil {
			return err
		}

		header, err := client.HeaderByNumber(context.Background(), nil)
		if err != nil {
			return err
		}

		receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash(v.Txid))
		if err != nil {
			return err
		}

		confirmations := header.Number.Int64() - receipt.BlockNumber.Int64()

		if confirmations > 8 {
			v.Status = "done"
			v.BlockConfirmatios = strconv.FormatInt(confirmations, 10)

			err := gorm.DB.Save(&v).Error
			if err != nil {
				return err
			}
			continue
		}

		v.BlockConfirmatios = strconv.FormatInt(confirmations, 10)

		err = gorm.DB.Save(&v).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func (g GormTransaction) FindTransactions(items int, status string) ([]wallet_cli.Transaction, error) {
	var query []models.Transaction

	var err error
	if status == "all" {
		err = gorm.DB.Order("id desc").Limit(items).Find(&query).Error
	} else {
		err = gorm.DB.Order("id desc").Limit(items).Find(&query, "status = ?", status).Error
	}

	if err != nil {
		return nil, err
	}

	var transactions []wallet_cli.Transaction
	for _, v := range query {
		var cur models.Currency
		var addr models.Address

		err := gorm.DB.First(&cur, "id = ?", v.CurrencyID).Error
		if err != nil {
			return nil, err
		}
		err = gorm.DB.First(&addr, "id = ?", v.AddressID).Error
		if err != nil {
			return  nil, err
		}

		currency := wallet_cli.NewCurrency(cur.ID, cur.Symbol, cur.Network, cur.Type, cur.URI, cur.ContractAddress)
		address := wallet_cli.NewAddressWithFields(addr.ID, addr.Code, addr.Derivation)

		transactions = append(transactions, *wallet_cli.NewTransactionWithFields(v.Txid, v.Amount, v.Fee, v.Status, v.BlockConfirmatios, v.ToAddress, currency, address))
	}

	return transactions, nil
}

func NewGormTransaction() *GormTransaction {
	return &GormTransaction{}
}

