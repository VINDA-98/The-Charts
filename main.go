package main

import (
	"github.com/VINDA-98/The-Charts/api"
	"log"
)

// @Title  The_Charts
// @Description  MyGO
// @Author  WeiDa  2023/6/25 17:59
// @Update  WeiDa  2023/6/25 17:59

func main() {
	// 初始化路由
	router := api.InitRouter()
	// 启动 API 服务
	// 新增：curl -H "Content-type: application/json" -d '{"username": "张三", "points": 25}' localhost:8080/rank
	// 获取：curl -H "Content-type: application/json" localhost:8080/rank/张三
	// 排行：curl -H "Content-type: application/json" localhost:8080/rank
	log.Fatal(router.Run("localhost:8080"))
}
