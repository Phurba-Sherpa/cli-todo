/*
Copyright © 2025 phurba
*/
package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tri",
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
var dataFile string

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	home, err := os.UserHomeDir()

	if err != nil {
		log.Panicln("Unable to detect the home director. Please set the data file using --dataFile")
	}

	defaultPath := home + string(os.PathSeparator) + "personal" + string(os.PathSeparator) + "tri" + string(os.PathSeparator) + ".tridos.json"

	rootCmd.PersistentFlags().StringVar(&dataFile, "datafile", defaultPath, "Data file to store todos")
}
