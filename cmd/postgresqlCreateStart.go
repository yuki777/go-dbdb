/*
Copyright © 2022 Yuki Adachi <yuki777@gmail.com>
*/
package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var postgresqlCreateStartCmd = &cobra.Command{
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
		postgresqlCreateStart(cmd)
	},
}

func init() {
	postgresqlCreateStartCmd.PersistentFlags().String("name", "", "Name for database (required)")
	postgresqlCreateStartCmd.PersistentFlags().String("version", "", "Version for database (required)")
	postgresqlCreateStartCmd.PersistentFlags().String("port", "", "Port for database (required)")
	postgresqlCreateStartCmd.MarkPersistentFlagRequired("name")
	postgresqlCreateStartCmd.MarkPersistentFlagRequired("version")
	postgresqlCreateStartCmd.MarkPersistentFlagRequired("port")
}
