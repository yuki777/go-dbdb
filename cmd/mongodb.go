/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"
	"os/exec"
	"syscall"
	"time"

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

func mongodbCreateStart(cmd *cobra.Command) {
	log.Println(getCurrentFuncName(), "called")
	mongodbCreate(cmd)
	mongodbStart(cmd)
}

func mongodbCreate(cmd *cobra.Command) {
	log.Println(getCurrentFuncName(), "called")
	dbdbBaseDir := dbdbBaseDir()

	optName := cmd.Flag("name").Value.String()
	optVersion := cmd.Flag("version").Value.String()
	optPort := cmd.Flag("port").Value.String()

	versionDir := dbdbBaseDir + "/mongodb/versions/" + optVersion
	os.MkdirAll(versionDir, 0755)
	os.Chdir(versionDir)

	downloadFilePart := "mongodb-" + optVersion + "-" + getOS()

	dataDir := versionDir + "/datadir/" + optName

	if exists(dataDir) {
		log.Println(dataDir + " directory is already exist")
		os.Exit(1)
	}

	if isRunningPort(optPort) {
		log.Println(optPort, "is already in use")
		os.Exit(1)
	}

	getUrlFileAs("https://dbdb.project8.jp/mongodb/"+downloadFilePart+".tar.gz", downloadFilePart+".tar.gz")
	os.MkdirAll(dataDir, 0755)

	extractFile(versionDir, downloadFilePart)

	portFile := dataDir + "/mongodb.port.init"
	fileWrite(portFile, optPort)
	log.Println("mongodb.port.init:", portFile)

	confFile := dataDir + "/mongod.conf"
	fileWrite(confFile, "#mongod.conf\n")
	log.Println("mongod.conf:", confFile)

	log.Println(optName, "MongoDB database successfully created.")
	printUsage(optName, optVersion, optPort)
}

func mongodbStart(cmd *cobra.Command) {
	log.Println(getCurrentFuncName(), "called")
	dbdbBaseDir := dbdbBaseDir()

	optName := cmd.Flag("name").Value.String()

	dataDir := getDataDirByName(optName, "mongodb")

	if notExists(dataDir) {
		log.Println(dataDir + " directory is NOT exist")
		os.Exit(1)
	}

	version := getVersionByDataDir(dataDir, optName, "mongodb")

	dbPort := getPortByName(optName, "mongodb")

	if isRunningPort(dbPort) {
		log.Println(dbPort, "is already in use")
		os.Exit(1)
	}

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

func mongodbStop(cmd *cobra.Command, checkPort bool) {
	log.Println(getCurrentFuncName(), "called")
	optName := cmd.Flag("name").Value.String()

	dataDir := getDataDirByName(optName, "mongodb")

	if notExists(dataDir) {
		log.Println(dataDir + " directory is NOT exist")
		os.Exit(1)
	}

	dbPort := getPortByName(optName, "mongodb")
	if checkPort && isNotRunningPort(dbPort) {
		log.Println(dbPort, "is NOT available")
		os.Exit(1)
	}

	pid := pidStringToPidInt(fileRead(dataDir + "/mongodb.pid"))
	syscall.Kill(pid, syscall.SIGTERM)
	remove(dataDir + "mongodb.pid")

	copy(dataDir+"/mongodb.port", dataDir+"/mongodb.port.last")

	remove(dataDir + "/mongodb.port")

	time.Sleep(1 * time.Second) // wait to close the port

	log.Println(optName, "MongoDB database successfully stopped.")
}

func mongodbRestart(cmd *cobra.Command) {
	log.Println(getCurrentFuncName(), "called")
	mongodbStop(cmd, false)
	mongodbStart(cmd)
}

func mongodbDelete(cmd *cobra.Command) {
	log.Println(getCurrentFuncName(), "called")
	dbdbDelete(cmd, "mongodb")
}
