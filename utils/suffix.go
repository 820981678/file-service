package utils

import (
	"bufio"
	"io"
	"os"
	"strings"
)

var SuffixMap map[string]string

var SuffixMap_F map[string]string

var suffixValueIndex = 100

func Init_suffix() {
	LoadLogger.Infoln("load suffix start...")

	SuffixMap = make(map[string]string)
	SuffixMap_F = make(map[string]string)

	file, err := os.Open("suffix.properties")
	if err != nil {
		LoadLogger.Fatalln(err)
	}
	defer file.Close()

	buff := bufio.NewReader(file)
	for {
		line, err := buff.ReadString('\n')
		line = strings.Replace(line, "\n", "", -1)
		var temp []string = strings.Split(line, "=")
		if len(temp) != 2 {
			LoadLogger.Fatalln("pei zhi chan shu error")
		}

		//check
		if !strings.HasPrefix(temp[0], ".") || len(temp[1]) != 3 {
			LoadLogger.Fatalln("pei zhi chan shu ge shi error")
		}

		SuffixMap[temp[0]] = temp[1]

		if io.EOF == err {
			break
		}
	}

	SuffixMap_F = make(map[string]string)
	for k, v := range SuffixMap {
		SuffixMap_F[v] = k
	}

	LoadLogger.Infoln("suffix load success")
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
