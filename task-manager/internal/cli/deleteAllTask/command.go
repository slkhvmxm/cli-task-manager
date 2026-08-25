package deletealltask

import (
	"github.com/slkhvmxm/cli-task-manager/task-manager/internal/cli/models"
	"github.com/slkhvmxm/cli-task-manager/task-manager/internal/utils"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deleteAll [task name]",
		Short: "delete all tasks",
		RunE: func(cmd *cobra.Command, args []string) error {
			var t models.Task
			var tasks []models.Task

			if err := t.SaveTask(tasks); err != nil {
				return err
			}

			utils.PrintTable(tasks)
			return nil
		},
	}
	return cmd
}
