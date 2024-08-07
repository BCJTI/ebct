package ebct

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

var client *Client

func init() {

	client = NewClient("empresacws",
		"123456",
		"",
		true,
	)
}

func TestLogin(t *testing.T) {
	var err error

	client, err = NewClientToken("epsadmin4675", "P7aesZtQGb4nyT7UzhkMfXYYBmYinQCwbEcRbON1", "9912435103", "0073881236", "LOGISTICAREVERSA", false)
	//client, err = NewClientToken("phoenexcargo", "pCXd7tgr7dGRXkWnlZ2FGsxGMMQ2QjeSmwmSvQ2p", "9912352167", "0069295816", "CARTAOPOSTAGEM", false)

	spew.Dump(client)
	assert.NoError(t, err)
}
