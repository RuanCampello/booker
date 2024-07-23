package utils

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func ExtractIDFromSelection(selection string) string {
	var id string
	fmt.Sscanf(selection, "%*s %*s %*s %*s (ID: %s)", &id)
	return id
}

func PromptForStatus() string {
	prompt := promptui.Select{
		Label: "Select status",
		Items: []string{"read", "reading", "to be read"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Println("Error selecting status:", err)
		return ""
	}

	return result
}
