/*
Copyright © 2023 Yuki Adachi <yuki777@gmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var mysqlStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start mysql server",
	Long:  `...`,
	Run: func(cmd *cobra.Command, args []string) {
		mysqlStart(cmd)
	},
}

func init() {
	mysqlStartCmd.PersistentFlags().String("name", "", "Name for database (required)")
	mysqlStartCmd.MarkPersistentFlagRequired("name")
}
