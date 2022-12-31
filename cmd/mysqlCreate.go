/*
Copyright © 2022 Yuki Adachi <yuki777@gmail.com>
*/
package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var mysqlCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create mysql server",
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
		mysqlCreate(cmd)
	},
}

func init() {
	mysqlCreateCmd.PersistentFlags().String("name", "", "Name for database (required)")
	mysqlCreateCmd.PersistentFlags().String("version", "", "Version for database (required)")
	mysqlCreateCmd.PersistentFlags().String("port", "", "Port for database (required)")
	mysqlCreateCmd.MarkPersistentFlagRequired("name")
	mysqlCreateCmd.MarkPersistentFlagRequired("version")
	mysqlCreateCmd.MarkPersistentFlagRequired("port")
}
