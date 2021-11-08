package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"github.com/hyperyuri/wallet-cli/pkg/gorm/repository"
	wallet_cli "github.com/hyperyuri/wallet-cli/pkg/wallet"
	"github.com/hyperyuri/wallet-cli/pkg/wallet/address"
)

var addressCmd = &cobra.Command{
	Use:   "add",
	Short: "Create a new address for a Wallet",
	Run: func(cmd *cobra.Command, args []string) {
		uuid, err := cmd.Flags().GetString("wallet")
		if err != nil {
			log.Fatalln(err)
		}
		password, err := cmd.Flags().GetString("pass")
		if err != nil {
			log.Fatalln(err)
		}
		coin, err := cmd.Flags().GetString("coin")
		if err != nil {
			log.Fatalln(err)
		}

		createAddress(uuid, password, coin)
	},
}

func init(){
	rootCmd.AddCommand(addressCmd)

	addressCmd.Flags().StringP("wallet", "w", "", "wallet uuid")
	addressCmd.Flags().StringP("coin", "c", "", "coin name (ETH, BNB)")
	addressCmd.Flags().StringP("pass", "p", "", "wallet password")
}

func createAddress(uuid, pass, coin string){
	var walletSvc wallet_cli.WalletStorage
	walletSvc = repository.NewGormWallet()

	wall, err :=  walletSvc.LIstWalletByUUID(uuid)
	if err != nil {
		log.Fatalln(err)
	}

	if err := wall.VerifyPassword(pass); err != nil {
		log.Fatalln("password is wrong!")
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