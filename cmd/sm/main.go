package main

import (
	"context"
	"log"

	"github.com/oneElectron/script_manager/internal/smgithub"
	xdg "github.com/twpayne/go-xdg/v6"
)

var DIRS *xdg.BaseDirectorySpecification

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var err error

	DIRS, err = xdg.NewBaseDirectorySpecification();
	if err != nil {
		panic("Could not get the base directory spec")
	}
}

func main() {
	Execute()
}

func initGithub() {
	ctx := context.Background()
	token := viperConf.GetString("GithubAuthToken")
	if token == "PLACEHOLDER" {
		smgithub.UnauthenticatedLogin(nil)
	}

	smgithub.Login(ctx, token, nil)
}
