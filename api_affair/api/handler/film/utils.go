package film

import (
	"api_affair/utils"
	"strconv"
)

func GetOffset(offset string) int32 {
	if offset == utils.GetDefaultValue() {
		return utils.GetDefaultOffset()
	}
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		return utils.GetDefaultOffset()
	}
	return int32(offsetInt)
}
func GetLimit(limit string) int32 {
	if limit == utils.GetDefaultValue() {
		return utils.GetDefaultLimit()
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return utils.GetDefaultLimit()
	}
	return int32(limitInt)
}

func GetReleaseYear(releaseYear string) uint32 {
	if releaseYear == utils.GetDefaultValue() {
		return utils.GetDefaultReleaseYear()
	}
	releaseYearValue, err := strconv.Atoi(releaseYear)
	if err != nil {
		return utils.GetDefaultReleaseYear()
	}
	return uint32(releaseYearValue)
}
