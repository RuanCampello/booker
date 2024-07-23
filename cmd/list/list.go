package list

import (
	"booker/db"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

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

	for _, book := range books {
		url := fmt.Sprintf("https://openlibrary.org%s.json", book.ID)
		response, err := http.Get(url)
		if err != nil {
			fmt.Println("Error fetching book details:", err)
			continue
		}
		defer response.Body.Close()

		var details BookDetails
		if err := json.NewDecoder(response.Body).Decode(&details); err != nil {
			fmt.Println("Error decoding book details:", err)
			continue
		}

		fmt.Printf("Title: %s\nAuthor: %s\nStatus: %s\n",
			details.Title, book.Author, book.Status)
	}
}
