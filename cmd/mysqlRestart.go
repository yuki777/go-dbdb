/*
Copyright © 2022 Yuki Adachi <yuki777@gmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var mysqlRestartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart mysql server",
	Long:  `...`,
	Run: func(cmd *cobra.Command, args []string) {
		mysqlRestart(cmd)
	},
}

func init() {
	mysqlRestartCmd.PersistentFlags().String("name", "", "Name for database (required)")
	mysqlRestartCmd.MarkPersistentFlagRequired("name")
}
