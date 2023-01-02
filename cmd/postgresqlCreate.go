/*
Copyright © 2022 Yuki Adachi <yuki777@gmail.com>
*/
package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var postgresqlCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create postgresql server",
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
		postgresqlCreate(cmd)
	},
}

func init() {
	postgresqlCreateCmd.PersistentFlags().String("name", "", "Name for database (required)")
	postgresqlCreateCmd.PersistentFlags().String("version", "", "Version for database (required)")
	postgresqlCreateCmd.PersistentFlags().String("port", "", "Port for database (required)")
	postgresqlCreateCmd.MarkPersistentFlagRequired("name")
	postgresqlCreateCmd.MarkPersistentFlagRequired("version")
	postgresqlCreateCmd.MarkPersistentFlagRequired("port")
}
