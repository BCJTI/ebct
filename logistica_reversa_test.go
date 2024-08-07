package ebct

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/fiorix/wsdl2go/soap"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func TestSolicitarPostagemReversa2(t *testing.T) {
	const SolicitarPostagemReversaInput = `{
    "codAdministrativo": "17000190",
    "codigo_servico": "04677",
    "cartao": "0067599079",
    "destinatario": {
      "nome": "João Silva",
      "logradouro": "São Paulo",
      "numero": "123",
      "complemento": "ap 123",
      "bairro": "jardim",
      "referencia": "alto",
      "cidade": "São Paulo",
      "uf": "SP",
      "cep": "01234567",
      "ddd": "11",
      "telefone": "99999999",
      "ddd_celular": "11",
      "celular": "999999999",
      "email": "abc@abc.com",
      "identificacao": "aaaaa",
      "ciencia_conteudo_proibido": "1",
      "dnec": {
        "dnec_mensagem": "ccccc",
        "dnec_resultset": "dddd",
        "dnec_retorno": 100,
        "dnec_total": 3.0
      }
    },
    "coletas_solicitadas": [
      {
        "tipo": "A",
        "id_cliente": "1133566",
        "valor_declarado": "100.50",
        "descricao": "Documentos importantes",
        "cklist": "Item 1, Item 2",
        "remetente": {
          "nome": "Maria Souza",
          "logradouro": "Rio de Janeiro",
          "numero": "123",
          "complemento": "ap 123",
          "bairro": "lapa",
          "referencia": "baixo",
          "cidade": "Rio de Janeiro",
          "uf": "RJ",
          "cep": "21000000",
          "ddd": "21",
          "telefone": "99999999",
          "email": "rio@rio.com.br",
          "ciencia_conteudo_proibido": "aaaaa",
          "identificacao": "bbbb",
          "ddd_celular": "21",
          "celular": "999999999",
          "sms": "S",
          "restricao_anac": "1",
          "documento_estrangeiro": ""
        },
        "produto": [
          {
            "codigo": "116600063",
            "tipo": "0",
            "qtd": "1"
          }
        ],
        "numero": "string",
        "ag": "string",
        "cartao": "string",
        "servico_adicional": "string",
        "ar": 100,
        "obj_col": [{
          "item": "1",
          "desc": "Certidão de Nascimento",
          "entrega": "Normal",
          "num": "001",
          "id": "555"
        }]
      }
    ],
"xmlnsSer":"http://service.logisticareversa.correios.com.br/"}
`

	sendObjects := SolicitarPostagemReversa{}
	var bPack = []byte(SolicitarPostagemReversaInput)
	var err error
	err = json.Unmarshal(bPack, &sendObjects)
	if err != nil {
		fmt.Println(err)
	}

	httpClient := &http.Client{
		Timeout: time.Second * 30,
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
		},
	}

	username := "empresacws"
	password := "123456"

	httpClient.Transport = &transportWithBasicAuth{
		Transport: httpClient.Transport,
		Username:  username,
		Password:  password,
	}

	soapClient := &soap.Client{
		URL:           "https://apphom.correios.com.br/logisticaReversaWS/logisticaReversaService/logisticaReversaWS",
		Namespace:     "http://service.logisticareversa.correios.com.br/",
		Envelope:      "http://schemas.xmlsoap.org/soap/envelope/",
		ThisNamespace: "tns",
		Config:        httpClient,
		UserAgent:     "Go-SOAP-Client/1.2",
		Ctx:           context.Background(),
	}
	soapClient.ExcludeActionNamespace = true

	service := NewLogisticaReversaWS(soapClient)

	requestReturn, err := service.SolicitarPostagemReversa(&sendObjects)
	if err != nil {
		fmt.Println("Erro ao solicitar postagem reversa:", err)
	} else {
		fmt.Println("Resposta da postagem reversa:", Serialize(requestReturn))
	}

	assert.NoError(t, err)
}

type transportWithBasicAuth struct {
	Transport http.RoundTripper
	Username  string
	Password  string
}

func (t *transportWithBasicAuth) RoundTrip(req *http.Request) (*http.Response, error) {
	req.SetBasicAuth(t.Username, t.Password)
	return t.Transport.RoundTrip(req)
}

func TestSolicitarPostagemReversa(t *testing.T) {
	const SolicitarPostagemReversaInput = `{
    "codAdministrativo": "17000190",
    "codigo_servico": "04677",
    "cartao": "0067599079",
    "destinatario": {
      "nome": "João Silva",
      "logradouro": "São Paulo",
      "numero": "123",
      "complemento": "ap 123",
      "bairro": "jardim",
      "referencia": "alto",
      "cidade": "São Paulo",
      "uf": "SP",
      "cep": "01234567",
      "ddd": "11",
      "telefone": "99999999",
      "ddd_celular": "11",
      "celular": "999999999",
      "email": "abc@abc.com",
      "identificacao": "aaaaa",
      "ciencia_conteudo_proibido": "1",
      "dnec": {
        "dnec_mensagem": "ccccc",
        "dnec_resultset": "dddd",
        "dnec_retorno": 100,
        "dnec_total": 3.0
      }
    },
    "coletas_solicitadas": [
      {
        "tipo": "A",
        "id_cliente": "1133566",
        "valor_declarado": "100.50",
        "descricao": "Documentos importantes",
        "cklist": "Item 1, Item 2",
        "remetente": {
          "nome": "Maria Souza",
          "logradouro": "Rio de Janeiro",
          "numero": "123",
          "complemento": "ap 123",
          "bairro": "lapa",
          "referencia": "baixo",
          "cidade": "Rio de Janeiro",
          "uf": "RJ",
          "cep": "21000000",
          "ddd": "21",
          "telefone": "99999999",
          "email": "rio@rio.com.br",
          "ciencia_conteudo_proibido": "aaaaa",
          "identificacao": "bbbb",
          "ddd_celular": "21",
          "celular": "999999999",
          "sms": "S",
          "restricao_anac": "1",
          "documento_estrangeiro": ""
        },
        "produto": [
          {
            "codigo": "116600063",
            "tipo": "0",
            "qtd": "1"
          }
        ],
        "numero": "string",
        "ag": "string",
        "cartao": "string",
        "servico_adicional": "string",
        "ar": 100,
        "obj_col": [{
          "item": "1",
          "desc": "Certidão de Nascimento",
          "entrega": "Normal",
          "num": "001",
          "id": "555"
        }]
      }
    ],
"xmlnsSer":"http://service.logisticareversa.correios.com.br/"}
`

	sendRequest := &SolicitarPostagemReversa{}
	var bPack = []byte(SolicitarPostagemReversaInput)
	var err error
	err = json.Unmarshal(bPack, sendRequest)
	if err != nil {
		spew.Dump(err)
	}

	requestReturn, err := client.PostSolicitarPostagemReversa(sendRequest)
	spew.Dump(requestReturn)
	assert.NoError(t, err)
}
