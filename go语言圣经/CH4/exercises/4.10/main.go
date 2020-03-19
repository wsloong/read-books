// 练习 4.10： 修改issues程序，根据问题的时间进行分类，比如不到一个月的、不到一年的、超过一年。

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues: \n", result.TotalCount)

	var mouthIssues []*Issue
	var lessYearIssues []*Issue
	var moreYearIssues []*Issue

	lastMouth := time.Now().AddDate(0, -1, 0)
	lastYear := time.Now().AddDate(1, 0, 0)

	for _, item := range result.Items {

		if item.CreatedAt.After(lastMouth) {
			mouthIssues = append(mouthIssues, item)
			continue
		}

		if item.CreatedAt.Before(lastMouth) && item.CreatedAt.After(lastYear) {
			lessYearIssues = append(lessYearIssues, item)
			continue
		}

		moreYearIssues = append(moreYearIssues, item)
	}

	fmt.Println("last mouth issues: ")
	for _, item := range mouthIssues {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}

	fmt.Println("last year issues:")

	for _, item := range lessYearIssues {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}

	fmt.Println("a year ago issues")

	for _, item := range moreYearIssues {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}
