/*
Copyright © 2022 Yuki Adachi <yuki777@gmail.com>
*/
package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var mongodbCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create mongodb server",
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
		mongodbCreate(cmd)
	},
}

func init() {
	mongodbCreateCmd.PersistentFlags().String("name", "", "Name for database (required)")
	mongodbCreateCmd.PersistentFlags().String("version", "", "Version for database (required)")
	mongodbCreateCmd.PersistentFlags().String("port", "", "Port for database (required)")
	mongodbCreateCmd.MarkPersistentFlagRequired("name")
	mongodbCreateCmd.MarkPersistentFlagRequired("version")
	mongodbCreateCmd.MarkPersistentFlagRequired("port")
}
