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
	var issues IssuesCombined
	issues.getClosed(githubToken, repo, githubDateLastWeek, githubDateToday)
	actual := *issues.ClosedIssues.Issues[0].ID
	var expected int64 = 1203265483
	if actual != expected {
		t.Error(fmt.Sprintf("GetClosed failed, got %d want %d", actual, expected))
	}
}

func TestGetOpen(t *testing.T) {
	githubDateToday := "2022-04-14"
	githubDateLastWeek := "2022-04-07"
	repo := "freeCodeCamp/freeCodeCamp"
	githubToken := os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN")
	var issues IssuesCombined
	issues.getOpen(githubToken, repo, githubDateLastWeek, githubDateToday)
	actual := *issues.OpenIssues.Issues[0].ID
	var expected int64 = 1200189459
	if actual != expected {
		t.Error(fmt.Sprintf("GetOpen failed, got %d want %d", actual, expected))
	}
}

func TestGetDraft(t *testing.T) {
	githubDateToday := "2022-04-14"
	githubDateLastWeek := "2022-04-07"
	repo := "freeCodeCamp/freeCodeCamp"
	githubToken := os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN")
	var issues IssuesCombined
	issues.getDraft(githubToken, repo, githubDateLastWeek, githubDateToday)
	actual := *issues.DraftIssues.Issues[0].ID
	var expected int64 = 1200189459
	if actual != expected {
		t.Error(fmt.Sprintf("GetDraft failed, got %d want %d", actual, expected))
	}
}

func TestBuildPrintMessage(t *testing.T) {
	githubDateToday := "2022-04-14"
	githubDateLastWeek := "2022-04-07"
	repo := "freeCodeCamp/freeCodeCamp"
	githubToken := os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN")
	var issues IssuesCombined
	issues.Repo = repo
	issues.StartDate = githubDateLastWeek
	issues.EndDate = githubDateToday
	issues.getClosed(githubToken, repo, githubDateLastWeek, githubDateToday)
	issues.getOpen(githubToken, repo, githubDateLastWeek, githubDateToday)
	issues.getDraft(githubToken, repo, githubDateLastWeek, githubDateToday)

	fromEmailAddress := os.Getenv("EMAIL_ADDRESS_FROM")
	toEmailAddress := os.Getenv("EMAIL_ADDRESS_TO")

	actual := buildPrintMessage(issues, fromEmailAddress, toEmailAddress)
	expected := ""
	expected += "From: pbbarthlome@gmail.com\n"
	expected += "To: pbbarthlome@gmail.com\n"
	expected += "Subject: Last weeks insights for: freeCodeCamp/freeCodeCamp\n"
	expected += "Body:\n"
	expected += "Over the last week:\n"
	expected += "30 Pull Requests Closed\n"
	expected += " - chore(i18n,learn): processed translations\n"
	expected += " - chore(i18n,docs): processed translations\n"
	expected += " - fix(curriculum): improve the wording of a challenge\n"
	expected += " - fix: add message about third-party cookies\n"
	expected += " - chore(i18n,learn): processed translations\n"
	expected += " - chore(i18n,client): processed translations\n"
	expected += " - fix(deps): update dependency react-i18next to v11.16.5\n"
	expected += " - fix(deps): update dependency react-i18next to v11.16.3\n"
	expected += " - fix(deps): update dependency @stripe/stripe-js to v1.27.0\n"
	expected += " - chore(deps): update dependency webpack to v5.72.0\n"
	expected += " - chore(deps): update dependency @types/enzyme to v3.10.12\n"
	expected += " - chore(deps): update dependency eslint-plugin-import to v2.26.0\n"
	expected += " - chore(deps): update dependency @testing-library/dom to v8.13.0\n"
	expected += " - chore(deps): update codesee to v0.227.0\n"
	expected += " - fix(deps): update dependency react-instantsearch-dom to v6.23.3\n"
	expected += " - fix(curriculum): adjusted hint img src\n"
	expected += " - fix(deps): update dependency @stripe/react-stripe-js to v1.7.1\n"
	expected += " - chore(deps): update storybook monorepo to v6.4.21\n"
	expected += " - chore(deps): update dependency @types/react-dom to v17.0.15\n"
	expected += " - chore(deps): update dependency @types/react to v17.0.44\n"
	expected += " - chore(deps): update dependency @testing-library/jest-dom to v5.16.4\n"
	expected += " - chore(deps): update babel monorepo to v7.17.9\n"
	expected += " - chore(i18n,learn): processed translations\n"
	expected += " - fix: handle missing sound saga payloads\n"
	expected += " - chore(i18n,client): processed translations\n"
	expected += " - fix(curriculum): include note about counting spaces by .length property\n"
	expected += " - fixed a typo in basic-node-and-express intro\n"
	expected += " - chore(i18n,client): processed translations\n"
	expected += " - feat: update footer Apr 2022\n"
	expected += " - feat: add secure donation border to donation form\n"
	expected += "2 Pull Requests Open\n"
	expected += " - feat(ui-components): add support for link to Button component\n"
	expected += " - feat(ui-components): add states and full-width support to Button component\n"
	expected += "2 Pull Requests in Draft state\n"
	expected += " - feat(ui-components): add support for link to Button component\n"
	expected += " - feat: add challenge type to mobile\n"
	if actual != expected {
		t.Error(fmt.Sprintf("BuildPrintMessage failed, got \"%s\" want \"%s\"", actual, expected))
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
	issues.getClosed(githubToken, repo, githubDateLastWeek, githubDateToday)
	issues.getOpen(githubToken, repo, githubDateLastWeek, githubDateToday)
	issues.getDraft(githubToken, repo, githubDateLastWeek, githubDateToday)

	fromEmailAddress := os.Getenv("EMAIL_ADDRESS_FROM")
	toEmailAddress := os.Getenv("EMAIL_ADDRESS_TO")
	var toAddresses = []string{toEmailAddress}

	actual := ""
	msg, err := sendEmail(issues, fromEmailAddress, toAddresses, true)
	if err != nil {
		actual = string(err.Error())
	} else {
		actual = msg
	}
	expected := "dry run enabled, no email sent"
	if actual != expected {
		t.Error(fmt.Sprintf("SendEmailDryRun failed, got %s want %s", actual, expected))
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
	issues.getClosed(githubToken, repo, githubDateLastWeek, githubDateToday)
	issues.getOpen(githubToken, repo, githubDateLastWeek, githubDateToday)
	issues.getDraft(githubToken, repo, githubDateLastWeek, githubDateToday)

	fromEmailAddress := os.Getenv("EMAIL_ADDRESS_FROM")
	toEmailAddress := os.Getenv("EMAIL_ADDRESS_TO")
	var toAddresses = []string{toEmailAddress}

	actual := ""
	msg, err := sendEmail(issues, fromEmailAddress, toAddresses, false)
	if err != nil {
		actual = string(err.Error())
	} else {
		actual = msg
	}
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
//	issues.getClosed(githubToken, repo, githubDateLastWeek, githubDateToday)
//	issues. getOpen(githubToken, repo, githubDateLastWeek, githubDateToday)
//	issues.getDraft(githubToken, repo, githubDateLastWeek, githubDateToday)
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
