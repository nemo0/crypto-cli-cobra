/*
Copyright © 2022 Subha Chanda <subhachanda88@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "crypto-cli",
	Short: "Get cryptocurrency prices and market data",
	Long: `To get the price of a cryptocurrency, use the price command.
	Get the price of a cryptocurrency. For currencies with spaces, example terra lune, use - instead of spaces. Example: terra-luna`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) { 
		fmt.Println("Hello World")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra-crypto-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().BoolP("markets", "m", false, "Get the markets listings of a cryptocurrency")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


