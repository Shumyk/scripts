package model

import (
	"fmt"
	"time"

	util "shumyk/kdeploy/cmd/util"
)

// ImageOptionFormat : 2006-01-02 15:04:05     7d639e...     [tags]
const ImageOptionFormat = "%v%v%v%v%v"

type ImageOption struct {
	Created time.Time
	Tags    []string
	Digest  string
}

func (o ImageOption) Stringify() string {
	return fmt.Sprintf(
		ImageOptionFormat,
		util.Date(o.Created),
		util.Divider,
		util.TrimDigestPrefix(o.Digest),
		util.Divider,
		util.JoinComma(o.Tags),
	)
}

func Stringify(inputs []ImageOption) []string {
	return util.SliceMapping(inputs, ImageOption.Stringify)
}
