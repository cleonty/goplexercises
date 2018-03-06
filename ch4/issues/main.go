package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/cleonty/gopl/ch4/github"
)

func getCategory(issue *github.Issue) string {
	now := time.Now()
	created := issue.CreatedAt
	if now.Year() != created.Year() {
		return "Более года назад"
	}
	if now.Month() != created.Month() {
		return "В прошлом месяце и ранее"
	}
	return "В этом месяце"
}

func main() {
	var lastCategory string
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d тем\n", result.TotalCount)
	for _, item := range result.Items {
		category := getCategory(item)
		if category != lastCategory {
			fmt.Printf("%s\n", category)
			lastCategory = category
		}
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
