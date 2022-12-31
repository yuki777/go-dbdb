/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"
	"os/exec"
	"syscall"

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
			os.Exit(0)
		}
	},
}

func init() {
	mongodbCmd.AddCommand(mongodbCreateCmd)
	mongodbCmd.AddCommand(mongodbStartCmd)
	mongodbCmd.AddCommand(mongodbStopCmd)
	mongodbCmd.AddCommand(mongodbRestartCmd)
	mongodbCmd.AddCommand(mongodbDeleteCmd)
	mongodbCmd.AddCommand(mongodbCreateStartCmd)
}

func mongodbCreate(cmd *cobra.Command) {
	dbdbBaseDir := dbdbBaseDir()

	optName := cmd.Flag("name").Value.String()
	optVersion := cmd.Flag("version").Value.String()
	optPort := cmd.Flag("port").Value.String()

	versionDir := dbdbBaseDir + "/mongodb/versions/" + optVersion
	os.MkdirAll(versionDir, 0755)
	os.Chdir(versionDir)

	downloadFilePart := "mongodb-" + optVersion + "-" + getOS()

	dataDir := versionDir + "/datadir/" + optName
	exitIfExistDir(dataDir)

	exitIfRunningPort(optPort)

	getUrlFileAs("https://dbdb.project8.jp/mongodb/"+downloadFilePart+".tar.gz", downloadFilePart+".tar.gz")
	os.MkdirAll(dataDir, 0755)

	extractFile(versionDir, downloadFilePart)

	portFile := versionDir + "/datadir/" + optName + "/mongodb.port.init"
	fileWrite(portFile, optPort)
	log.Println("mongodb.port.init:", portFile)

	confFile := versionDir + "/datadir/" + optName + "/mongod.conf"
	fileWrite(confFile, "#mongod.conf\n")
	log.Println("mongod.conf:", confFile)

	log.Println(optName, "MongoDB database successfully created.")
	printUsage(optName, optVersion, optPort)
}

func mongodbStart(cmd *cobra.Command) {
	dbdbBaseDir := dbdbBaseDir()

	optName := cmd.Flag("name").Value.String()

	dataDir := getDataDirByName(optName, "mongodb")
	exitIfNotExistDir(dataDir)

	version := getVersionByDataDir(dataDir, optName, "mongodb")

	dbPort := getPortByName(optName, "mongodb")
	exitIfRunningPort(dbPort)

	versionDir := dbdbBaseDir + "/mongodb/versions/" + version

	mongodbCmd := exec.Command(
		versionDir+"/basedir/bin/mongod",
		"--config", dataDir+"/mongod.conf",
		"--dbpath", dataDir,
		"--logpath", dataDir+"/mongodb.log",
		"--pidfilepath", dataDir+"/mongodb.pid",
		"--port", dbPort,
		"--fork",
	)

	log.Println("mongodbCmd: " + mongodbCmd.String())
	mongodbCmd.Run()

	portFile := dataDir + "/mongodb.port"
	log.Println("portFile:", portFile)
	fileWrite(portFile, dbPort)

	confFile := dataDir + "/mongod.conf"
	log.Println("Your config file is located:", confFile)

	log.Println(optName, "MongoDB database successfully started.")
}

func mongodbStop(cmd *cobra.Command, ignoreError bool) {
	dbdbBaseDir := dbdbBaseDir()

	optName := cmd.Flag("name").Value.String()

	dataDir := getDataDirByName(optName, "mongodb")
	exitIfNotExistDir(dataDir)

	version := getVersionByDataDir(dataDir, optName, "mongodb")

	dbPort := getPortByName(optName, "mongodb")
	if !ignoreError {
		exitIfNotRunningPort(dbPort)
	}

	dbSocket := "/tmp/dbdb_mongodb_" + dbPort + ".sock"
	log.Println("dbSocket", dbSocket)

	versionDir := dbdbBaseDir + "/mongodb/versions/" + version
	log.Println("versionDir", versionDir)

	pid := pidStringToPidInt(fileRead(dataDir + "/mongodb.pid"))
	syscall.Kill(pid, syscall.SIGTERM)
	remove(dataDir + "mongodb.pid")

	copyFile(dataDir+"/mongodb.port", dataDir+"/mongodb.port.last")

	remove(dataDir + "/mongodb.port")

	log.Println(optName, "MongoDB database successfully stopped.")
}

func mongodbDelete(cmd *cobra.Command) {
	optName := cmd.Flag("name").Value.String()

	dataDir := getDataDirByName(optName, "mongodb")
	exitIfNotExistDir(dataDir)

	dbPort := getPortByName(optName, "mongodb")
	exitIfRunningPort(dbPort)

	remove(dataDir)
	log.Println("data directory deleted. ", dataDir)

	log.Println(optName, "MongoDB database successfully deleted.")
}
