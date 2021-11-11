package models

import "time"

type Transaction struct {
	ID uint `gorm:"primaryKey"`
	Txid string
	Amount string
	Fee string
	Status string
	BlockConfirmatios string
	ToAddress string
	CurrencyID uint
	AddressID uint
	CreatedAt time.Time
	UpdatedAt    time.Time
}