package cmd

import (
	"fmt"
	"github.com/hyperyuri/wallet-cli/pkg/gorm/repository"
	wallet_cli "github.com/hyperyuri/wallet-cli/pkg/wallet"
	"github.com/hyperyuri/wallet-cli/pkg/wallet/mnemonic"
	"github.com/spf13/cobra"
	"log"
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
		uuid, err := cmd.Flags().GetString("wallet")
		if err != nil {
			log.Fatalln(err)
		}
		password, err := cmd.Flags().GetString("pass")
		if err != nil {
			log.Fatalln(err)
		}
		currency, err := cmd.Flags().GetString("currency")
		if err != nil {
			log.Fatalln(err)
		}
		network, err := cmd.Flags().GetString("network")
		if err != nil {
			log.Fatalln(err)
		}

		wallet, err := walletRepo.LIstWalletByUUID(uuid)
		if err != nil {
			log.Fatalln(err)
		}

		if err := wallet.VerifyPassword(password); err != nil {
			log.Fatalln("password is wrong!")
		}

		var currSvc wallet_cli.CurrencyStorage
		currSvc = repository.NewGormCurrency()
		c, err := currSvc.GetCurrency(network, currency)
		if err != nil {
			log.Fatalln("cannot get currency", err)
		}

		fmt.Printf("Uuid:     %s\n", wallet.Uuid())
		fmt.Printf("Menmonic: %s\n", wallet.Mnemonic().Code())

		addresses, err := addressRepo.GetAddresses(wallet)
		if err != nil {
			log.Fatalln("cannot get adresses")
		}

		for i, v := range addresses {
			b, err := balanceSvc.GetBalance(&v, c)
			if err != nil {
				log.Fatalf("cannot get balance: %v\n", err)
			}

			fmt.Printf("---------- ADDRESS: %v ----------\n", i + 1)
			fmt.Printf("Currency:            %s\n", c.Symbol())
			fmt.Printf("Network:             %s\n", c.Network())
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

	wallCmd.Flags().StringP("pass", "p", "", "wallet password")
	listWallCmd.Flags().StringP( "wallet", "w", "","waallet uuid")
	listWallCmd.Flags().StringP( "currency", "c", "","currency name")
	listWallCmd.Flags().StringP( "network", "n", "","currency network")
	listWallCmd.Flags().StringP("pass", "p", "","wallet password")
}

func createUser(password string) {
	wall, err := wallet_cli.NewWallet(password, mnemonic.NewService())
	if err != nil {
		log.Fatalln(err)
	}

	err = walletRepo.SaveWallet(wall)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Uuid:     %v\n", wall.Uuid())
	fmt.Printf("Mnemonic: %v\n", wall.Mnemonic().Code())
}
