# Go 排行榜实验

[English](README.md)

Go 排行榜接口原型，定义了提交分数和查询排名的入口，原设计计划用 Redis 保存排名数据。

| 预设接口 | 用途 |
| --- | --- |
| `POST /rank` | 提交用户分数。 |
| `GET /rank/:username` | 查询指定用户的排名。 |
| `GET /rank` | 查看排行榜。 |

[`main.go`](main.go) 是服务入口。[`api/route.go`](api/route.go) 声明了 Gin 路由，注释里保留处理函数草稿。目前只有路由定义，处理函数和 Redis 接入仍在计划中。
