package model

import (
	"github.com/google/go-containerregistry/pkg/v1/google"
	util "shumyk/kdeploy/cmd/util"
)

func ImageOptionsOfTags(manifests map[string]google.ManifestInfo) []ImageOption {
	var options ImageOptions = util.MapToSliceMapping(manifests, ImageOptionOfManifest)
	return options.Sorted()
}

func ImageOptionOfManifest(digest string, manifest google.ManifestInfo) ImageOption {
	return ImageOption{
		Created: manifest.Created,
		Tags:    manifest.Tags,
		Digest:  digest,
	}
}

func ImageOptionsOfPrevImages(inputs []PreviousImage) ImageOptions {
	var imageOptions ImageOptions = util.SliceMapping(inputs, ImageOptionOfPrevImage)
	return imageOptions.Sorted()
}

func ImageOptionOfPrevImage(prevImage PreviousImage) ImageOption {
	return ImageOption{
		Created: prevImage.Deployed,
		Tags:    []string{prevImage.Tag},
		Digest:  prevImage.Digest,
	}
}
