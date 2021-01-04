## 什么是Go Modules
Go modules是Go语言的依赖解决方案，正式于Go1.14推荐生产上使用。
Go modules目前集成在Go的工具链中，只要安装了Go，就可以使用Go modules了。解决的问题：

1. Go语言长久依赖的依赖管理问题
2. 淘汰现有的GoPATH的使用模式
3. 统一社区中的其他依赖管理工具（提供迁移功能）
## Go Modules基本使用
生成go.mod文件
```shell
go mod init
```
下载go.mod文件中指明的所有依赖
```shell
go mod download
```
### 提供的环境变量
在Go modules中有如下常用环境变量，通过go env命令查看
```shell
go env
```
* GO111MODULE
  GO111MODULE环境变量是Go modules的开关。参数值有：
  auto : 只要项目包含了go mod 文件Go modules就启用
  on : 启用Go modules，推荐设置
  off : 禁用go modules

* GOPROXY
  这个环境变量主要是用于设置Go模块代理，作用是用于使Go在后续拉取模块版本时能够脱离传统的VCS方式，直接通过镜像站点来快速拉取。
  设置国内的Go模块代理：
```shell
go env -w GOPROXY=https://goproxy.cn,direct
```
GOPROXY的值是一个以英文逗号,分割的GO模块代理列表，允许设置多个模块代理。
* direct
  direct是一个特殊指示符，指示Go回源到模块版本的源地址去抓取
* GONOPROXY/GONOSUMDB/GOPRIVATE
  这三个环境变量都是用在当前项目依赖了私有模块，例如像公司私有git仓库，又或是github中的私有库，都是属于私有模块，都是要进行设置的，否则会拉取失败。

## 初始化项目


### 参考
https://blog.csdn.net/RA681t58CJxsgCkJ31/article/details/104568182/