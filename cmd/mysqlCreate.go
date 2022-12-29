/*
Copyright © 2022 Yuki Adachi <yuki777@gmail.com>
*/
package cmd

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var mysqlCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create mysql server",
	Long:  `...`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("mysqlCreate called")

		dbdbBaseDir := dbdbBaseDir()
		log.Println("dbdbBaseDir: " + dbdbBaseDir)

		myOS := getOS()
		log.Println("myOS: " + myOS)

		optName := cmd.Flag("name").Value.String()
		optVersion := cmd.Flag("version").Value.String()
		optPort := cmd.Flag("port").Value.String()
		log.Println("optName: " + optName)
		log.Println("optVersion : " + optVersion)
		log.Println("optPort: " + optPort)

		dbUser := "_dbdb_mysql"
		log.Println("dbUser: " + dbUser)
		dbSocket := "/tmp/dbdb_mysql_" + optPort + ".sock"
		log.Println("dbSocket: " + dbSocket)

		dir := dbdbBaseDir + "/mysql/versions/" + optVersion
		log.Println("dir: " + dir)

		os.MkdirAll(dir, 0755)

		beforeDir, err := os.Getwd()
		if err != nil {
		}
		log.Println("Before directory: " + beforeDir)

		err = os.Chdir(dir)
		if err != nil {
			log.Println(err)
		}

		afterDir, err := os.Getwd()
		if err != nil {
		}
		log.Println("After directory: " + afterDir)

		downloadFilePart := "mysql-" + optVersion + "-" + myOS
		log.Println("downloadFilePart: " + downloadFilePart)

		checkDir := dir + "/datadir/" + optName
		log.Println("checkDir: " + checkDir)
		exitIfExistDir(checkDir)

		exitIfRunningPort(optPort)

		getUrlFileAs("https://dbdb.project8.jp/mysql/"+downloadFilePart+".tar.gz", downloadFilePart+".tar.gz")
		os.MkdirAll(dir+"/datadir/"+optName, 0755)

		extractFile(dir, downloadFilePart)

		err = os.Chdir(dbdbBaseDir)
		if err != nil {
			log.Println(err)
		}

		// mysqld initialize
		mysqldCmd := exec.Command(
			dir+"/basedir/bin/mysqld",
			"--no-defaults",
			"--initialize-insecure",
			"--user="+dbUser,
			"--port="+optPort,
			"--socket="+dbSocket,
			"--basedir="+dir+"/basedir",
			"--plugin-dir="+dir+"/basedir/lib/plugin",
			"--datadir="+dir+"/datadir/"+optName,
			"--log-error="+dir+"/datadir/"+optName+"/mysqld.err",
			"--pid-file="+dir+"/datadir/"+optName+"/mysql.pid",
		)

		log.Println("mysqldCmd: " + mysqldCmd.String())

		var stdout bytes.Buffer
		var stderr bytes.Buffer
		mysqldCmd.Stdout = &stdout
		mysqldCmd.Stderr = &stderr
		mysqldErr := mysqldCmd.Run()
		if mysqldErr != nil {
			log.Println("stdout: " + stdout.String())
			log.Println(fmt.Sprint(mysqldErr) + ": " + stderr.String())
			panic(mysqldErr)
		}

		// mysql.port.init
		mysqlPortFile, err := os.Create(dir + "/datadir/" + optName + "/mysql.port.init")
		if err != nil {
			panic(err)
		}
		defer mysqlPortFile.Close()
		_, err = mysqlPortFile.WriteString(optPort)
		if err != nil {
			panic(err)
		}

		// my.cnf
		myCnf, err := os.Create(dir + "/datadir/" + optName + "/my.cnf")
		if err != nil {
			panic(err)
		}
		defer myCnf.Close()
		myCnfText := "[mysqld]\nbind-address = 127.0.0.1"
		_, err = myCnf.WriteString(myCnfText)
		if err != nil {
			panic(err)
		}
		log.Println("my.cnf is here. " + dir + "/datadir/" + optName + "/my.cnf")

		err = os.Chdir(dbdbBaseDir)
		if err != nil {
			log.Println(err)
		}

		printUsage(optName, optVersion, optPort)
	},
}

func init() {
	//rootCmd.AddCommand(mysqlCreateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	mysqlCreateCmd.PersistentFlags().String("name", "", "Name for database (required)")
	mysqlCreateCmd.PersistentFlags().String("version", "", "Version for database (required)")
	mysqlCreateCmd.PersistentFlags().String("port", "", "Port for database (required)")

	mysqlCreateCmd.MarkPersistentFlagRequired("name")
	mysqlCreateCmd.MarkPersistentFlagRequired("version")
	mysqlCreateCmd.MarkPersistentFlagRequired("port")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mysqlCreateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}