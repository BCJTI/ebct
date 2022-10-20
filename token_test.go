package ebct

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

var client *Client

func init() {

	client = NewClient(
		"eyJhbGciOiJSUzUxMiJ9.eyJhbWJpZW50ZSI6IkhPTU9MT0dBQ0FPIiwiaWQiOiJ0ZXN0ZWludCIsInBmbCI6IlBKIiwicGpJbnRlcm5hY2lvbmFsIjo0MjM3NDg0MCwiYXBpIjpbODAsMzYsNTcsNTY0XSwiY2FydGFvLXBvc3RhZ2VtIjp7Im51bWVybyI6IjAwNzM0MjM3NzciLCJjb250cmF0byI6Ijk5OTIxNTc5NzAiLCJkciI6MzYsImFwaSI6WzgwLDU3LDgzXX0sImlwIjoiMTAuMC4yMDQuMTAsMTAuMC4yMDQuMTAiLCJpYXQiOjE2NjYyNzY1MTcsImlzcyI6InRva2VuLXNlcnZpY2UiLCJleHAiOjE2NjYzNjI5MTcsImp0aSI6ImU0Yjk2OGFiLWQ3NTAtNDFmZC04MzRmLWNhMDQ4MzQwMmJkZCJ9.q98V9r8PtyGvKuQjPlagQ3SHgY52butxavCynC3tV7Tb3ttBzHC6bHAvCLH5xSRow8KP9YuYiDi9UVlA01jsoiINL1gGydOIf63pGK9lcROll_QBs0_Y50rfxxfxQ4LsOfAC8QZH8UvPb2slhJ4c6pjQZ79rjnqUtXmH13O-AJp3YOJNhlLLRn1hJY7RjbKRiXAdgg3QH6MaV4XtEfbUdVL4CGt6bamB3pvvQEodhRP2aO7Oeb-mkKrEzGbmP8lPlR4TJNRdHeTnMqCJmfY51NkDvH6a9_R_HttBQNK2vv59Csk2xenxUQ7iAtPfyw3RHFi5iZf7qRd0h7-ESDLUZg",
		true,
	)
}

func TestLogin(t *testing.T) {
	var err error

	client, err = NewClientToken("testeint", "NppSSY:4d:%x_b$)GQ7j-^}I5JF3MDY7i|2]yoz9", "0073423777", true)
	spew.Dump(client)
	assert.NoError(t, err)
}
