/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/josenaldo/fc-arquitetura-hexagonal-jom/adapters/cli"
	"github.com/spf13/cobra"
)

var (
	action       string
	productId    string
	productName  string
	productPrice float64
)

// cliCmd represents the cli command
var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Run(productService, action, productId, productName, productPrice)
	},
}

func init() {
	rootCmd.AddCommand(cliCmd)
	rootCmd.Flags().StringVarP(&action, "action", "a", "enabled", "Enable or disable a product")
	rootCmd.Flags().StringVarP(&productId, "id", "i", "", "Product ID")
	rootCmd.Flags().StringVarP(&productName, "name", "n", "", "Product name")
	rootCmd.Flags().Float64VarP(&productPrice, "price", "p", 0.0, "Product price")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cliCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cliCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
