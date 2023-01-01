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
	// postgresqlCmd.AddCommand(postgresqlStartCmd)
	// postgresqlCmd.AddCommand(postgresqlStopCmd)
	// postgresqlCmd.AddCommand(postgresqlRestartCmd)
	// postgresqlCmd.AddCommand(postgresqlDeleteCmd)
	// postgresqlCmd.AddCommand(postgresqlCreateStartCmd)
}

func postgresqlCreate(cmd *cobra.Command) {
	dbdbBaseDir := dbdbBaseDir()

	optName := cmd.Flag("name").Value.String()
	optVersion := cmd.Flag("version").Value.String()
	optPort := cmd.Flag("port").Value.String()

	versionDir := dbdbBaseDir + "/postgresql/versions/" + optVersion
	os.MkdirAll(versionDir, 0755)
	os.Chdir(versionDir)

	downloadFilePart := "postgresql-" + optVersion + "-" + getOS()

	dataDir := versionDir + "/datadir/" + optName
	exitIfExistDir(dataDir)

	exitIfRunningPort(optPort)

	getUrlFileAs("https://dbdb.project8.jp/postgresql/"+downloadFilePart+".tar.gz", downloadFilePart+".tar.gz")
	os.MkdirAll(dataDir, 0755)

	extractFile(versionDir, downloadFilePart)

	postgresqlInstallForLinux(versionDir + "/basedir/bin")

	initCmd := exec.Command(
		versionDir+"/basedir/bin/initdb",
		"--pgdata="+dataDir,
		"--username=postgres",
		"--encoding=UTF-8",
		"--locale=en_US.UTF-8",
	)
	log.Println("initCmd: " + initCmd.String())
	initCmd.Run()

	portFile := versionDir + "/datadir/" + optName + "/postgresql.port.init"
	fileWrite(portFile, optPort)
	log.Println("postgresql.port.init:", portFile)

	confFile := dataDir + "/postgresql.conf"
	log.Println("postgresql.conf is here", confFile)

	log.Println(optName, "PostgreSQL database successfully created.")
	printUsage(optName, optVersion, optPort)
}

func postgresqlInstallForLinux(versionDir string) {
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
