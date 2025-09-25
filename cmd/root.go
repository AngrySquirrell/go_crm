package cmd

import (
	"fmt"
	"go_crm/internal/storage"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var store storage.Storer

var rootCmd = &cobra.Command{
	Use:   "go_crm",
	Short: "Un gestionnaire de contacts simple",
	Long:  "Mini-CRM CLI pour gérer vos contacts avec différents backends de stockage",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Erreur lors de la lecture du fichier de configuration: %v\n", err)
		os.Exit(1)
	}

	storageType := viper.GetString("storage.type")
	var err error

	switch storageType {
	case "gorm":
		dbPath := viper.GetString("storage.database_path")
		store, err = storage.NewGORMStore(dbPath)
	case "json":
		jsonPath := viper.GetString("storage.json_path")
		store = storage.NewJSONStoreWithPath(jsonPath)
	case "memory":
		store = storage.NewMemoryStore()
	default:
		fmt.Printf("Type de stockage non supporté: %s\n", storageType)
		os.Exit(1)
	}

	if err != nil {
		fmt.Printf("Erreur lors de l'initialisation du stockage: %v\n", err)
		os.Exit(1)
	}
}
