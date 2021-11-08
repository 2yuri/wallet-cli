package repository

import (
	"wallet-cli/pkg/gorm"
	"wallet-cli/pkg/gorm/models"
	wallet_cli "wallet-cli/pkg/wallet"
)

type GormWallet struct {
}

func (g GormWallet) LIstWalletByUUID(uuid string) (*wallet_cli.Wallet, error) {
	query := models.Wallet{
	 	Uuid: uuid,
	}
	err := gorm.DB.Find(&query).Error
	if err != nil {
		return nil, err
	}

	return wallet_cli.NewWalletWithFields(
		int64(query.ID), query.Uuid, query.Password, wallet_cli.NewMnemonic(query.Mnemonic), nil),nil
}

func NewGormWallet() *GormWallet {
	return &GormWallet{}
}

func (g GormWallet) SaveWallet(wallet *wallet_cli.Wallet) error {
	return gorm.DB.Create(&models.Wallet{
		Uuid:     wallet.Uuid(),
		Password: wallet.Password(),
		Mnemonic: wallet.Mnemonic().Code(),
	}).Error
}

func (g GormWallet) ListWallets() ([]wallet_cli.Wallet, error) {
	var query []models.Wallet
	err := gorm.DB.Find(&query).Error
	if err != nil {
		return nil, err
	}


	var wallets []wallet_cli.Wallet
	for _, v := range query {
		wallets = append(wallets, *wallet_cli.NewWalletWithFields(
			int64(v.ID), v.Uuid, v.Password, wallet_cli.NewMnemonic(v.Mnemonic), nil))
	}

	return wallets, nil
}
