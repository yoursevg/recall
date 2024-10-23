package cmd

import (
	"github.com/spf13/cobra"
	"strings"
)

// searchCmd представляет команду поиска
var searchCmd = &cobra.Command{
	Use:   "search [keyword]",
	Short: "Поиск команды по ключевому слову",
	Long:  `Позволяет искать сохранённые команды по ключевым словам, тегам или именам.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		//keyword := args[0]
		//commands, err := storage.GetCommands()
		//if err != nil {
		//	fmt.Println("Error reading commands:", err)
		//	return
		//}

		//found := false
		//for _, c := range commands {
		//	if contains(c.Name, keyword) || contains(c.Description, keyword) {
		//		fmt.Printf("Name: %s\nCommand: %s\nDescription: %s\nTags: %v\n\n", c.Name, c.Command, c.Description, c.Tags)
		//		found = true
		//	}
		//}

		//if !found {
		//	fmt.Println("No commands found for the given keyword.")
		//}
	},
}

func contains(source, keyword string) bool {
	return len(source) > 0 && (source == keyword || containsIgnoreCase(source, keyword))
}

func containsIgnoreCase(source, keyword string) bool {
	return len(source) >= len(keyword) && strings.Contains(source, keyword)
}

func init() {
	rootCmd.AddCommand(searchCmd) // Добавляем подкоманду search в rootCmd
}
