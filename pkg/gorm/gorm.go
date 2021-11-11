package gorm

import (
	"github.com/hyperyuri/wallet-cli/pkg/gorm/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)

var DB *gorm.DB

func init(){
	db, err := gorm.Open("sqlite3", "./db.db")
	if err != nil {
		log.Fatalf("cannot connect to db: %v\n", err)
	}

	DB = db

	DB.AutoMigrate(&models.Wallet{})
	DB.AutoMigrate(&models.Address{})
	DB.AutoMigrate(&models.Transaction{})
	DB.AutoMigrate(&models.Currency{})

	if err := DB.First(&models.Currency{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			RunSeeds()
			return
		}

		log.Fatalln("cannot run seeds: ", err)
	}
}

func RunSeeds(){
	tx := DB.Begin()

	if err := CreateCurrency(tx, "ETH", "ETH", "https://mainnet.infura.io/v3/bf5b520aafff4ec8abf9fb1983391a3a", "native", ""); err != nil {
		tx.Rollback()
		log.Fatalln("cannot run seeds: ", err)
	}
	if err := CreateCurrency(tx,"USDT", "ETH", "https://mainnet.infura.io/v3/bf5b520aafff4ec8abf9fb1983391a3a", "erc20", "0xdAC17F958D2ee523a2206206994597C13D831ec7"); err != nil {
		tx.Rollback()
		log.Fatalln("cannot run seeds: ", err)
	}

	if err := CreateCurrency(tx,"BNB", "BSC", "https://bsc-dataseed.binance.org", "native", ""); err != nil {
		tx.Rollback()
		log.Fatalln("cannot run seeds: ", err)
	}

	if err := CreateCurrency(tx,"USDT", "BSC", "https://bsc-dataseed.binance.org", "erc20", "0x55d398326f99059ff775485246999027b3197955"); err != nil {
		tx.Rollback()
		log.Fatalln("cannot run seeds: ", err)
	}

	if err := CreateCurrency(tx,"BNB", "TEST", "https://data-seed-prebsc-1-s1.binance.org:8545", "native", ""); err != nil {
		tx.Rollback()
		log.Fatalln("cannot run seeds: ", err)
	}


	tx.Commit()
}

func CreateCurrency(tx *gorm.DB, symbol, network, uri, tokenType, contractAddress string) error {
	return tx.Create(&models.Currency{
		Symbol:          symbol,
		Network:         network,
		URI: uri,
		Type:            tokenType,
		ContractAddress: contractAddress,
	}).Error
}