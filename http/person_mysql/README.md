# welcome样例

welcome样例，是一个restapi的样例。

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

参数          | 说明
------------ | ------------------
host         | 数据库主机地址
port         | 数据库端口
username     | 用户名
password     | 密码
MaxOpenConns | 最多可开连接
MaxIdleConns | 最多空闲连接
database     | 数据库名


## 1.2 恢复依赖库

```
godep restore
```

### 1.3 数据库建表

```sql
CREATE DATABASE `sample`;

CREATE TABLE `person` (
  `name` varchar(255) NOT NULL,
  `phone` varchar(45) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
```


## 2 编译运行

```
  $ godep go build
  $ ./person_mysql serve
```

或者

```
  $ godep go run main.go serve
```

## 4 测试

### 4.1 单元测试

	$ cd test
	$ go test -v --test.run TestCreate
	$ go test -v --test.run TestFind
	$ go test -v --test.run TestUpdate
	$ go test -v --test.run TestDelete

### 4.2 性能测试

	$ cd benchmark
	$ go run benchmark.go -m RequestCreate
	$ go run benchmark.go -m RequestFind
	$ go run benchmark.go -m RequestUpdate
	$ go run benchmark.go -m RequestDelete
