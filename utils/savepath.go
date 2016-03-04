package utils

import (
	"strings"

	"github.com/satori/go.uuid"

	"file-service/global"
)

func GetSavePath(filehz string) (string, string, bool) {
	id := uuid.NewV4().String()
	id = strings.Replace(id, "-", "", -1)

	ok, suffix_id := GetSuffixValue(filehz)
	if !ok {
		return "", "", false
	}

	id = id + suffix_id
	return id, global.ApplicationConfig.SaveRootPath + id + filehz, true
}
