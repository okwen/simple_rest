package api

import (
	"simple_rest/api/controller/dbaccess"
	"simple_rest/api/controller/demo"
	"simple_rest/api/controller/fixme"
	
	"simple_rest/api/controller/user"

	"github.com/gin-gonic/gin"
)

// BindRouting : 綁定路由與操作方法
func BindRouting(router *gin.Engine) {

	// 在路由中區分版本號
	v1 := router.Group("/v1")
	{
		v1.GET("/get", demo.Getting)
		v1.POST("/post", demo.Postting)
		v1.GET("/deposit", fixme.GetDepostit)
		v1.GET("/user", dbaccess.GetUser)
		v1.GET("/user/example", dbaccess.GetUser)

		
		v1.POST("/user/create", user.CreatetUser)
		v1.POST("/user/delete", user.DeletetUser)
		v1.POST("/user/pwd/change", user.ChangeUser)
	}

}
