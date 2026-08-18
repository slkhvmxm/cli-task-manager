package mark

import (
	"fmt"
	"strconv"
	"time"

	"github.com/slkhvmxm/cli-task-manager/task-manager/internal/cli/models"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "markTask [task_id] [task_status]",
		Short: "mark task status",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			var t models.Task

			if len(args) == 0 {
				return fmt.Errorf("task id and task description is required")
			}

			tasks, err := t.LoadTasks()
			if err != nil {
				return err
			}
			task_id, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			task_status := args[1]

			for i := range tasks {
				if tasks[i].ID == task_id {
					tasks[i].Status = task_status
					now := time.Now().UTC().Format("2006-01-02 15:04:05")
					tasks[i].UpdatedAt = now
				}
			}

			if err := t.SaveTask(tasks); err != nil {
				return fmt.Errorf("error creating JSON: %w", err)
			}

			return nil
		},
	}
	return cmd
}
