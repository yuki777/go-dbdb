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
	log.Println(getCurrentFuncName(), "called")
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
	if exists(dataDir) {
		log.Println(dataDir + " directory is already exist")
		os.Exit(1)
	}

	if isRunningPort(optPort) {
		log.Println(optPort, "is already in use")
		os.Exit(1)
	}

	getUrlFileAs("https://dbdb.project8.jp/mysql/"+downloadFilePart+".tar.gz", downloadFilePart+".tar.gz")
	os.MkdirAll(dataDir, 0755)

	extractFile(versionDir, downloadFilePart)

	createCmd := exec.Command(
		versionDir+"/basedir/bin/mysqld",
		"--no-defaults",
		"--initialize-insecure",
		"--user="+dbUser,
		"--port="+optPort,
		"--socket="+dbSocket,
		"--basedir="+versionDir+"/basedir",
		"--plugin-dir="+versionDir+"/basedir/lib/plugin",
		"--datadir="+dataDir,
		"--log-error="+dataDir+"/mysqld.err",
		"--pid-file="+dataDir+"/mysql.pid",
	)

	log.Println("createCmd:", createCmd.String())
	createCmd.Run()

	portFile := dataDir + "/mysql.port.init"
	fileWrite(portFile, optPort)
	log.Println("mysql.port.init:", portFile)

	confFile := dataDir + "/my.cnf"
	fileWrite(confFile, "[mysqld]\nbind-address = 127.0.0.1\n")
	log.Println("my.cnf:", confFile)

	log.Println(optName, "MySQL database successfully created.")
	printUsage(optName, optVersion, optPort)
}

func mysqlStart(cmd *cobra.Command) {
	log.Println(getCurrentFuncName(), "called")
	dbdbBaseDir := dbdbBaseDir()

	optName := cmd.Flag("name").Value.String()

	dataDir := getDataDirByName(optName, "mysql")

	if notExists(dataDir) {
		log.Println(dataDir + " directory is NOT exist")
		os.Exit(1)
	}

	version := getVersionByDataDir(dataDir, optName, "mysql")

	dbPort := getPortByName(optName, "mysql")

	if isRunningPort(dbPort) {
		log.Println(dbPort, "is already in use")
		os.Exit(1)
	}

	versionDir := dbdbBaseDir + "/mysql/versions/" + version

	dbUser := "_dbdb_mysql"
	dbSocket := "/tmp/dbdb_mysql_" + dbPort + ".sock"

	startCmd := exec.Command(
		versionDir+"/basedir/bin/mysqld",
		"--defaults-file="+dataDir+"/my.cnf",
		"--daemonize",
		"--user="+dbUser,
		"--port="+dbPort,
		"--socket="+dbSocket,
		"--basedir="+versionDir+"/basedir",
		"--plugin-dir="+versionDir+"/basedir/lib/plugin",
		"--datadir="+dataDir,
		"--log-error="+dataDir+"/mysqld.err",
		"--pid-file="+dataDir+"/mysql.pid",
	)

	log.Println("startCmd:", startCmd.String())
	startCmd.Run()

	portFile := dataDir + "/mysql.port"
	log.Println("portFile:", portFile)
	fileWrite(portFile, dbPort)

	confFile := dataDir + "/my.cnf"
	log.Println("Your config file is located:", confFile)

	log.Println(optName, "MySQL database successfully started.")
}

func mysqlCreateStart(cmd *cobra.Command) {
	log.Println(getCurrentFuncName(), "called")
	mysqlCreate(cmd)
	mysqlStart(cmd)
}

func mysqlStop(cmd *cobra.Command, checkPort bool) {
	log.Println(getCurrentFuncName(), "called")
	dbdbBaseDir := dbdbBaseDir()

	optName := cmd.Flag("name").Value.String()

	dataDir := getDataDirByName(optName, "mysql")

	if notExists(dataDir) {
		log.Println(dataDir + " directory is NOT exist")
		os.Exit(1)
	}

	version := getVersionByDataDir(dataDir, optName, "mysql")

	dbPort := getPortByName(optName, "mysql")
	if checkPort && isNotRunningPort(dbPort) {
		log.Println(dbPort, "is NOT available")
		os.Exit(1)
	}

	dbSocket := "/tmp/dbdb_mysql_" + dbPort + ".sock"

	versionDir := dbdbBaseDir + "/mysql/versions/" + version

	stopCmd := exec.Command(
		versionDir+"/basedir/bin/mysqladmin",
		"--user=root",
		"--host=localhost",
		"--port="+dbPort,
		"--socket="+dbSocket,
		"shutdown",
	)
	log.Println("stopCmd", stopCmd.String())
	stopCmd.Run()

	copy(dataDir+"/mysql.port", dataDir+"/mysql.port.last")

	remove(dataDir + "/mysql.port")

	log.Println(optName, "MySQL database successfully stopped.")
}

func mysqlRestart(cmd *cobra.Command) {
	log.Println(getCurrentFuncName(), "called")
	mysqlStop(cmd, false)
	mysqlStart(cmd)
}

func mysqlDelete(cmd *cobra.Command) {
	log.Println(getCurrentFuncName(), "called")
	dbdbDelete(cmd, "mysql")
}
