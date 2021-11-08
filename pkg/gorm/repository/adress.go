package repository

import (
	"github.com/hyperyuri/wallet-cli/pkg/gorm"
	"github.com/hyperyuri/wallet-cli/pkg/gorm/models"
	wallet_cli "github.com/hyperyuri/wallet-cli/pkg/wallet"
)

type GormAddress struct {
	
}

func NewGormAddress() *GormAddress {
	return &GormAddress{}
}

func (g GormAddress) SaveAddress(wallet *wallet_cli.Wallet, address *wallet_cli.Address) error {
	return gorm.DB.Create(&models.Address{
		Code:       address.Code(),
		Derivation: address.Derivation(),
		Network:   address.Network(),
		WalletID:   uint(wallet.Id()),
	}).Error
}

func (g GormAddress) GetAdresses(wallet *wallet_cli.Wallet) ([]wallet_cli.Address, error) {
	var query []models.Address
	err := gorm.DB.Find(&query, "wallet_id = ?", wallet.Id()).Error
	if err != nil {
		return nil, err
	}

	var addresses []wallet_cli.Address
	for _, v := range query {
		addresses = append(addresses, *wallet_cli.NewAddressWithFields(v.Code, v.Derivation, v.Network))
	}

	return addresses, nil
}
