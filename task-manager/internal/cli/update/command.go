package update

import (
	"fmt"
	"strconv"
	"time"

	"github.com/slkhvmxm/cli-task-manager/task-manager/internal/cli/models"
	"github.com/spf13/cobra"
)

// updateTaskCmd represents the updateTask command
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "updateTask [task id] [task new description]",
		Short: "A brief description of your command",
		Args:  cobra.MaximumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {

			var t models.Task

			// Take all exists tasks
			tasks, err := t.LoadTasks()
			if err != nil {
				return err
			}

			if len(args) == 0 {
				return fmt.Errorf("task id and task description is required")
			}

			id, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			for i := range tasks {
				if tasks[i].ID == id {
					tasks[i].Description = args[1]
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
