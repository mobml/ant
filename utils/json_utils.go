package utils

import (
	"encoding/json"
	"github.com/mobml/ant/models"
	"os"
)

func ReadJSON(file string) ([]models.Task, error) {
	var tasks []models.Task

	if _, err := os.Stat(file); os.IsNotExist(err) {
		f, err := os.Create(file)
		if err != nil {
			return nil, err
		}
		f.Close()
		return tasks, nil
	}

	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func WriteJSON(file string, tasks []models.Task) error {
	data, err := json.Marshal(tasks)

	if err != nil {
		return err
	}

	err = os.WriteFile(file, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
