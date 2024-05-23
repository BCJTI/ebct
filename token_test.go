package ebct

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

var client *Client

func init() {

	client = NewClient(
		"eyJhbGciOiJSUzUxMiJ9.eyJhbWJpZW50ZSI6IlBST0RVQ0FPIiwiaWQiOiJlcHNhZG1pbjQ2NzUiLCJwZmwiOiJQSiIsImNucGoiOiIwNTMxNzcwODAwMDE5NCIsImFwaSI6WzU2M10sImNvbnRyYXRvIjp7Im51bWVybyI6Ijk5MTI0MzUxMDMiLCJkciI6NzIsImFwaSI6WzI3LDM0LDM1LDQxLDc2LDc4LDg3LDU2Niw1ODddfSwiaXAiOiIxNzcuMTYyLjkxLjE4NSwgMTkyLjE2OC4xLjEzMSIsImlhdCI6MTcxNjMyNDI2MiwiaXNzIjoidG9rZW4tc2VydmljZSIsImV4cCI6MTcxNjQxMDY2MiwianRpIjoiYTBhYmJlMWMtZDAyMi00ODZjLTk5MzItNDA2OTQwMTkzNzcxIn0.WSwPpPJpGXAOjI7BBVecMk5ls4HkwzB01oFWAGR6AnSlg2Zf1t-HbdmPPFayaX9_4S5rlo5ZolkTp38n_3MR4nU7HbpbTU3WCIODvR_pZ2NtirS10Dd0t3re5PKTzPnZsCk_wsoVXCHSHOjHg2E8q5TQvIcclEmKvgoUnSqtybrrjl2riA_H3MUd4NudD5dYAWnf6Pod9gS84Ivy3wPk4R4xZgfzDLOd3_6ishdDB9Dc5rHhm3UAs7dlpOzVa45eZltOqldFCmFRyi70zbw643RE9dbTXEVWXFtFWPa6FuPayqiaSquj0X-HwYl5f2EhZ9PARNNoIcwoXjvdswleUQ",
		true,
	)
}

func TestLogin(t *testing.T) {
	var err error

	//client, err = NewClientToken("epsadmin4675", "correiofox123", "9912435103", true)
	client, err = NewClientToken("05317708000194", "correiofox123", "9912435103", true)
	spew.Dump(client)
	assert.NoError(t, err)
}
