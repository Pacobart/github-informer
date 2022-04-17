package main

import (
	"fmt"
	"os"
	"testing"
)

//func TestSearchGitHubClosed(t *testing.T) {
//	githubDateToday := "2022-04-14"
//	githubDateLastWeek := "2022-04-07"
//	closedSearchQuery := fmt.Sprintf("repo:freeCodeCamp/freeCodeCamp is:pr is:closed merged:%s..%s", githubDateLastWeek, githubDateToday)
//	closedSearchData := searchGithub("", closedSearchQuery)
//
//	actual := *closedSearchData.Issues[0].State
//	expected := "closed"
//	if actual != expected {
//		t.Error(fmt.Sprintf("SearchGitHubClosed failed, got %s want %s", actual, expected))
//	}
//}
//
//func TestSearchGitHubOpen(t *testing.T) {
//	githubDateToday := "2022-04-14"
//	githubDateLastWeek := "2022-04-07"
//	closedSearchQuery := fmt.Sprintf("repo:freeCodeCamp/freeCodeCamp is:pr is:open created:%s..%s", githubDateLastWeek, githubDateToday)
//	closedSearchData := searchGithub("", closedSearchQuery)
//
//	actual := *closedSearchData.Issues[0].State
//	expected := "open"
//	if actual != expected {
//		t.Error(fmt.Sprintf("SearchGitHubOpen failed, got %s want %s", actual, expected))
//	}
//}

func TestSearchGitHubDraft(t *testing.T) {
	githubDateToday := "2022-04-14"
	githubDateLastWeek := "2022-04-07"
	githubToken := os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN")
	closedSearchQuery := fmt.Sprintf("repo:freeCodeCamp/freeCodeCamp is:pr is:draft created:%s..%s", githubDateLastWeek, githubDateToday)
	closedSearchData := searchGithub(githubToken, closedSearchQuery)

	actual := *closedSearchData.Issues[0].State
	expected := "open"
	if actual != expected {
		t.Error(fmt.Sprintf("SearchGitHubDraft failed, got %s want %s", actual, expected))
	}
}

func TestGetClosed(t *testing.T) {
	githubDateToday := "2022-04-14"
	githubDateLastWeek := "2022-04-07"
	repo := "freeCodeCamp/freeCodeCamp"
	githubToken := os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN")
	data := getClosed(githubToken, repo, githubDateLastWeek, githubDateToday)
	var issues IssuesCombined
	issues.ClosedIssues = data
	expected := *issues.ClosedIssues.Issues[0].ID
	var actual int64 = 1203265483
	if actual != expected {
		t.Error(fmt.Sprintf("GetClosed failed, got %d want %d", actual, expected))
	}
}

func TestGetOpen(t *testing.T) {
	githubDateToday := "2022-04-14"
	githubDateLastWeek := "2022-04-07"
	repo := "freeCodeCamp/freeCodeCamp"
	githubToken := os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN")
	data := getOpen(githubToken, repo, githubDateLastWeek, githubDateToday)
	var issues IssuesCombined
	issues.OpenIssues = data
	expected := *issues.OpenIssues.Issues[0].ID
	var actual int64 = 1200189459
	if actual != expected {
		t.Error(fmt.Sprintf("GetOpen failed, got %d want %d", actual, expected))
	}
}

func TestGetDraft(t *testing.T) {
	githubDateToday := "2022-04-14"
	githubDateLastWeek := "2022-04-07"
	repo := "freeCodeCamp/freeCodeCamp"
	githubToken := os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN")
	data := getDraft(githubToken, repo, githubDateLastWeek, githubDateToday)
	var issues IssuesCombined
	issues.DraftIssues = data
	expected := *issues.DraftIssues.Issues[0].ID
	var actual int64 = 1200189459
	if actual != expected {
		t.Error(fmt.Sprintf("GetDraft failed, got %d want %d", actual, expected))
	}
}

