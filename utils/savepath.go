package utils

import (
	"strings"

	"github.com/satori/go.uuid"

	"file-service/global"
)

func GetSavePath(filehz string) (string, string) {
	id := uuid.NewV4().String()

	id = strings.Replace(id, "-", "", -1)

	return id, global.ApplicationConfig.SaveRootPath + id + filehz
}
