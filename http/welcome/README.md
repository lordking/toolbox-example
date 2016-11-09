# welcome样例

welcome样例，是一个restapi的样例。

## 1 编译运行前准备

## 1.1 配置说明

所有配置文件均放置在config目录下，内容以YAML格式存放。

```
  +- config
      |
      +---- http.yaml
```

###  HTTP配置文件说明: http.json

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

## 1.2 恢复依赖库

```
godep restore
```

## 2 运行样例

编译
```
  $ godep go build
  $ ./welcome
```

直接运行
```
  $ godep go run main.go
```

## 3 单元测试

#### 3.1 单元测试

	$ cd test
	$ godep go test -v -test.run TestHello

#### 3.2 性能测试

	$ cd benchmark
	$ go run benchmark.go -m RequestHello