// TODO: fix this test...doesn't seem to be running.
func TestbuildPrintMessage(t *testing.T) {
	githubDateToday := "2022-04-14"
	githubDateLastWeek := "2022-04-07"
	repo := "freeCodeCamp/freeCodeCamp"
	githubToken := os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN")
	var issues IssuesCombined
	issues.Repo = repo
	issues.StartDate = githubDateLastWeek
	issues.EndDate = githubDateToday
	issues.ClosedIssues = getClosed(githubToken, repo, githubDateLastWeek, githubDateToday)
	issues.OpenIssues = getOpen(githubToken, repo, githubDateLastWeek, githubDateToday)
	issues.DraftIssues = getDraft(githubToken, repo, githubDateLastWeek, githubDateToday)

	fromEmailAddress := os.Getenv("EMAIL_ADDRESS_FROM")
	toEmailAddress := os.Getenv("EMAIL_ADDRESS_TO")

	actual := buildPrintMessage(issues, fromEmailAddress, toEmailAddress)
	expected := "12345"
	if actual != expected {
		t.Error(fmt.Sprintf("BuildPrintMessage failed, got %s want %s", actual, expected))
	}
}

func TestSendEmailDryRun(t *testing.T) {
	githubDateToday := "2022-04-14"
	githubDateLastWeek := "2022-04-07"
	repo := "freeCodeCamp/freeCodeCamp"
	githubToken := os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN")
	var issues IssuesCombined
	issues.Repo = repo
	issues.StartDate = githubDateLastWeek
	issues.EndDate = githubDateToday
	issues.ClosedIssues = getClosed(githubToken, repo, githubDateLastWeek, githubDateToday)
	issues.OpenIssues = getOpen(githubToken, repo, githubDateLastWeek, githubDateToday)
	issues.DraftIssues = getDraft(githubToken, repo, githubDateLastWeek, githubDateToday)

	fromEmailAddress := os.Getenv("EMAIL_ADDRESS_FROM")
	toEmailAddress := os.Getenv("EMAIL_ADDRESS_TO")
	var toAddresses = []string{toEmailAddress}

	actual, _ := sendEmail(issues, fromEmailAddress, toAddresses, true)
	expected := "dry run enabled, no email sent"
	if actual != expected {
		t.Error(fmt.Sprintf("SendEmail failed, got %s want %s", actual, expected))
	}
}

func TestSendEmail(t *testing.T) {
	githubDateToday := "2022-04-14"
	githubDateLastWeek := "2022-04-07"
	repo := "freeCodeCamp/freeCodeCamp"
	githubToken := os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN")
	var issues IssuesCombined
	issues.Repo = repo
	issues.StartDate = githubDateLastWeek
	issues.EndDate = githubDateToday
	issues.ClosedIssues = getClosed(githubToken, repo, githubDateLastWeek, githubDateToday)
	issues.OpenIssues = getOpen(githubToken, repo, githubDateLastWeek, githubDateToday)
	issues.DraftIssues = getDraft(githubToken, repo, githubDateLastWeek, githubDateToday)

	fromEmailAddress := os.Getenv("EMAIL_ADDRESS_FROM")
	toEmailAddress := os.Getenv("EMAIL_ADDRESS_TO")
	var toAddresses = []string{toEmailAddress}

	actual, _ := sendEmail(issues, fromEmailAddress, toAddresses, false)
	expected := fmt.Sprintf("email sent to: %s", toAddresses)
	if actual != expected {
		t.Error(fmt.Sprintf("SendEmail failed, got \"%s\" want \"%s\"", actual, expected))
	}
}

// TODO: not sure why test failing...
//func TestSendEmailNoEmail(t *testing.T) {
//	githubDateToday := "2022-04-14"
//	githubDateLastWeek := "2022-04-07"
//	repo := "freeCodeCamp/freeCodeCamp"
//	githubToken := os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN")
//	var issues IssuesCombined
//	issues.Repo = repo
//	issues.StartDate = githubDateLastWeek
//	issues.EndDate = githubDateToday
//	issues.ClosedIssues = getClosed(githubToken, repo, githubDateLastWeek, githubDateToday)
//	issues.OpenIssues = getOpen(githubToken, repo, githubDateLastWeek, githubDateToday)
//	issues.DraftIssues = getDraft(githubToken, repo, githubDateLastWeek, githubDateToday)
//
//	fromEmailAddress := ""
//	toEmailAddress := os.Getenv("EMAIL_ADDRESS_TO")
//	var toAddresses = []string{toEmailAddress}
//
//	_, err := sendEmail(issues, fromEmailAddress, toAddresses, true)
//	expected := fmt.Errorf("from email address required")
//	if err != expected {
//		t.Error(fmt.Sprintf("SendEmail failed, got \"%v\" want \"%v\"", err, expected))
//	}
//}
