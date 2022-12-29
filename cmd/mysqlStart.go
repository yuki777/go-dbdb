/*
Copyright © 2022 Yuki Adachi <yuki777@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var mysqlStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start mysql database",
	Long:  `...`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mysqlStart called")
		// --defaults-file=
	},
}

func init() {
	//rootCmd.AddCommand(mysqlStartCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	mysqlStartCmd.PersistentFlags().String("name", "", "Name for database (required)")
	mysqlStartCmd.PersistentFlags().String("port", "", "Port for database")

	mysqlStartCmd.MarkPersistentFlagRequired("name")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mysqlStartCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
