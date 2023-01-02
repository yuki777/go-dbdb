/*
Copyright © 2022 Yuki Adachi <yuki777@gmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var redisDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete redis server",
	Long:  `...`,
	Run: func(cmd *cobra.Command, args []string) {
		redisDelete(cmd)
	},
}

func init() {
	redisDeleteCmd.PersistentFlags().String("name", "", "Name for database (required)")
	redisDeleteCmd.MarkPersistentFlagRequired("name")
}
