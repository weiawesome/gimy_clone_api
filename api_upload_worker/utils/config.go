package utils

const (
	replicas = 1

	minIOUseSSL = false

	originalFileDescription = "-original-file"
)

func GetDefaultReplicas() int {
	return replicas
}
func GetMinIOUseSSL() bool {
	return minIOUseSSL
}
func GetOriginalFileDescription() string {
	return originalFileDescription
}
