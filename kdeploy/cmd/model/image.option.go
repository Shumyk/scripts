package model

import (
	"time"

	util "shumyk/kdeploy/cmd/util"
)

type ImageOption struct {
	Created time.Time
	Tags    []string
	Digest  string
}

func (o ImageOption) String() string {
	return util.FormatImageOption(o.Created, o.Digest, o.Tags...)
}

type ImageOptions []ImageOption

func (opts ImageOptions) Stringify() []string {
	return util.SliceMapping(opts, ImageOption.String)
}
