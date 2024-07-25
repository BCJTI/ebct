package ebct

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

var client *Client

func init() {

	client = NewClient(
		"eyJhbGciOiJSUzUxMiJ9.eyJhbWJpZW50ZSI6IlBST0RVQ0FPIiwiaWQiOiJwaG9lbmV4Y2FyZ28iLCJwZmwiOiJQSiIsImNucGoiOiIxMDI1NzYwMjAwMDE4MiIsImFwaSI6WzU2M10sImNhcnRhby1wb3N0YWdlbSI6eyJudW1lcm8iOiIwMDY5Mjk1ODE2IiwiY29udHJhdG8iOiI5OTEyMzUyMTY3IiwiZHIiOjcyLCJhcGkiOlsyNywzNCwzNSwzNiwzNyw0MSw3Niw3OCw4MCw4Myw4Nyw1NjYsNTg3XX0sImlwIjoiMTc5LjEwMS4yMDguMjEzLCAxOTIuMTY4LjEuMTMwIiwiaWF0IjoxNzIxOTM1NjE4LCJpc3MiOiJ0b2tlbi1zZXJ2aWNlIiwiZXhwIjoxNzIyMDIyMDE4LCJqdGkiOiJiMDAwYTJjZS00YmI5LTQ4ZDktOWI0OS1hNDlkM2M1YzNjNzcifQ.UUn7QA-DP1qxwwlfHvJL1JBZy3E3uEe8AVWUI-KrtxaawN83Nt4_92xF2WS3TIJ4fXhqxUMwc9E5LdIeKb0qcCVrj5h5Prza7Owmc9UAudbra08fxstNLfBmhlydWBhIkn7GpyPRRKqk5QsN9FGwgcDwLj9R0iF8k7kFS4uJiZKg3vISamzlJtqpIKURs4XDoFT-ORLpqf3ebbfIbJ24bH-jtTbpqr4rpWuvm0LY66xBPgbAZyrl1KKQafO8WiAJiStX5TIaD74V3kctjYH9gR4wf7Y1ZCY45JFGjdNdOnr_RxMCb10dG9Vp83L89Sw8nnfoQKjU-ntDqUcRlbZl2w",
		false,
	)
}

func TestLogin(t *testing.T) {
	var err error

	//client, err = NewClientToken("05317708000194", "0ywD1PYCdkzojfPUGAfifFyTIHFOhTZUoGZgOVpP", "9912435103", "AUTH", true)
	client, err = NewClientToken("phoenexcargo", "pCXd7tgr7dGRXkWnlZ2FGsxGMMQ2QjeSmwmSvQ2p", "0069295816", "CARTAOPOSTAGEM", false)

	spew.Dump(client)
	assert.NoError(t, err)
}
