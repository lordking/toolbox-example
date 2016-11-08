SQLite
=======

使用SQLite数据库实现的简单的增、删除、修改、查询例子。

## 运行前的准备

## 恢复依赖库

```
godep restore
```

### 编译前安装数据库驱动

一般情况下无需安装，如果没有，需提前安装好。

```
# Mac
brew install sqlite3
```

### 数据库文件

本例person.db已经创建完毕，可以直接使用。如果需要创建，按此下SQL脚本。

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
# Mac下编译
godep go build --tags "libsqlite3 darwin"
```
