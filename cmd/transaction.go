package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

var transactionCmd = &cobra.Command{
	Use:   "transaction",
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
		status, err := cmd.Flags().GetString("status")
		if err != nil {
			log.Fatalln(err)
		}
		items, err := cmd.Flags().GetString("items")
		if err != nil {
			log.Fatalln(err)
		}

		getTransactions(uuid, password, status, items)
	},
}

func init(){
	rootCmd.AddCommand(transactionCmd)

	transactionCmd.Flags().StringP("wallet", "w", "", "wallet uuid")
	transactionCmd.Flags().StringP("pass", "p", "", "wallet password")
	transactionCmd.Flags().StringP("status", "s", "", "transaction status")
	transactionCmd.Flags().StringP("items", "i", "", "total of transaction to list")
}

func getTransactions(uuid, pass, status, items string) {
	wall, err :=  walletRepo.LIstWalletByUUID(uuid)
	if err != nil {
		log.Fatalln(err)
	}
	if err := wall.VerifyPassword(pass); err != nil {
		log.Fatalln("password is wrong!")
	}

	total, err := strconv.Atoi(items)
	if err != nil {
		log.Println("cannot convert total")
	}

	if err := trxRepo.UpdateTranscations(); err != nil {
		log.Fatalln("cannot verify transactions", err)
	}

	trxs, err := trxRepo.FindTransactions(total, status)
	if err != nil {
		log.Println("cannot convert total")
	}

	for i, v := range trxs {
		fmt.Printf("---------- Transaction: %v ----------\n", i + 1)
		fmt.Printf("TxID:                %s\n", v.Txid())
		fmt.Printf("Status:              %s\n", v.Status())
		fmt.Printf("Confirmation Blocks: %s\n", v.BlockConfirmatios())
		fmt.Println("-------------------------------")
	}

	if len(trxs) == 0 {
		fmt.Println("cannot find any transactions with your search")
	}

}