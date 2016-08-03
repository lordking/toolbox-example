# person_mongo样例

person_mongo样例，一个最基础的mongodb + http rest api的样例

## 1. 安装说明

运行之前，先安装GO语言环境

## 2 配置说明

所有配置文件均放置在config目录下，内容以JSON格式存放。

```
+- config
    |
    |---- http.json
    +---- db.json
```

### 2.1 HTTP配置文件: http.json

参数       | 说明
-------- | ------------------
http     | HTTP端口。
https    | HTTPS端口，不能与HTTP相同。
ssl_cert | HTTPS需要的证书文件的相对路径。
ssl_key  | HTTPS需要的公钥文件的相对路径。

ssl_cert和ssl_key的生成方式是：

```
  $ go run $GOROOT/src/crypto/tls/generate_cert.go --host="localhost"
```

#### 2.2 DB配置文件: mongo.json

```json
{
	"url": "127.0.0.1",
	"database": "sample"
}
```

连接字符串的格式是

`[mongodb://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]`

## 3 编译运行

```
  $ godep go build
  $ ./person_mongo
```

或者

```
  $ godep go run main.go
```

## 4 单元测试

### 4.1 单元测试

	$ cd test
	$ go test -v --test.run TestCreate
	$ go test -v --test.run TestFind
	$ go test -v --test.run TestUpdate
	$ go test -v --test.run TestDelete

### 4.2 性能测试

mongodb 缺少数据库连接池，性能测试不通过

	$ cd benchmark
	$ go run benchmark.go -m RequestCreate
	$ go run benchmark.go -m RequestFind
	$ go run benchmark.go -m RequestUpdate
	$ go run benchmark.go -m RequestDelete
