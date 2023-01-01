/*
Copyright © 2022 Yuki Adachi <yuki777@gmail.com>
*/
package cmd

import (
	"time"

	"github.com/spf13/cobra"
)

var mongodbDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete mongodb server",
	Long:  `...`,
	Run: func(cmd *cobra.Command, args []string) {
		mongodbStop(cmd, false)
		time.Sleep(1 * time.Second) // Waiting for the port to close
		mongodbDelete(cmd)
	},
}

func init() {
	mongodbDeleteCmd.PersistentFlags().String("name", "", "Name for database (required)")
	mongodbDeleteCmd.MarkPersistentFlagRequired("name")
}
