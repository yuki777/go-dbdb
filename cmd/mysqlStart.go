/*
Copyright © 2022 Yuki Adachi <yuki777@gmail.com>
*/
package cmd

import (
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

var mysqlStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start mysql server",
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

		port := getPortByName(optName)
		log.Println("port:", port)
		exitIfRunningPort(port)

		versionDir := dbdbBaseDir + "/mysql/versions/" + version
		log.Println("versionDir:", versionDir)

		dbPort := getPortByName(optName)
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

		mysqlPortFile := dataDir + "/mysql.port"
		log.Println("mysqlPortFile:", mysqlPortFile)
		fileWrite(mysqlPortFile, dbPort)

		myConfFile := dataDir + "/my.cnf"
		log.Println("Your config file is located:", myConfFile)

		log.Println(optName, "MySQL database successfully started.")
	},
}

func init() {
	//rootCmd.AddCommand(mysqlCreateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	mysqlStartCmd.PersistentFlags().String("name", "", "Name for database (required)")

	mysqlStartCmd.MarkPersistentFlagRequired("name")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mysqlCreateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
