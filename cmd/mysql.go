/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var mysqlCmd = &cobra.Command{
	Use:   "mysql",
	Short: "..",
	Long:  `...`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Println("The name argument is required")
			cmd.Help()
			os.Exit(0)
		}
	},
}

func init() {
	mysqlCmd.AddCommand(mysqlCreateCmd)
	mysqlCmd.AddCommand(mysqlStartCmd)
	mysqlCmd.AddCommand(mysqlStopCmd)
	mysqlCmd.AddCommand(mysqlRestartCmd)
	mysqlCmd.AddCommand(mysqlDeleteCmd)
	mysqlCmd.AddCommand(mysqlCreateStartCmd)
}

func mysqlCreate(cmd *cobra.Command) {
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
}

func mysqlStart(cmd *cobra.Command) {
	dbdbBaseDir := dbdbBaseDir()

	optName := cmd.Flag("name").Value.String()

	dataDir := getDataDirByName(optName, "mysql")
	exitIfNotExistDir(dataDir)

	version := getVersionByDataDir(dataDir, optName, "mysql")

	dbPort := getPortByName(optName, "mysql")
	exitIfRunningPort(dbPort)

	versionDir := dbdbBaseDir + "/mysql/versions/" + version

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

	portFile := dataDir + "/mysql.port"
	log.Println("portFile:", portFile)
	fileWrite(portFile, dbPort)

	confFile := dataDir + "/my.cnf"
	log.Println("Your config file is located:", confFile)

	log.Println(optName, "MySQL database successfully started.")
}

func mysqlStop(cmd *cobra.Command, ignoreError bool) {
	dbdbBaseDir := dbdbBaseDir()

	optName := cmd.Flag("name").Value.String()

	dataDir := getDataDirByName(optName, "mysql")
	exitIfNotExistDir(dataDir)

	version := getVersionByDataDir(dataDir, optName, "mysql")

	dbPort := getPortByName(optName, "mysql")
	if !ignoreError {
		exitIfNotRunningPort(dbPort)
	}

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

	remove(dataDir + "/mysql.port")

	log.Println(optName, "MySQL database successfully stopped.")
}

func mysqlDelete(cmd *cobra.Command) {
	optName := cmd.Flag("name").Value.String()

	dataDir := getDataDirByName(optName, "mysql")
	exitIfNotExistDir(dataDir)

	dbPort := getPortByName(optName, "mysql")
	exitIfRunningPort(dbPort)

	remove(dataDir)
	log.Println("data directory deleted. ", dataDir)

	log.Println(optName, "MySQL database successfully deleted.")
}
