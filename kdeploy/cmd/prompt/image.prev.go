package cmd

import "time"

type PrevImage struct {
	Tag      string
	Digest   string
	Deployed time.Time
}

func PrevImageOf(tag, digest string) PrevImage {
	return PrevImage{
		Tag:      tag,
		Digest:   digest,
		Deployed: time.Now(),
	}
}

func PrevImageToOptions(p []PrevImage) (o []ImageOption) {
	for _, v := range p {
		o = append(o, ImageOption{v.Deployed, []string{v.Tag}, v.Digest})
	}
	return
}
