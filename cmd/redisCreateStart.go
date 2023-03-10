/*
Copyright © 2023 Yuki Adachi <yuki777@gmail.com>
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
			log.Println("Error: Invalid arguments. use these characters strings [0-9a-zA-Z-_.] for `name`" + optName)
			cmd.Usage()
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		redisCreateStart(cmd)
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
