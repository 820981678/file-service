package utils

import (
	"github.com/satori/go.uuid"
	"os"
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

	t := strings.Replace(time.Now().Format("2006-01-02"), "-", "", -1)
	dirpath := ApplicationConfig.SaveRootPath + t + "/"
	err := CheckDirAndCreate(dirpath)
	if err != nil {
		return "", "", false
	}

	id = id + suffix_id + t
	return id, dirpath + id + filehz, true
}

func GetReadPath(fileid string) string {
	suffix := GetSuffixValue_F(fileid[32:35])
	t := fileid[35:]
	filepath := ApplicationConfig.SaveRootPath + t + "/" + fileid + suffix
	return filepath
}

func CheckDirAndCreate(path string) error {
	info, err := os.Stat(path)
	temp := true
	if err != nil {
		temp = os.IsExist(err)
	} else {
		temp = info.IsDir()
	}

	if !temp {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}
