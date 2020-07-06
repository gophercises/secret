package cmd

import (
	"Gophercizes/secret/students/jbimbert/vault"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list -k encodingKey -f fileVault keyname",
	Short: "List all keys in the vault",
	Run: func(cmd *cobra.Command, args []string) {
		fv := &vault.FileVault{Key: encodingKey, File: filename}
		keys, err := fv.List()
		if err != nil {
			log.Fatal(err)
		}
		for _, k := range keys {
			fmt.Println(k)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
