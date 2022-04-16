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
	closedSearchData := searchGithub(githubToken, closedSearchQuery) //authenticated

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
	data := getClosed("", repo, githubDateLastWeek, githubDateToday) //unathenticated
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
	data := getOpen("", repo, githubDateLastWeek, githubDateToday) //unathenticated
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
	data := getDraft("", repo, githubDateLastWeek, githubDateToday) //unathenticated
	var issues IssuesCombined
	issues.DraftIssues = data
	expected := *issues.DraftIssues.Issues[0].ID
	var actual int64 = 1200189459
	if actual != expected {
		t.Error(fmt.Sprintf("GetDraft failed, got %d want %d", actual, expected))
	}
}
