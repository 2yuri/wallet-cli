package cmd

import (
	"fmt"
	wallet_cli "github.com/hyperyuri/wallet-cli/pkg/wallet"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"log"
	"math/big"
	"strconv"
	"time"
)

var sendCmd = &cobra.Command{
	Use:   "send",
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
		budge, err := cmd.Flags().GetString("amount")
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

		sendTransaction(uuid, password, from, to, budge, currency, network)
	},
}

func init(){
	rootCmd.AddCommand(sendCmd)

	sendCmd.Flags().StringP("wallet", "w", "", "wallet uuid")
	sendCmd.Flags().StringP("pass", "p", "", "wallet password")
	sendCmd.Flags().StringP("from", "f", "", "wallet address")
	sendCmd.Flags().StringP("to", "t", "", "adress")
	sendCmd.Flags().StringP("amount", "a", "", "amount to be sended")
	sendCmd.Flags().StringP("currency", "c", "", "currency")
	sendCmd.Flags().StringP("network", "n", "", "network")
}

func sendTransaction(uuid, pass, from, to, amount, currency, network string){
	wall, err :=  walletRepo.LIstWalletByUUID(uuid)
	if err != nil {
		log.Fatalln(err)
	}
	if err := wall.VerifyPassword(pass); err != nil {
		log.Fatalln("password is wrong!")
	}

	amountInt, ok := new(big.Float).SetString(amount)
	if !ok {
		log.Fatalln("cannot convert your amount")
	}

	cur, err := currRepo.GetCurrency(network, currency)
	if err != nil {
		log.Fatalln(err)
	}

	addr, err := addressRepo.GetAddressByCode(from, wall.Id())
	if err != nil {
		log.Fatalln(err)
	}

	trx := wallet_cli.NewTransaction(amount, to, cur, addr)

	fee, err := trx.GetFee(trxSvc)
	if err != nil {
		log.Fatalln(err)
	}

	feeInt, ok := new(big.Float).SetString(fee)
	if !ok {
		log.Fatalln("cannot get your balance")
	}

	if	feeInt.Cmp(amountInt) > 0{
		log.Fatalln("cannot continue with operation, fee is bigger than amount")
	}

	balance, err := balanceSvc.GetBalance(addr, cur)
	if err != nil {
		log.Fatalln(err)
	}

	balanceInt, ok := new(big.Float).SetString(balance.Confimated())
	if !ok {
		log.Fatalln("cannot get your balance")
	}

	coinType := "BNB"

	if cur.Network() == "ETH" {
		coinType = "ETH"
	}

	switch getFeeOption(fee, coinType) {
	case 1:
		if	new(big.Float).Add(amountInt, feeInt).Cmp(balanceInt) > 0 {
			log.Fatalln("insufficient founds for this operation!")
		}

		transaction, err := trx.SendTransaction(trxSvc, wall.Mnemonic().Code())
		if err != nil {
			log.Fatalln("cannot complete transaction: ", err)
		}

		if err != trxRepo.CreateTransaction(transaction) {
			log.Fatalln("cannot save at database")
		}

		fmt.Println("transaction succeed!!")
		fmt.Printf("transactionID: %v\n", transaction.Txid())
	case 2:
		log.Fatalln("operation cancelled.")
	default:
		log.Fatalln("invalid option")
	}

}

func getFeeOption(fee string, coin string) int {
	fmt.Printf("Your transaction fee is %v %v\n\n", fee, coin)
	fmt.Println("1 - Accept discounting fee from balance.")
	fmt.Printf("2 - Cancel transaction.\n\n")

	ticker := time.NewTicker(5 * time.Second)
	var result string

	go func() {
		var err error
		prompt := promptui.Prompt{
			Label: "Select your option",
		}
		result, err = prompt.Run()
		if err != nil {
			log.Fatalf("Prompt failed %v\n", err)
		}
	}()

	<-ticker.C

	if result == "" {
		log.Fatalln("time out!")
	}


	i, err := strconv.Atoi(result)
	if err != nil {
		log.Fatalln("value is not a number", err)
	}

	return i
}