/*
Copyright © 2025 Phurba Sherpa <phurba1404@gmail.com>
This is an awesome CLI app in Go Lang.
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/phurba-sherpa/tri/todo"
	"github.com/spf13/cobra"
)

var (
	doneOpt bool
	allOpt  bool
)

func listTodo(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(dataFile)
	if err != nil {
		log.Fatalf(dataFile, "Error reading content", err)
	}

	sort.Sort(todo.ByPri(items))

	w := tabwriter.NewWriter(os.Stdout, 3, 0, 2, ' ', 0)
	defer w.Flush()

	// ** if items display record
	if len(items) > 0 {
		fmt.Fprintln(w, "SNo\tPriority\tName\tStatus\t")
		for _, item := range items {
			if allOpt || (item.Done == doneOpt) {
				fmt.Fprintln(w, item.Label()+"\t"+item.PrettyP()+"\t"+item.Text+"\t"+item.PrettyD()+"\t")
			}
		}
	} else {
		fmt.Println("No todos to display")
	}
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List is used to dispaly a list of todo",
	Long:  `Display all todos`,
	Run:   listTodo,
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVar(&doneOpt, "done", false, "Show 'done' todos")
	listCmd.Flags().BoolVar(&allOpt, "all", false, "Show all todos")
}
