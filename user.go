package wallet_cli

import (
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	id int64
	uuid string
	password string
	mnemonic *Mnemonic
	wallets	[]*Wallet
}

func (u *User) Id() int64 {
	return u.id
}

func (u *User) Uuid() string {
	return u.uuid
}

func (u *User) Password() string {
	return u.password
}

func (u *User) Mnemonic() *Mnemonic {
	return u.mnemonic
}

func (u *User) Wallets() []*Wallet {
	return u.wallets
}

func NewUserCompelted(id int64, uuid string, password string, mnenonics *Mnemonic, wallets []*Wallet) *User {
	return &User{id: id, uuid: uuid, password: password, mnemonic: mnenonics, wallets: wallets}
}

func NewUser(password string, generator MnemonicGenerator) (*User, error) {
	newUser := &User{id: 1, uuid: uuid.New().String()}
	mnm, err := generator.Generate()
	if err != nil {
		return nil, err
	}

	newUser.mnemonic = mnm

	err = newUser.hashPassword(password)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (u *User) hashPassword(password string) error {
	p, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}
	u.password = string(p)
	return nil
}

func (u *User) VerifyPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.password), []byte(providedPassword))
	if err != nil {
		return errors.New("invalid password")
	}
	return nil
}

func (u *User) AddWallet(wallet *Wallet) {
	u.wallets = append(u.wallets, wallet)
}

