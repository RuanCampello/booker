package list

import (
	"booker/db"
	"fmt"
	"os"

	"github.com/alexeyco/simpletable"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all books in your bookshelf",
	Run: func(cmd *cobra.Command, args []string) {
		listBooks()
	},
}

func Execute() {
	rootCmd := &cobra.Command{Use: "booker"}
	rootCmd.AddCommand(ListCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error executing command:", err)
		os.Exit(1)
	}
}

func listBooks() {
	books, err := db.ReadBooks()
	if err != nil {
		fmt.Println("Error reading books from database:", err)
		return
	}

	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Title"},
			{Align: simpletable.AlignCenter, Text: "Author"},
			{Align: simpletable.AlignRight, Text: "Status"},
		},
	}

	var cells [][]*simpletable.Cell
	for _, book := range books {
		status := lime(book.Status)
		if book.Status == "reading" {
			status = lavender(book.Status)
		} else if book.Status == "read" {
			status = purple(book.Status)
		}

		cells = append(cells, []*simpletable.Cell{
			{Text: book.Title},
			{Text: book.Author},
			{Text: status, Align: simpletable.AlignRight},
		})
	}

	table.Body = &simpletable.Body{
		Cells: cells,
	}

	table.Footer = &simpletable.Footer{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Span: 3, Text: "That's your bookshelf"},
		},
	}
	table.SetStyle(simpletable.StyleUnicode)
	table.Println()
}
