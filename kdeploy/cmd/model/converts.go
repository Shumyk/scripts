package model

import (
	"github.com/google/go-containerregistry/pkg/v1/google"
)

func ImageOptionOfManifest(digest string, manifest google.ManifestInfo) ImageOption {
	return ImageOption{
		Created: manifest.Created,
		Tags:    manifest.Tags,
		Digest:  digest,
	}
}

func ImageOptionOfPrevImage(prevImage PrevImage) ImageOption {
	return ImageOption{
		Created: prevImage.Deployed,
		Tags:    []string{prevImage.Tag},
		Digest:  prevImage.Digest,
	}
}
