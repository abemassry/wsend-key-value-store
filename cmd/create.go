/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// UID is the access token
var UID string

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new key value store",
	Long: `Initialize a key value store with a name
wkv create --name="foo"
optional pass in a uid
wkv create --name="foo" --uid="0123456789abcdef"

if uid is not passed in wkv will search for the uid in the wsend install path
incase wsend is installed
`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("create called")
		fmt.Println("store-name:")
		fmt.Println(StoreName)
		contents := []byte("{}")

		if UID == "" {
			homeDir, err := os.UserHomeDir()
			if err != nil {
				fmt.Println(fmt.Errorf("error reading homedir ~"))
				fmt.Println(fmt.Errorf(err.Error()))
				os.Exit(1)
			}
			wsendDir := homeDir + "/.wsend"
			uidPath := wsendDir + "/.id"
			id, err := os.ReadFile(uidPath)
			if err != nil {
				fmt.Println(fmt.Errorf("error reading id file"))
				fmt.Println(fmt.Errorf(err.Error()))
				os.Exit(1)
			}
			UID = string(id)
		}
		if UID == "" {
			fmt.Println(fmt.Errorf("Could not find uid"))
			os.Exit(1)
		}
		UID = strings.TrimSpace(UID)
		fmt.Println(UID)

		extraParams := map[string]string{
			"uid": UID,
		}
		formDataContentType, request, err := UploadNoFile("https://wsend.net/upload_cli", extraParams, StoreName, contents)
		if err != nil {
			fmt.Println(fmt.Errorf(err.Error()))
		}
		request.Header.Add("Content-Type", formDataContentType)
		client := &http.Client{}
		resp, err := client.Do(request)
		if err != nil {
			fmt.Println(fmt.Errorf(err.Error()))
		} else {
			var bodyContent []byte
			fmt.Println(resp.StatusCode)
			fmt.Println(resp.Header)
			resp.Body.Read(bodyContent)
			resp.Body.Close()
			fmt.Println(bodyContent)
		}

	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	createCmd.Flags().StringVarP(&StoreName, "name", "n", "", "name of the key value store")
	createCmd.Flags().StringVarP(&UID, "uid", "u", "", "access token")

}

// UploadNoFile uploads a file without the file existing on the filesystem
func UploadNoFile(uri string, params map[string]string, name string, contents []byte) (string, *http.Request, error) {

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	for key, val := range params {
		_ = writer.WriteField(key, val)
	}

	part, err := writer.CreateFormFile("filehandle", name)
	if err != nil {
		return "", nil, err
	}
	part.Write(contents)

	formDataContentType := writer.FormDataContentType()

	err = writer.Close()
	if err != nil {
		return "", nil, err
	}
	req, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return "", nil, err
	}

	return formDataContentType, req, nil
}
