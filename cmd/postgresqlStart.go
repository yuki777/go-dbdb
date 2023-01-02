/*
Copyright © 2023 Yuki Adachi <yuki777@gmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var postgresqlStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start postgresql server",
	Long:  `...`,
	Run: func(cmd *cobra.Command, args []string) {
		postgresqlStart(cmd)
	},
}

func init() {
	postgresqlStartCmd.PersistentFlags().String("name", "", "Name for database (required)")
	postgresqlStartCmd.MarkPersistentFlagRequired("name")
}
