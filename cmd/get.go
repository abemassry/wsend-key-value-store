/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// Action is optional and is either print, download, or dump, default dump
var Action string

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
		a file`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get(StoreLink)
		if resp.StatusCode != 200 {
			fmt.Println(fmt.Errorf("could not get link, status code is %d", resp.StatusCode))
		}
		defer resp.Body.Close()
		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		dataStore := buf.String()
		var keyValStore map[string]string
		err = json.Unmarshal([]byte(dataStore), &keyValStore)
		if err != nil {
			fmt.Println(fmt.Errorf(err.Error()))
			os.Exit(1)
		}
		valueLink := keyValStore[Key]
		if Action == "" || Action == "dump" {
			resp, err = http.Get(valueLink)
			if err != nil {
				fmt.Println(fmt.Errorf(err.Error()))
				os.Exit(1)
			}
			defer resp.Body.Close()
			if resp.StatusCode != 200 {
				fmt.Println(fmt.Errorf("could not get link, status code is %d", resp.StatusCode))
			}
			buf = new(bytes.Buffer)
			buf.ReadFrom(resp.Body)
			valueContent := buf.String()
			fmt.Println(valueContent)
		} else if Action == "print" {
			fmt.Println(valueLink)
		} else if Action == "download" {
			// Create the file
			splited := strings.Split(valueLink, "/")
			fileName := splited[4]
			out, err := os.Create(fileName)
			if err != nil {
				fmt.Println(fmt.Errorf(err.Error()))
				os.Exit(1)
			}
			defer out.Close()

			// Get the data
			resp, err := http.Get(valueLink)
			if err != nil {
				fmt.Println(fmt.Errorf(err.Error()))
				os.Exit(1)
			}
			defer resp.Body.Close()

			// Write the body to file
			_, err = io.Copy(out, resp.Body)
			if err != nil {
				fmt.Println(fmt.Errorf(err.Error()))
				os.Exit(1)
			}
			fmt.Println("Downloaded: " + fileName)
		} else {
			fmt.Println("Action not found, must be print, download, or dump, default dump")
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
	getCmd.Flags().StringVarP(&StoreLink, "store-link", "n", "", "link of the key value store")

	getCmd.Flags().StringVarP(&Key, "key", "k", "", "name of the key")
	getCmd.Flags().StringVarP(&Action, "action", "a", "", "action to print, download, dump")

	getCmd.MarkFlagRequired("store-link")
	getCmd.MarkFlagRequired("key")

}
