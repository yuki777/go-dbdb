/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// mysqlCmd represents the mysql command
var mysqlCmd = &cobra.Command{
	Use:   "mysql",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("The name argument is required")
			cmd.Help()
			os.Exit(0)
		}
		fmt.Println("mysql called")
	},
}

func init() {
	// rootCmd.AddCommand(mysqlCmd)
	mysqlCmd.AddCommand(createCmd)
	// mysqlCmd.AddCommand(startCmd)
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
