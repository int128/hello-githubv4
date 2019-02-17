package main

import (
	"context"
	"github.com/pkg/errors"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
	"log"
	"os"
)

func run(ctx context.Context, client *githubv4.Client) error {
	{
		var q struct {
			Repository struct {
				PullRequests struct {
					Nodes []struct {
						Number int
					}
				} `graphql:"pullRequests(states: $states, first: 1, headRefName: $head, baseRefName: $base)"`
			} `graphql:"repository(owner: $owner, name: $repo)"`
		}
		v := map[string]interface{}{
			"owner": githubv4.String("octocat"),
			"repo":  githubv4.String("Spoon-Knife"),
			"states": []githubv4.PullRequestState{
				githubv4.PullRequestStateOpen,
				githubv4.PullRequestStateClosed,
				githubv4.PullRequestStateMerged,
			},
			"head": githubv4.String("example"),
			"base": githubv4.String("master"),
		}
		if err := client.Query(ctx, &q, v); err != nil {
			return errors.Wrapf(err, "error while query")
		}
		log.Printf("%+v", q)
	}
	return nil
}

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")})
	hc := oauth2.NewClient(ctx, ts)
	client := githubv4.NewClient(hc)
	if err := run(ctx, client); err != nil {
		log.Fatalf("Error: %+v", err)
	}
}
