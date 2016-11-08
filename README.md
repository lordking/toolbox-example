toolbox example
===================

使用toolbox的样例代码

## 1 编译环境准备

### 1.1 Go语言环境

```bash
# 以root用户身份登录
# 下载
$ wget https://storage.googleapis.com/golang/go1.7.1.linux-amd64.tar.gz
$ tar xzvf go1.7.1.linux-amd64.tar.gz
$ mv go /usr/local

# 设置GO环境
$ vi /etc/profie.d/go_env.sh

# 添加如下代码
export GOROOT=/usr/local/go

PATH=$PATH:$GOROOT/bin
export PATH

# root用户身份操作，或切换成普通用户
$ su digger

# 设置GO PATH
$ vi ~/.bash_profile

# 添加如下代码
export GOPATH=$HOME/go-project

PATH=$PATH:$HOME/bin:$GOPATH/bin
export PATH

# 生效GO PATH配置
$ source ~/.bash_profile
```

### 1.2 Godep安装

Godep是一个go语言库包管理工具。

```bash
# 安装命令
$ go get github.com/tools/godep
```

### 1.3 项目工程安装方式

* 把项目路径加入到$GOPATH/src
* 依赖的项目和项目本身都应该是个git仓库
* 目录结构例如

```
$GOPATH
 |-src
 |  |-github.com
 |     |-lordking
 |        |-toolbox-example
 |-pkg
 |-bin

```

## 2 Godep的基本用法介绍


保存依赖库
```
godep save
```

恢复依赖库
```
godep restore
```

编译运行
```
godep go run main.go
godep go build
godep go install
godep go test
```

## 3. 样例说明

请打开每个样例的文件夹，都有说明文档

* log 日志处理的样例

* database\mongo mongodb的样例
* database\mysql mysql的样例
* database\redis redis的样例
* database\sqlite3 sqlite的样例

* http\welcome 一个最小的http rest api服务的样例
* http\person_mongo 一个最基础的mongodb + http rest api的样例
* http\person_mysql 一个最基础的mysql + http rest api的样例
* http\blog 一个最基础的mongodb + http rest api + webapp的样例
