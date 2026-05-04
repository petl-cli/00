package commands

import "github.com/spf13/cobra"

var userTaskListsCmd = &cobra.Command{
	Use:   "user-task-lists",
	Short: "",
}

func init() {
	rootCmd.AddCommand(userTaskListsCmd)
}
