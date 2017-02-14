package drainer

import (
	"fmt"
	"os"
	"strconv"

	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type GitHubDrainer struct {
	token     string
	prNumber  int
	ownerName string
	repoName  string
	client    *github.Client
}

func NewGitHubDrainer() (*GitHubDrainer, error) {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("environment variable GITHUB_TOKEN is not found")
	}

	prNumberStr := os.Getenv("CIRCLE_PR_NUMBER")

	if prNumberStr == "" {
		pullRequestURL := os.Getenv("CI_PULL_REQUEST")

		xs := strings.Split(pullRequestURL, "/")
		prNumberStr = xs[len(xs)-1]
	}

	if prNumberStr == "" {
		return nil, fmt.Errorf("environment variable CIRCLE_PR_NUMBER is not found")
	}

	prNumber, err := strconv.Atoi(prNumberStr)

	if err != nil {
		return nil, fmt.Errorf("cannot convert environment variable CIRCLE_PR_NUMBER: %s to int", prNumberStr)
	}

	u := os.Getenv("CIRCLE_PROJECT_USERNAME")
	r := os.Getenv("CIRCLE_PROJECT_REPONAME")

	if u == "" || r == "" {
		return nil, fmt.Errorf("environment variables CIRCLE_PROJECT_USERNAME or CIRCLE_PROJECT_REPONAME are not found")
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)

	d := &GitHubDrainer{
		token:     token,
		prNumber:  prNumber,
		ownerName: u,
		repoName:  r,
		client:    client,
	}

	return d, nil
}

func (d GitHubDrainer) Drain(message string) error {
	_, _, err := d.client.Issues.CreateComment(
		d.ownerName,
		d.repoName,
		d.prNumber,
		&github.IssueComment{
			Body: &message,
		},
	)

	if err != nil {
		return err
	}

	return nil
}
