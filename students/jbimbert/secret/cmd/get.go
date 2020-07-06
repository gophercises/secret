package cmd

import (
	"fmt"
	"log"

	"Gophercizes/secret/students/jbimbert/vault"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get -k encodingKey -f fileVault keyname",
	Short: "Get a key stored in the vault",
	Run: func(cmd *cobra.Command, args []string) {
		fv := &vault.FileVault{Key: encodingKey, File: filename}
		value, err := fv.Get(args[0])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(value)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
