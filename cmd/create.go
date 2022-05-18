/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new key value store",
	Long: `Initialize a key value store with a name
wkv create --name="foo"
`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("create called")
		fmt.Println("store-name:")
		fmt.Println(StoreName)
		contents := []byte("{}")
		extraParams := map[string]string{
			"uid": "uid",
		}
		request, err := newfileUploadRequest("https://wsend.net/upload_cli", extraParams, StoreName, contents)
		if err != nil {
			fmt.Fatal(err)
		}
		client := &http.Client{}
		resp, err := client.Do(request)
		if err != nil {
			fmt.Fatal(err)
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
}

// UploadNoFile uploads a file without the file existing on the filesystem
func UploadNoFile(uri string, params map[string]string, name string, contents []byte) (*http.Request, error) {

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("filehandle", name)
	if err != nil {
		return nil, err
	}
	part.Write(contents)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	return http.NewRequest("POST", uri, body)
}
