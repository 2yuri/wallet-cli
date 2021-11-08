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
	Short: "ETH repository cli",
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] == "" {
			return
		}

		createUser(args[0])
	},
}

var listWallCmd = &cobra.Command{
	Use:   "user",
	Short: "ETH repository cli",
	Run: func(cmd *cobra.Command, args []string) {
		var balanceSvc wallet_cli.BalanceInfo
		balanceSvc = balance.NewBalanceService()

		var walletSvc wallet_cli.WalletStorage
		walletSvc = repository.NewGormWallet()
		wallets, err :=  walletSvc.ListWallets()
		if err != nil {
			log.Fatalln(err)
		}

		var addressSvc wallet_cli.AddressStorage
		addressSvc = repository.NewGormAddress()

		for _, user := range wallets {
			fmt.Printf("Uuid:     %s\n", user.Uuid())
			fmt.Printf("Menmonic: %s\n", user.Mnemonic().Code())

			addresses, err := addressSvc.GetAdresses(&user)
			if err != nil {
				continue
			}

			for _, v := range addresses {
				b, err := balanceSvc.GetBalance(&v)
				if err != nil {
					log.Fatalln(err)
				}

				fmt.Printf("---------- WALLET: %v ----------\n", 1)
				fmt.Printf("Network:             %s\n", v.Network())
				fmt.Printf("Address:             %s\n", v.Code())
				fmt.Printf("Balance Confirmed:   %s\n", b.Confimated())
				fmt.Printf("Balance Unconfirmed: %s\n", b.Unconfirmed())
				fmt.Println("-------------------------------")
			}
		}

	},
}

func init(){
	rootCmd.AddCommand(wallCmd)
	rootCmd.AddCommand(listWallCmd)

	wallCmd.Flags().String("pass", "p", "yoyur pass")
	listWallCmd.Flags().String("list", "l", "list user")
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
}
