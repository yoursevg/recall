package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// Определение корневой команды
var rootCmd = &cobra.Command{
	Use:   "recall", // Основное имя CLI
	Short: "Recall is a CLI tool for storing and recalling terminal commands",
	Long: `Recall allows you to save, search, and quickly execute your favorite terminal commands
without having to remember them all. It makes terminal work more efficient.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Recall! Use 'recall --help' to see available commands.")
	},
}

// Execute запускает rootCmd и обрабатывает ошибки
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("config", "c", "", "Config file (default is $HOME/.recall.yaml)")
}
