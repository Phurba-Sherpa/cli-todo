/*
Copyright © 2025 Phurba Sherpa <phurba1404@gmail.com>
This is an awesome CLI app in Go Lang.
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// doneCmd represents the done command
func done(cmd *cobra.Command, args []string) {}

var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "A brief description of your command",
	Long:  "",
	Run:   done,
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
