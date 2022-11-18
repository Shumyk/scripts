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

func ImageOptionOf(m google.ManifestInfo, d string) ImageOption {
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

func Stringify(options []ImageOption) []string {
	result := make([]string, len(options))
	for i, option := range options {
		result[i] = option.Stringify()
	}
	return result
}

func ImageOptionsOfTags(tags *google.Tags) []ImageOption {
	results := make([]ImageOption, len(tags.Manifests))
	position := 0
	for digest, manifest := range tags.Manifests {
		results[position] = ImageOptionOf(manifest, digest)
		position++
	}
	return sorted(results)
}

func ImageOptionsOfPrevImages(prevs []PrevImage) []ImageOption {
	results := make([]ImageOption, len(prevs))
	for i, prev := range prevs {
		results[i] = ImageOption{
			prev.Deployed,
			[]string{prev.Tag},
			prev.Digest,
		}
	}
	return sorted(results)
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
