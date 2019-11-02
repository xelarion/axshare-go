package api

import (
	"axshare_go/api/v1"
	"github.com/gin-gonic/gin"
)

func SetV1Router(router *gin.Engine) {
	r := router.Group("/api/v1")
	{
		r.GET("/axures/:id", v1.GetAxureWebInfo)

		r.POST("/user/login", v1.Login)
		r.POST("/user/logout", v1.Logout)
		r.GET("/user/info", v1.GetUserInfo)

		r.GET("/axure_groups", v1.GetAxureGroups)

		r.GET("/axure_groups/:axure_group_id/axures", v1.GetAxures)
		r.GET("/axure_groups/:axure_group_id/axure/:id", v1.GetAxure)
		r.POST("/axure_groups/:axure_group_id/axures", v1.CreateAxure)
		r.PUT("/axure_groups/:axure_group_id/axure/:id", v1.UpdateAxure)

		r.GET("/axure_groups/:axure_group_id/axures/:axure_id/attachments", v1.GetAttachments)

		r.GET("/attachments", v1.GetAllAttachments)

		r.GET("/upload/token", v1.CreateUploadToken)
	}
}
