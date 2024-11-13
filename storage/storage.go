package storage

import (
	"encoding/json"
	"errors"
	"os"
	"os/user"
	"path/filepath"
	"recall/models"
)

// Функция для получения пути к файлу commands.json в папке .recall
func getFilePath() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	recallDir := filepath.Join(usr.HomeDir, ".recall")

	// Проверка и создание папки .recall, если она не существует
	if _, err := os.Stat(recallDir); os.IsNotExist(err) {
		err := os.Mkdir(recallDir, 0755)
		if err != nil {
			return "", err
		}
	}

	return filepath.Join(recallDir, "commands.json"), nil
}

// SaveCommand сохраняет команду в файл
func SaveCommand(cmd models.Command) error {
	filePath, err := getFilePath()
	if err != nil {
		return err
	}

	var commands []models.Command
	file, err := os.ReadFile(filePath)
	if err == nil {
		err = json.Unmarshal(file, &commands)
		if err != nil {
			return err
		}
	}

	commands = append(commands, cmd)

	fileData, err := json.MarshalIndent(commands, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, fileData, 0644)
}

// GetCommands возвращает список всех команд
func GetCommands() ([]models.Command, error) {
	filePath, err := getFilePath()
	if err != nil {
		return nil, err
	}

	var commands []models.Command
	file, err := os.ReadFile(filePath)
	if err != nil {
		return commands, errors.New("could not read file")
	}

	err = json.Unmarshal(file, &commands)
	if err != nil {
		return nil, err
	}

	return commands, nil
}

func FindCommandByAlias(alias string) (string, error) {
	filePath, err := getFilePath()
	if err != nil {
		return "", err
	}

	var commands []models.Command
	file, err := os.ReadFile(filePath)
	if err == nil {
		err = json.Unmarshal(file, &commands)
		if err != nil {
			return "", err
		}
	}

	for _, c := range commands {
		if c.Alias == alias {
			return c.Command, nil
		}
	}

	return "", errors.New("command not found")
}
