# file-service
使用golang编写的文件服务

###使用第三方包<br/>
`github.com/satori/go.uuid`

###配置文件
- config.json 
主要配置系统的启动运行环境
```
{
	"ListenAddr": ":8080",
	"SaveRootPath": "/home/xxx/fileservice/"
}
```
ListenAddr：服务开启端口 <br/>
SaveRootPath：文件存放根目录 <br/>

- suffix.properties
主要配置服务器接收的文件类型
```
.jpg=100
```
.jpg：文件后缀 <br/>
100：后缀编号，编号采用3个字符表示，编号保证全局唯一性。
