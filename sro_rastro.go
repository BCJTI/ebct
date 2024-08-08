package ebct

import (
	"errors"
	"os"
	"time"
)

const sroRastro = "srorastro/" + Version + "objetos"

type TimeStrRastro struct {
	time.Time
}

// UnmarshalJSON decodifica o JSON para o tipo TimeStr
func (t *TimeStrRastro) UnmarshalJSON(data []byte) error {
	os.Setenv("TZ", "America/Sao_Paulo")

	str := string(data)
	if str == "null" || str == "" {
		*t = TimeStrRastro{time.Time{}}
		return nil
	}
	parsedTime, err := time.Parse(`"2006-01-02T15:04:05"`, str)
	if err != nil {
		parsedTime2, err2 := time.Parse(`"2006-01-02"`, str)
		if err2 != nil {
			return err2
		}
		parsedTime = parsedTime2.Add(time.Hour * 3)
	}
	*t = TimeStrRastro{parsedTime}
	return nil
}

// MarshalJSON codifica o tipo TimeStr para JSON
func (t TimeStrRastro) MarshalJSON() ([]byte, error) {
	if t.Time.IsZero() {
		return []byte("null"), nil
	}
	return []byte(`"` + t.Format("2006-01-02T15:04:05") + `"`), nil
}

type RastroRequest struct {
	CodigosObjetos []string `json:"codigosObjetos,omitempty"`
	Resultado      *string  `json:"resultado,omitempty"`
}

type SroRastro struct {
	Versao        string   `json:"versao,omitempty"`
	Quantidade    int      `json:"quantidade,omitempty"`
	Objetos       []Objeto `json:"objetos,omitempty"`
	TipoResultado string   `json:"tipoResultado,omitempty"`
}

type Objeto struct {
	CodObjeto         string        `json:"codObjeto,omitempty"`
	TipoPostal        TipoPostal    `json:"tipoPostal,omitempty"`
	Mensagem          string        `json:"mensagem,omitempty"`
	DtPrevista        TimeStrRastro `json:"dtPrevista,omitempty"`
	MultiVolume       string        `json:"multiVolume,omitempty"`
	Afterschedule     string        `json:"afterschedule,omitempty"`
	Volume            int           `json:"volume,omitempty"`
	ValorRecebido     string        `json:"valorRecebido,omitempty"`
	EntregaProgramada TimeStrRastro `json:"entregaProgramada,omitempty"`
	PrazoTratamento   int           `json:"prazoTratamento,omitempty"`
	Contrato          string        `json:"contrato,omitempty"`
	CodAdm            string        `json:"codAdm,omitempty"`
	CartaoPostagem    string        `json:"cartaoPostagem,omitempty"`
	CNPJ              string        `json:"CNPJ,omitempty"`
	Largura           float64       `json:"largura,omitempty"`
	Comprimento       float64       `json:"comprimento,omitempty"`
	Altura            float64       `json:"altura,omitempty"`
	Diametro          float64       `json:"diametro,omitempty"`
	Peso              float64       `json:"peso,omitempty"`
	Formato           string        `json:"formato,omitempty"`
	CelRemetSMS       string        `json:"celRemetSMS,omitempty"`
	CelDestSMS        string        `json:"celDestSMS,omitempty"`
	ValorDeclarado    interface{}   `json:"valorDeclarado,omitempty"`
	Eventos           []Evento      `json:"eventos,omitempty"`
	Servico           Servico       `json:"servico,omitempty"`
}

type TipoPostal struct {
	Sigla     string `json:"sigla,omitempty"`
	Nome      string `json:"nome,omitempty"`
	Categoria string `json:"categoria,omitempty"`
	Familia   string `json:"familia,omitempty"`
}

