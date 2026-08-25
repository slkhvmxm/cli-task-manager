package cmd

import (
	"os"

	addtask "github.com/slkhvmxm/cli-task-manager/task-manager/internal/cli/addTask"
	deletealltask "github.com/slkhvmxm/cli-task-manager/task-manager/internal/cli/deleteAllTask"
	deletetask "github.com/slkhvmxm/cli-task-manager/task-manager/internal/cli/deleteTask"
	listtask "github.com/slkhvmxm/cli-task-manager/task-manager/internal/cli/listTask"
	"github.com/slkhvmxm/cli-task-manager/task-manager/internal/cli/mark"
	updatetask "github.com/slkhvmxm/cli-task-manager/task-manager/internal/cli/updateTask"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cli-task-manager",
	Short: "CLI tool for managing your tasks",
	Long:  `You can add, update, delete and list all of your tasks. Also, you can check tasks status`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(addtask.NewCommand())
	rootCmd.AddCommand(updatetask.NewCommand())
	rootCmd.AddCommand(listtask.NewCommand())
	rootCmd.AddCommand(deletetask.NewCommand())
	rootCmd.AddCommand(deletealltask.NewCommand())
	rootCmd.AddCommand(mark.NewCommand())
}
