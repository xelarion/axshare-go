package v1

import (
	"github.com/gin-gonic/gin"
)

func RouterV1(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.POST("/user/login", Login)
		v1.POST("/user/logout", Logout)
		v1.GET("/user/info", GetUserInfo)

		v1.GET("/axure_groups", GetAxureGroups)

		v1.GET("/axure_groups/:axure_group_id/axures", GetAxures)
		v1.GET("/axure_groups/:axure_group_id/axure/:id", GetAxure)
		v1.POST("/axure_groups/:axure_group_id/axures", CreateAxure)
		v1.PUT("/axure_groups/:axure_group_id/axure/:id", UpdateAxure)

		v1.GET("/axure_groups/:axure_group_id/axures/:axure_id/attachments", GetAttachments)

		v1.GET("/upload/token", CreateUploadToken)
	}
}
