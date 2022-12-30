/*
Copyright © 2022 Yuki Adachi <yuki777@gmail.com>
*/
package cmd

import (
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

var mysqlStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start mysql server",
	Long:  `...`,
	Run: func(cmd *cobra.Command, args []string) {
		dbdbBaseDir := dbdbBaseDir()

		optName := cmd.Flag("name").Value.String()

		dataDir := getDataDirByName(optName, "mysql")
		exitIfNotExistDir(dataDir)

		version := getVersionByDataDir(dataDir, optName, "mysql")

		port := getPortByName(optName)
		exitIfRunningPort(port)

		versionDir := dbdbBaseDir + "/mysql/versions/" + version

		dbPort := getPortByName(optName)
		dbUser := "_dbdb_mysql"
		dbSocket := "/tmp/dbdb_mysql_" + dbPort + ".sock"

		mysqldCmd := exec.Command(
			versionDir+"/basedir/bin/mysqld",
			"--defaults-file="+dataDir+"/my.cnf",
			"--daemonize",
			"--user="+dbUser,
			"--port="+dbPort,
			"--socket="+dbSocket,
			"--basedir="+versionDir+"/basedir",
			"--plugin-dir="+versionDir+"/basedir/lib/plugin",
			"--datadir="+versionDir+"/datadir/"+optName,
			"--log-error="+versionDir+"/datadir/"+optName+"/mysqld.err",
			"--pid-file="+versionDir+"/datadir/"+optName+"/mysql.pid",
		)

		log.Println("mysqldCmd: " + mysqldCmd.String())
		mysqldCmd.Run()

		mysqlPortFile := dataDir + "/mysql.port"
		log.Println("mysqlPortFile:", mysqlPortFile)
		fileWrite(mysqlPortFile, dbPort)

		myConfFile := dataDir + "/my.cnf"
		log.Println("Your config file is located:", myConfFile)

		log.Println(optName, "MySQL database successfully started.")
	},
}

func init() {
	mysqlStartCmd.PersistentFlags().String("name", "", "Name for database (required)")
	mysqlStartCmd.MarkPersistentFlagRequired("name")
}
