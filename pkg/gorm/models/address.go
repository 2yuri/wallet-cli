package models

type Address struct {
	ID uint `gorm:"primaryKey"`
	Code string
	Derivation string
	Network string
	WalletID uint
}
