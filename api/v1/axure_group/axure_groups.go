package axure_group

import (
	"axshare_go/internal/db"
	"axshare_go/internal/models"
	"github.com/gin-gonic/gin"
	gorsp "github.com/standard-rsp/gorsp"
	"net/http"
)

func FetchList(c *gin.Context) {
	var axureGroups []models.AxureGroup
	db.AxshareDb.Model(&models.AxureGroup{}).Find(&axureGroups)
	c.JSON(http.StatusOK, gorsp.RspPagData(
		FormatList(axureGroups), gorsp.OK,
		gorsp.NewMessage("", ""),
		gorsp.NewPaginate(1, 102, 10)))
}
