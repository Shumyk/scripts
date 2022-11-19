package model

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
