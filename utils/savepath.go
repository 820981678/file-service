package utils

import (
	"github.com/satori/go.uuid"
	"strings"
	"time"
)

func GetSavePath(filehz string) (string, string, bool) {
	id := uuid.NewV4().String()
	id = strings.Replace(id, "-", "", -1)

	ok, suffix_id := GetSuffixValue(filehz)
	if !ok {
		return "", "", false
	}

	t := strings.Replace(time.Now().Format("2006-01-01"), "-", "", -1)
	id = id + suffix_id + t
	return id, ApplicationConfig.SaveRootPath + id + filehz, true
}
