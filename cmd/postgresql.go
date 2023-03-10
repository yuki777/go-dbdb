/*
Copyright © 2023 Yuki Adachi <yuki777@gmail.com>
*/
package cmd

import (
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var postgresqlCmd = &cobra.Command{
	Use:   "postgresql",
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
	postgresqlCmd.AddCommand(postgresqlCreateCmd)
	postgresqlCmd.AddCommand(postgresqlStartCmd)
	postgresqlCmd.AddCommand(postgresqlStopCmd)
	postgresqlCmd.AddCommand(postgresqlRestartCmd)
	postgresqlCmd.AddCommand(postgresqlDeleteCmd)
	postgresqlCmd.AddCommand(postgresqlCreateStartCmd)
}

func postgresqlCreate(cmd *cobra.Command) {
	log.Println(getCurrentFuncName(), "called")
	dbdbBaseDir := dbdbBaseDir()

	optName := cmd.Flag("name").Value.String()
	optVersion := cmd.Flag("version").Value.String()
	optPort := cmd.Flag("port").Value.String()

	versionDir := dbdbBaseDir + "/postgresql/versions/" + optVersion
	os.MkdirAll(versionDir, 0755)
	os.Chdir(versionDir)

	downloadFilePart := "postgresql-" + optVersion + "-" + getOS()

	dataDir := versionDir + "/datadir/" + optName
	if exists(dataDir) {
		log.Println(dataDir + " directory is already exist")
		os.Exit(1)
	}

	if isRunningPort(optPort) {
		log.Println(optPort, "is already in use")
		os.Exit(1)
	}

	getUrlFileAs("https://dbdb.project8.jp/postgresql/"+downloadFilePart+".tar.gz", downloadFilePart+".tar.gz")
	os.MkdirAll(dataDir, 0755)

	extractFile(versionDir, downloadFilePart)

	postgresqlInstallForLinux(versionDir)

	createCmd := exec.Command(
		versionDir+"/basedir/bin/initdb",
		"--pgdata="+dataDir,
		"--username=postgres",
		"--encoding=UTF-8",
		"--locale=en_US.UTF-8",
	)
	log.Println("createCmd:", createCmd.String())
	createCmd.Run()

	portFile := dataDir + "/postgresql.port.init"
	fileWrite(portFile, optPort)
	log.Println("postgresql.port.init:", portFile)

	confFile := dataDir + "/postgresql.conf"
	log.Println("postgresql.conf is here", confFile)

	log.Println(optName, "PostgreSQL database successfully created.")
	printUsage(optName, optVersion, optPort)
}

func postgresqlInstallForLinux(versionDir string) {
	log.Println(getCurrentFuncName(), "called")
	if getOS() != "linux" {
		return
	}

	_, err := os.Stat(versionDir + "/basedir/bin")
	if err == nil {
		return
	}

	if os.IsNotExist(err) {
		os.Chdir(versionDir + "/basedir")

		// configure
		configureCmd := exec.Command(
			"./configure",
			"--prefix="+versionDir+"/basedir",
		)
		log.Println("configureCmd:", configureCmd.String())
		configureCmd.Run()

		// make
		makeCmd := exec.Command(
			"make",
		)
		log.Println("makeCmd:", makeCmd.String())
		makeCmd.Run()

		// make install
		makeInstallCmd := exec.Command(
			"make",
			"install",
		)
		log.Println("makeInstallCmd:", makeInstallCmd.String())
		makeInstallCmd.Run()

		// rm
		rmCmd := exec.Command(
			"rm",
			"-fr",
			"config",
			"contrib",
			"doc",
			"src",
		)
		log.Println("rmCmd:", rmCmd.String())
		rmCmd.Run()
	}
}

func postgresqlStart(cmd *cobra.Command) {
	log.Println(getCurrentFuncName(), "called")
	dbdbBaseDir := dbdbBaseDir()

	optName := cmd.Flag("name").Value.String()

	dataDir := getDataDirByName(optName, "postgresql")

	if notExists(dataDir) {
		log.Println(dataDir + " directory is NOT exist")
		os.Exit(1)
	}

	version := getVersionByDataDir(dataDir, optName, "postgresql")

	dbPort := getPortByName(optName, "postgresql")

	if isRunningPort(dbPort) {
		log.Println(dbPort, "is already in use")
		os.Exit(1)
	}

	versionDir := dbdbBaseDir + "/postgresql/versions/" + version

	startCmd := exec.Command(
		versionDir+"/basedir/bin/pg_ctl",
		"--pgdata", dataDir,
		"--log", dataDir+"/postgres.log",
		"-w",
		"-o '-p "+dbPort+"'",
		"start",
	)

	log.Println("startCmd:", startCmd.String())
	startCmd.Run()

	portFile := dataDir + "/postgresql.port"
	log.Println("portFile:", portFile)
	fileWrite(portFile, dbPort)

	confFile := dataDir + "/postgresql.conf"
	log.Println("Your config file is located:", confFile)

	log.Println(optName, "PostgreSQL database successfully started.")
}

func postgresqlStop(cmd *cobra.Command, checkPort bool) {
	log.Println(getCurrentFuncName(), "called")
	dbdbBaseDir := dbdbBaseDir()

	optName := cmd.Flag("name").Value.String()

	dataDir := getDataDirByName(optName, "postgresql")

	if notExists(dataDir) {
		log.Println(dataDir + " directory is NOT exist")
		os.Exit(1)
	}

	version := getVersionByDataDir(dataDir, optName, "postgresql")

	dbPort := getPortByName(optName, "postgresql")
	if checkPort && isNotRunningPort(dbPort) {
		log.Println(dbPort, "is NOT available")
		os.Exit(1)
	}

	versionDir := dbdbBaseDir + "/postgresql/versions/" + version

	stopCmd := exec.Command(
		versionDir+"/basedir/bin/pg_ctl",
		"--pgdata", dataDir,
		"--log", dataDir+"/postgres.log",
		"-w",
		"-o '-p "+dbPort+"'",
		"stop",
	)
	log.Println("stopCmd", stopCmd.String())
	stopCmd.Run()

	copy(dataDir+"/postgresql.port", dataDir+"/postgresql.port.last")

	remove(dataDir + "/postgresql.port")

	log.Println(optName, "PostgreSQL database successfully stopped.")
}

func postgresqlCreateStart(cmd *cobra.Command) {
	log.Println(getCurrentFuncName(), "called")
	postgresqlCreate(cmd)
	postgresqlStart(cmd)
}

func postgresqlRestart(cmd *cobra.Command) {
	log.Println(getCurrentFuncName(), "called")
	postgresqlStop(cmd, false)
	postgresqlStart(cmd)
}

func postgresqlDelete(cmd *cobra.Command) {
	log.Println(getCurrentFuncName(), "called")
	dbdbDelete(cmd, "postgresql")
}
