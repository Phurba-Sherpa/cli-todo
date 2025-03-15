/*
Copyright © 2025 Phurba Sherpa <phurba1404@gmail.com>
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

	sort.Sort(todo.ByPri(items))
	if err != nil {
		fmt.Printf("file path: %s\n", dataFile)
		log.Fatal(err)
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 3, 0, 2, ' ', 0)
	fmt.Fprintln(w, "SNo\tPriority\tName\tStatus\t")
	for _, item := range items {
		if allOpt || (item.Done == doneOpt) {
			fmt.Fprintln(w, item.Label()+"\t"+item.PrettyP()+"\t"+item.Text+"\t"+item.PrettyD()+"\t")
		}
	}
	w.Flush()
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
