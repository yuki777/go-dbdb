/*
Copyright © 2022 Yuki Adachi <yuki777@gmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var postgresqlStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop postgresql server",
	Long:  `...`,
	Run: func(cmd *cobra.Command, args []string) {
		postgresqlStop(cmd, true)
	},
}

func init() {
	postgresqlStopCmd.PersistentFlags().String("name", "", "Name for database (required)")
	postgresqlStopCmd.MarkPersistentFlagRequired("name")
}
