/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

// mysqlCmd represents the mysql command
var mysqlCmd = &cobra.Command{
	Use:   "mysql",
	Short: "..",
	Long:  `...`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Println("The name argument is required")
			cmd.Help()
			os.Exit(0)
		}
		log.Println("mysql called")
	},
}

func init() {
	mysqlCmd.AddCommand(mysqlCreateCmd)
	mysqlCmd.AddCommand(mysqlStartCmd)
	mysqlCmd.AddCommand(mysqlStopCmd)
	// mysqlCmd.AddCommand(mysqlRestartCmd)
	// mysqlCmd.AddCommand(mysqlStatusCmd)
	// mysqlCmd.AddCommand(mysqlConnectCmd)
	mysqlCmd.AddCommand(mysqlDeleteCmd)
}
