/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start called")

		currentDir := currentDir()
		fmt.Println("currentDir: " + currentDir)

		os := getOS()
		fmt.Println("os: " + os)

		optName := cmd.Flag("name").Value.String()
		optVersion := cmd.Flag("version").Value.String()
		optPort := cmd.Flag("port").Value.String()
		fmt.Println("optName: " + optName)
		fmt.Println("optVersion : " + optVersion)
		fmt.Println("optPort: " + optPort)

		fileName := "mysql-" + optVersion + "-" + os
		fmt.Println("fileName: " + fileName)

		dir := currentDir + "/versions/" + optVersion
		fmt.Println("dir: " + dir)

		// TODO...
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	startCmd.PersistentFlags().String("name", "", "Name for database (required)")
	startCmd.PersistentFlags().String("version", "", "Version for database (required)")
	startCmd.PersistentFlags().String("port", "", "Port for database (required)")

	startCmd.MarkPersistentFlagRequired("name")
	startCmd.MarkPersistentFlagRequired("version")
	startCmd.MarkPersistentFlagRequired("port")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
