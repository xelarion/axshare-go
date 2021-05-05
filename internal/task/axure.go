package task

import (
	"axshare_go/internal/db"
	"axshare_go/internal/models"
	"github.com/sirupsen/logrus"
)

func cleanDeletedAxureAttachments(axureId uint) error {
	axure := models.Axure{}
	db.AxshareDb.Unscoped().Find(&axure, axureId)
	if axure.ID == 0 {
		return nil
	}

	var attachments []models.Attachment
	db.AxshareDb.Where("axure_id = ?", axure.ID).Find(&attachments)
	for _, attachment := range attachments {
		if err := attachment.CleanAxureFileDir(); err != nil {
			logrus.Error("cleanDeletedAxureAttachments: ", err)
		}
	}

	return nil
}
