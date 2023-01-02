/*
Copyright © 2022 Yuki Adachi <yuki777@gmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var redisRestartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart redis server",
	Long:  `...`,
	Run: func(cmd *cobra.Command, args []string) {
		redisRestart(cmd)
	},
}

func init() {
	redisRestartCmd.PersistentFlags().String("name", "", "Name for database (required)")
	redisRestartCmd.MarkPersistentFlagRequired("name")
}
