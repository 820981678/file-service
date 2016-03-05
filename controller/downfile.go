package controller

import (
	"file-service/utils"
	"fmt"
	"net/http"
)

func DownFile(response http.ResponseWriter, request *http.Request) {
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
	suffix := utils.GetSuffixValue_F(fileid[32:35])

	filepath := utils.ApplicationConfig.SaveRootPath + fileid + suffix

	fmt.Println("<--- down file fileid: " + fileid + " suffix: " + suffix)

	if !(suffix == ".jpg" || suffix == ".jpeg" || suffix == ".png" || suffix == ".bmp" || suffix == ".jif") {
		response.Header().Add("content-disposition", "attachment;filename="+fileid+suffix)
	}
	http.ServeFile(response, request, filepath)
}
