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
	// rootCmd.AddCommand(mysqlCmd)
	mysqlCmd.AddCommand(mysqlCreateCmd)
	mysqlCmd.AddCommand(mysqlStartCmd)
	// mysqlCmd.AddCommand(stopCmd)
	// mysqlCmd.AddCommand(restartCmd)
	// mysqlCmd.AddCommand(statusCmd)
	// mysqlCmd.AddCommand(connectCmd)
	// mysqlCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mysqlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mysqlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
