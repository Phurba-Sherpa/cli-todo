/*
Copyright © 2025 Phurba Sherpa <phurba1404@gmail.com>
This is an awesome CLI app in Go Lang.
*/
package cmd

import (
	"log"

	"github.com/phurba-sherpa/tri/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func createTodo(cmd *cobra.Command, args []string) {
	fp := viper.GetString("datafile")
	items, err := todo.ReadItems(fp)
	if err != nil {
		log.Fatal(err)
	}

	for _, arg := range args {
		item := todo.Item{
			Text: arg,
		}
		item.SetPriority(priority)
		items = append(items, item)
	}

	if err := todo.SaveItem(fp, items); err != nil {
		log.Fatal(err)
	}
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new todo",
	Long:  `Add will create a new to do item to the list`,
	Run:   createTodo,
}

var priority int

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority: 1, 2, 3")
}
