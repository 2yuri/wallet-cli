package models

import "time"

type Transaction struct {
	ID uint `gorm:"primaryKey"`
	Txid string
	Amount string
	Fee string
	Status string
	BlockHash string
	BlockConfirmatios string
	ToAddress string
	AddressID uint
	CreatedAt time.Time
	UpdatedAt    time.Time
}