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
	var q struct {
		Viewer struct {
			Login     githubv4.String
			CreatedAt githubv4.DateTime
		}
	}
	if err := client.Query(ctx, &q, nil); err != nil {
		return errors.Wrapf(err, "error while query")
	}
	log.Printf("%+v", q)
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
