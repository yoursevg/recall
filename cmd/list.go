package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"recall/storage"
)

// listCmd представляет команду вывода всех команд
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Вывод всех сохранённых команд",
	Long:  `Позволяет выводить список всех команд, сохранённых в хранилище.`,
	Run: func(cmd *cobra.Command, args []string) {
		commands, err := storage.GetCommands()
		if err != nil {
			fmt.Println("Error reading commands:", err)
			return
		}

		if len(commands) == 0 {
			fmt.Println("No commands found.")
			return
		}

		fmt.Println("Saved Commands:")
		for _, c := range commands {
			if c.Alias != "" {
				fmt.Printf("Command: %s\nDescription: %s\nAlias: %s\n", c.Command, c.Description, c.Alias)
			} else {
				fmt.Printf("Command: %s\nDescription: %s\n", c.Command, c.Description)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
