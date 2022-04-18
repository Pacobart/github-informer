package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"sync"
	"text/template"
	"time"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type IssuesCombined struct {
	Repo         string
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

func (issues *IssuesCombined) getClosed(authToken string, repo string, startDate string, endDate string) *github.IssuesSearchResult {
	closedSearchQuery := fmt.Sprintf("repo:%s is:pr is:closed merged:%s..%s", repo, startDate, endDate)
	closedSearchData := searchGithub(authToken, closedSearchQuery)
	issues.ClosedIssues = closedSearchData
	return closedSearchData
}

func (issues *IssuesCombined) getOpen(authToken string, repo string, startDate string, endDate string) *github.IssuesSearchResult {
	openSearchQuery := fmt.Sprintf("repo:%s is:pr is:open created:%s..%s", repo, startDate, endDate)
	openSearchData := searchGithub(authToken, openSearchQuery)
	issues.OpenIssues = openSearchData
	return openSearchData
}

func (issues *IssuesCombined) getDraft(authToken string, repo string, startDate string, endDate string) *github.IssuesSearchResult {
	draftSearchQuery := fmt.Sprintf("repo:%s is:pr is:draft created:%s..%s", repo, startDate, endDate)
	draftSearchData := searchGithub(authToken, draftSearchQuery)
	issues.DraftIssues = draftSearchData
	return draftSearchData
}

func buildPrintMessage(issues IssuesCombined, fromAddress string, toAddress string) string {
	printText := ""
	printText += fmt.Sprintf("From: %s\n", fromAddress)
	printText += fmt.Sprintf("To: %s\n", toAddress)
	printText += fmt.Sprintf("Subject: Last weeks insights for: %s\n", issues.Repo)
	printText += "Body:\n"
	printText += "Over the last week:\n"
	printText += fmt.Sprintf("%d Pull Requests Closed\n", len(issues.ClosedIssues.Issues))
	for _, v := range issues.ClosedIssues.Issues {
		printText += fmt.Sprintf(" - %s\n", *v.Title)
	}
	printText += fmt.Sprintf("%d Pull Requests Open\n", len(issues.OpenIssues.Issues))
	for _, v := range issues.OpenIssues.Issues {
		printText += fmt.Sprintf(" - %s\n", *v.Title)
	}
	printText += fmt.Sprintf("%d Pull Requests in Draft state\n", len(issues.DraftIssues.Issues))
	for _, v := range issues.DraftIssues.Issues {
		printText += fmt.Sprintf(" - %s\n", *v.Title)
	}

	return printText
}

func sendEmail(issues IssuesCombined, fromAddress string, toAddresses []string, dryrun bool) (string, error) {
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	authPassword := os.Getenv("SMTP_PASSWORD")

	if fromAddress != "" {
		auth := smtp.PlainAuth("", fromAddress, authPassword, smtpHost)

		t, _ := template.ParseFiles("email_template.html")

		var body bytes.Buffer

		mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
		body.Write([]byte(fmt.Sprintf("Subject: Weekly GitHub Pull Request Data for: %s \n%s\n\n", issues.Repo, mimeHeaders)))

		closedIssueshtml := ""
		for _, v := range issues.ClosedIssues.Issues {
			closedIssueshtml += fmt.Sprintf("<li><a href=\"%s\">%s</a></li>", *v.HTMLURL, *v.Title)
		}

		openIssueshtml := ""
		for _, v := range issues.OpenIssues.Issues {
			openIssueshtml += fmt.Sprintf("<a href=\"%s\"><li>%s</a></li>", *v.HTMLURL, *v.Title)
		}

		draftIssueshtml := ""
		for _, v := range issues.DraftIssues.Issues {
			draftIssueshtml += fmt.Sprintf("<a href=\"%s\"><li>%s</a></li>", *v.HTMLURL, *v.Title)
		}

		t.Execute(&body, struct {
			ClosedIssueCount int
			OpenIssueCount   int
			DraftIssueCount  int
			ClosedIssues     string
			OpenIssues       string
			DraftIssues      string
		}{
			ClosedIssueCount: len(issues.ClosedIssues.Issues),
			OpenIssueCount:   len(issues.OpenIssues.Issues),
			DraftIssueCount:  len(issues.DraftIssues.Issues),
			ClosedIssues:     closedIssueshtml,
			OpenIssues:       openIssueshtml,
			DraftIssues:      draftIssueshtml,
		})

		if !dryrun {
			err := smtp.SendMail(smtpHost+":"+smtpPort, auth, fromAddress, toAddresses, body.Bytes())
			if err != nil {
				return "", fmt.Errorf(err.Error())
			}
			return fmt.Sprintf("email sent to: %s", toAddresses), nil
		} else {
			return "dry run enabled, no email sent", nil
		}

	} else {
		return "", fmt.Errorf("from email address required")
	}
}

func main() {
	currentTime := time.Now()
	githubDateToday := currentTime.Format("2006-01-02")
	githubDateLastWeek := currentTime.Add(-(7 * 24) * time.Hour).Format("2006-01-02")
	fmt.Printf("using daterange: %s - %s\n", githubDateLastWeek, githubDateToday)

	githubToken := os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN")
	fromEmailAddress := os.Getenv("EMAIL_ADDRESS_FROM")
	toEmailAddress := os.Getenv("EMAIL_ADDRESS_TO")
	repo := "freeCodeCamp/freeCodeCamp"

	var wg sync.WaitGroup

	var issues IssuesCombined
	issues.Repo = repo
	issues.StartDate = githubDateLastWeek
	issues.EndDate = githubDateToday

	wg.Add(1)
	go func() {
		defer wg.Done()
		issues.getClosed(githubToken, repo, githubDateLastWeek, githubDateToday)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		issues.getOpen(githubToken, repo, githubDateLastWeek, githubDateToday)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		issues.getDraft(githubToken, repo, githubDateLastWeek, githubDateToday)
	}()
	wg.Wait()

	wg.Add(1)
	go func() {
		defer wg.Done()
		var toAddresses = []string{toEmailAddress}
		emailStatus, err := sendEmail(issues, fromEmailAddress, toAddresses, false)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(emailStatus)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		printMessage := buildPrintMessage(issues, fromEmailAddress, toEmailAddress)
		fmt.Println(printMessage)
	}()

	wg.Wait()
}
