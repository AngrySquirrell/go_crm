package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lister tous les contacts",
	Run: func(cmd *cobra.Command, args []string) {
		contacts := store.GetAll()

		if len(contacts) == 0 {
			fmt.Println("Aucun contact trouvé")
			return
		}

		fmt.Println("Liste des contacts:")
		for _, contact := range contacts {
			contact.Display()
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
