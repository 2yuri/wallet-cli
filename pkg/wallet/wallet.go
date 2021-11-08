package wallet

import (
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type WalletStorage interface {
	SaveWallet(wallet *Wallet) error
	ListWallets() ([]Wallet, error)
	LIstWalletByUUID(uuid string) (*Wallet, error)
}

type Wallet struct {
	id int64
	uuid string
	password string
	mnemonic *Mnemonic
	adresses	[]*Address
}

func NewWalletWithFields(id int64, uuid string, password string, mnemonic *Mnemonic, adresses []*Address) *Wallet {
	return &Wallet{id: id, uuid: uuid, password: password, mnemonic: mnemonic, adresses: adresses}
}

func (u *Wallet) Id() int64 {
	return u.id
}

func (u *Wallet) Uuid() string {
	return u.uuid
}

func (u *Wallet) Password() string {
	return u.password
}

func (u *Wallet) Mnemonic() *Mnemonic {
	return u.mnemonic
}

func (u *Wallet) Adresses() []*Address {
	return u.adresses
}

func NewWallet(password string, generator MnemonicGenerator) (*Wallet, error) {
	newUser := &Wallet{id: 1, uuid: uuid.New().String()}
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

func (u *Wallet) hashPassword(password string) error {
	p, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}
	u.password = string(p)
	return nil
}

func (u *Wallet) VerifyPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.password), []byte(providedPassword))
	if err != nil {
		return errors.New("invalid password")
	}
	return nil
}

func (u *Wallet) AddAddress(wallet *Address) {
	u.adresses = append(u.adresses, wallet)
}

