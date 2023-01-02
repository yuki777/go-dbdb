/*
Copyright © 2022 Yuki Adachi <yuki777@gmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var mongodbDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete mongodb server",
	Long:  `...`,
	Run: func(cmd *cobra.Command, args []string) {
		mongodbDelete(cmd)
	},
}

func init() {
	mongodbDeleteCmd.PersistentFlags().String("name", "", "Name for database (required)")
	mongodbDeleteCmd.MarkPersistentFlagRequired("name")
}
