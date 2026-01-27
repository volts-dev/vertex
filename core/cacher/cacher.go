package cacher

import (
	"github.com/volts-dev/cacher"
	_ "github.com/volts-dev/cacher/memory"
)

var (
	defaultCacher cacher.ICacher
)

type CacheBlock = cacher.CacheBlock

func Default() cacher.ICacher {
	if defaultCacher == nil {
		defaultCacher, _ = cacher.New("memory")
	}
	defaultCacher.Active(true)
	return defaultCacher
}
