SQLite
=======

## 安装依赖

```
godep restore
```

## 配置数据库

### Mac

brew install sqlite3

### 建表

```sql

CREATE TABLE "person" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "name" TEXT,
  "phone" TEXT
)
```

## 运行样例

```
godep go run main.go person.go
```

## 编译

```
godep go build --tags "libsqlite3 darwin"
```
