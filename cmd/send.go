package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"github.com/hyperyuri/wallet-cli/pkg/gorm/repository"
	wallet_cli "github.com/hyperyuri/wallet-cli/pkg/wallet"
)

var sendCmd = &cobra.Command{
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
		from, err := cmd.Flags().GetString("from")
		if err != nil {
			log.Fatalln(err)
		}
		to, err := cmd.Flags().GetString("to")
		if err != nil {
			log.Fatalln(err)
		}
		budge, err := cmd.Flags().GetString("budge")
		if err != nil {
			log.Fatalln(err)
		}


		sendTransaction(uuid, password, from, to, budge)
	},
}

func init(){
	rootCmd.AddCommand(sendCmd)

	sendCmd.Flags().StringP("wallet", "w", "", "wallet uuid")
	sendCmd.Flags().StringP("pass", "p", "", "wallet password")
	sendCmd.Flags().StringP("from", "f", "", "wallet address")
	sendCmd.Flags().StringP("to", "t", "", "adress")
	sendCmd.Flags().StringP("budge", "b", "", "value")
}

func sendTransaction(uuid, pass, from, to, budge string){
	var walletSvc wallet_cli.WalletStorage
	walletSvc = repository.NewGormWallet()

	wall, err :=  walletSvc.LIstWalletByUUID(uuid)
	if err != nil {
		log.Fatalln(err)
	}

	if err := wall.VerifyPassword(pass); err != nil {
		log.Fatalln("password is wrong!")
	}
	
	log.Fatalln(wall)


}