package main

import (
	"io"
	"time"
)

type Artifact interface {
	Title() string
	Creators() []string
	Created() time.Time
}

// Text представляет собой газеты и журналы
type Text interface {
	Pages() int
	Words() int
	PageSize() int
}

type Audio interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string // mp4, etc.int
	Resolution() (x, y int)
}

type Streamer interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string
}
