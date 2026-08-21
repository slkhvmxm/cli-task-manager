package listtask

import (
	"github.com/slkhvmxm/cli-task-manager/task-manager/internal/cli/models"
	"github.com/slkhvmxm/cli-task-manager/task-manager/internal/utils"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "listTask [task name]",
		Short: "A brief description of your command",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var t models.Task

			// Take all exists tasks
			tasks, err := t.LoadTasks()
			if err != nil {
				return err
			}

			var sortedTasks []models.Task

			if len(args) == 0 {
				utils.PrintTable(tasks)
				return nil
			}

			switch args[0] {
			case "done":
				for i := range tasks {
					if tasks[i].Status == "done" {
						sortedTasks = append(sortedTasks, tasks[i])
						utils.PrintTable(sortedTasks)
					}
				}
			case "in-progress":
				for i := range tasks {
					if tasks[i].Status == "in-progress" {
						sortedTasks = append(sortedTasks, tasks[i])
						utils.PrintTable(sortedTasks)
					}
				}
			case "comleted":
				for i := range tasks {
					if tasks[i].Status == "completed" {
						sortedTasks = append(sortedTasks, tasks[i])
						utils.PrintTable(sortedTasks)
					}
				}
			}
			return nil
		},
	}
	return cmd
}
