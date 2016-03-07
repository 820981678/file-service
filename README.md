# file-service
使用golang编写的文件服务

###使用第三方包<br/>
`github.com/satori/go.uuid`

###配置文件
- config.json <br/>
主要配置系统的启动运行环境
```
{
	"ListenAddr": ":8080",
	"SaveRootPath": "/home/xxx/fileservice/"
}
```
ListenAddr：服务开启端口 <br/>
SaveRootPath：文件存放根目录 <br/>

- suffix.properties <br/>
主要配置服务器接收的文件类型
```
.jpg=100
```
.jpg：文件后缀 <br/>
100：后缀编号，编号采用3个字符表示，编号保证全局唯一性。

###使用方式
本服务采用http方式进行文件的上传及下载

- /upfile 上传文件 <br/>
参数名称：file <br/>
参数值：需要上传的文件 <br/>
上传成功后系统将返回json，json字符串格式如下：
```
{
	"code": 0,
	"fileid": "XXXXXXXX"
}
```

- /downfile 下载文件 <br/>
参数名称：fileid <br/>
参数值：上传文件时系统返回的json中的fileid字值
