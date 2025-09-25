package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update [id]",
	Short: "Mettre à jour un contact",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("ID invalide")
			return
		}

		name, _ := cmd.Flags().GetString("name")
		email, _ := cmd.Flags().GetString("email")

		if name == "" && email == "" {
			fmt.Println("Au moins un champ (nom ou email) doit être fourni")
			return
		}

		if _, exists := store.GetByID(id); !exists {
			fmt.Printf("Contact avec l'ID %d non trouvé\n", id)
			return
		}

		if err := store.Update(id, name, email); err != nil {
			fmt.Printf("Erreur lors de la mise à jour: %v\n", err)
			return
		}

		fmt.Printf("Contact %d mis à jour avec succès\n", id)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringP("name", "n", "", "Nouveau nom du contact")
	updateCmd.Flags().StringP("email", "e", "", "Nouveau email du contact")
}
