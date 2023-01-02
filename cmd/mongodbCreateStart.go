/*
Copyright © 2022 Yuki Adachi <yuki777@gmail.com>
*/
package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var mongodbCreateStartCmd = &cobra.Command{
	Use:   "create-start",
	Short: "Try create and start mongodb server",
	Long:  `...`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		optName := cmd.Flag("name").Value.String()
		if !validateOptName(optName) {
			log.Println("Error: Invalid arguments. use these characters strings [0-9a-zA-Z-_.] for `name`" + optName)
			cmd.Usage()
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		mongodbCreateStart(cmd)
	},
}

func init() {
	mongodbCreateStartCmd.PersistentFlags().String("name", "", "Name for database (required)")
	mongodbCreateStartCmd.PersistentFlags().String("version", "", "Version for database (required)")
	mongodbCreateStartCmd.PersistentFlags().String("port", "", "Port for database (required)")
	mongodbCreateStartCmd.MarkPersistentFlagRequired("name")
	mongodbCreateStartCmd.MarkPersistentFlagRequired("version")
	mongodbCreateStartCmd.MarkPersistentFlagRequired("port")
}
