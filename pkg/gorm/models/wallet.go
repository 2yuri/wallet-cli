package models

type Wallet struct {
	ID uint `gorm:"primaryKey"`
	Uuid string
	Password string
	Mnemonic string
	Addresses []Address
}