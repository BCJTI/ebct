package ebct

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

var client *Client

func init() {

	client = NewClient(
		"eyJhbGciOiJSUzUxMiJ9.eyJhbWJpZW50ZSI6IlBST0RVQ0FPIiwiaWQiOiJwaG9lbmV4Y2FyZ28iLCJwZmwiOiJQSiIsImNucGoiOiIxMDI1NzYwMjAwMDE4MiIsImFwaSI6WzU2M10sImNhcnRhby1wb3N0YWdlbSI6eyJudW1lcm8iOiIwMDY5Mjk1ODE2IiwiY29udHJhdG8iOiI5OTEyMzUyMTY3IiwiZHIiOjcyLCJhcGkiOlsyNywzNCwzNSwzNiwzNyw0MSw3Niw3OCw4MCw4Myw4Nyw1NjYsNTg3XX0sImlwIjoiMTc5LjEwMS4yMTYuMjE0LCAxOTIuMTY4LjEuMTMyIiwiaWF0IjoxNzIxNzQ2MDM0LCJpc3MiOiJ0b2tlbi1zZXJ2aWNlIiwiZXhwIjoxNzIxODMyNDM0LCJqdGkiOiJjYWQyMTk2OC03MGRmLTRkYjUtOGFhNS0wNWJjMTk1NDFmNGMifQ.yeISSHn_uZsn4lSkFuPjiFuPhFyF_WSpzObGfngR2TEAcyTq99mS-sYgfdyxNV8YPf5hL6Y-FbS0SIeBY7UllntdEGFofyLW7lmGKQIuVybqfjeocGbKL15Cw-c3kbMY8trU5JQcUge_bSVsrjsZzMOBRCmUEPfRD0zxYyJEHKDxHmdyUfEE29_WF3tJ0BbcTx68Dzt7qH2XgMKfmY7VBK-PybS0I7DC1ZlC3DiPFc1-deLmgCNKeCj094jdlesRJu6hLOCw80QegrlVhkNGQDERO0okVPYenoO-FoOssZ7I5rSmxtz4WFJmn-Twz4QR5kcB0RwxpuCk8kYLdh9RXA",
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
