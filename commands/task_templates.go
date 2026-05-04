package commands

import "github.com/spf13/cobra"

var taskTemplatesCmd = &cobra.Command{
	Use:   "task-templates",
	Short: "",
}

func init() {
	rootCmd.AddCommand(taskTemplatesCmd)
}
