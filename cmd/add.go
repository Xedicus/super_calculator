/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"calculator/internals/math"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	x, y float64
)

// addCmd représente la commande add
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Addition de x et y",
	Long:  `Addition de x et y`,
	Run: func(cmd *cobra.Command, args []string) {
		ret := math.Addition(x, y)
		fmt.Printf("%f + %f = %f\n", x, y, ret)
	},
}

func init() {
	// Assurez-vous que rootCmd est déjà défini dans main.go
	addCmd.Flags().Float64VarP(&x, "x", "x", 0, "première valeur")
	addCmd.Flags().Float64VarP(&y, "y", "y", 0, "deuxième valeur")
}

