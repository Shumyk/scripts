package model

import (
	"github.com/google/go-containerregistry/pkg/v1/google"
	util "shumyk/kdeploy/cmd/util"
)

func ImageOptionsOfTags(manifests map[string]google.ManifestInfo) []ImageOption {
	options := util.MapToSliceMapping(manifests, ImageOptionOfManifest)
	return Sorted(options)
}

func ImageOptionOfManifest(digest string, manifest google.ManifestInfo) ImageOption {
	return ImageOption{
		Created: manifest.Created,
		Tags:    manifest.Tags,
		Digest:  digest,
	}
}

func ImageOptionsOfPrevImages(inputs []PrevImage) []ImageOption {
	imageOptions := util.SliceMapping(inputs, ImageOptionOfPrevImage)
	return Sorted(imageOptions)
}

func ImageOptionOfPrevImage(prevImage PrevImage) ImageOption {
	return ImageOption{
		Created: prevImage.Deployed,
		Tags:    []string{prevImage.Tag},
		Digest:  prevImage.Digest,
	}
}
