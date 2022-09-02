/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// descriptCatCmd represents the descriptCat command
var descriptCatCmd = &cobra.Command{
	Use:   "descriptCat",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		b, err := os.ReadFile("ascii_art/cat.txt")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print(string(b))
	},
}

func init() {
	rootCmd.AddCommand(descriptCatCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// descriptCatCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// descriptCatCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
