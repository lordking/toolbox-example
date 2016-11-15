mysql
=======

使用MySQL数据库实现的简单的增、删除、修改、查询例子。

## 运行前准备

* 需安装MySQL或者MariaDB数据库。
* `config.yaml`，配置文件，配置数据库、日志。

## 恢复依赖库

```
godep restore
```

## 创建样例数据库和表。

```sql
CREATE DATABASE `sample`;

USE sample

CREATE TABLE `person` (
  `name` varchar(255) NOT NULL,
  `phone` varchar(45) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
```

## 运行样例

```
godep go run main.go person.go
```
