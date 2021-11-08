package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"wallet-cli/pkg/gorm/repository"
	wallet_cli "wallet-cli/pkg/wallet"
	"wallet-cli/pkg/wallet/address"
)

var addressCmd = &cobra.Command{
	Use:   "add",
	Short: "Create a new address",
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] == "" {
			return
		}

		createAddress(args[1], args[3])
	},
}

func init(){
	rootCmd.AddCommand(addressCmd)

	addressCmd.Flags().String("user", "u", "pass user uuid")
	addressCmd.Flags().String("coin", "c", "coin name (ETH, BNB)")
}

func createAddress(uuid string, coin string){
	var walletSvc wallet_cli.WalletStorage
	walletSvc = repository.NewGormWallet()

	wall, err :=  walletSvc.LIstWalletByUUID(uuid)
	if err != nil {
		log.Fatalln(err)
	}

	var addressSvc wallet_cli.AddressStorage
	addressSvc = repository.NewGormAddress()

	addr, err :=  addressSvc.GetAdresses(wall)
	if err != nil {
		log.Fatalln(err)
	}

	var derivation int
	if len(addr) > 0 {
		derivation, err = strconv.Atoi(addr[len(addr)-1].Derivation())
		if err != nil {
			log.Fatalln(err)
		}
	}

	add, err := wallet_cli.NewAddress(wall.Mnemonic().Code(), strconv.Itoa(derivation + 1), coin, address.NewAddressService())
	if err != nil {
		log.Fatalln(err)
	}


	err = addressSvc.SaveAddress(wall, add)
	if err != nil {
		log.Fatalln(err)
	}
}