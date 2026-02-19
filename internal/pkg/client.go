package pkg

import (
	"context"
	"fmt"
	"os"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

func CreateClient() (githubv4.Client, error) {
	envVarName := "VANITYPROX_GITHUB_TOKEN"
	token := os.Getenv(envVarName)
	if token == "" {
		return githubv4.Client{}, fmt.Errorf("%s is not defined", envVarName)
	}
	src := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	httpClient := oauth2.NewClient(context.Background(), src)
	client := githubv4.NewClient(httpClient)
	return *client, nil
}
