package controller

import (
	"file-service/utils"
	"fmt"
	"net/http"
	"os"
	"path"
	"time"
)

func DownFile(response http.ResponseWriter, request *http.Request) {
	start := time.Now().UnixNano()

	err := request.ParseForm()
	if err != nil {
		response.Write([]byte(`{"code":-1}`))
		return
	}

	fileid := request.Form.Get("fileid")
	if len(fileid) != 43 {
		response.Write([]byte(`"code":-1, "msg":"fileid error"`))
		return
	}

	filepath := utils.GetReadPath(fileid)
	//check file
	_, err = os.Stat(filepath)
	if err != nil {
		if !os.IsExist(err) {
			response.Write([]byte(`"code":-1, "msg":"file not found"`))
			return
		}
	}

	suffix := path.Ext(filepath)
	if !(suffix == ".jpg" || suffix == ".jpeg" || suffix == ".png" || suffix == ".bmp" || suffix == ".jif") {
		response.Header().Add("content-disposition", "attachment;filename="+fileid+suffix)
	}

	getpath := (time.Now().UnixNano() - start) / 1000000

	http.ServeFile(response, request, filepath)
	fmt.Printf("<--- down file fileid: %s, suffix: %s, getpath: %dms -> writefile: %dms \n",
		fileid, suffix, getpath, (time.Now().UnixNano()-start)/1000000)
}
