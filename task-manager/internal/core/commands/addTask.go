package cli_task_manager

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// JSON with all tasks
const taskFile = "tasks.json"

// list for all tasks
var tasks []Task

// addTaskCmd represents the addTask command
var addTaskCmd = &cobra.Command{
	Use:   "addTask [task name]",
	Short: "A brief description of your command",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		// Take all exists tasks
		loadTasks()

		if len(args) == 0 {
			return fmt.Errorf("task description is required")
		}

		// create task with description take from args
		task := Task{
			ID:          getMaxID(),
			Description: args[0],
			Status:      "created",
			CreatedAt:   time.Now(),
			UpdatedAt:   nil,
		}

		//append new task to tasks
		tasks = append(tasks, task)

		//write task struct into JSON
		if err := saveTask(); err != nil {
			return fmt.Errorf("error creating JSON: %w", err)
		}
		fmt.Printf("task added successfully (ID: %d)", task.ID)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addTaskCmd)
}

func getMaxID() int {
	if len(tasks) == 0 {
		return 1
	}
	maxID := 0
	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}

	return maxID + 1
}

func loadTasks() {
	data, err := os.ReadFile(taskFile)
	if err != nil {
		tasks = []Task{}
		return
	}
	json.Unmarshal(data, &tasks)
}

func saveTask() error {
	jsonData, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(taskFile, jsonData, 0644)
}
