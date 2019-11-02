package pg

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/ogsapi/ogs-go"
	"strconv"
)

var defaultPerPage = 20

func Paginate(gormDB *gorm.DB, page, perPage int) (*gorm.DB, ogs.BasePaginate) {
	if page <= 0 {
		page = 1
	}
	if perPage == 0 {
		perPage = defaultPerPage
	}

	totalCount := 0
	gormDB.Count(&totalCount)

	offset := perPage * (page - 1)
	gormDB = gormDB.Limit(perPage).Offset(offset)

	return gormDB, ogs.NewPaginate(page, totalCount, perPage)
}

func PaginateGin(gormDB *gorm.DB, c *gin.Context) (*gorm.DB, ogs.BasePaginate) {
	page, _ := strconv.ParseInt(c.Query("page"), 10, 64)
	perPage, _ := strconv.ParseInt(c.Query("per_page"), 10, 64)
	return Paginate(gormDB, int(page), int(perPage))
}
