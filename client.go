package ebct

import (
	"encoding/base64"
	"fmt"
)

type Client struct {
	Token            string
	Sandbox          bool
	BasicAuth        string
	Package          string
	ExpiryDate       string
	IdCorreios       string
	DsContract       string
	DsCartaoPostagem string
	DsNamespace      string
	FlSoap           bool
}

func NewClientToken(user, pass, contract, cartaopostagem, tpPost string, sandbox bool) (*Client, error) {

	c := &Client{
		BasicAuth:        fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(user+":"+pass))),
		Package:          contract,
		Sandbox:          sandbox,
		IdCorreios:       user,
		DsContract:       contract,
		DsCartaoPostagem: cartaopostagem,
	}

	tmpLoginToken := LoginToken{
		Numero: contract,
	}

	model := new(TokenResponse)

	var err error

	switch tpPost {
	case "AUTH":
		err = c.Post(tokenAuth, nil, nil, model)
	case "CONTRATO":
		err = c.Post(tokenAuthContrato, tmpLoginToken, nil, model)
	case "CARTAOPOSTAGEM":
		tmpLoginToken.Numero = cartaopostagem
		err = c.Post(tokenAuthCartaoPostagem, tmpLoginToken, nil, model)
	case "LOGISTICAREVERSA":
		err = nil
	}

	if err == nil {
		if *model.Token != "" {
			c.Token = *model.Token
			c.ExpiryDate = *model.ExpiraEm
		}
	}

	return c, err
}

func NewClient(user, pass, token string, sandbox bool) *Client {
	return &Client{
		BasicAuth: fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(user+":"+pass))),
		Token:     token,
		Sandbox:   sandbox,
	}
}

func (c *Client) GetToken() string {
	if c.Token != "" {
		return fmt.Sprintf("Bearer %s", c.Token)
	}
	return ""
}

func (c *Client) GetEndpoint() string {

	if c.FlSoap {
		if c.Sandbox {
			return SandBoxUrlSoap
		}

		return ProductionUrlSoap
	}
	if c.Sandbox {
		return SandBoxUrl
	}

	return ProductionUrl

}
