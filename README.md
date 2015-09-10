# 开发环境

## 配置 golang 环境变量

将最新的golang下载解压到 $HOME/go 目录下

配置一下环境变量到 `.bashrc` 或 `.zshrc`, GOPATH 是你的工作目录

```bash
export GOROOT=$HOME/go
export GOPATH=$HOME/workspace/Go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```

## 获取项目代码到本地, 并安装依赖

```bash
$ go get github.com/shelmesky/rtmonitor_dashboard
$ go get github.com/astaxie/beego
$ go get github.com/beego/bee
```
