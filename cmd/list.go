/*
Copyright © 2025 Phurba Sherpa <phurba1404@gmail.com>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/phurba-sherpa/tri/todo"
	"github.com/spf13/cobra"
)

func listTodo(cmd *cobra.Command, args []string) {
	item, err := todo.ReadItems("tridos.json")

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("%+v", item)
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
