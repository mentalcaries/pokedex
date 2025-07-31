package pokecache

import (
	"time"
)

const cacheInterval = 5 * time.Minute // 5 minute interval

var ApiCache Cache = NewCache(cacheInterval)