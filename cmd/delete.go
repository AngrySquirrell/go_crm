package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Supprimer un contact",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("ID invalide")
			return
		}

		if _, exists := store.GetByID(id); !exists {
			fmt.Printf("Contact avec l'ID %d non trouvé\n", id)
			return
		}

		if err := store.Remove(id); err != nil {
			fmt.Printf("Erreur lors de la suppression: %v\n", err)
			return
		}

		fmt.Printf("Contact %d supprimé avec succès\n", id)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
