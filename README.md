# file-service
使用golang编写的文件服务

###使用第三方包<br/>
`github.com/satori/go.uuid`

###配置文件
根目录下 config.json
```
{
	"ListenAddr": ":8080",
	"SaveRootPath": "/home/xxx/fileservice/"
}
```
ListenAddr：服务开启端口 <br/>
SaveRootPath：文件存放根目录
