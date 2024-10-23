package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"recall/models"
	"recall/storage"
	"strings"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Добавляет команду в хранилище",
	Run: func(cmd *cobra.Command, args []string) {
		var builder strings.Builder
		for _, key := range args {
			builder.WriteString(key + " ")
		}
		command := builder.String()

		newCmd := models.Command{
			Command:     command,
			Description: "New command",
		}

		err := storage.SaveCommand(newCmd)
		if err != nil {
			fmt.Println("Error saving command:", err)
		} else {
			fmt.Println("Command added successfully!")
		}
	},
}

func init() {
	addCmd.Flags().StringP("name", "n", "", "Имя команды")
	addCmd.Flags().StringSliceP("tags", "t", []string{}, "Теги для команды")
	rootCmd.AddCommand(addCmd)
}
