package cmd

import (
	"Gophercizes/secret/students/jbimbert/vault"
	"log"

	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set -k encodingKey -f fileVault keyname keyvalue",
	Short: "Store a key/value in the vault",
	Run: func(cmd *cobra.Command, args []string) {
		fv := &vault.FileVault{Key: encodingKey, File: filename}
		err := fv.Set(args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
