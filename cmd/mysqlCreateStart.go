/*
Copyright © 2023 Yuki Adachi <yuki777@gmail.com>
*/
package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var mysqlCreateStartCmd = &cobra.Command{
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
		mysqlCreateStart(cmd)
	},
}

func init() {
	mysqlCreateStartCmd.PersistentFlags().String("name", "", "Name for database (required)")
	mysqlCreateStartCmd.PersistentFlags().String("version", "", "Version for database (required)")
	mysqlCreateStartCmd.PersistentFlags().String("port", "", "Port for database (required)")
	mysqlCreateStartCmd.MarkPersistentFlagRequired("name")
	mysqlCreateStartCmd.MarkPersistentFlagRequired("version")
	mysqlCreateStartCmd.MarkPersistentFlagRequired("port")
}
