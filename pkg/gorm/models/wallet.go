package models

import "time"

type Wallet struct {
	ID uint `gorm:"primaryKey"`
	Uuid string
	Password string
	Mnemonic string
	Addresses []Address
	CreatedAt time.Time
	UpdatedAt    time.Time
}