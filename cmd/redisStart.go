/*
Copyright © 2022 Yuki Adachi <yuki777@gmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var redisStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start redis server",
	Long:  `...`,
	Run: func(cmd *cobra.Command, args []string) {
		redisStart(cmd)
	},
}

func init() {
	redisStartCmd.PersistentFlags().String("name", "", "Name for database (required)")
	redisStartCmd.MarkPersistentFlagRequired("name")
}
