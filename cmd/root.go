/*
Copyright © 2025 Phurba Sherpa <phurba1404@gmail.com>
This is an awesome CLI app in Go Lang.
*/
package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

/*
1. Config varies with environment mac, window, linux
2. code remains same
3. 12 Factors App stores config in ENV
*/

var (
	cfgFile  string
	dataFile string
)

func initConfig() {
	viper.SetConfigName(".tri")
	viper.AddConfigPath("$HOME")
	viper.AutomaticEnv()

	// If a config file is found, read it in
	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file: ", viper.ConfigFileUsed())
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	home, err := os.UserHomeDir()

	if err != nil {
		log.Panicln("Unable to detect the home director. Please set the data file using --dataFile")
	}

	defaultPath := home + string(os.PathSeparator) + "personal" + string(os.PathSeparator) + "tri" + string(os.PathSeparator) + ".tridos.json"
	rootCmd.PersistentFlags().StringVar(&dataFile, "datafile", defaultPath, "Data file to store todos")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default to $HOME/.tri.yaml)")
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tri",
	Short: "An awesome CLI App",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
