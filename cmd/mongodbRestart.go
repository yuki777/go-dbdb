/*
Copyright © 2022 Yuki Adachi <yuki777@gmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var mongodbRestartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart mongodb server",
	Long:  `...`,
	Run: func(cmd *cobra.Command, args []string) {
		mongodbRestart(cmd)
	},
}

func init() {
	mongodbRestartCmd.PersistentFlags().String("name", "", "Name for database (required)")
	mongodbRestartCmd.MarkPersistentFlagRequired("name")
}
