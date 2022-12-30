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
		log.Println("dbdbBaseDir: " + dbdbBaseDir)

		optName := cmd.Flag("name").Value.String()
		log.Println("optName: " + optName)

		dataDir := getDataDirByName(optName, "mysql")
		log.Println("dataDir:", dataDir)
		exitIfNotExistDir(dataDir)

		version := getVersionByDataDir(dataDir, optName, "mysql")
		log.Println("version:", version)

		dbPort := getPortByName(optName)
		log.Println("dbPort:", dbPort)
		exitIfNotRunningPort(dbPort)

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
