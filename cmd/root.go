package cmd

import (
	"fmt"
	"github.com/hyperyuri/wallet-cli/pkg/gorm/repository"
	wallet_cli "github.com/hyperyuri/wallet-cli/pkg/wallet"
	"github.com/hyperyuri/wallet-cli/pkg/wallet/balance"
	"github.com/hyperyuri/wallet-cli/pkg/wallet/transaction"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "repository-cli",
	Short: "ETH repository cli",
	Long: "vtnc",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to ETH-HDWALLET")
	},
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".my-calc" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".my-calc")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

var balanceSvc wallet_cli.BalanceInfo
var trxSvc wallet_cli.TransactionActions


var currRepo wallet_cli.CurrencyStorage
var walletRepo wallet_cli.WalletStorage
var addressRepo wallet_cli.AddressStorage
var trxRepo wallet_cli.TransactionStorage


func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.my-calc.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")


	balanceSvc = balance.NewBalanceService()
	currRepo = repository.NewGormCurrency()
	walletRepo = repository.NewGormWallet()
	addressRepo = repository.NewGormAddress()
	trxSvc = transaction.NewTransactionService()
	trxRepo = repository.NewGormTransaction()
}


func Execute(){
	cobra.CheckErr(rootCmd.Execute())
}
