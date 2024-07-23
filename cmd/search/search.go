package search

import (
	"booker/db"
	"booker/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var SearchCmd = &cobra.Command{
	Use:   "search [book name]",
	Short: "Search for books by name",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		searchBooks(args[0])
	},
}

func Execute() {
	rootCmd := &cobra.Command{Use: "booker"}
	rootCmd.AddCommand(SearchCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error executing command:", err)
		os.Exit(1)
	}
}

type OpenLibraryResponse struct {
	Docs []struct {
		Key    string   `json:"key"`
		Title  string   `json:"title"`
		Author []string `json:"author_name"`
	} `json:"docs"`
}

func searchBooks(query string) {
	url := fmt.Sprintf("https://openlibrary.org/search.json?q=%s", url.QueryEscape(query))
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		fmt.Println("Error: receiving status code", response.StatusCode)
		return
	}

	var result OpenLibraryResponse
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}

	var books []string
	bookMap := make(map[string]string)

	for _, doc := range result.Docs[:10] {
		var author string

		if len(doc.Author) > 0 {
			author = doc.Author[0]
		}
		displayText := fmt.Sprintf("%s by %s (ID: %s)", doc.Title, author, doc.Key)
		books = append(books, displayText)
		bookMap[displayText] = doc.Key
	}

	prompt := promptui.Select{
		Label: "Select a book to add to your bookshelf",
		Items: books,
		Size:  10,
	}

	_, selected, err := prompt.Run()
	if err != nil {
		fmt.Println("Error selecting book:", err)
		return
	}

	selectId, ok := bookMap[selected]
	if !ok {
		fmt.Println("Error: selected book ID not found.")
		return
	}

	var selectedBookAuthor, selectedTitle string
	for _, doc := range result.Docs {
		if doc.Key == selectId {
			selectedTitle = doc.Title
			selectedBookAuthor = doc.Author[0]
			break
		}
	}

	status := utils.PromptForStatus()
	db.AddBookToDB(selectId, status, selectedBookAuthor, selectedTitle)
}
