package cmd

import (
	"fmt"
	"os"
	"os/user"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	sshDir  string
)

var sshtoolsCmd = &cobra.Command{
	Use:   "sshtools",
	Short: "SSH tools",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the sshtoolsCmd.
func Execute() {
	if err := sshtoolsCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	currUser, err := user.Current()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// cobra.OnInitialize(initConfig(currUser))
	// sshtoolsCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", path.Join(currUser.HomeDir, ".sshtools"), "config file")
	sshtoolsCmd.PersistentFlags().StringVarP(&sshDir, "ssh-dir", "s", path.Join(currUser.HomeDir, ".ssh"), "ssh root dir")
	viper.BindPFlag("ssh-dir", sshtoolsCmd.PersistentFlags().Lookup("ssh-dir"))
}

// initConfig reads in config file and ENV variables if set.
// TODO config file not working, investigate
// func initConfig(currUser *user.User) func() {
// 	configFunc := func() {
// 		home := currUser.HomeDir
// 		if cfgFile != "" {
// 			// Use config file from the flag.
// 			viper.SetConfigFile(cfgFile)
// 		} else {
// 			// Search config in home directory with name ".sshtools" (without extension).
// 			viper.AddConfigPath(home)
// 			viper.SetConfigName(".sshtools")
// 		}
//
// 		viper.AutomaticEnv() // read in environment variables that match
//
// 		// If a config file is found, read it in.
// 		if err := viper.ReadInConfig(); err == nil {
// 			fmt.Println("Using config file:", viper.ConfigFileUsed())
// 		}
// 	}
// 	return configFunc
// }
