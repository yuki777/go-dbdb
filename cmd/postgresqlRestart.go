/*
Copyright © 2023 Yuki Adachi <yuki777@gmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var postgresqlRestartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart postgresql server",
	Long:  `...`,
	Run: func(cmd *cobra.Command, args []string) {
		postgresqlRestart(cmd)
	},
}

func init() {
	postgresqlRestartCmd.PersistentFlags().String("name", "", "Name for database (required)")
	postgresqlRestartCmd.MarkPersistentFlagRequired("name")
}
