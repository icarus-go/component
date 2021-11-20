package constant

type ContentType string

const (
	JSON       ContentType = "application/json;charset=utf-8"
	XML        ContentType = "application/xml;charset=utf-8"
	FormData   ContentType = "multipart/form-data;charset=utf-8"
	URLEncode  ContentType = "application/x-www-form-urlencoded;charset=utf-8"
	FileStream ContentType = "application/octet-stream"
)

func (c ContentType) Value() string {
	return string(c)
}
