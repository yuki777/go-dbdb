/*
Copyright © 2023 Yuki Adachi <yuki777@gmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var redisStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop redis server",
	Long:  `...`,
	Run: func(cmd *cobra.Command, args []string) {
		redisStop(cmd, true)
	},
}

func init() {
	redisStopCmd.PersistentFlags().String("name", "", "Name for database (required)")
	redisStopCmd.MarkPersistentFlagRequired("name")
}
