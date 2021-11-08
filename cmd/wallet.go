package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"wallet-cli/pkg/gorm/repository"
	wallet_cli "wallet-cli/pkg/wallet"
	"wallet-cli/pkg/wallet/balance"
	"wallet-cli/pkg/wallet/mnemonic"
)

var wallCmd = &cobra.Command{
	Use:   "create",
	Short: "Create your wallet",
	Run: func(cmd *cobra.Command, args []string) {
		password, err := cmd.Flags().GetString("pass")
		if err != nil {
			log.Fatalln(err)
		}

		createUser(password)
	},
}

var listWallCmd = &cobra.Command{
	Use:   "list",
	Short: "List your wallet by UUID",
	Run: func(cmd *cobra.Command, args []string) {
		uuid, err := cmd.Flags().GetString("uuid")
		if err != nil {
			log.Fatalln(err)
		}
		password, err := cmd.Flags().GetString("pass")
		if err != nil {
			log.Fatalln(err)
		}

		var balanceSvc wallet_cli.BalanceInfo
		balanceSvc = balance.NewBalanceService()

		var walletSvc wallet_cli.WalletStorage
		walletSvc = repository.NewGormWallet()
		wallet, err := walletSvc.LIstWalletByUUID(uuid)
		if err != nil {
			log.Fatalln(err)
		}

		if err := wallet.VerifyPassword(password); err != nil {
			log.Fatalln("password is wrong!")
		}

		var addressSvc wallet_cli.AddressStorage
		addressSvc = repository.NewGormAddress()

		fmt.Printf("Uuid:     %s\n", wallet.Uuid())
		fmt.Printf("Menmonic: %s\n", wallet.Mnemonic().Code())

		addresses, err := addressSvc.GetAdresses(wallet)
		if err != nil {
			log.Fatalln("cannot get adresses")
		}

		for _, v := range addresses {
			b, err := balanceSvc.GetBalance(&v)
			if err != nil {
				log.Fatalf("cannot get balance: %v\n", err)
			}

			fmt.Printf("---------- WALLET: %v ----------\n", 1)
			fmt.Printf("Network:             %s\n", v.Network())
			fmt.Printf("Address:             %s\n", v.Code())
			fmt.Printf("Balance Confirmed:   %s\n", b.Confimated())
			fmt.Printf("Balance Unconfirmed: %s\n", b.Unconfirmed())
			fmt.Println("-------------------------------")
		}

	},
}

func init() {
	rootCmd.AddCommand(wallCmd)
	rootCmd.AddCommand(listWallCmd)

	wallCmd.Flags().StringP("pass", "p", "", "wallet pass")
	listWallCmd.Flags().StringP( "uuid", "u", "","waallet uuid")
	listWallCmd.Flags().StringP("pass", "p", "","wallet password")
}

func createUser(password string) {
	wall, err := wallet_cli.NewWallet(password, mnemonic.NewService())
	if err != nil {
		log.Fatalln(err)
	}

	var walletSvc wallet_cli.WalletStorage
	walletSvc = repository.NewGormWallet()

	err = walletSvc.SaveWallet(wall)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Uuid:     %v\n", wall.Uuid())
	log.Printf("Mnemonic: %v\n", wall.Mnemonic().Code())
}
