package ebct

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

var client *Client

func init() {

	client = NewClient(
		"eyJhbGciOiJSUzUxMiJ9.eyJhbWJpZW50ZSI6IkhPTU9MT0dBQ0FPIiwiaWQiOiIwNTMxNzcwODAwMDE5NCIsInBmbCI6IlBKIiwiY25waiI6IjA1MzE3NzA4MDAwMTk0IiwiYXBpIjpbNTYzXSwiaXAiOiIxNzkuMTAxLjIwOC4yNTQsIDE5Mi4xNjguMS44IiwiaWF0IjoxNzE2ODE1MzA2LCJpc3MiOiJ0b2tlbi1zZXJ2aWNlIiwiZXhwIjoxNzE2OTAxNzA2LCJqdGkiOiI0YjFmM2Q0Zi1mZmEwLTQxZGItYmY5Yi02ODA4MDZkZGZlNmUifQ.Mk3WgcG5dpMC-ibr8YxxAX17oibl16TbjQVjqZccLR2ZofAziwAq_U26tGRN8mlhj4Ee9TYYAVmN7Te67c_KCDT0pjVgll51PVeCsZj_wMqPpnpiu_4VCCF61WxiLAWc0Lvb0XaZ5lUt8GOpI2ZMNiVLpOkoePrz3BpNUmcmIZxac30qvPKFmMZoWrVykludrmFdA9se58VDonPsQ_L1rfZYjA0MxClwbiBgfnC_1f2GvfAGe2ge6iL8sogQ-R_GeuF8ncNwh0CRTq_iCUIAn7pl84DhqhujnD0x8o77xZRTkmXQuhJXnS-r9ismjlUdfZ10kN-isgRLKPHeBz4RBw",
		true,
	)
}

func TestLogin(t *testing.T) {
	var err error

	client, err = NewClientToken("05317708000194", "0ywD1PYCdkzojfPUGAfifFyTIHFOhTZUoGZgOVpP", "9912435103", "AUTH", true)

	spew.Dump(client)
	assert.NoError(t, err)
}
