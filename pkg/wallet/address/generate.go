package address

import (
	"errors"
	"fmt"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"regexp"
)

type AddressService struct {
}

func NewAddressService() *AddressService {
	return &AddressService{}
}

func (a *AddressService) Generate(men string, derivation string) (string, error){
	wallet, err := hdwallet.NewFromMnemonic(men)
	if err != nil {
		return "", err
	}

	path := hdwallet.MustParseDerivationPath(fmt.Sprintf("m/44'/60'/0'/0/%s", derivation))
	account, err := wallet.Derive(path, false)
	if err != nil {
		return "", err
	}

	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	if !re.MatchString(account.Address.Hex()) {
		return "", errors.New("cannot generate new address")
	}


	return account.Address.Hex(), nil
}

