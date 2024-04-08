package redis

const (
	minCacheMinute = 30
	maxCacheMinute = 60
)

func GetMinCacheMinute() int {
	return minCacheMinute
}

func GetMaxCacheMinute() int {
	return maxCacheMinute
}
