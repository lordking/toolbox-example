# person_mongo样例

person_mongo样例，一个最基础的mongodb + http rest api的样例

## 1 编译运行前准备

## 1.1 配置说明

所有配置文件均放置在config目录下，内容以YAML格式存放。

```
  +- config
      |
      +---- config.yaml
```

###  `http`配置说明

参数     | 说明
------- | ------------------
http    | HTTP端口。
sslport | HTTPS端口，不能与HTTP相同。
sslcert | HTTPS需要的证书文件的相对路径。
sslkey  | HTTPS需要的公钥文件的相对路径。

ssl_cert和ssl_key的生成方式是：

```
  $ go run $GOROOT/src/crypto/tls/generate_cert.go --host="localhost"
```

#### `database`配置说明

参数      | 说明
-------- | ------------------
url      | 数据库URL连接字符串
database | 数据库名

连接字符串的格式是

`[mongodb://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]`

## 1.2 恢复依赖库

```
godep restore
```

## 2 编译运行

```
  $ godep go build
  $ ./person_mongo serve
```

或者

```
  $ godep go run main.go serve
```

## 3 测试

### 3.1 单元测试

	$ cd test
	$ go test -v --test.run TestCreate
	$ go test -v --test.run TestFind
	$ go test -v --test.run TestUpdate
	$ go test -v --test.run TestDelete

### 3.2 性能测试

	$ cd benchmark
	$ go run benchmark.go -m RequestCreate
	$ go run benchmark.go -m RequestFind
	$ go run benchmark.go -m RequestUpdate
	$ go run benchmark.go -m RequestDelete
