/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "calculator",
	Short: "This is a simple calculator",
	Long: `This is a simple calculator application.
It supports Addition, Subtraction, Division, Multiplication.
Specify the numbers to use with --x or --y.
Default values for x and y are 0.
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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
	// Ajout des commandes comme addCmd à rootCmd
	rootCmd.AddCommand(addCmd)

	// Définir les flags persistants pour rootCmd (si nécessaire)
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.truc.yaml)")

	// Définir les flags locaux pour rootCmd (si nécessaire)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
