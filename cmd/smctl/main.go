package main

import (
	"context"

	"github.com/oneElectron/script_manager/internal/github_connection"
	xdg "github.com/twpayne/go-xdg/v6"
)

var DIRS *xdg.BaseDirectorySpecification

func init() {
	var err error

	DIRS, err = xdg.NewBaseDirectorySpecification();
	if err != nil {
		panic("Could not get the base directory spec")
	}
}

func main() {
	initGithub()
	Execute()
}

func initGithub() {
	ctx := context.Background()
	token := viperConf.GetString("GithubAuthToken")
	if token == "PLACEHOLDER" {
		github_connection.UnauthenticatedLogin(nil)
	}
	if err := github_connection.Login(ctx, token, nil); err != nil {
		github_connection.UnauthenticatedLogin(nil)
	}
}
