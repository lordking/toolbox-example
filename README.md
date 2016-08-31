toolbox example
===================

使用toolbox的样例代码

## 1 开发环境安装与配置

### 1.1 安装go语言环境

#### 下载golang

从 https://golang.org 下载安装包。

```
tar xzvf go1.6.1.linux-amd64.tar.gz $HOME/go
```

#### 设置环境变量

export GOROOT=$HOME/go
export GOPATH=$HOME/go-project
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

### 1.2 安装与使用Godep

建议采用godep做包管理。所有样例均使用godep。

#### 下载安装

```
go get github.com/tools/godep
```

#### 项目工程要求

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

#### 保存依赖库配置

```
godep save
```

#### 恢复依赖库

```
godep restore
```

#### 编译运行

```
godep go run main.go
godep go build
godep go install
godep go test
```

## 2. 样例说明

请打开每个样例的文件夹，都有说明文档

* log 日志处理的样例
* database\mongo mongodb的样例
* database\mysql mysql的样例
* database\redis redis的样例
* http\welcome 一个最小的http rest api服务的样例
* http\person_mongo 一个最基础的mongodb + http rest api的样例
* http\person_mysql 一个最基础的mysql + http rest api的样例
* http\blog 一个最基础的mongodb + http rest api + webapp的样例
