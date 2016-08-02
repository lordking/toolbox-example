mysql
=======

## 安装依赖

```
godep restore
```

## 配置数据库

```sql
CREATE DATABASE `sample`

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
