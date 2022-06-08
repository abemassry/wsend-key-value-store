/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// StoreLink is the link of the key value store
var StoreLink string

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

wkv store --store-link="https://wsnd.io/IdGzDoh/foo" --key="bar" --value="baz" --type="string"
    --store-link name of the key value store container
    --key name of the key inside key value store
    --value either a string (default) or a file based on the --type
      in either case a file is uploaded which allows the string data to be
      incredbily long and is referenced by a URL pointing to the file.
    --type string or file
value is either a string (default) or a file specified by --type="file"
if a file is specified the path is either absolute or the default is the
current directory
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
		var contents []byte
		var err error
		if Type == "" || Type == "string" {
			contents = []byte(Value)
		} else if Type == "file" {
			contents, err = ioutil.ReadFile(Value)
			if err != nil {
				fmt.Println(fmt.Errorf(err.Error()))
				os.Exit(1)
			}
		}
		valueLink := ""

		extraParams := map[string]string{
			"uid": UID,
		}
		formDataContentType, request, err := UploadNoFile("https://wsend.net/upload_cli", extraParams, Key, contents)
		if err != nil {
			fmt.Println(fmt.Errorf(err.Error()))
			os.Exit(1)
		}
		request.Header.Add("Content-Type", formDataContentType)
		client := &http.Client{}
		resp, err := client.Do(request)
		if err != nil {
			fmt.Println(fmt.Errorf(err.Error()))
			os.Exit(1)
		} else {

			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(fmt.Errorf(err.Error()))
				os.Exit(1)
			}
			valueLink = string(bodyBytes)

			resp.Body.Close()
		}

		if valueLink == "" {
			fmt.Println("problem with upload contents")
			os.Exit(1)
		}

		splited := strings.Split(StoreLink, "/")
		storeName := splited[4]
		resp, err = http.Get(StoreLink)
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
		keyValStore[Key] = valueLink

		textContents, err := json.Marshal(keyValStore)
		if err != nil {
			fmt.Println(fmt.Errorf("error generating json"))
			fmt.Println(fmt.Errorf(err.Error()))
			os.Exit(1)
		}

		contents = []byte(textContents)
		extraParams = map[string]string{
			"uid":  UID,
			"link": StoreLink,
		}
		formDataContentType, request, err = UploadNoFile("https://wsend.net/update_cli", extraParams, storeName, contents)
		if err != nil {
			fmt.Println(fmt.Errorf(err.Error()))
			os.Exit(1)
		}
		request.Header.Add("Content-Type", formDataContentType)
		client = &http.Client{}
		resp, err = client.Do(request)
		if err != nil {
			fmt.Println(fmt.Errorf(err.Error()))
			os.Exit(1)
		}
		var bodyContent []byte
		resp.Body.Read(bodyContent)
		resp.Body.Close()
		os.Exit(0)
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
	storeCmd.Flags().StringVarP(&StoreLink, "store-link", "n", "", "link of the key value store")
	storeCmd.Flags().StringVarP(&Key, "key", "k", "", "name of the key")
	storeCmd.Flags().StringVarP(&Value, "value", "v", "", "either text or the name of a file")
	storeCmd.Flags().StringVarP(&Type, "type", "t", "string", "the type of value, defaults to string")
	storeCmd.Flags().StringVarP(&UID, "uid", "u", "", "access token")
	storeCmd.MarkFlagRequired("store-link")
	storeCmd.MarkFlagRequired("key")
	storeCmd.MarkFlagRequired("value")
}
