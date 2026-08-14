package models

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// JSON with all tasks
const taskFile = "tasks.json"

type Task struct {
	ID          int        `json:"id"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

func (t *Task) UpdateTask(description string) {
	t.Description = description
}

func (t Task) LoadTasks() ([]Task, error) {

	var tasks []Task

	if _, err := os.Stat(taskFile); os.IsNotExist(err) {
		emptyTasks := []Task{}
		data, _ := json.MarshalIndent(tasks, "", " ")
		err = os.WriteFile(taskFile, data, 0644)
		if err != nil {
			return tasks, fmt.Errorf("file %s does not exist", taskFile)
		}
		return emptyTasks, nil
	}

	data, err := os.ReadFile(taskFile)
	if err != nil {
		return tasks, fmt.Errorf("smth wrong: %w", err)
	}
	json.Unmarshal(data, &tasks)
	return tasks, nil
}

func (t Task) SaveTask(tasks []Task) error {
	jsonData, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(taskFile, jsonData, 0644)
}
