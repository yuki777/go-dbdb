/*
Copyright © 2022 Yuki Adachi <yuki777@gmail.com>
*/
package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var redisCreateStartCmd = &cobra.Command{
	Use:   "create-start",
	Short: "Try create and start mysql server",
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
		redisStart(cmd)
	},
}

func init() {
	redisCreateStartCmd.PersistentFlags().String("name", "", "Name for database (required)")
	redisCreateStartCmd.PersistentFlags().String("version", "", "Version for database (required)")
	redisCreateStartCmd.PersistentFlags().String("port", "", "Port for database (required)")
	redisCreateStartCmd.MarkPersistentFlagRequired("name")
	redisCreateStartCmd.MarkPersistentFlagRequired("version")
	redisCreateStartCmd.MarkPersistentFlagRequired("port")
}
