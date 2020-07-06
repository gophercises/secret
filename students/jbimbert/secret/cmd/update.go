package cmd

import (
	"Gophercizes/secret/students/jbimbert/vault"
	"log"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update -k encodingKey -f fileVault keyname keyvalue",
	Short: "Update the value of an existing key in the vault",
	Run: func(cmd *cobra.Command, args []string) {
		fv := &vault.FileVault{Key: encodingKey, File: filename}
		err := fv.Update(args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
