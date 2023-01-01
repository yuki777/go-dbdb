/*
Copyright © 2022 Yuki Adachi <yuki777@gmail.com>
*/
package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var redisCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create redis server",
	Long:  `...`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		optName := cmd.Flag("name").Value.String()
		if !validateOptName(optName) {
			log.Println("Error: Invalid arguments. use string, number and -_. for --name=" + optName)
			cmd.Usage()
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		redisCreate(cmd)
	},
}

func init() {
	redisCreateCmd.PersistentFlags().String("name", "", "Name for database (required)")
	redisCreateCmd.PersistentFlags().String("version", "", "Version for database (required)")
	redisCreateCmd.PersistentFlags().String("port", "", "Port for database (required)")
	redisCreateCmd.MarkPersistentFlagRequired("name")
	redisCreateCmd.MarkPersistentFlagRequired("version")
	redisCreateCmd.MarkPersistentFlagRequired("port")
}
