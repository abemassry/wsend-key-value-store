/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// storeCmd represents the store command
var storeCmd = &cobra.Command{
	Use:   "store",
	Short: "store a value based on a key",
	Long: `Store a value based on a key.

A key is a string and value can be anything but defaults to a string and is
stored as a URL

wkv store --key="foo" --value="bar" --type="string"
value is either a string (default) or a file specified by --type="file"
if a file is specified the path is either absolute or the default is the
current directory`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("store called")
	},
}

func init() {
	rootCmd.AddCommand(storeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// storeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// storeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	storeCmd.Flags().String("key", "", "name of the key")
	storeCmd.Flags().String("value", "", "either text or the name of a file")
	storeCmd.Flags().String("type", "string", "the type of value, defaults to string")
	storeCmd.MarkFlagRequired("key")
	storeCmd.MarkFlagRequired("value")
}
