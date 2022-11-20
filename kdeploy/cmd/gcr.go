package cmd

import (
	. "shumyk/kdeploy/cmd/util"
	"strings"

	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/google"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

// TODO: to config
const (
	DefaultRegistry = "us.gcr.io"
	Repository      = ""
)

func ListRepoImages(ch chan<- *google.Tags) {
	_, err := google.NewGcloudAuthenticator()
	Laugh(err)

	registry := name.WithDefaultRegistry(DefaultRegistry)
	// todo refactor rep + ms
	repo, err := name.NewRepository(Repository+microservice, registry)
	Laugh(err)

	keychain := google.WithAuthFromKeychain(auth)
	tags, err := google.List(repo, keychain)
	Laugh(err)

	ch <- tags
}

func ListRepos() (results []string) {
	registry, err := name.NewRegistry(DefaultRegistry)
	Laugh(err)

	authOption := remote.WithAuthFromKeychain(auth)
	repos, err := remote.Catalog(ctx, registry, authOption)
	Laugh(err)

	return filterRepos(repos)
}

func filterRepos(reposRaw []string) (results []string) {
	for _, repoRaw := range reposRaw {
		if strings.HasPrefix(repoRaw, Repository) {
			repo := strings.TrimPrefix(repoRaw, Repository)
			results = append(results, repo)
		}
	}
	return
}
