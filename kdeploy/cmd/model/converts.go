package model

import (
	"github.com/google/go-containerregistry/pkg/v1/google"
	util "shumyk/kdeploy/cmd/util"
)

func ImageOptionsOfTags(manifests map[string]google.ManifestInfo) ImageOptions {
	return util.MapToSliceMapping(manifests, ImageOptionOfManifest)
}

func ImageOptionOfManifest(digest string, manifest google.ManifestInfo) ImageOption {
	return ImageOption{
		Created: manifest.Created,
		Tags:    manifest.Tags,
		Digest:  digest,
	}
}

func ImageOptionsOfPrevImages(inputs PreviousImages) ImageOptions {
	return util.SliceMapping(inputs, ImageOptionOfPrevImage)
}

func ImageOptionOfPrevImage(prevImage PreviousImage) ImageOption {
	return ImageOption{
		Created: prevImage.Deployed,
		Tags:    []string{prevImage.Tag},
		Digest:  prevImage.Digest,
	}
}
