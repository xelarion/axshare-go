package v1

import (
	"axshare_go/api/v1/attachment"
	"axshare_go/api/v1/axure"
	"axshare_go/api/v1/axure_group"
	"axshare_go/api/v1/qiniu"
	"axshare_go/api/v1/user"
	"github.com/gin-gonic/gin"
)

func RouterV1(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.POST("/user/login", user.Authenticate)
		v1.POST("/user/logout", user.DestroyAuthorization)
		v1.GET("/user/info", user.GetInfo)

		v1.GET("/axure_groups", axure_group.FetchList)

		v1.GET("/axure_groups/:axure_group_id/axures", axure.GetAxures)
		v1.GET("/axure_groups/:axure_group_id/axure/:id", axure.GetAxure)
		v1.POST("/axure_groups/:axure_group_id/axures", axure.CreateAxure)
		v1.PUT("/axure_groups/:axure_group_id/axure/:id", axure.UpdateAxure)

		v1.GET("/axure_groups/:axure_group_id/axures/:axure_id/attachments", attachment.GetAttachments)

		v1.GET("/upload/token", qiniu.CreateUploadToken)
	}
}
