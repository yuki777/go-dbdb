/*
Copyright © 2022 Yuki Adachi <yuki777@gmail.com>
*/
package cmd

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-dbdb",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-dbdb.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.AddCommand(mysqlCmd)
	// rootCmd.AddCommand(postgresqlCmd)
	// rootCmd.AddCommand(redisCmd)
	rootCmd.AddCommand(mongodbCmd)
}

func dbdbBaseDir() string {
	dbdbBaseDir := ""
	xdgDataHome := os.Getenv("XDG_DATA_HOME")
	if xdgDataHome == "" {
		currentUser, err := user.Current()
		if err != nil {
			os.Exit(1)
		}

		homeDir := currentUser.HomeDir
		dbdbBaseDir = homeDir + "/.local/share/dbdb"
	} else {
		dbdbBaseDir = xdgDataHome + "/dbdb"
	}

	os.MkdirAll(dbdbBaseDir, 0755)
	os.Chdir(dbdbBaseDir)

	return dbdbBaseDir
}

func getUname() string {
	uname, err := exec.Command("uname", "-s").Output()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	return string(uname)
}

func getOS() string {
	uname := getUname()
	if strings.HasPrefix(strings.ToLower(uname), "linux") {
		return "linux"
	} else if strings.HasPrefix(strings.ToLower(uname), "darwin") {
		return "macos"
	} else {
		log.Println("unknown os")
		os.Exit(1)
		return "unknown"
	}
}

func exitIfExistDir(checkDir string) {
	if _, err := os.Stat(checkDir); !os.IsNotExist(err) {
		log.Println(checkDir + " directory is already exist")
		os.Exit(1)
	}
}

func exitIfNotExistDir(checkDir string) {
	if _, err := os.Stat(checkDir); os.IsNotExist(err) {
		log.Println(checkDir + " directory is NOT exist")
		os.Exit(1)
	}
}

func exitIfRunningPort(port string) {
	cmd := exec.Command("nc", "-z", "127.0.0.1", port)
	cmd.Run()
	exitCode := cmd.ProcessState.ExitCode()

	if exitCode == 0 {
		log.Println(port + " is already in use")
		os.Exit(1)
	}
}

func getUrlFileAs(url string, saveAs string) {
	log.Println("url: " + url)
	log.Println("saveAs: " + saveAs)

	if _, err := os.Stat(saveAs); !os.IsNotExist(err) {
		log.Println(saveAs + " is already exist")
		return
	}

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	reader := response.Body

	file, err := os.Create(saveAs)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, reader)
	if err != nil {
		panic(err)
	}
}

func extractFile(dir string, filepart string) {
	if _, err := os.Stat(dir + "/basedir"); !os.IsNotExist(err) {
		log.Println(dir + "/basedir directory is already exist")
		return
	}

	log.Println("Extracting..." + filepart)
	os.MkdirAll(dir+"/basedir", 0755)
	os.Chdir(dir + "/basedir")

	cpCmd := exec.Command("cp", dir+"/"+filepart+".tar.gz", dir+"/basedir/")
	cpCmd.Run()
	cpExitCode := cpCmd.ProcessState.ExitCode()
	if cpExitCode != 0 {
		log.Println("Unknown error on cp")
		os.Exit(1)
	}

	tarCmd := exec.Command("tar", "zxf", filepart+".tar.gz", "--strip-components", "1")
	tarCmd.Run()
	tarExitCode := tarCmd.ProcessState.ExitCode()
	if tarExitCode != 0 {
		log.Println("Unknown error on tar")
		os.Exit(1)
	}

	rmCmd := exec.Command("rm", "-f", filepart+".tar.gz")
	rmCmd.Run()
	rmExitCode := rmCmd.ProcessState.ExitCode()
	if rmExitCode != 0 {
		log.Println("Unknown error on tar")
		os.Exit(1)
	}
}

func printUsage(optName string, optVersion string, optPort string) {
	prefix := os.Args[0]
	log.Println("")
	log.Println("# Start")
	log.Println(prefix + " mysql start --name=" + optName)
	log.Println("")
	log.Println("# Stop")
	log.Println(prefix + " mysql stop --name=" + optName)
	log.Println("")
	log.Println("# Restart")
	log.Println(prefix + " mysql restart --name=" + optName)
	log.Println("")
	log.Println("# Status")
	log.Println(prefix + " mysql status --name=" + optName)
	log.Println("")
	log.Println("# Connect")
	log.Println(prefix + " mysql connect --name=" + optName)
	log.Println("")
	log.Println("# Delete")
	log.Println(prefix + " mysql delete --name=" + optName)
	log.Println("")
}

func getDataDirByName(optName string) string {
	dbdbBaseDir := dbdbBaseDir()
	pattern := dbdbBaseDir + "/mysql/versions/*/datadir/" + optName
	files, err := filepath.Glob(pattern)
	if len(files) != 1 {
		log.Println("data directory not found.", pattern)
		panic(err)
	}

	return files[0]
}

func getPortByName(optName string) string {
	dataDir := getDataDirByName(optName)
	mysqlPortInitFile := dataDir + "/mysql.port.init"

	bytes, err := ioutil.ReadFile(mysqlPortInitFile)
	if err != nil {
		log.Println("unknown error on mysql.port.init file", mysqlPortInitFile)
		panic(err)
	}

	return string(bytes)
}

func removeDir(dir string) {
	err := os.RemoveAll(dir)
	if err != nil {
		log.Println("unknown error on removeDir", dir)
		panic(err)
	}
}

func validateOptName(optName string) bool {
	if !regexp.MustCompile(`^[0-9a-zA-Z-_.]+$`).MatchString(optName) {
		return false
	}
	return true
}
