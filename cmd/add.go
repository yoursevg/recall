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
		var alias string

		for _, key := range args {
			builder.WriteString(key + " ")
		}
		command := builder.String()

		alias = cmd.Flags().Lookup("alias").Value.String()

		newCmd := models.Command{
			Command:     command,
			Alias:       alias,
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
	addCmd.Flags().StringP("alias", "a", "", "Save command for running in future: recall run [your alias]")
	addCmd.Flags().StringSliceP("tags", "t", []string{}, "Tags for command")
	rootCmd.AddCommand(addCmd)
}
