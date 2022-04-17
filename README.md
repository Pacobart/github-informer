# github-informer

[![tests](https://github.com/Pacobart/github-informer/actions/workflows/actions.yml/badge.svg)](https://github.com/Pacobart/github-informer/actions/workflows/actions.yml)

## Prompt:


Using the language of your choice, write code that will use the GitHub API to retrieve a summary of all opened, closed, and in draft pull requests in the last week for a given repository and send a summary email to a configurable email address. Choose any public target GitHub repository you like that has had at least 3 pull requests in the last week. Format the content email as you see fit, with the goal to allow the reader to easily digest the events of the past week. If sending email is not an option, then please print to console the details of the email you would send (From, To, Subject, Body).


Push your solution to a fresh GitHub or BitBucket repository and share the repository with GitHub user `Andrew-kim-sp` at least 1 business day prior to your scheduled “final” interview. Reply to this email with a link to the repo once done.


During the next interview, we ask that you share your screen and demonstrate the solution. We will also talk through your code and the design considerations you took in the implementation.

## Details

App looks for environment variables:
- `GITHUB_PERSONAL_ACCESS_TOKEN` = gitlab personal access token
- `EMAIL_ADDRESS_FROM` = email address to send summary from
- `EMAIL_ADDRESS_TO` = email address to send summary to
- `SMTP_PASSEORD` = password associated with `EMAIL_ADDRESS_FROM`


To send emails using gmail's smtp server, allow less secure apps: https://myaccount.google.com/lesssecureapps

Makefile created to debug, test, build etc.

## Todo

- email
- tests


## References

- https://docs.github.com/en/search-github/searching-on-github/searching-issues-and-pull-requests


## email template:

```
From: parker
To: parker
Suject: Last weeks insights
Body:
Over the last week:
X Pull Requests Merged
  - results list
X Pull Requests Open
  - results list
X Pull Requests in Draft state
  - results list
```