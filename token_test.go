package ebct

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

var client *Client

func init() {

	client = NewClient(
		"eyJhbGciOiJSUzUxMiJ9.eyJhbWJpZW50ZSI6IkhPTU9MT0dBQ0FPIiwiaWQiOiIwNTMxNzcwODAwMDE5NCIsInBmbCI6IlBKIiwiY25waiI6IjA1MzE3NzA4MDAwMTk0IiwiYXBpIjpbNTYzXSwiaXAiOiIxNzkuMTAxLjIxMy4xODcsIDE5Mi4xNjguMS44IiwiaWF0IjoxNzE3NDIyMDAzLCJpc3MiOiJ0b2tlbi1zZXJ2aWNlIiwiZXhwIjoxNzE3NTA4NDAzLCJqdGkiOiJmZDc3YTQ4MC0zYzk2LTQ4NmUtOWJhYi00NTUzYmI3Y2QyMWQifQ.O3wU-E2LI5pfnKirU9SiLrVEu1fudcDVlrGpTrcyujO_vDFm6mEkXivfCdLOJ3ZK4_U607J6XfxdW1lmf5f2DMwD8g0GA9HDQzJF81cDmn6PZtnKw4dqYWwXQ8OihQf-B1xMROJCSZ3fjUiNuRLjW9GCN_gQnVYi-A7AVOgWIo_Jj7aCazpRhN8tmUCleiuMq_1DWjHvZ46UuEyb6HbzSMgRbBNB9CeZxCAn7e0Nn_mh361KnigoldfZBm-dSsQ03gDHvPV-HBqou7BuoAndNDnqtJ0MxeGtvHdul5TOejAVwOS2frAb2aO-5hQyFOtl83ocKlteBtzB4pYJc64pdQ",
		true,
	)
}

func TestLogin(t *testing.T) {
	var err error

	client, err = NewClientToken("05317708000194", "0ywD1PYCdkzojfPUGAfifFyTIHFOhTZUoGZgOVpP", "9912435103", "AUTH", true)

	spew.Dump(client)
	assert.NoError(t, err)
}
