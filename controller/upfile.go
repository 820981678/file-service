package controller

import (
	"file-service/utils"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"time"
)

func UpFile(response http.ResponseWriter, request *http.Request) {
	var header string = ""
	for _, v := range utils.ApplicationDomain {
		header += (v + ",")
	}
	header = string([]rune(header)[ : len(header)])
	response.Header().Add("Access-Control-Allow-Origin", header)
	
	start := time.Now().UnixNano()

	file, fileh, error := request.FormFile("file")
	if error != nil {
		fmt.Println("jie xie file error")
		response.Write([]byte(`{"code":-1}`))
		return
	}
	defer file.Close()

	uptime := (time.Now().UnixNano() - start) / 1000000

	filename := fileh.Filename
	filehz := path.Ext(filename)

	id, savefilename, ok := utils.GetSavePath(filehz)
	if !ok {
		response.Write([]byte(`{"code":-1, "message": "bu bei zhi chi de wen jian"}`))
		return
	}

	uf, err := os.Create(savefilename)
	//uf, err := os.OpenFile("D://111/111.jpg", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		response.Write([]byte(`{"code":-1, "msg": "save file error"}`))
		return
	}
	defer uf.Close()

	pathtime := (time.Now().UnixNano() - start) / 1000000

	_, err = io.Copy(uf, file)
	if err != nil {
		response.Write([]byte(`{"code":-1, "msg": "save file error"}`))
		return
	}

	fmt.Printf("---> up file name: %s, suffix: %s, upusetime: %dms -> pathtime: %dms -> iocopytime: %dms \n",
		filename, filehz, uptime, pathtime, (time.Now().UnixNano()-start)/1000000)

	response.Write([]byte("{\"code\":0, \"fileid\":\"" + id + "\"}"))
}

func UpTest(rw http.ResponseWriter, r *http.Request) {

	io.WriteString(rw, "<html><head><title>上传测试</title></head><body><form action='../upfile' method=\"post\" enctype=\"multipart/form-data\"><label>上传图片</label><input type=\"file\" name='file'  /><br/><label><input type=\"submit\" value=\"上传图片\"/></label></form></body></html>")

}
