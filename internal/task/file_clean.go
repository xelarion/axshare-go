package task

import (
	"axshare_go/internal/db"
	"axshare_go/internal/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

// 清理原型组织陈旧的解压文件(html)
func cleanAxureOldAttachments() {
	// 一周前（只清理一周内没有更新过的附件）
	aWeekAgo := time.Now().AddDate(0, 0, -7).UTC()

	var axures []models.Axure
	db.AxshareDb.FindInBatches(&axures, 50, func(tx *gorm.DB, batch int) error {
		for _, axure := range axures {
			latestAttachment := axure.LatestAttachment()
			var oldAttachments []models.Attachment
			// 该原型 非最新的、成功解压的 历史附件
			db.AxshareDb.Where("axure_id = ?", axure.ID).
				Where("id < ?", latestAttachment.ID).
				Where("release_status = ?", models.AttachmentReleaseStatusSuccessful).
				Where("updated_at < ?", aWeekAgo).
				Find(&oldAttachments)

			for _, attachment := range oldAttachments {
				if err := attachment.CleanAxureFileDir(); err != nil {
					logrus.Error(err)
				}
			}
		}

		return nil
	})
}
