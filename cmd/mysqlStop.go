/*
Copyright © 2022 Yuki Adachi <yuki777@gmail.com>
*/
package cmd

import (
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

var mysqlStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop mysql server",
	Long:  `...`,
	Run: func(cmd *cobra.Command, args []string) {
		dbdbBaseDir := dbdbBaseDir()

		optName := cmd.Flag("name").Value.String()

		dataDir := getDataDirByName(optName, "mysql")
		exitIfNotExistDir(dataDir)

		version := getVersionByDataDir(dataDir, optName, "mysql")

		dbPort := getPortByName(optName)
		exitIfNotRunningPort(dbPort)

		dbSocket := "/tmp/dbdb_mysql_" + dbPort + ".sock"

		versionDir := dbdbBaseDir + "/mysql/versions/" + version

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

		source := dataDir + "/mysql.port"
		dest := dataDir + "/mysql.port.last"
		copyFile(source, dest)

		removeDir(dataDir + "/mysql.port")

		log.Println(optName, "MySQL database successfully stopped.")
	},
}

func init() {
	mysqlStopCmd.PersistentFlags().String("name", "", "Name for database (required)")
	mysqlStopCmd.MarkPersistentFlagRequired("name")
}
