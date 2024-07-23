package main

import (
	"booker/cmd/list"
	"booker/cmd/search"
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "booker",
	Short: "A CLI tool for managing your bookshelf",
}

func main() {
	rootCmd.AddCommand(search.SearchCmd)
	rootCmd.AddCommand(list.ListCmd)

	// execute the root command
	if err := rootCmd.Execute(); err != nil {
		log.Fatal("Error executing command:", err)
	}
}
