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

func ImageOptionOfManifest(digest string, manifest google.ManifestInfo) ImageOption {
	return ImageOption{
		Created: manifest.Created,
		Tags:    manifest.Tags,
		Digest:  digest,
	}
}

// ImageOptionOfPrevImage TODO: converters
func ImageOptionOfPrevImage(prevImage PrevImage) ImageOption {
	return ImageOption{
		prevImage.Deployed,
		[]string{prevImage.Tag},
		prevImage.Digest,
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

func Stringify(inputs []ImageOption) []string {
	return util.SliceMapping(inputs, ImageOption.Stringify)
}

func ImageOptionsOfTags(tags *google.Tags) []ImageOption {
	options := util.MapToSliceMapping(tags.Manifests, ImageOptionOfManifest)
	return sorted(options)
}

func ImageOptionsOfPrevImages(inputs []PrevImage) []ImageOption {
	imageOptions := util.SliceMapping(inputs, ImageOptionOfPrevImage)
	return sorted(imageOptions)
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
