package task

var registerTasks = map[string]interface{}{
	"release_attachment":              releaseAttachment,
	"clean_deleted_axure_attachments": cleanDeletedAxureAttachments,
}
