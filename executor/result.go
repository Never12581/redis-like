package executor

type Result interface {
	getResult() interface{}
	getError() error
	hasError() bool
	getAttachments() map[string]string
	getAttachment(key string) string
	getAttachmentOrDefaultValue(key, defaultValue string) string
}
