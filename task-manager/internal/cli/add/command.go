package add

import (
	"fmt"
	"time"

	"github.com/slkhvmxm/cli-task-manager/task-manager/internal/cli/models"
	"github.com/slkhvmxm/cli-task-manager/task-manager/internal/cli/utils"
	"github.com/spf13/cobra"
)

// addTaskCmd represents the addTask command

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "addTask [task name]",
		Short: "A brief description of your command",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var t models.Task

			// Take all exists tasks
			tasks, err := t.LoadTasks()
			if err != nil {
				return err
			}

			if len(args) == 0 {
				return fmt.Errorf("task description is required")
			}

			// create task with description take from args
			task := models.Task{
				ID:          utils.GetMaxID(tasks),
				Description: args[0],
				Status:      "created",
				CreatedAt:   time.Now(),
				UpdatedAt:   nil,
			}

			//append new task to tasks
			tasks = append(tasks, task)

			//write task struct into JSON
			if err := t.SaveTask(tasks); err != nil {
				return fmt.Errorf("error creating JSON: %w", err)
			}
			fmt.Printf("task added successfully (ID: %d)", task.ID)
			return nil
		},
	}
	return cmd
}
