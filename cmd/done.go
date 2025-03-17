/*
Copyright © 2025 Phurba Sherpa <phurba1404@gmail.com>
This is an awesome CLI app in Go Lang.
*/
package cmd

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/phurba-sherpa/tri/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// doneCmd represents the done command
func doneRun(cmd *cobra.Command, args []string) {
	fp := viper.GetString("datafile")
	items, err := todo.ReadItems(fp)
	i, err := strconv.Atoi(args[0])

	if err != nil {
		log.Fatalln(args[0], "is not valid label\n", err)
	}

	if i > 0 && i < len(items) {
		items[i-1].Done = true
		fmt.Printf("%q %v", items[i-1].Text, items[i-1].Done)
		sort.Sort(todo.ByPri(items))
		todo.SaveItem(fp, items)
	} else {
		log.Println(i, "Doesn't match any item")
	}

}

var doneCmd = &cobra.Command{
	Use:     "done",
	Aliases: []string{"do"},
	Short:   "Mark Item as Done",
	Run:     doneRun,
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
