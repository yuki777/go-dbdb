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

var redisCmd = &cobra.Command{
	Use:   "redis",
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
	redisCmd.AddCommand(redisCreateCmd)
	redisCmd.AddCommand(redisStartCmd)
	redisCmd.AddCommand(redisStopCmd)
	redisCmd.AddCommand(redisRestartCmd)
	redisCmd.AddCommand(redisDeleteCmd)
	redisCmd.AddCommand(redisCreateStartCmd)
}

func redisCreate(cmd *cobra.Command) {
	log.Println(getCurrentFuncName(), "called")
	dbdbBaseDir := dbdbBaseDir()

	optName := cmd.Flag("name").Value.String()
	optVersion := cmd.Flag("version").Value.String()
	optPort := cmd.Flag("port").Value.String()

	versionDir := dbdbBaseDir + "/redis/versions/" + optVersion
	os.MkdirAll(versionDir, 0755)
	os.Chdir(versionDir)

	downloadFilePart := "redis-" + optVersion + "-" + getOS()

	dataDir := versionDir + "/datadir/" + optName
	if exists(dataDir) {
		log.Println(dataDir + " directory is already exist")
		os.Exit(1)
	}

	if isRunningPort(optPort) {
		log.Println(optPort, "is already in use")
		os.Exit(1)
	}

	getUrlFileAs("https://dbdb.project8.jp/redis/"+downloadFilePart+".tar.gz", downloadFilePart+".tar.gz")
	os.MkdirAll(dataDir, 0755)

	extractFile(versionDir, downloadFilePart)

	if notExists(versionDir + "/basedir/src/redis-server") {
		os.Chdir(versionDir + "/basedir")
		makeCmd := exec.Command(
			"make",
		)
		log.Println("makeCmd:", makeCmd.String())
		makeCmd.Run()
	}

	if notExists(dataDir + "/redis.conf") {
		copy(versionDir+"/basedir/redis.conf", dataDir+"/redis.conf")
	}
	log.Println("redis.conf:", dataDir+"/redis.conf")

	portFile := dataDir + "/redis.port.init"
	fileWrite(portFile, optPort)
	log.Println("redis.port.init:", portFile)

	log.Println(optName, "Redis database successfully created.")
	printUsage(optName, optVersion, optPort)
}

func redisStart(cmd *cobra.Command) {
	log.Println(getCurrentFuncName(), "called")
	dbdbBaseDir := dbdbBaseDir()

	optName := cmd.Flag("name").Value.String()

	dataDir := getDataDirByName(optName, "redis")

	if notExists(dataDir) {
		log.Println(dataDir + " directory is NOT exist")
		os.Exit(1)
	}

	version := getVersionByDataDir(dataDir, optName, "redis")

	dbPort := getPortByName(optName, "redis")

	if isRunningPort(dbPort) {
		log.Println(dbPort, "is already in use")
		os.Exit(1)
	}

	versionDir := dbdbBaseDir + "/redis/versions/" + version

	startCmd := exec.Command(
		versionDir+"/basedir/src/redis-server",
		dataDir+"/redis.conf",
		"--port", dbPort,
		"--dir", dataDir,
		"--pidfile", dataDir+"/redis.pid",
		"--daemonize", "yes",
	)

	log.Println("startCmd:", startCmd.String())
	startCmd.Run()

	portFile := dataDir + "/redis.port"
	log.Println("portFile:", portFile)
	fileWrite(portFile, dbPort)

	confFile := dataDir + "/redis.conf"
	log.Println("Your config file is located:", confFile)

	log.Println(optName, "Redis database successfully started.")
}

func redisCreateStart(cmd *cobra.Command) {
	log.Println(getCurrentFuncName(), "called")
	redisCreate(cmd)
	redisStart(cmd)
}

func redisRestart(cmd *cobra.Command) {
	redisStop(cmd, false)
	redisStart(cmd)
}

func redisStop(cmd *cobra.Command, checkPort bool) {
	log.Println(getCurrentFuncName(), "called")
	dbdbBaseDir := dbdbBaseDir()

	optName := cmd.Flag("name").Value.String()

	dataDir := getDataDirByName(optName, "redis")

	if notExists(dataDir) {
		log.Println(dataDir + " directory is NOT exist")
		os.Exit(1)
	}

	version := getVersionByDataDir(dataDir, optName, "redis")

	dbPort := getPortByName(optName, "redis")
	if checkPort && isNotRunningPort(dbPort) {
		log.Println(dbPort, "is NOT available")
		os.Exit(1)
	}

	versionDir := dbdbBaseDir + "/redis/versions/" + version

	stopCmd := exec.Command(
		versionDir+"/basedir/src/redis-cli",
		"-p", dbPort,
		"shutdown",
	)
	log.Println("stopCmd", stopCmd.String())
	stopCmd.Run()

	copy(dataDir+"/redis.port", dataDir+"/redis.port.last")

	remove(dataDir + "/redis.port")

	log.Println(optName, "Redis database successfully stopped.")
}

func redisDelete(cmd *cobra.Command) {
	log.Println(getCurrentFuncName(), "called")
	dbdbDelete(cmd, "redis")
}
