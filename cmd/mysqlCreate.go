/*
Copyright © 2022 Yuki Adachi <yuki777@gmail.com>
*/
package cmd

import (
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var mysqlCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create mysql server",
	Long:  `...`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		optName := cmd.Flag("name").Value.String()
		if !validateOptName(optName) {
			log.Println("Error: Invalid arguments. use string, number and -_. for --name=" + optName)
			cmd.Usage()
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		dbdbBaseDir := dbdbBaseDir()

		optName := cmd.Flag("name").Value.String()
		optVersion := cmd.Flag("version").Value.String()
		optPort := cmd.Flag("port").Value.String()

		dbUser := "_dbdb_mysql"
		dbSocket := "/tmp/dbdb_mysql_" + optPort + ".sock"

		versionDir := dbdbBaseDir + "/mysql/versions/" + optVersion
		os.MkdirAll(versionDir, 0755)
		os.Chdir(versionDir)

		downloadFilePart := "mysql-" + optVersion + "-" + getOS()

		dataDir := versionDir + "/datadir/" + optName
		exitIfExistDir(dataDir)

		exitIfRunningPort(optPort)

		getUrlFileAs("https://dbdb.project8.jp/mysql/"+downloadFilePart+".tar.gz", downloadFilePart+".tar.gz")
		os.MkdirAll(dataDir, 0755)

		extractFile(versionDir, downloadFilePart)

		mysqldCmd := exec.Command(
			versionDir+"/basedir/bin/mysqld",
			"--no-defaults",
			"--initialize-insecure",
			"--user="+dbUser,
			"--port="+optPort,
			"--socket="+dbSocket,
			"--basedir="+versionDir+"/basedir",
			"--plugin-dir="+versionDir+"/basedir/lib/plugin",
			"--datadir="+versionDir+"/datadir/"+optName,
			"--log-error="+versionDir+"/datadir/"+optName+"/mysqld.err",
			"--pid-file="+versionDir+"/datadir/"+optName+"/mysql.pid",
		)

		log.Println("mysqldCmd: " + mysqldCmd.String())
		mysqldCmd.Run()

		portFile := versionDir + "/datadir/" + optName + "/mysql.port.init"
		fileWrite(portFile, optPort)
		log.Println("mysql.port.init:", portFile)

		confFile := versionDir + "/datadir/" + optName + "/my.cnf"
		fileWrite(confFile, "[mysqld]\nbind-address = 127.0.0.1\n")
		log.Println("my.cnf:", confFile)

		log.Println(optName, "MySQL database successfully created.")
		printUsage(optName, optVersion, optPort)
	},
}

func init() {
	mysqlCreateCmd.PersistentFlags().String("name", "", "Name for database (required)")
	mysqlCreateCmd.PersistentFlags().String("version", "", "Version for database (required)")
	mysqlCreateCmd.PersistentFlags().String("port", "", "Port for database (required)")

	mysqlCreateCmd.MarkPersistentFlagRequired("name")
	mysqlCreateCmd.MarkPersistentFlagRequired("version")
	mysqlCreateCmd.MarkPersistentFlagRequired("port")
}
