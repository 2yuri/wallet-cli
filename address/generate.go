package address

import (
	"fmt"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

type AddressService struct {

}

func NewAddressService() *AddressService {
	return &AddressService{}
}

func (a *AddressService) Generate(men string, derivation string, network string) (string, error){
	wallet, err := hdwallet.NewFromMnemonic(men)
	if err != nil {
		return "", err
	}

	path := hdwallet.MustParseDerivationPath(fmt.Sprintf("m/44'/%s'/0'/0/%s", network, derivation))
	account, err := wallet.Derive(path, false)
	if err != nil {
		return "", err
	}

	return account.Address.Hex(), nil
}

