package model

import "time"

type PreviousImage struct {
	Tag      string
	Digest   string
	Deployed time.Time
}

func PrevImageOf(tag, digest string) PreviousImage {
	return PreviousImage{
		Tag:      tag,
		Digest:   digest,
		Deployed: time.Now(),
	}
}
