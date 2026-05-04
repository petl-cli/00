package commands

import "github.com/spf13/cobra"

var auditLogApiCmd = &cobra.Command{
	Use:   "audit-log-api",
	Short: "",
}

func init() {
	rootCmd.AddCommand(auditLogApiCmd)
}
