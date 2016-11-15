redis
=======

实现了Redis的set、get、del、publish/subscrbe的简单例子

## 运行前准备

* 需安装redis。
* `config.yaml`，配置文件，配置数据库、日志。

## 恢复依赖库

```
godep restore
```

## 运行样例

```
godep go run main.go person.go
```
