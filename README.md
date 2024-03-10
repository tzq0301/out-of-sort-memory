# out of sort memory

## 触发场景

1. `SELECT` 列中包含 JSON 字段
2. JSON 值大于 sort buffer size
3. `ORDER BY` 的列不是索引

## 解决方法

为 `ORDER BY` 的列加上索引，这样就不需要使用 sort buffer 进行排序了

### 优化

使用“延迟查询”进行优化

由于回表操作先于 limit-offset 操作完成，因此以下第一条 SQL 会回表 1010 次，而第二条 SQL 只会回表 10 次

```mysql
SELECT * FROM test FORCE INDEX (create_time) ORDER BY create_time DESC LIMIT 10 OFFSET 1000;

SELECT * FROM test JOIN (SELECT id FROM test ORDER BY create_time DESC LIMIT 10 OFFSET 1000) AS t ON test.id = t.id;
```

## 根本原因

在 MySQL 8.0.20 之前，当一个排序操作中涉及比 `TINYBLOB` 或 `BLOB` 更大的类型时，只将行的 ID 放在 sort buffer 中，这导致了大量的回表； 而 MySQL 在 8.0.20 时做了一个优化，即对于上述列，直接将内容放到 sort buffer 中，可以极大地提高性能，但同时也可能会造成像这样的 out of sort memory 的 bug

## 参考资料

1. https://stackoverflow.com/questions/76557503/mysql-8-0-33-error-when-selecting-json-column-out-of-sort-memory-consider-inc
2. https://bugs.mysql.com/bug.php?id=103225