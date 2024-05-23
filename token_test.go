package ebct

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

var client *Client

func init() {

	client = NewClient(
		"eyJhbGciOiJSUzUxMiJ9.eyJhbWJpZW50ZSI6IkhPTU9MT0dBQ0FPIiwiaWQiOiIwNTMxNzcwODAwMDE5NCIsInBmbCI6IlBKIiwiY25waiI6IjA1MzE3NzA4MDAwMTk0IiwiaXAiOiIxNzkuMTAxLjIwOC4yNTQsIDE5Mi4xNjguMS44IiwiaWF0IjoxNzE2NDgwOTg5LCJpc3MiOiJ0b2tlbi1zZXJ2aWNlIiwiZXhwIjoxNzE2NTY3Mzg5LCJqdGkiOiJjODIxYWQ0MS04MWNiLTRhY2ItYjhmYi1kYzdmODYwM2Q3OTYifQ.yx_DBU6Bq0uFUrZsZT3SdUMYSCgJ-pQY-WcT2d0bawlC8wY7BuQCN0FwhRly_ReC34Xyz3Mx3GZXDq0xdw6QjZfKS18tQi9yzJ8rQ2LaCnjerZG9b49vxlEEbB3GzSFs7d8NamiXC1C-e2qNfVeyIqDwhOqsjYEwBjgulE1AwNKTa_LG--6lXejv2ZwABeyUpVXrQp7i5mBGsI1XzlFP9omDdkQI5biryB0wDDOZBz3USzXCoH9FS2jIEv2SvFfza30LR3-gWrifsdECbtNL-wtaMlAZqDyhy5iaxEFKwkebXgrxbQ658bmOrQaRle5LZMXF6fygbSv__4lXKBvp9A",
		true,
	)
}

func TestLogin(t *testing.T) {
	var err error

	//client, err = NewClientToken("epsadmin4675", "correiofox123", "9912435103", true)
	client, err = NewClientToken("05317708000194", "0ywD1PYCdkzojfPUGAfifFyTIHFOhTZUoGZgOVpP", "9912435103", "AUTH", true)
	spew.Dump(client)
	assert.NoError(t, err)
}
