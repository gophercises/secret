package cmd

import (
	"Gophercizes/secret/students/jbimbert/vault"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete -k encodingKey -f fileVault keyname",
	Short: "Delete a key/value from the vault",
	Run: func(cmd *cobra.Command, args []string) {
		fv := &vault.FileVault{Key: encodingKey, File: filename}
		_, err := fv.Delete(args[0])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("key", args[0], "deleted")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
