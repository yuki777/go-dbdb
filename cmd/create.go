/*
Copyright © 2022 Yuki Adachi <yuki777@gmail.com>
*/
package cmd

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")

		currentDir := currentDir()
		fmt.Println("currentDir: " + currentDir)

		myOS := getOS()
		fmt.Println("myOS: " + myOS)

		optName := cmd.Flag("name").Value.String()
		optVersion := cmd.Flag("version").Value.String()
		optPort := cmd.Flag("port").Value.String()
		fmt.Println("optName: " + optName)
		fmt.Println("optVersion : " + optVersion)
		fmt.Println("optPort: " + optPort)

		dbUser := "_dbdb_mysql"
		fmt.Println("dbUser: " + dbUser)
		dbSocket := "/tmp/dbdb_mysql_" + optPort + ".sock"
		fmt.Println("dbSocket: " + dbSocket)

		dir := currentDir + "/versions/" + optVersion
		fmt.Println("dir: " + dir)

		os.MkdirAll(dir, 0755)

		beforeDir, err := os.Getwd()
		if err != nil {
		}
		fmt.Println("Before directory: " + beforeDir)

		err = os.Chdir(dir)
		if err != nil {
			fmt.Println(err)
		}

		afterDir, err := os.Getwd()
		if err != nil {
		}
		fmt.Println("After directory: " + afterDir)

		downloadFilePart := "mysql-" + optVersion + "-" + myOS
		fmt.Println("downloadFilePart: " + downloadFilePart)

		checkDir := dir + "/datadir/" + optName
		fmt.Println("checkDir: " + checkDir)
		exitIfExistDir(checkDir)

		exitIfRunningPort(optPort)

		getUrlFileAs("https://dbdb.project8.jp/mysql/"+downloadFilePart+".tar.gz", downloadFilePart+".tar.gz")
		os.MkdirAll(dir+"/datadir/"+optName, 0755)

		extractFile(dir, downloadFilePart)

		err = os.Chdir(currentDir)
		if err != nil {
			fmt.Println(err)
		}

		// mysqld initialize
		mysqldCmd := exec.Command(
			dir+"/basedir/bin/mysqld",
			"--initialize-insecure",
			"--no-defaults",
			"--user="+dbUser,
			"--port="+optPort,
			"--socket="+dbSocket,
			"--basedir="+dir+"/basedir",
			"--plugin-dir="+dir+"/basedir/lib/plugin",
			"--datadir="+dir+"/datadir/"+optName,
			"--log-error="+dir+"/datadir/"+optName+"/mysqld.err",
			"--pid-file="+dir+"/datadir/"+optName+"/mysql.pid",
		)

		fmt.Println("mysqldCmd: " + mysqldCmd.String())

		var stdout bytes.Buffer
		var stderr bytes.Buffer
		mysqldCmd.Stdout = &stdout
		mysqldCmd.Stderr = &stderr
		mysqldErr := mysqldCmd.Run()
		if mysqldErr != nil {
			fmt.Println("stdout: " + stdout.String())
			fmt.Println(fmt.Sprint(mysqldErr) + ": " + stderr.String())
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
		fmt.Println("my.cnf is here. " + dir + "/datadir/" + optName + "/my.cnf")

		err = os.Chdir(currentDir)
		if err != nil {
			fmt.Println(err)
		}

		printUsage(optName, optVersion, optPort)
	},
}

func init() {
	//rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	createCmd.PersistentFlags().String("name", "", "Name for database (required)")
	createCmd.PersistentFlags().String("version", "", "Version for database (required)")
	createCmd.PersistentFlags().String("port", "", "Port for database (required)")

	createCmd.MarkPersistentFlagRequired("name")
	createCmd.MarkPersistentFlagRequired("version")
	createCmd.MarkPersistentFlagRequired("port")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
