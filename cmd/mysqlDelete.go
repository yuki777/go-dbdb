/*
Copyright © 2022 Yuki Adachi <yuki777@gmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var mysqlDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete mysql server",
	Long:  `...`,
	Run: func(cmd *cobra.Command, args []string) {
		mysqlStop(cmd, false)
		mysqlDelete(cmd)
	},
}

func init() {
	mysqlDeleteCmd.PersistentFlags().String("name", "", "Name for database (required)")
	mysqlDeleteCmd.MarkPersistentFlagRequired("name")
}
