package api

import (
	"axshare_go/api/v1"
	"github.com/gin-gonic/gin"
	"github.com/xandercheung/acct"
)

func SetV1Router(router *gin.Engine) {
	r := router.Group("/api/v1")
	{
		r.POST("/sign_in", acct.Handler.SignIn)

		r.GET("/axures/:id", v1.GetAxureWebInfo)
		r.GET("/config/public", v1.GetPublicConfig)

		r.Use(acct.TokenAuthMiddleware())

		accounts := r.Group("/accounts")
		{
			accounts.GET("", acct.Handler.FetchAccounts)
			accounts.POST("", acct.Handler.CreateAccount)
			accounts.GET("/:id", acct.Handler.FetchAccount)
			accounts.POST("/:id", acct.Handler.UpdateAccount)
			accounts.DELETE("/:id", acct.Handler.DestroyAccount)
			accounts.POST("/:id/password", acct.Handler.UpdateAccountPassword)
			accounts.POST("/:id/reset_password", v1.ResetAccountPassword)
		}

		r.GET("/account/info", acct.Handler.FetchCurrentAccountInfo)

		r.GET("/axure_groups", v1.GetAxureGroups)
		r.POST("/axure_groups", v1.CreateAxureGroup)
		r.GET("/axure_group/:id", v1.GetAxureGroup)
		r.POST("/axure_group/:id", v1.UpdateAxureGroup)

		r.GET("/axure_groups/:axure_group_id/axures", v1.GetAxures)
		r.GET("/axure_groups/:axure_group_id/axure/:id", v1.GetAxure)
		r.POST("/axure_groups/:axure_group_id/axures", v1.CreateAxure)
		r.PUT("/axure_groups/:axure_group_id/axure/:id", v1.UpdateAxure)
		r.DELETE("/axure_groups/:axure_group_id/axure/:id", v1.DestroyAxure)

		r.GET("/axure_groups/:axure_group_id/axures/:axure_id/attachments", v1.GetAttachments)

		r.GET("/attachments", v1.GetAllAttachments)

		r.GET("/upload/token", v1.CreateUploadToken)
		r.GET("/config", v1.GetConfig)
		r.POST("/config", v1.UpdateConfig)
	}
}
