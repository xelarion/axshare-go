package task

import (
	"axshare_go/internal/db"
	"axshare_go/internal/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

func ReleaseAllAxures() {
	db.AxshareDb.Model(&models.Attachment{}).
		Where("release_status = ?", models.AttachmentReleaseStatusSuccessful).
		UpdateColumns(map[string]interface{}{
			"release_status": models.AttachmentReleaseStatusCleaned,
			"updated_at":     time.Now().UTC(),
		})

	var axures []models.Axure
	db.AxshareDb.FindInBatches(&axures, 50, func(tx *gorm.DB, batch int) error {
		for _, axure := range axures {
			attachment := axure.LatestAttachment()
			if err := releaseAttachment(attachment.ID); err != nil {
				logrus.Error(err)
			}
		}

		return nil
	})

}
