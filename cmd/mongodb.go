/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var mongodbCmd = &cobra.Command{
	Use:   "mongodb",
	Short: "..",
	Long:  `...`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Println("The name argument is required")
			cmd.Help()
			os.Exit(1)
		}
		log.Println("mongodb called")
	},
}

func init() {
	mongodbCmd.AddCommand(mongodbCreateCmd)
	// mongodbCmd.AddCommand(mongodbStartCmd)
	// mongodbCmd.AddCommand(mongodbStopCmd)
	// mongodbCmd.AddCommand(mongodbRestartCmd)
	// mongodbCmd.AddCommand(mongodbStatusCmd)
	// mongodbCmd.AddCommand(mongodbConnectCmd)
	// mongodbCmd.AddCommand(mongodbDeleteCmd)
}

func mongodbCreate(cmd *cobra.Command) {

}
