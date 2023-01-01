/*
Copyright © 2022 Yuki Adachi <yuki777@gmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var mysqlStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop mysql server",
	Long:  `...`,
	Run: func(cmd *cobra.Command, args []string) {
		mysqlStop(cmd, true)
	},
}

func init() {
	mysqlStopCmd.PersistentFlags().String("name", "", "Name for database (required)")
	mysqlStopCmd.MarkPersistentFlagRequired("name")
}
