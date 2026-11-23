# The Charts · Go 排行榜 API 草稿

[English](README.md)

这是一个 Go 练习项目，设计了提交分数和查询排名的 HTTP API。原方案计划用 Redis 保存和查询排行数据。

| 预设接口 | 用途 |
| --- | --- |
| `POST /rank` | 提交用户分数。 |
| `GET /rank/:username` | 查询指定用户的排名。 |
| `GET /rank` | 查看排行榜。 |

[`main.go`](main.go) 是服务入口草稿。[`api/route.go`](api/route.go) 声明了 Gin 路由，并保留了注释中的处理逻辑草稿。当前代码尚未接入处理函数或 Redis，[`go.mod`](go.mod) 的模块路径与 `main.go` 中的导入路径也不一致。
