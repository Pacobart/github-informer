package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type IssuesSearchResult struct {
	Total             *int
	IncompleteResults *bool
	Issues            []*Issue
}

type Issue []struct {
	ID        *int64
	Number    *int
	State     *string
	Locked    *bool
	Title     *string
	Body      *string
	ClosedAt  *time.Time
	CreatedAt *time.Time
	UpdatedAt *time.Time
	URL       *string
	HTMLURL   *string
}

type IssuesCombined struct {
	StartDate    string
	EndDate      string
	ClosedIssues *github.IssuesSearchResult
	OpenIssues   *github.IssuesSearchResult
	DraftIssues  *github.IssuesSearchResult
}

func searchGithub(authToken string, searchQuery string) *github.IssuesSearchResult {
	var ctx = context.Background()
	var client = github.NewClient(nil)
	if authToken != "" {
		token := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: authToken},
		)
		authClient := oauth2.NewClient(ctx, token)
		client = github.NewClient(authClient)
	}
	opt := &github.SearchOptions{}
	results, _, err := client.Search.Issues(ctx, searchQuery, opt)
	if _, ok := err.(*github.RateLimitError); ok {
		log.Println("hit rate limit")
	}
	if err != nil {
		log.Fatal(err)
	}
	return results
}

func getClosed(authToken string, repo string, startDate string, endDate string) *github.IssuesSearchResult {
	closedSearchQuery := fmt.Sprintf("repo:%s is:pr is:closed merged:%s..%s", repo, startDate, endDate)
	closedSearchData := searchGithub(authToken, closedSearchQuery)
	return closedSearchData
}

func getOpen(authToken string, repo string, startDate string, endDate string) *github.IssuesSearchResult {
	openSearchQuery := fmt.Sprintf("repo:%s is:pr is:open created:%s..%s", repo, startDate, endDate)
	openSearchData := searchGithub(authToken, openSearchQuery)
	return openSearchData
}

func getDraft(authToken string, repo string, startDate string, endDate string) *github.IssuesSearchResult {
	draftSearchQuery := fmt.Sprintf("repo:%s is:pr is:draft created:%s..%s", repo, startDate, endDate)
	draftSearchData := searchGithub(authToken, draftSearchQuery)
	return draftSearchData
}

func createMessaging(issues IssuesCombined, emailAddress string, printToScreen bool, sendEmail bool) {
	if printToScreen {
		fmt.Println("From: githubsummary@somesite.com")
		fmt.Printf("To: %s\n", emailAddress)
		fmt.Println("Subject: Last weeks insights")
		fmt.Println("Body:")
		fmt.Println("Over the last week:")
		fmt.Printf("%d Pull Requests Closed\n", len(issues.ClosedIssues.Issues))
		for _, v := range issues.ClosedIssues.Issues {
			fmt.Printf(" - %s\n", *v.Title)
		}
		fmt.Printf("%d Pull Requests Open\n", len(issues.OpenIssues.Issues))
		for _, v := range issues.OpenIssues.Issues {
			fmt.Printf(" - %s\n", *v.Title)
		}
		fmt.Printf("%d Pull Requests in Draft state\n", len(issues.DraftIssues.Issues))
		for _, v := range issues.DraftIssues.Issues {
			fmt.Printf(" - %s\n", *v.Title)
		}
	}

	if sendEmail {
		fmt.Println("Will send email after I write logic.")
	}
}

func main() {
	currentTime := time.Now()
	githubDateToday := currentTime.Format("2006-01-02")
	githubDateLastWeek := currentTime.Add(-(7 * 24) * time.Hour).Format("2006-01-02")
	fmt.Printf("using daterange: %s - %s\n", githubDateLastWeek, githubDateToday)

	githubToken := os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN")
	emailAddress := os.Getenv("EMAIL_ADDRESS_TO")
	repo := "freeCodeCamp/freeCodeCamp"

	var issues IssuesCombined
	issues.StartDate = githubDateLastWeek
	issues.EndDate = githubDateToday
	issues.ClosedIssues = getClosed(githubToken, repo, githubDateLastWeek, githubDateToday)
	issues.OpenIssues = getOpen(githubToken, repo, githubDateLastWeek, githubDateToday)
	issues.DraftIssues = getDraft(githubToken, repo, githubDateLastWeek, githubDateToday)
	createMessaging(issues, emailAddress, true, false)
}
