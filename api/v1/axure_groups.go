package v1

import (
	"axshare_go/internal/db"
	"axshare_go/internal/models"
	"axshare_go/internal/task"
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/gin-gonic/gin"
	"github.com/ogsapi/ogs-go"
	"net/http"
)

func GetAxureGroups(c *gin.Context) {
	var axureGroups []models.AxureGroup
	db.AxshareDb.Model(&models.AxureGroup{}).Find(&axureGroups)
	c.JSON(http.StatusOK, ogs.RspOKWithPaginate(
		ogs.BlankMessage(),
		axureGroups,
		ogs.NewPaginate(1, 101, 10)))
}

func testMechineryTask() {
	var addTask0 tasks.Signature
	addTask0 = tasks.Signature{
		Name: "add",
		Args: []tasks.Arg{
			{
				Type:  "int64",
				Value: 1,
			},
			{
				Type:  "int64",
				Value: 1,
			},
		},
	}
	_ = task.Send(&addTask0)
}
