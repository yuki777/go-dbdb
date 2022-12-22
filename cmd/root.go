/*
Copyright © 2022 Yuki Adachi <yuki777@gmail.com>
*/
package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
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
	// rootCmd.AddCommand(mongodbCmd)
}

func currentDir() string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return pwd
}

func getUname() string {
	uname, err := exec.Command("uname", "-s").Output()
	if err != nil {
		fmt.Println(err)
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
		os.Exit(1)
		return "unknown"
	}
}

func exitIfExistDir(checkDir string) {
	if _, err := os.Stat(checkDir); !os.IsNotExist(err) {
		fmt.Println(checkDir + " directory is already exist")
		os.Exit(1)
	}
}

func exitIfRunningPort(port string) {
	cmd := exec.Command("nc", "-z", "127.0.0.1", port)
	cmd.Run()
	exitCode := cmd.ProcessState.ExitCode()

	if exitCode == 0 {
		fmt.Println(port + " is already in use")
		os.Exit(1)
	}
}

func getUrlFileAs(url string, saveAs string) {
	fmt.Println("url: " + url)
	fmt.Println("saveAs: " + saveAs)

	if _, err := os.Stat(saveAs); !os.IsNotExist(err) {
		fmt.Println(saveAs + " is already exist")
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
		fmt.Println(dir + "/basedir directory is already exist")
		return
	}

	fmt.Println("Extracting..." + filepart)
	os.MkdirAll(dir+"/basedir", 0755)
	os.Chdir(dir + "/basedir")

	cpCmd := exec.Command("cp", dir+"/"+filepart+".tar.gz", dir+"/basedir/")
	cpCmd.Run()
	cpExitCode := cpCmd.ProcessState.ExitCode()
	if cpExitCode != 0 {
		fmt.Println("Unknown error on cp")
		os.Exit(1)
	}

	tarCmd := exec.Command("tar", "zxf", filepart+".tar.gz", "--strip-components", "1")
	tarCmd.Run()
	tarExitCode := tarCmd.ProcessState.ExitCode()
	if tarExitCode != 0 {
		fmt.Println("Unknown error on tar")
		os.Exit(1)
	}

	rmCmd := exec.Command("rm", "-f", filepart+".tar.gz")
	rmCmd.Run()
	rmExitCode := rmCmd.ProcessState.ExitCode()
	if rmExitCode != 0 {
		fmt.Println("Unknown error on tar")
		os.Exit(1)
	}
}

func printUsage(optName string, optVersion string, optPort string) {
	currentDir := currentDir()

	prefix := currentDir + "/go-dbdb"
	fmt.Println("")
	fmt.Println("# Start")
	fmt.Println(prefix + " start --name=" + optName)
	fmt.Println("")
	fmt.Println("# Stop")
	fmt.Println(prefix + " stop --name=" + optName)
	fmt.Println("")
	fmt.Println("# Restart")
	fmt.Println(prefix + " restart --name=" + optName)
	fmt.Println("")
	fmt.Println("# Status")
	fmt.Println(prefix + " status --name=" + optName)
	fmt.Println("")
	fmt.Println("# Connect")
	fmt.Println(prefix + " connect --name=" + optName)
	fmt.Println("")
	fmt.Println("# Delete")
	fmt.Println(prefix + " delete --name=" + optName)
	fmt.Println("")
}
