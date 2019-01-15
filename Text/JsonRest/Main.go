package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"./github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	jstext, err := json.MarshalIndent(result.Items[0], " ", "  ")
	fmt.Printf("%s", jstext)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}

}
