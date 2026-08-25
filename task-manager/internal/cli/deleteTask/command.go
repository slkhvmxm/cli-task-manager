package deletetask

import (
	"fmt"
	"strconv"

	"github.com/slkhvmxm/cli-task-manager/task-manager/internal/cli/models"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deleteTask [task name]",
		Short: "delete task by ID",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var t models.Task

			// Take all exists tasks
			tasks, err := t.LoadTasks()
			if err != nil {
				return err
			}

			id_task_to_delete, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			updated_tasks := deleteTask(tasks, id_task_to_delete)

			if err := t.SaveTask(updated_tasks); err != nil {
				return fmt.Errorf("error creating JSON: %w", err)
			}
			return nil
		},
	}
	return cmd
}

func deleteTask(tasks []models.Task, id int) []models.Task {

	result := make([]models.Task, 0, len(tasks))

	for _, task := range tasks {
		if task.ID != id {
			result = append(result, task)
		}
	}
	return result
}
