toolbox example
===================

使用toolbox的样例代码

## godep

本项目使用godep做包管理，所以需先安装godep

### 安装

```
go get github.com/golang/tools
```

### 开发工程配置要求

* 把项目路径加入到GOPATH
* 依赖的项目和项目本身都应该是个git仓库
* 目录结构是

```
toolbox-example
 |-src
 |  |-github.com
 |     |-lordking
 |        |-toolbox-example
 |-pkg
 |-bin

```

### 保存依赖库配置

```
godep save
```

### 恢复依赖库

```
godep restore
```

### 编译运行

```
godep go run main.go
godep go build
godep go install
godep go test
```

## 样例工程说明

### 日志

日志处理的样例。

### 数据库

数据库处理的样例。目前有mysql、mongodb

### HTTP服务

HTTP服务的样例。
