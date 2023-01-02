/*
Copyright © 2023 Yuki Adachi <yuki777@gmail.com>
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
		mysqlDelete(cmd)
	},
}

func init() {
	mysqlDeleteCmd.PersistentFlags().String("name", "", "Name for database (required)")
	mysqlDeleteCmd.MarkPersistentFlagRequired("name")
}
