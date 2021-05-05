package models

import (
	"encoding/json"
	"errors"
)

type AttachmentReleaseStatus int

const (
	AttachmentReleaseStatusPending AttachmentReleaseStatus = iota
	AttachmentReleaseStatusSuccessful
	AttachmentReleaseStatusFailed
	AttachmentReleaseStatusCleaned
)

func (c AttachmentReleaseStatus) MarshalJSON() ([]byte, error) {
	var status string
	switch c {
	case AttachmentReleaseStatusPending:
		status = "pending"
	case AttachmentReleaseStatusSuccessful:
		status = "successful"
	case AttachmentReleaseStatusFailed:
		status = "failed"
	case AttachmentReleaseStatusCleaned:
		status = "cleaned"
	}

	return json.Marshal(status)
}

func (c *AttachmentReleaseStatus) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"pending"`:
		*c = AttachmentReleaseStatusPending
	case `"successful"`:
		*c = AttachmentReleaseStatusSuccessful
	case `"failed"`:
		*c = AttachmentReleaseStatusFailed
	case `"cleaned"`:
		*c = AttachmentReleaseStatusCleaned
	default:
		return errors.New("unknown account status")
	}
	return nil
}
