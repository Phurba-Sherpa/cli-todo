/*
Copyright © 2025 Phurba Sherpa <phurba1404@gmail.com>
*/
package cmd

import (
	"log"

	"github.com/phurba-sherpa/tri/todo"
	"github.com/spf13/cobra"
)

func createTodo(cmd *cobra.Command, args []string) {
	items := []todo.Item{}
	for _, arg := range args {
		item := todo.Item{
			Text: arg,
		}
		items = append(items, item)
	}

	if err := todo.SaveItem("tridos.json", items); err != nil {
		log.Fatal(err)
	}
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new todo",
	Long:  `Add will create a new to do item to the list`,
	Run:   createTodo,
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
