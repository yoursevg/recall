package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// runCmd представляет команду выполнения
var runCmd = &cobra.Command{
	Use:   "run [command name]",
	Short: "Выполнение сохранённой команды",
	Long:  `Позволяет выполнять команды, сохранённые в хранилище, по их названию.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		commandName := args[0]
		//commands, err := storage.GetCommands()
		//if err != nil {
		//	fmt.Println("Error reading commands:", err)
		//	return
		//}

		fmt.Printf("Command '%s' not found.\n", commandName)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
