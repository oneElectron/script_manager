package github_connection

import (
	"context"

	"github.com/google/go-github/v70/github"
)


func ListGists(ctx context.Context) ([]*github.Gist, error) {
	gists, _, err := Client.Gists.List(ctx, "oneElectron", &github.GistListOptions{})

	if err != nil {
		return gists, err
	}

	return gists, nil
}

func ReadGist(ctx context.Context, gistID *string) (*github.Gist, error) {
	gist, _, err := Client.Gists.Get(ctx, *gistID)
	if err != nil {
		return gist, err;
	}

	return gist, nil;
}

func CreateGist(ctx context.Context, gistID string, description string, contents map[string]string, public bool) (*github.Gist, error) {
	gist := createGist(contents, description, public)

	ogist, _, err := Client.Gists.Create(ctx, &gist)

	return ogist, err
}

func createGist(contents map[string]string, description string, public bool) github.Gist {
	c := mapmap(contents, func(k string, v string) (github.GistFilename, github.GistFile) {
		gistFile := github.GistFile{
			Content: &v,
		}

		return github.GistFilename(k), gistFile
	})

	gist := github.Gist{
		Description: &description,
		Public: &public,
		Files: c,
	}

	return gist
}

func EditGist(ctx context.Context, gistID string, contents map[string]string, description string, public bool) (*github.Gist, error) {
	gist := createGist(contents, description, public)

	ogist, _, err := Client.Gists.Edit(ctx, gistID, &gist)

	return ogist, err
}

func mapmap[T comparable, U any, V comparable, W any](m map[T]U, mapper func(T, U) (V, W)) map[V]W {
	output := make(map[V]W)

	for key, value := range m {
		k, v := mapper(key, value)

		output[k] = v
	}

	return output
}
