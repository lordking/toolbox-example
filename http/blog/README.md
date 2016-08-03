# blog样例

blog样例，是mongodb + restapi + webapp的样例。

## 1\. 安装说明

运行之前，除go语言环境外先要安装js库。步骤如下：

- 如果没有nodejs，到如下地址下载安装。

  <https://nodejs.org>

- 如果没有bower, 安装。

  $ sudo npm install -g bower

- 安装js库

  $ cd blog $ bower install

## 2 配置说明

所有配置文件均放置在config目录下，内容以JSON格式存放。

```
  +- config
      |
      |---- http.json
      +---- db.json
```

### 2.1 HTTP配置文件: http.json

参数      | 说明
-------- | ------------------
http     | HTTP端口。
https    | HTTPS端口，不能与HTTP相同。
ssl_cert | HTTPS需要的证书文件的相对路径。
ssl_key  | HTTPS需要的公钥文件的相对路径。

ssl_cert和ssl_key的生成方式是：

```
  $ go run $GOROOT/src/crypto/tls/generate_cert.go --host="localhost"
```

### 2.2 DB配置文件: mongo.json

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
  $ ./blog
```

成功后，可访问如下地址:

<http://localhost:8000/login.html>

用户名/密码: admin/admin

## 4 单元测试

### 4.1 用户登录接口测试

```
  $ cd test
  $ go test -v -test.run TestLogin
```

如果单元测试运行成功，将会在终端上的打印输出中获得token。这将用于后面的单元测试。 打开test.go文件，修改token的值。如：

```
  token = "57884dba17a06faba180e46a"
```

#### 4.2 创建日志

```
  $ go test -v -test.run TestCreate
```

#### 4.3 查询日志

```
  $ go test -v -test.run TestFind
```

#### 4.4 修改日志

在运行测试之前，先通过之前创建或者查询的测试用例获取一个id。然后打开test.go文件，修改update_id的值。如：

```
  update_id := "57884d1a17a06faba180e468"
```

再运行一下测试

```
  $ go test -v -test.run TestUpdate
```

#### 4.5 删除日志

在运行测试之前，先通过之前创建或者查询的测试用例获取一个id。然后打开test.go文件，找到delete_id的值。如：

```
  delete_id = "57884d1a17a06faba180e468"
```

再运行一下测试

```
  $ go test -v -test.run TestDelete
```
