package repository

import (
	"github.com/hyperyuri/wallet-cli/pkg/gorm"
	"github.com/hyperyuri/wallet-cli/pkg/gorm/models"
	wallet_cli "github.com/hyperyuri/wallet-cli/pkg/wallet"
)

type GormCurrency struct {
}

func NewGormCurrency() *GormCurrency {
	return &GormCurrency{}
}

func (g GormCurrency) GetCurrency(net, sym string) (*wallet_cli.Currency, error) {
	var query models.Currency

	err := gorm.DB.Find(&query, "symbol = ? and network = ?", sym, net).Error
	if err != nil {
		return nil, err
	}

	return wallet_cli.NewCurrency(query.ID, query.Symbol, query.Network, query.Type, query.URI, query.ContractAddress), nil
}