package axure_group

import (
	"axshare_go/internal/db"
	"axshare_go/internal/models"
	"github.com/gin-gonic/gin"
	go_rsp "github.com/standard-rsp/go-rsp"
	"net/http"
)

func FetchList(c *gin.Context) {
	var axureGroups []models.AxureGroup
	db.AxshareDb.Model(&models.AxureGroup{}).Find(&axureGroups)
	c.JSON(http.StatusOK, go_rsp.RspPagData(
		FormatList(axureGroups), go_rsp.OK,
		go_rsp.NewBaseMessage("", ""),
		go_rsp.BasePaginate{CurrentPage: 1, TotalPages: 20, TotalCount: 201, PerPage: 10}))
}
