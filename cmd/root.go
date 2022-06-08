/*
Copyright © 2022 Abe Massry a@abemassry.com

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wkv",
	Short: "wsend-key-value-store",
	Long: `wsend key value store is command line tool to store
a value based on a key some examples include:

wkv create --name="foo"
To create a store

wkv store --store-link="https://wsnd.io/IdGzDoh/foo" --key="bar" --value="baz" --type="string"
this will store the value "bar" at the key "foo" and it's of type string
which is the default if type was "file" then it would attempt to upload the
file specified. In either case a file always gets uploaded because the string
value can be very large and it makes more sense to be flexible.

wkv get --store-link="https://wsnd.io/IdGzDoh/foo" --key="bar"
will print the contents of the value to stdout

wkv remove --store-link="https://wsnd.io/IdGzDoh/foo" --key="bar"
to remove the key and associated value and file.
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.wsend-key-value-store.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
