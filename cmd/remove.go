/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "delete a key and value",
	Long: `Delete a key and value

Deletes a key and the value, as well as the backing file of the value.

wkv remove --store-link="https://wsnd.io/IdGzDoh/foo" --key="bar" --uid="0123456789abcdef"

	--store-link name of the key value store container
	--key name of the key inside key value store
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
		splited := strings.Split(StoreLink, "/")
		storeName := splited[4]
		resp, err := http.Get(StoreLink)
		if err != nil {
			fmt.Println(fmt.Errorf(err.Error()))
			os.Exit(1)
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			fmt.Println(fmt.Errorf("could not get link, status code is %d", resp.StatusCode))
			os.Exit(1)
		}
		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		dataStore := buf.String()
		var keyValStore map[string]string
		err = json.Unmarshal([]byte(dataStore), &keyValStore)
		if err != nil {
			fmt.Println(fmt.Errorf("store link invalid json"))
			fmt.Println(fmt.Errorf(err.Error()))
			os.Exit(1)
		}

		fileToRemove := keyValStore[Key]
		err = DeleteFile(fileToRemove)
		if err != nil {
			fmt.Println(fmt.Errorf("error deleting file"))
			fmt.Println(fmt.Errorf(err.Error()))
			os.Exit(1)
		}

		delete(keyValStore, Key)

		textContents, err := json.Marshal(keyValStore)
		if err != nil {
			fmt.Println(fmt.Errorf("error generating json"))
			fmt.Println(fmt.Errorf(err.Error()))
			os.Exit(1)
		}

		contents := []byte(textContents)
		extraParams := map[string]string{
			"uid":  UID,
			"link": StoreLink,
		}
		formDataContentType, request, err := UploadNoFile("https://wsend.net/update_cli", extraParams, storeName, contents)
		if err != nil {
			fmt.Println(fmt.Errorf(err.Error()))
			os.Exit(1)
		}
		request.Header.Add("Content-Type", formDataContentType)
		client := &http.Client{}
		resp, err = client.Do(request)
		if err != nil {
			fmt.Println(fmt.Errorf(err.Error()))
			os.Exit(1)
		}
		var bodyContent []byte
		resp.Body.Read(bodyContent)
		resp.Body.Close()
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	removeCmd.Flags().StringVarP(&StoreLink, "store-link", "n", "", "link of the key value store")
	removeCmd.Flags().StringVarP(&Key, "key", "k", "", "name of the key")
	removeCmd.Flags().StringVarP(&UID, "uid", "u", "", "access token")
	removeCmd.MarkFlagRequired("store-link")
	removeCmd.MarkFlagRequired("key")
}

// DeleteFile deletes a file that is referred to by Value
func DeleteFile(fileToRemove string) error {
	data := url.Values{
		"uid":  {UID},
		"link": {fileToRemove},
	}

	_, err := http.PostForm("https://wsend.net/delete_cli", data)

	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
		return err
	}

	return nil
}
