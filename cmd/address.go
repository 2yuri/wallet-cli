package cmd

import (
	"fmt"
	wallet_cli "github.com/hyperyuri/wallet-cli/pkg/wallet"
	"github.com/hyperyuri/wallet-cli/pkg/wallet/address"
	"github.com/spf13/cobra"
	"log"
	"strconv"
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

		createAddress(uuid, password)
	},
}

func init(){
	rootCmd.AddCommand(addressCmd)

	addressCmd.Flags().StringP("wallet", "w", "", "wallet uuid")
	addressCmd.Flags().StringP("pass", "p", "", "wallet password")
}

func createAddress(uuid, pass string){
	wall, err :=  walletRepo.LIstWalletByUUID(uuid)
	if err != nil {
		log.Fatalln(err)
	}

	if err := wall.VerifyPassword(pass); err != nil {
		log.Fatalln("password is wrong!")
	}

	addr, err :=  addressRepo.GetAddresses(wall)
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

	add, err := wallet_cli.NewAddress(wall.Mnemonic().Code(), strconv.Itoa(derivation + 1), address.NewAddressService())
	if err != nil {
		log.Fatalln(err)
	}

	err = addressRepo.SaveAddress(wall, add)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Address: %v\n", add.Code())
}