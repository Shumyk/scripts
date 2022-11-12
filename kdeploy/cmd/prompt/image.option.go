package cmd

import (
	"fmt"
	"sort"
	"time"

	util "shumyk/kdeploy/cmd/util"

	"github.com/google/go-containerregistry/pkg/v1/google"
)

type ImageOption struct {
	Created time.Time
	Tags    []string
	Digest  string
}

func of(m google.ManifestInfo, d string) ImageOption {
	return ImageOption{
		Created: m.Created,
		Tags:    m.Tags,
		Digest:  d,
	}
}

func (o ImageOption) Stringify() string {
	return fmt.Sprintf(
		"%v%v%v%v%v",
		util.Date(o.Created), util.DIVIDER,
		util.TrimDigestPrefix(o.Digest), util.DIVIDER,
		util.ToString(o.Tags),
	)
}

func Stringify(options []ImageOption) (res []string) {
	for _, option := range options {
		res = append(res, option.Stringify())
	}
	return
}

func ImageOptions(tags *google.Tags) (options []ImageOption) {
	for digest, manifest := range tags.Manifests {
		options = append(options, of(manifest, digest))
	}
	return sorted(options)
}

func sorted(options []ImageOption) []ImageOption {
	sort.SliceStable(options, sortByCreated(options))
	return options
}

func sortByCreated(options []ImageOption) func(i, j int) bool {
	return func(i, j int) bool {
		return options[i].Created.After(options[j].Created)
	}
}
