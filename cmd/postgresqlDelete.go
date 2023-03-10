/*
Copyright © 2023 Yuki Adachi <yuki777@gmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var postgresqlDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete postgresql server",
	Long:  `...`,
	Run: func(cmd *cobra.Command, args []string) {
		postgresqlDelete(cmd)
	},
}

func init() {
	postgresqlDeleteCmd.PersistentFlags().String("name", "", "Name for database (required)")
	postgresqlDeleteCmd.MarkPersistentFlagRequired("name")
}
