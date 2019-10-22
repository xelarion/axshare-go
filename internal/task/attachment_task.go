package task

import (
	"axshare_go/internal/db"
	"axshare_go/internal/models"
	"axshare_go/internal/utils"
)

func releaseAttachment(attachmentId uint) error {
	defer utils.RecoverProgram(
		"happened when releaseAttachment, attachmentId: ", attachmentId)

	tx := db.AxshareDb.Begin()
	attachment := models.Attachment{}
	db.AxshareDb.First(&attachment, attachmentId)
	webLink := deployAxure(attachment.DownloadUrl(), attachment.GenFileName())
	db.AxshareDb.Model(&attachment).Update("link", webLink)
	tx.Commit()
	return nil
}
