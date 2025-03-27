package smgithub

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/cli/oauth"
	"github.com/google/go-github/v70/github"
	"github.com/pkg/browser"
)

const CLIENT_ID = "Iv23liEjj0fbjwLQLDY0"

var SCOPES = []string{"gist"}

var Client *github.Client

func Login(ctx context.Context, token string, httpClient *http.Client) error {
	Client = github.NewClient(httpClient).WithAuthToken(token)
	_, _, err := Client.Users.Get(ctx, "")

	return err
}

func UnauthenticatedLogin(httpClient *http.Client) {
	Client = github.NewClient(httpClient)
}

func IsLoggedIn(ctx context.Context) bool {
	if Client == nil {
		return false
	}
	if _, err := GetUsername(ctx); err != nil {
		return false
	}

	return true
}

func GetUsername(ctx context.Context) (string, error) {
	username, _, err := Client.Users.Get(ctx, "")
	if err != nil {
		return "", err
	}

	return username.GetName(), nil
}

// Shamelessly copied from GitHub's cli
func AuthFlow(isInteractive bool) string {
	w := os.Stderr

	host, err := oauth.NewGitHubHost("https://github.com")
	if err != nil {
		panic(err)
	}
	flow := &oauth.Flow{
		Host:         host,
		ClientID:     CLIENT_ID,
		ClientSecret: "",                          // only applicable to web app flow
		CallbackURI:  "http://127.0.0.1/callback", // only applicable to web app flow
		Scopes:       SCOPES,
		DisplayCode: func(code, verificationURL string) error {
			fmt.Fprintf(w, "! First copy your one-time code: %s\n", code)
			return nil
		},
		BrowseURL: func(authURL string) error {
			if u, err := url.Parse(authURL); err == nil {
				if u.Scheme != "http" && u.Scheme != "https" {
					return fmt.Errorf("invalid URL: %s", authURL)
				}
			} else {
				return err
			}

			if !isInteractive {
				fmt.Fprintf(w, "Open this URL to continue in your web browser: %s\n", authURL)
				return nil
			}

			fmt.Fprintf(w, "Press Enter to open %s in your browser... ", authURL)
			_ = waitForEnter(os.Stdin)

			if err := browser.OpenURL(authURL); err != nil {
				fmt.Fprintf(w, "! Failed opening a web browser at %s\n", authURL)
				fmt.Fprintf(w, "  %s\n", err)
				fmt.Fprint(w, "  Please try entering the URL in your browser manually\n")
			}
			return nil
		},
	}

	accessToken, err := flow.DetectFlow()
	if err != nil {
		panic(err)
	}

	return accessToken.Token
}

func waitForEnter(r io.Reader) error {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	return scanner.Err()
}
