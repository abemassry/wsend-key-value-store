/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Key is the key in the key value store
var Key string

// Value is the value in the key value store
var Value string

// Type is optional and can be a string or file
var Type string

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
		fmt.Println("")
		fmt.Println("key:")
		fmt.Println(Key)
		fmt.Println("")
		fmt.Println("value:")
		fmt.Println(Value)
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
	storeCmd.Flags().StringVarP(&Key, "key", "k", "", "name of the key")
	storeCmd.Flags().StringVarP(&Value, "value", "v", "", "either text or the name of a file")
	storeCmd.Flags().StringVarP(&Type, "type", "t", "string", "the type of value, defaults to string")
	storeCmd.MarkFlagRequired("key")
	storeCmd.MarkFlagRequired("value")
}
