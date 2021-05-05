package task

import (
	"axshare_go/internal/db"
	"axshare_go/internal/models"
	"github.com/sirupsen/logrus"
)

func releaseAttachment(attachmentId uint) error {
	attachment := models.FindAttachment(attachmentId)
	if attachment.ID == 0 {
		return nil
	}

	webLink, err := deployAxure(attachment.DownloadUrl(), attachment.GenFileName())
	if err != nil {
		logrus.Error("releaseAttachment error, id: ", attachment.ID, ", error: ", err.Error())
		db.AxshareDb.Model(&attachment).Updates(models.Attachment{
			ReleaseStatus: models.AttachmentReleaseStatusFailed,
			ReleaseError:  err.Error(),
		})
	} else {
		db.AxshareDb.Model(&attachment).Updates(models.Attachment{
			ReleaseStatus: models.AttachmentReleaseStatusSuccessful,
			Link:          webLink,
		})
	}

	return nil
}
