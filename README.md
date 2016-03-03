# file-service
使用golang编写的文件服务，

使用第三方包
`github.com/satori/go.uuid`

配置文件：config.json
```
{
	"ListenAddr": ":8080",
	"SaveRootPath": "/home/xxx/fileservice/"
}
```
ListenAddr：服务开启端口
SaveRootPath：文件存放根目录