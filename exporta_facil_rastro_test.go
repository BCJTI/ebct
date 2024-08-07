package ebct

import (
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

const RastRemessaInputCancel = `{
	  "idSolicitacao": 378290723,
	  "dtEvento": "2024-06-03",
	  "idCorreios": "05317708000194",
	  "motivoCancelamento": "cancelado por motivo qualquer"
	}`

const RastRemessaInput = `{
	  "idSolicitacao": 378290723,
	  "evento": 4,
	  "dtEvento": "2024-06-03",
	  "etiqueta": "EF2024US378290723",
	  "observacao": "remessa teste fox",
	  "idCorreios": "05317708000194"
	}`

func TestGetExpFacilRemessa(t *testing.T) {

	expFacilRemessa := &RastreamentoRemessaGet{
		CnpjOperador: "05317708000194",
	}

	requestReturn, err := client.GetExpFacilRemessa(expFacilRemessa)
	spew.Dump(requestReturn)
	assert.NoError(t, err)
}

func TestGetExpFacilEvento(t *testing.T) {

	requestReturn, err := client.GetExpFacilEvento()
	spew.Dump(requestReturn)
	assert.NoError(t, err)
}

func TestCancelExpFacilRastro(t *testing.T) {

	rastRemessaCancel := &RastreamentoRemessaInputCancel{}
	var bPack = []byte(RastRemessaInputCancel)
	var err error
	err = json.Unmarshal(bPack, rastRemessaCancel)
	if err != nil {
		spew.Dump(err)
	}

	requestReturn, err := client.CancelExpFacilRastro(rastRemessaCancel)
	spew.Dump(requestReturn)
	assert.NoError(t, err)
}

func TestAddExpFacilRastro(t *testing.T) {

	rastRemessaPost := &RastreamentoRemessaInput{}
	var bPack = []byte(RastRemessaInput)
	var err error
	err = json.Unmarshal(bPack, rastRemessaPost)
	if err != nil {
		spew.Dump(err)
	}

	requestReturn, err := client.AddExpFacilRastro(rastRemessaPost)

	spew.Dump(requestReturn)
	assert.NoError(t, err)
}
