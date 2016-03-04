package utils

import (
	"fmt"	
)

var SuffixMap map[string]string

var SuffixMap_F map[string]string

var suffixValueIndex = 100

func init() {
	SuffixMap = make(map[string]string)
	SuffixMap[".jpg"] = "100"
	SuffixMap[".jpeg"] = "101"
	SuffixMap[".png"] = "102"
	SuffixMap[".go"] = "103"
	SuffixMap[".exe"] = "104"
	SuffixMap[".rar"] = "105"

	SuffixMap_F = make(map[string]string)
	for k, v := range SuffixMap {
		SuffixMap_F[v] = k
	}

	fmt.Println(SuffixMap)
	fmt.Println(SuffixMap_F)
}

func GetSuffixValue(key string) (bool, string) {
	if _, ok := SuffixMap[key]; ok {
		return true, SuffixMap[key]
	}

	return false, ""
}

func GetSuffixValue_F(key string) string {
	return SuffixMap_F[key]
}

