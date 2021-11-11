package models

import "time"

type Address struct {
	ID uint `gorm:"primaryKey"`
	Code string
	Derivation string
	WalletID uint
	CreatedAt time.Time
	UpdatedAt    time.Time
}
