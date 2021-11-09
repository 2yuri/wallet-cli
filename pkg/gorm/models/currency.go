package models

type Currency struct {
	ID uint `gorm:"primaryKey"`
	Symbol string
	Network string
	Type string
	URI string
	ContractAddress string
}