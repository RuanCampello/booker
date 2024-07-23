package main

import (
	cmd "booker/cmd/search"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "booker",
	Short: "A CLI tool for managing your bookshelf",
}

func main() {
	cmd.Execute()
}
