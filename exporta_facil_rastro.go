package ebct

import (
	"errors"
	"time"
)

const exportaFacilRemessa = ExportaFacil + Version + "remessas"
const exportaFacilRastros = ExportaFacil + Version + "rastros"
const exportaFacilEventos = ExportaFacil + Version + "eventos"

// GetExpFacilRemessa
// RastreamentoRemessaGet define a estrutura do JSON de entrada
type RastreamentoRemessaGet struct {
	Etiqueta     *string `json:"etiqueta,omitempty"`
	CnpjOperador string  `json:"cnpjOperador,omitempty"`
}

// GetExpFacilRemessa
// RastreamentoRemessaResult define a estrutura do JSON de resposta
type RastreamentoRemessaResult struct {
	IdSolicitacao int    `json:"idSolicitacao"`
	CnpjOperador  string `json:"cnpjOperador"`
	Etiqueta      string `json:"etiqueta"`
}

// GetExpFacilEvento
// RastreamentoEventoResult define a estrutura do JSON de resposta
type RastreamentoEventoResult struct {
	Id        int    `json:"id"`
	Descricao string `json:"descricao"`
}

// RastreamentoRemessaInputCancel define a estrutura do JSON de entrada
type RastreamentoRemessaInputCancel struct {
	IdSolicitacao      int    `json:"idSolicitacao"`
	DtEvento           string `json:"dtEvento"`
	IdCorreios         string `json:"idCorreios"`
	MotivoCancelamento string `json:"motivoCancelamento"`
}

// RastreamentoRemessaResponseCancel define a estrutura do JSON de entrada
type RastreamentoRemessaResponseCancel struct {
	IdSolicitacao      int    `json:"idSolicitacao"`
	IdEvento           int    `json:"idEvento"`
	DataRegistro       string `json:"dataRegistro"`
	DtEvento           string `json:"dtEvento"`
	Observacao         string `json:"observacao"`
	IdCorreios         string `json:"idCorreios"`
	MotivoCancelamento string `json:"motivoCancelamento"`
	DataCancelamento   string `json:"dataCancelamento"`
}

// RastreamentoRemessaInput define a estrutura do JSON de entrada
type RastreamentoRemessaInput struct {
	IdSolicitacao int    `json:"idSolicitacao"`
	Evento        int    `json:"evento"`
	DtEvento      string `json:"dtEvento"`
	Etiqueta      string `json:"etiqueta"`
	Observacao    string `json:"observacao"`
	IdCorreios    string `json:"idCorreios"`
}

// RastreamentoRemessaResponse define a estrutura do JSON de resposta
type RastreamentoRemessaResponse struct {
	IdSolicitacao      int        `json:"idSolicitacao"`
	IdEvento           int        `json:"idEvento"`
	DataRegistro       string     `json:"dataRegistro"`
	DtEvento           string     `json:"dtEvento"`
	Observacao         string     `json:"observacao"`
	IdCorreios         *string    `json:"idCorreios"`
	MotivoCancelamento *string    `json:"motivoCancelamento,omitempty"`
	DataCancelamento   *time.Time `json:"dataCancelamento,omitempty"`
}

func (c *Client) GetExpFacilRemessa(req *RastreamentoRemessaGet) ([]RastreamentoRemessaResult, error) {

	var err error

	if req.CnpjOperador == "" {
		err = errors.New("the CnpjOperador must be a valid")
		return nil, err
	}

	rastRemess := []RastreamentoRemessaResult{}
	err = c.Get(exportaFacilRemessa, req, nil, &rastRemess)
	return rastRemess, err
}

func (c *Client) GetExpFacilEvento() ([]RastreamentoEventoResult, error) {

	var err error

	rastEvento := []RastreamentoEventoResult{}
	err = c.Get(exportaFacilEventos, nil, nil, &rastEvento)
	return rastEvento, err
}

func (c *Client) CancelExpFacilRastro(unitcodes *RastreamentoRemessaInputCancel) (*RastreamentoRemessaResponseCancel, error) {

	newExpFacilRastroReturn := &RastreamentoRemessaResponseCancel{}
	err := c.Put(exportaFacilRastros, unitcodes, nil, newExpFacilRastroReturn)
	return newExpFacilRastroReturn, err
}

func (c *Client) AddExpFacilRastro(unitcodes *RastreamentoRemessaInput) (*RastreamentoRemessaResponse, error) {

	newExpFacilRastroReturn := &RastreamentoRemessaResponse{}
	err := c.Post(exportaFacilRastros, unitcodes, nil, newExpFacilRastroReturn)
	return newExpFacilRastroReturn, err
}
