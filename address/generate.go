package address

import (
	"fmt"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	wallet_cli "wallet-cli"
)

type AddressService struct {

}

func NewAddressService() *AddressService {
	return &AddressService{}
}

func (a *AddressService) Generate(men string, derivation string, network string) (*wallet_cli.Address, error){
	wallet, err := hdwallet.NewFromMnemonic(men)
	if err != nil {
		return nil, err
	}

	path := hdwallet.MustParseDerivationPath(fmt.Sprintf("m/44'/%s'/0'/0/%s", network, derivation))
	account, err := wallet.Derive(path, false)
	if err != nil {
		return nil, err
	}

	return wallet_cli.NewAddress(account.Address.Hex()), nil
}

