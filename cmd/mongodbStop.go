/*
Copyright © 2022 Yuki Adachi <yuki777@gmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var mongodbStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop mongodb server",
	Long:  `...`,
	Run: func(cmd *cobra.Command, args []string) {
		mongodbStop(cmd, false)
	},
}

func init() {
	mongodbStopCmd.PersistentFlags().String("name", "", "Name for database (required)")
	mongodbStopCmd.MarkPersistentFlagRequired("name")
}
