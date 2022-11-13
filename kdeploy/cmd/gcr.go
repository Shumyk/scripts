package cmd

import (
	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/google"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

const (
	DEFAULT_REGISTRY = "us.gcr.io"
	REPOSITORY       = ""
)

func ListRepoImages(ch chan<- *google.Tags) {
	google.NewGcloudAuthenticator()
	repo, _ := name.NewRepository(REPOSITORY+microservice, name.WithDefaultRegistry(DEFAULT_REGISTRY))
	tags, _ := google.List(repo, google.WithAuthFromKeychain(authn.DefaultKeychain))
	ch <- tags
}
