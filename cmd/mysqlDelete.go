/*
Copyright © 2022 Yuki Adachi <yuki777@gmail.com>
*/
package cmd

import (
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

var mysqlDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete mysql server",
	Long:  `...`,
	Run: func(cmd *cobra.Command, args []string) {
		dbdbBaseDir := dbdbBaseDir()

		optName := cmd.Flag("name").Value.String()
		log.Println("optName: " + optName)

		dataDir := getDataDirByName(optName, "mysql")
		log.Println("dataDir:", dataDir)
		exitIfNotExistDir(dataDir)

		version := getVersionByDataDir(dataDir, optName, "mysql")
		log.Println("version:", version)

		dbPort := getPortByName(optName)
		log.Println("dbPort:", dbPort)

		dbSocket := "/tmp/dbdb_mysql_" + dbPort + ".sock"

		versionDir := dbdbBaseDir + "/mysql/versions/" + version
		log.Println("versionDir:", versionDir)

		mysqlAdminCmd := exec.Command(
			versionDir+"/basedir/bin/mysqladmin",
			"--user=root",
			"--host=localhost",
			"--port="+dbPort,
			"--socket="+dbSocket,
			"shutdown",
		)
		log.Println("mysqldCmd: " + mysqlAdminCmd.String())
		mysqlAdminCmd.Run()

		exitIfNotExistDir(dataDir)
		exitIfRunningPort(dbPort)

		removeDir(dataDir)
		log.Println("data directory deleted. ", dataDir)

		log.Println(optName, "MySQL database successfully deleted.")
	},
}

func init() {
	//rootCmd.AddCommand(mysqlCreateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	mysqlDeleteCmd.PersistentFlags().String("name", "", "Name for database (required)")

	mysqlDeleteCmd.MarkPersistentFlagRequired("name")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mysqlCreateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
