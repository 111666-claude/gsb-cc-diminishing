# cc-diminishing

控制递减：按类别分别计数、窗口内逐次减半、免疫窗口与查询代价，只用 Go 标准库。

```
go run ./cmd/cc-diminishing --sample=curve
go run ./cmd/cc-diminishing --sample=category
go run ./cmd/cc-diminishing --sample=window
go run ./cmd/cc-diminishing --sample=scan
go test ./...
go vet ./...
```

## 口径（README 为准）

- **递减表**：同一单位同一控制类别，在 `windowMs` 内第 1 次是满时长，
  第 2 次是一半，第 3 次是四分之一，第 4 次及以后是 0（免疫）。
  时长按整除取整。
- **类别独立**：`stun` 与 `root` 各自计数，互不影响。
- **窗口重置**：距上次施加超过 `windowMs` 时计数清零，下一次又是满时长。
- **免疫**：返回 0 的那一次把免疫结束时间设为 `now + immuneMs`；
  免疫期内再次施加一律返回 0，且不延长免疫结束时间。
- **不变量**：同一串施加序列重复执行结果相同；计数只增不减（窗口内）。
- **代价**：`Apply` 不许扫描全部单位，`Scanned` 不随单位数线性增长。
- **规模**：10 万单位、每秒 10 万次施加，单次摊还 O(1)。

## 输出契约（不改格式）

```
first=.. second=.. third=.. fourth=..
root=..
second=..
scanned=..
```

## 目录

```
control.go              递减、类别、窗口与免疫
control_test.go         单元测试
cmd/cc-diminishing/     命令行入口
```
