/*
Copyright © 2022 Yuki Adachi <yuki777@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

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

		myOS := getOS()
		fmt.Println("myOS: " + myOS)

		optName := cmd.Flag("name").Value.String()
		optVersion := cmd.Flag("version").Value.String()
		optPort := cmd.Flag("port").Value.String()
		fmt.Println("optName: " + optName)
		fmt.Println("optVersion : " + optVersion)
		fmt.Println("optPort: " + optPort)

		dbUser := "_dbdb_mysql"
		fmt.Println("dbUser: " + dbUser)
		dbSocket := "/tmp/dbdb_mysql_" + optPort + ".sock"
		fmt.Println("dbSocket: " + dbSocket)

		dir := currentDir + "/versions/" + optVersion
		fmt.Println("dir: " + dir)

		os.MkdirAll(dir, 0755)

		beforeDir, err := os.Getwd()
		if err != nil {
		}
		fmt.Println("Before directory: " + beforeDir)

		err = os.Chdir(dir)
		if err != nil {
			fmt.Println(err)
		}

		afterDir, err := os.Getwd()
		if err != nil {
		}
		fmt.Println("After directory: " + afterDir)

		downloadFilePart := "mysql-" + optVersion + "-" + myOS
		fmt.Println("downloadFilePart: " + downloadFilePart)

		checkDir := dir + "/datadir/" + optName
		fmt.Println("checkDir: " + checkDir)
		exitIfExistDir(checkDir)

		exitIfRunningPort(optPort)

		// TODO
		getUrlFileAs("https://dbdb.project8.jp/mysql/"+downloadFilePart+".tar.gz", downloadFilePart+".tar.gz")
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
