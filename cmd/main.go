package main

import (
	"fmt"
	"log"
	wallet_cli "wallet-cli"
	"wallet-cli/address"
	"wallet-cli/mnemonic"
)

func main(){
	user, err := wallet_cli.NewUser("123456", mnemonic.NewService())
	if err != nil {
		log.Fatalln(err)
	}

	wall, err := wallet_cli.NewWallet(user.Mnemonic().Code(), "0", "60", address.NewAddressService())
	user.AddWallet(wall)

	wall1, err := wallet_cli.NewWallet(user.Mnemonic().Code(), "1", "60", address.NewAddressService())
	user.AddWallet(wall1)

	fmt.Printf("Uuid: %s\n", user.Uuid())
	fmt.Printf("Menmonic: %s\n", user.Mnemonic().Code())

	for i, v := range user.Wallets() {
		fmt.Printf("---------- WALLET: %v ----------\n", i + 1)
		fmt.Println("Network: ETH")
		fmt.Printf("Address: %s\n", v.Address().Code())
		fmt.Println("-------------------------------")
	}
}
