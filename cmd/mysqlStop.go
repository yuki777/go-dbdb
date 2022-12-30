/*
Copyright © 2022 Yuki Adachi <yuki777@gmail.com>
*/
package cmd

import (
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var mysqlStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop mysql server",
	Long:  `...`,
	Run: func(cmd *cobra.Command, args []string) {

		dbdbBaseDir := dbdbBaseDir()
		log.Println("dbdbBaseDir: " + dbdbBaseDir)

		optName := cmd.Flag("name").Value.String()
		log.Println("optName: " + optName)

		dataDir := getDataDirByName(optName, "mysql")
		log.Println("dataDir:", dataDir)
		exitIfNotExistDir(dataDir)

		version := getVersionByDataDir(dataDir, optName, "mysql")
		log.Println("version:", version)

		dbPort := getPortByName(optName)
		log.Println("dbPort:", dbPort)
		exitIfNotRunningPort(dbPort)

		dbSocket := "/tmp/dbdb_mysql_" + dbPort + ".sock"

		versionDir := dbdbBaseDir + "/mysql/versions/" + version
		log.Println("versionDir:", versionDir)

		// $dir/basedir/bin/mysqladmin --user=root --host=localhost --port=$optPort --socket=$optSocket shutdown
		mysqldCmd := exec.Command(
			versionDir+"/basedir/bin/mysqladmin",
			"--user=root",
			"--host=localhost",
			"--port="+dbPort,
			"--socket="+dbSocket,
			"shutdown",
		)
		log.Println("mysqldCmd: " + mysqldCmd.String())
		mysqldCmd.Run()

		// [ -f "$dir/datadir/$optName/mysql.port" ] && cp $dir/datadir/$optName/mysql.port $dir/datadir/$optName/mysql.port.last
		inputFile, err := os.Open(dataDir + "/mysql.port")
		if err != nil {
			log.Println("unknown error on inputFile:", inputFile)
		}
		defer inputFile.Close()
		outputFile, err := os.Create(dataDir + "/mysql.port.last")
		if err != nil {
			log.Println("unknown error on outputFile:", outputFile)
		}
		defer outputFile.Close()
		_, err = io.Copy(outputFile, inputFile)
		if err != nil {
			log.Println("unknown error on copy:", inputFile, outputFile)
		}

		// [ -f "$dir/datadir/$optName/mysql.port" ] && rm -f $dir/datadir/$optName/mysql.port
		removeDir(dataDir + "/mysql.port")

		log.Println(optName, "MySQL database successfully stopped.")
	},
}

func init() {
	//rootCmd.AddCommand(mysqlCreateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	mysqlStopCmd.PersistentFlags().String("name", "", "Name for database (required)")

	mysqlStopCmd.MarkPersistentFlagRequired("name")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mysqlCreateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
