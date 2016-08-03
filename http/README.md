# webseed

以最佳实践、极简方式提供HTTP服务的范例。

## 1. 样例说明

- welcome: 最基本的rest api样例
- person_mongo: 使用mongodb作为数据库的rest api样例
- person_mysql: 使用mysql作为数据库的restful样例
- blog: 使用mongodb作为数据库、用rest api作为服务模型的WebApp样例

## 2. 安装及配置说明

### 2.1 安装go语言环境

从 https://golang.org 下载安装包，并安装。


### 2.2 配置Go语言环境

设置GOPATH

GOPATH路径根据各人情况设置。可以统一设置在同一个目录下，也可以为不同的项目设置不同的GOPATH

	$ mkdir ~/Documents/workspace/go-project
	$ cd ~/Documents/workspace/go-project
	$ export GOPATH=$GOPATH:~/Documents/workspace/go-project

### 2.3 安装godep

	$ go get github.com/tools/godep

### 2.4 下载代码

	$ git clone https://git.oschina.net/gfound/goutils.git
	$ git clone https://git.oschina.net/gfound/goutils-example.git

### 2.5 运行样例

#### 2.5.1 普通样例

	$ cd [样例名称]
	$ gpdep go run [样例名称].go

#### 2.5.2 blog样例

blog样例，是restAPI + webapp的样例。所以运行之前，先要安装js库。

安装js库的步骤如下：

1.如果没有nodejs，到如下地址下载安装。

	https://nodejs.org

2.如果没有bower, 安装。

	$ sudo npm install -g bower

3.安装js库

	$ cd blog
	$ bower install

4.安装成功后，登录访问如下地址

	http://localhost:8000/login.html

用户名/密码: admin/admin

### 2.6 编译与运行

	$ godep go build
	$ ./[样例名称]

### 2.7 配置说明

所有配置文件均放置在config目录下，内容以JSON格式存放。

```
+- config
	|
	|---- http.json
	+---- db.json
```

#### 2.7.1 HTTP配置文件: http.json

|参数       |说明                   |
|----------|------------------------|
|http      | HTTP端口。 |
|https     | HTTPS端口，不能与HTTP相同。 |
|ssl_cert  | HTTPS需要的证书文件的相对路径。|
|ssl_key   | HTTPS需要的公钥文件的相对路径。|

ssl_cert和ssl_key的生成方式是：

	$ go run $GOROOT/src/crypto/tls/generate_cert.go --host="localhost"

#### 2.7.2 DB配置文件: db.json

|参数       |说明                     |
|----------|-------------------------|
|adapter   | 数据库类型，mongo, mysql  |

**mongodb**

```json
{
	"adapter": "mongo",
	"url": "127.0.0.1",
	"database": "sample"
}
```

连接字符串的格式是

`[mongodb://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]`

**mysql**

```json
{
	"adapter": "mysql",
	"host": "127.0.0.1",
	"port":"3306",
	"username":"root",
	"password":"",
    "MaxOpenConns":200,
    "MaxIdleConns":100,
	"database": "webseed"
}
```

## 3 单元测试

所有测试之前必须先启动服务，然后再进行测试

### 3.1 welcome

#### 3.1.1 单元测试

	$ cd test
	$ go test -v -test.run TestHello

#### 3.1.2 性能测试

	$ cd benchmark
	$ go run benchmark.go -m RequestHello

### 3.2 person_mongo

#### 3.2.1 单元测试

	$ cd test
	$ go test -v --test.run TestCreate
	$ go test -v --test.run TestFind
	$ go test -v --test.run TestUpdate
	$ go test -v --test.run TestDelete

#### 3.2.2 性能测试

	$ cd benchmark
	$ go run benchmark.go -m RequestCreate
	$ go run benchmark.go -m RequestFind
	$ go run benchmark.go -m RequestUpdate
	$ go run benchmark.go -m RequestDelete

### 3.3 person_mysql

#### 3.2.1 创建数据库和表

```sql
CREATE DATABASE `sample`

CREATE TABLE `person` (
  `name` varchar(255) NOT NULL,
  `phone` varchar(45) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
```

#### 3.2.2 单元测试

	$ cd test
	$ go test -v --test.run TestCreate
	$ go test -v --test.run TestFind
	$ go test -v --test.run TestUpdate
	$ go test -v --test.run TestDelete

#### 3.2.3 性能测试

	$ cd benchmark
	$ go run benchmark.go -m RequestCreate
	$ go run benchmark.go -m RequestFind
	$ go run benchmark.go -m RequestUpdate
	$ go run benchmark.go -m RequestDelete

### 3.4 blog

#### 3.4.1 用户登录接口测试

	$ cd test
	$ go test -v -test.run TestLogin

如果单元测试运行成功，将会在终端上的打印输出中获得token。这将用于后面的单元测试。
打开test.go文件，修改token的值。如：

```
token = "57884dba17a06faba180e46a"
```

#### 3.4.2 创建日志

	$ go test -v -test.run TestCreate

#### 3.4.3 查询日志

	$ go test -v -test.run TestFind

#### 3.4.4 修改日志

在运行测试之前，先通过之前创建或者查询的测试用例获取一个id。然后打开test.go文件，修改update_id的值。如：

```
update_id := "57884d1a17a06faba180e468"
```

再运行一下测试

	$ go test -v -test.run TestUpdate

#### 3.4.5 删除日志

在运行测试之前，先通过之前创建或者查询的测试用例获取一个id。然后打开test.go文件，找到delete_id的值。如：

```
delete_id = "57884d1a17a06faba180e468"
```

再运行一下测试

	$ go test -v -test.run TestDelete
