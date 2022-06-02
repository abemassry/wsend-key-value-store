/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "wsend-key-value store get a value based on a key",
	Long: `wsend-key-value store get a value based on a key
This command gets a value given a key and a store

wkv get --store-link="https://wsnd.io/IdGzDoh/foo" --key="bar"
    --store-link name of the key value store container
    --key name of the key inside key value store
    --action optional
      "print" prints the text content if this is a link, it prints the link
      "download" downloads the file to the current directory
      "dump" the default, dumps the contents of the linked file if the link is
      a file
    --uid is optionally passed in like in create`,
	Run: func(cmd *cobra.Command, args []string) {
		if UID == "" {
			UID = getUID()
		}
		if UID == "" {
			fmt.Println(fmt.Errorf("Could not find uid"))
			os.Exit(1)
		}
		UID = strings.TrimSpace(UID)
		contents := []byte(Value)
		valueLink := ""


		extraParams := map[string]string{
			"uid": UID,
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
