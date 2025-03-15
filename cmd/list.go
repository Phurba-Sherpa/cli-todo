/*
Copyright © 2025 Phurba Sherpa <phurba1404@gmail.com>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/phurba-sherpa/tri/todo"
	"github.com/spf13/cobra"
)

func listTodo(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(dataFile)

	if err != nil {
		fmt.Printf("file path: %s\n", dataFile)
		log.Fatal(err)
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.TabIndent|tabwriter.Debug)
	fmt.Fprintln(w, "Name\tPriority\t")
	for _, item := range items {
		fmt.Fprintln(w, strconv.Itoa(item.Priority)+"\t"+item.Text+"\t")
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
}
