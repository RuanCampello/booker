package list

type BookDetails struct {
	Title   string `json:"title"`
	Authors []struct {
		Author struct {
			Key string `json:"key"`
		} `json:"author"`
	} `json:"authors"`
}

type AuthorName struct {
	Name string `json:"name"`
}
