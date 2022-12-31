/*
Copyright © 2022 Yuki Adachi <yuki777@gmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var mongodbStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start mongodb server",
	Long:  `...`,
	Run: func(cmd *cobra.Command, args []string) {
		mongodbStart(cmd)
	},
}

func init() {
	mongodbStartCmd.PersistentFlags().String("name", "", "Name for database (required)")
	mongodbStartCmd.MarkPersistentFlagRequired("name")
}
