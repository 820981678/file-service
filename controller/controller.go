package controller

import(
	"fmt"
	"net/http"
	"os"
	"io"
	"path"

	_ "file-service/config"
	"file-service/utils"
)

func DownFile(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "D://liuyan.jpg")
}

func UpFile(response http.ResponseWriter, request *http.Request) {
	file, fileh, error := request.FormFile("file")
	if error != nil {
		fmt.Println("jie xie file error")
		response.Write([]byte(`{"code":-1}`))
		return
	}
	defer file.Close()

	filename := fileh.Filename;
	filehz := path.Ext(filename)//filename[strings.LastIndex(filename, ".") : ]
	fmt.Printf("upfile name: %s, hz: %s \n", filename, filehz)

	id, savefilename := utils.GetSavePath(filehz)
	uf, err := os.Create(savefilename)
	//uf, err := os.OpenFile("D://111/111.jpg", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("save file error")
		response.Write([]byte(`{"code":-1}`))
		return
	}
	defer uf.Close()

	io.Copy(uf, file);
	response.Write([]byte("{\"code\":0, \"fileid\":\"" + id + "\"}"))
}

func UpTest(rw http.ResponseWriter, r *http.Request) {

	io.WriteString(rw, "<html><head><title>上传测试</title></head><body><form action='../upfile' method=\"post\" enctype=\"multipart/form-data\"><label>上传图片</label><input type=\"file\" name='file'  /><br/><label><input type=\"submit\" value=\"上传图片\"/></label></form></body></html>")
}