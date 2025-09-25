package cmd

import (
	"fmt"
	"go_crm/internal/storage"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Ajouter un nouveau contact",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		email, _ := cmd.Flags().GetString("email")

		if name == "" || email == "" {
			fmt.Println("Les champs nom et email sont obligatoires")
			return
		}

		contact := storage.NewContact(store.GetNextID(), name, email)
		if err := store.Add(contact); err != nil {
			fmt.Printf("Erreur lors de l'ajout du contact: %v\n", err)
			return
		}

		fmt.Printf("Contact ajouté avec succès: %s (%s)\n", name, email)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("name", "n", "", "Nom du contact")
	addCmd.Flags().StringP("email", "e", "", "Email du contact")
	addCmd.MarkFlagRequired("name")
	addCmd.MarkFlagRequired("email")
}
