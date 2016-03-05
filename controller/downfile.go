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
	suffix := utils.GetSuffixValue_F(fileid[32:])
	filepath := utils.ApplicationConfig.SaveRootPath + fileid + suffix

	fmt.Println("<--- down file fileid: " + fileid + " suffix: " + suffix)

	http.ServeFile(response, request, filepath)
}
