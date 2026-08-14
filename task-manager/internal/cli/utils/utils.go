package utils

import (
	"github.com/slkhvmxm/cli-task-manager/task-manager/internal/cli/models"
)

func GetMaxID(tasks []models.Task) int {
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
