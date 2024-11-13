package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"recall/storage"
	"strings"
)

// runCmd представляет команду выполнения
var runCmd = &cobra.Command{
	Use:               "run [alias]",
	Short:             "Execute saved command",
	ValidArgsFunction: getCommandSuggestions,
	Long:              `Allows you to execute the storage saved command by their alias`,
	Args:              cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		alias := args[0]
		commandToExecute, err := storage.FindCommandByAlias(alias)
		if err != nil {
			return
		}

		// Разбиваем команду на аргументы
		cmdArgs := strings.Fields(commandToExecute)
		if len(cmdArgs) == 0 {
			fmt.Println("Empty command")
			return
		}

		// Создаем команду
		command := exec.Command(cmdArgs[0], cmdArgs[1:]...)

		// Привязываем stdout и stderr к текущему терминалу
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
		command.Stdin = os.Stdin

		// Запускаем команду
		if err := command.Run(); err != nil {
			fmt.Printf("Failed to execute command: %v\n", err)
		}
	},
}

func getCommandSuggestions(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	commands, err := storage.GetCommands()
	if err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	var suggestions []string
	for _, command := range commands {
		if command.Alias != "" {
			suggestions = append(suggestions, command.Alias)
		}
	}

	return suggestions, cobra.ShellCompDirectiveNoFileComp
}

func init() {
	rootCmd.AddCommand(runCmd)
}