type Evento struct {
	Codigo            string            `json:"codigo,omitempty"`
	Tipo              string            `json:"tipo,omitempty"`
	DtHrCriado        TimeStrRastro     `json:"dtHrCriado,omitempty"`
	Descricao         string            `json:"descricao,omitempty"`
	Estacao           string            `json:"estacao,omitempty"`
	Usuario           string            `json:"usuario,omitempty"`
	Carteiro          string            `json:"carteiro,omitempty"`
	Entregador        string            `json:"entregador,omitempty"`
	Latitude          string            `json:"latitude,omitempty"`
	Longitude         string            `json:"longitude,omitempty"`
	Detalhe           string            `json:"detalhe,omitempty"`
	DescCodigo        string            `json:"descCodigo,omitempty"`
	DtHrGravado       TimeStrRastro     `json:"dtHrGravado,omitempty"`
	Recebedor         Recebedor         `json:"recebedor,omitempty"`
	Unidade           Unidade           `json:"unidade,omitempty"`
	EntregadorExterno EntregadorExterno `json:"entregadorExterno,omitempty"`
	Contrato          string            `json:"contrato,omitempty"`
	CodAdm            string            `json:"codAdm,omitempty"`
	CartaoPostagem    string            `json:"cartaoPostagem,omitempty"`
	CNPJ              string            `json:"CNPJ,omitempty"`
	Remetente         Remetente         `json:"remetente,omitempty"`
	Destinatario      Destinatario      `json:"destinatario,omitempty"`
	DtLimiteRetirada  TimeStrRastro     `json:"dtLimiteRetirada,omitempty"`
	Comentario        string            `json:"comentario,omitempty"`
	CodLista          string            `json:"codLista,omitempty"`
	UnidadeDestino    Unidade           `json:"unidadeDestino,omitempty"`
}

type Recebedor struct {
	Nome       string `json:"nome,omitempty"`
	Documento  string `json:"documento,omitempty"`
	Celular    string `json:"celular,omitempty"`
	Email      string `json:"email,omitempty"`
	Comentario string `json:"comentario,omitempty"`
}

type Unidade struct {
	Nome     string   `json:"nome,omitempty"`
	CodSro   string   `json:"codSro,omitempty"`
	CodMcu   string   `json:"codMcu,omitempty"`
	Tipo     string   `json:"tipo,omitempty"`
	Endereco Endereco `json:"endereco,omitempty"`
	Codse    string   `json:"codse,omitempty"`
	Se       string   `json:"se,omitempty"`
}

type Endereco struct {
	Cep         string `json:"cep,omitempty"`
	Logradouro  string `json:"logradouro,omitempty"`
	Complemento string `json:"complemento,omitempty"`
	Numero      string `json:"numero,omitempty"`
	Bairro      string `json:"bairro,omitempty"`
	Cidade      string `json:"cidade,omitempty"`
	Uf          string `json:"uf,omitempty"`
	Pais        string `json:"pais,omitempty"`
	Telefone    string `json:"telefone,omitempty"`
}

type EntregadorExterno struct {
	Documento string `json:"documento,omitempty"`
	Nome      string `json:"nome,omitempty"`
}

type Remetente struct {
	Nome      string     `json:"nome,omitempty"`
	Documento string     `json:"documento,omitempty"`
	Endereco  Endereco   `json:"endereco,omitempty"`
	Telefones []Telefone `json:"telefones,omitempty"`
}

type Destinatario struct {
	Nome      string     `json:"nome,omitempty"`
	Documento string     `json:"documento,omitempty"`
	Email     string     `json:"email,omitempty"`
	Telefones []Telefone `json:"telefones,omitempty"`
	Endereco  Endereco   `json:"endereco,omitempty"`
}

type Telefone struct {
	Tipo   string `json:"tipo,omitempty"`
	Ddd    string `json:"ddd,omitempty"`
	Numero string `json:"numero,omitempty"`
}

type Servico struct {
	Codigo      string `json:"codigo,omitempty"`
	Ar          string `json:"ar,omitempty"`
	TipoAr      string `json:"tipoAr,omitempty"`
	Mp          string `json:"mp,omitempty"`
	Vd          string `json:"vd,omitempty"`
	EVizinho    string `json:"eVizinho,omitempty"`
	FotoFachada string `json:"fotoFachada,omitempty"`
	EControlada string `json:"eControlada,omitempty"`
	EInterativa string `json:"eInterativa,omitempty"`
	ELocker     string `json:"eLocker,omitempty"`
}

func (c *Client) GetSroRastroObjetos(codObjeto []string) (SroRastro, error) {

	sendRequest := &RastroRequest{}
	var err error

	rastRemess := SroRastro{}
	if len(codObjeto) == 0 {
		err = errors.New("the codObjeto must be a valid")
		return rastRemess, err
	}
	sendRequest.CodigosObjetos = codObjeto
	sendRequest.Resultado = StringPtr("T")

	err = c.Get(sroRastro, sendRequest, nil, &rastRemess)
	return rastRemess, err
}
