package axure_group

import (
	"axshare_go/internal/db"
	"axshare_go/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/ogsapi/ogs-go"
	"net/http"
)

func FetchList(c *gin.Context) {
	var axureGroups []models.AxureGroup
	db.AxshareDb.Model(&models.AxureGroup{}).Find(&axureGroups)
	c.JSON(http.StatusOK, ogs.RspOKWithPaginate(
		ogs.BlankMessage(),
		axureGroups,
		ogs.NewPaginate(1, 101, 10)))
}
