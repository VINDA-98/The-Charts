package api

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	// 新增用户排行信息
	r.POST("/rank")
	// 根据用户名查看用户排行信息
	r.GET("/rank/:username")
	// 查看排行榜
	r.GET("/rank")

	return r
}

//
//// 添加排行信息
//func handleAddRank(c *api.Context) {
//	// 提取客户端用户对象
//	var userJson redis.User
//	if err = c.ShouldBindJSON(&userJson); err != nil {
//		c.ResponseFailureArgs(err.Error())
//		return
//	}
//
//	// 添加用户
//	err = database.SaveUser(&userJson)
//	if err != nil {
//		c.ResponseFailureArgs(err.Error())
//		return
//	}
//
//	// 返回
//	c.ResponseSuccess(api.Json{"user": userJson})
//}
//
//// 获取用户排名信息
//func handleGetUserRank(c *api.Context) {
//	// 获取客户端传递过来的用户名
//	username := c.Param("username")
//
//	// 从redis根据用户名查询用户排行
//	user, err := database.GetUser(username)
//	if err != nil {
//		if err == redis.ErrNil {
//			c.ResponseFailureNotFound(fmt.Sprintf("%s 不存在", username))
//			return
//		}
//		c.ResponseFailure(err.Error())
//		return
//	}
//	c.ResponseSuccess(user)
//}
//
//// 获取排行榜信息
//func handleGetRank(c *api.Context) {
//	// 从redis获取排行榜
//	users, err := database.GetRank()
//	if err != nil {
//		if err == redis.ErrNil {
//			c.ResponseFailureNotFound("排行榜不存在")
//			return
//		}
//		c.ResponseFailure(err.Error())
//		return
//	}
//
//	// 返回排行榜信息
//	c.ResponseSuccess(users)
//}
