package ebct

import (
	"encoding/xml"
	"github.com/fiorix/wsdl2go/soap"
)

// Namespace was auto-generated from WSDL.
var Namespace = "http://service.logisticareversa.correios.com.br/"

// NewLogisticaReversaWS creates an initializes a LogisticaReversaWS.
func NewLogisticaReversaWS(cli *soap.Client) LogisticaReversaWS {
	return &logisticaReversaWS{cli}
}

// LogisticaReversaWS was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type LogisticaReversaWS interface {
	// AcompanharPedido was auto-generated from WSDL.
	AcompanharPedido(AcompanharPedido *AcompanharPedido) (*AcompanharPedidoResponse, error)

	// AcompanharPedidoPorData was auto-generated from WSDL.
	AcompanharPedidoPorData(AcompanharPedidoPorData *AcompanharPedidoPorData) (*AcompanharPedidoPorDataResponse, error)

	// CalcularDigitoVerificador was auto-generated from WSDL.
	CalcularDigitoVerificador(CalcularDigitoVerificador *CalcularDigitoVerificador) (*CalcularDigitoVerificadorResponse, error)

	// CancelarPedido was auto-generated from WSDL.
	CancelarPedido(CancelarPedido *CancelarPedido) (*CancelarPedidoResponse, error)

	// ConsultarResumoColeta was auto-generated from WSDL.
	ConsultarResumoColeta(ConsultarResumoColeta *ConsultarResumoColeta) (*ConsultarResumoColetaResponse, error)

	// RevalidarPrazoAutorizacaoPostagem was auto-generated from WSDL.
	RevalidarPrazoAutorizacaoPostagem(RevalidarPrazoAutorizacaoPostagem *RevalidarPrazoAutorizacaoPostagem) (*RevalidarPrazoAutorizacaoPostagemResponse, error)

	// SobreWebService was auto-generated from WSDL.
	SobreWebService(SobreWebService *SobreWebService) (*SobreWebServiceResponse, error)

	// SolicitarPostagemReversa was auto-generated from WSDL.
	SolicitarPostagemReversa(SolicitarPostagemReversa *SolicitarPostagemReversa) (*SolicitarPostagemReversaResponse, error)

	// SolicitarPostagemSimultanea was auto-generated from WSDL.
	SolicitarPostagemSimultanea(SolicitarPostagemSimultanea *SolicitarPostagemSimultanea) (*SolicitarPostagemSimultaneaResponse, error)

	// SolicitarRange was auto-generated from WSDL.
	SolicitarRange(SolicitarRange *SolicitarRange) (*SolicitarRangeResponse, error)

	// ValidarPostagemReversa was auto-generated from WSDL.
	ValidarPostagemReversa(ValidarPostagemReversa *ValidarPostagemReversa) (*ValidarPostagemReversaResponse, error)

	// ValidarPostagemSimultanea was auto-generated from WSDL.
	ValidarPostagemSimultanea(ValidarPostagemSimultanea *ValidarPostagemSimultanea) (*ValidarPostagemSimultaneaResponse, error)
}

// WebServiceFaultInfo was auto-generated from WSDL.
type WebServiceFaultInfo struct {
	Mensagem      *string `xml:"mensagem,omitempty" json:"mensagem,omitempty" yaml:"mensagem,omitempty"`
	Excecao       *string `xml:"excecao,omitempty" json:"excecao,omitempty" yaml:"excecao,omitempty"`
	Classificacao *string `xml:"classificacao,omitempty" json:"classificacao,omitempty" yaml:"classificacao,omitempty"`
	Causa         *string `xml:"causa,omitempty" json:"causa,omitempty" yaml:"causa,omitempty"`
	StackTrace    *string `xml:"stackTrace,omitempty" json:"stackTrace,omitempty" yaml:"stackTrace,omitempty"`
}

// AcompanharPedido was auto-generated from WSDL.
type AcompanharPedido struct {
	CodAdministrativo *string   `xml:"codAdministrativo,omitempty" json:"codAdministrativo,omitempty" yaml:"codAdministrativo,omitempty"`
	TipoBusca         *string   `xml:"tipoBusca,omitempty" json:"tipoBusca,omitempty" yaml:"tipoBusca,omitempty"`
	TipoSolicitacao   *string   `xml:"tipoSolicitacao,omitempty" json:"tipoSolicitacao,omitempty" yaml:"tipoSolicitacao,omitempty"`
	NumeroPedido      []*string `xml:"numeroPedido,omitempty" json:"numeroPedido,omitempty" yaml:"numeroPedido,omitempty"`
}

// AcompanharPedidoPorData was auto-generated from WSDL.
type AcompanharPedidoPorData struct {
	CodAdministrativo *string `xml:"codAdministrativo,omitempty" json:"codAdministrativo,omitempty" yaml:"codAdministrativo,omitempty"`
	TipoSolicitacao   *string `xml:"tipoSolicitacao,omitempty" json:"tipoSolicitacao,omitempty" yaml:"tipoSolicitacao,omitempty"`
	Data              *string `xml:"data,omitempty" json:"data,omitempty" yaml:"data,omitempty"`
}

// AcompanharPedidoPorDataResponse was auto-generated from WSDL.
type AcompanharPedidoPorDataResponse struct {
	AcompanharPedidoPorData *RetornoAcompanhamento `xml:"acompanharPedidoPorData,omitempty" json:"acompanharPedidoPorData,omitempty" yaml:"acompanharPedidoPorData,omitempty"`
}

// AcompanharPedidoResponse was auto-generated from WSDL.
type AcompanharPedidoResponse struct {
	AcompanharPedido *RetornoAcompanhamento `xml:"acompanharPedido,omitempty" json:"acompanharPedido,omitempty" yaml:"acompanharPedido,omitempty"`
}

// CalcularDigitoVerificador was auto-generated from WSDL.
type CalcularDigitoVerificador struct {
	Numero *string `xml:"numero,omitempty" json:"numero,omitempty" yaml:"numero,omitempty"`
}

// CalcularDigitoVerificadorResponse was auto-generated from WSDL.
type CalcularDigitoVerificadorResponse struct {
	CalcularDigitoVerificador *RetornoDigitoVerificador `xml:"calcularDigitoVerificador,omitempty" json:"calcularDigitoVerificador,omitempty" yaml:"calcularDigitoVerificador,omitempty"`
}

// CancelarPedido was auto-generated from WSDL.
type CancelarPedido struct {
	CodAdministrativo *string `xml:"codAdministrativo,omitempty" json:"codAdministrativo,omitempty" yaml:"codAdministrativo,omitempty"`
	NumeroPedido      *string `xml:"numeroPedido,omitempty" json:"numeroPedido,omitempty" yaml:"numeroPedido,omitempty"`
	Tipo              *string `xml:"tipo,omitempty" json:"tipo,omitempty" yaml:"tipo,omitempty"`
}

// CancelarPedidoResponse was auto-generated from WSDL.
type CancelarPedidoResponse struct {
	CancelarPedido *RetornoCancelamento `xml:"cancelarPedido,omitempty" json:"cancelarPedido,omitempty" yaml:"cancelarPedido,omitempty"`
}

// Coleta was auto-generated from WSDL.
type Coleta struct {
	Tipo            *string           `xml:"tipo,omitempty" json:"tipo,omitempty" yaml:"tipo,omitempty"`
	Id_cliente      *string           `xml:"id_cliente,omitempty" json:"id_cliente,omitempty" yaml:"id_cliente,omitempty"`
	Valor_declarado *string           `xml:"valor_declarado,omitempty" json:"valor_declarado,omitempty" yaml:"valor_declarado,omitempty"`
	Descricao       *string           `xml:"descricao,omitempty" json:"descricao,omitempty" yaml:"descricao,omitempty"`
	Cklist          *string           `xml:"cklist,omitempty" json:"cklist,omitempty" yaml:"cklist,omitempty"`
	Documento       []*string         `xml:"documento,omitempty" json:"documento,omitempty" yaml:"documento,omitempty"`
	Remetente       *RemetenteReversa `xml:"remetente,omitempty" json:"remetente,omitempty" yaml:"remetente,omitempty"`
	Produto         []*Produto        `xml:"produto,omitempty" json:"produto,omitempty" yaml:"produto,omitempty"`
}

// ColetaReversa was auto-generated from WSDL.
type ColetaReversa struct {
	Tipo              *string           `xml:"tipo,omitempty" json:"tipo,omitempty" yaml:"tipo,omitempty"`
	Id_cliente        *string           `xml:"id_cliente,omitempty" json:"id_cliente,omitempty" yaml:"id_cliente,omitempty"`
	Valor_declarado   *string           `xml:"valor_declarado,omitempty" json:"valor_declarado,omitempty" yaml:"valor_declarado,omitempty"`
	Descricao         *string           `xml:"descricao,omitempty" json:"descricao,omitempty" yaml:"descricao,omitempty"`
	Cklist            *string           `xml:"cklist,omitempty" json:"cklist,omitempty" yaml:"cklist,omitempty"`
	Documento         []*string         `xml:"documento,omitempty" json:"documento,omitempty" yaml:"documento,omitempty"`
	Remetente         *RemetenteReversa `xml:"remetente,omitempty" json:"remetente,omitempty" yaml:"remetente,omitempty"`
	Produto           []*Produto        `xml:"produto,omitempty" json:"produto,omitempty" yaml:"produto,omitempty"`
	Numero            *string           `xml:"numero,omitempty" json:"numero,omitempty" yaml:"numero,omitempty"`
	Ag                *string           `xml:"ag,omitempty" json:"ag,omitempty" yaml:"ag,omitempty"`
	Cartao            *string           `xml:"cartao,omitempty" json:"cartao,omitempty" yaml:"cartao,omitempty"`
	Servico_adicional *string           `xml:"servico_adicional,omitempty" json:"servico_adicional,omitempty" yaml:"servico_adicional,omitempty"`
	Ar                *int              `xml:"ar,omitempty" json:"ar,omitempty" yaml:"ar,omitempty"`
	Obj_col           []*ObjetoReversa  `xml:"obj_col,omitempty" json:"obj_col,omitempty" yaml:"obj_col,omitempty"`
	TypeAttrXSI       string            `xml:"xsi:type,attr,omitempty"`
	TypeNamespace     string            `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *ColetaReversa) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:coletaReversa"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://service.logisticareversa.correios.com.br/"
	}
}

// ColetaSimultanea was auto-generated from WSDL.
type ColetaSimultanea struct {
	Tipo            *string           `xml:"tipo,omitempty" json:"tipo,omitempty" yaml:"tipo,omitempty"`
	Id_cliente      *string           `xml:"id_cliente,omitempty" json:"id_cliente,omitempty" yaml:"id_cliente,omitempty"`
	Valor_declarado *string           `xml:"valor_declarado,omitempty" json:"valor_declarado,omitempty" yaml:"valor_declarado,omitempty"`
	Descricao       *string           `xml:"descricao,omitempty" json:"descricao,omitempty" yaml:"descricao,omitempty"`
	Cklist          *string           `xml:"cklist,omitempty" json:"cklist,omitempty" yaml:"cklist,omitempty"`
	Documento       []*string         `xml:"documento,omitempty" json:"documento,omitempty" yaml:"documento,omitempty"`
	Remetente       *RemetenteReversa `xml:"remetente,omitempty" json:"remetente,omitempty" yaml:"remetente,omitempty"`
	Produto         []*Produto        `xml:"produto,omitempty" json:"produto,omitempty" yaml:"produto,omitempty"`
	Obs             *string           `xml:"obs,omitempty" json:"obs,omitempty" yaml:"obs,omitempty"`
	Obj             *string           `xml:"obj,omitempty" json:"obj,omitempty" yaml:"obj,omitempty"`
	TypeAttrXSI     string            `xml:"xsi:type,attr,omitempty"`
	TypeNamespace   string            `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *ColetaSimultanea) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:coletaSimultanea"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://service.logisticareversa.correios.com.br/"
	}
}

// ColetasSolicitadas was auto-generated from WSDL.
type ColetasSolicitadas struct {
	Numero_pedido    *string            `xml:"numero_pedido,omitempty" json:"numero_pedido,omitempty" yaml:"numero_pedido,omitempty"`
	Controle_cliente *string            `xml:"controle_cliente,omitempty" json:"controle_cliente,omitempty" yaml:"controle_cliente,omitempty"`
	Historico        []*HistoricoColeta `xml:"historico,omitempty" json:"historico,omitempty" yaml:"historico,omitempty"`
	Objeto           []*ObjetoPostal    `xml:"objeto,omitempty" json:"objeto,omitempty" yaml:"objeto,omitempty"`
}

// ConsultarResumoColeta was auto-generated from WSDL.
type ConsultarResumoColeta struct {
	CodigoObjeto *string `xml:"codigoObjeto,omitempty" json:"codigoObjeto,omitempty" yaml:"codigoObjeto,omitempty"`
}

// ConsultarResumoColetaResponse was auto-generated from WSDL.
type ConsultarResumoColetaResponse struct {
	ConsultarResumoColeta *RetornoResumoColeta `xml:"consultarResumoColeta,omitempty" json:"consultarResumoColeta,omitempty" yaml:"consultarResumoColeta,omitempty"`
}

// Destinatario was auto-generated from WSDL.
type DestinatarioReversa struct {
	Nome                      *string `xml:"nome,omitempty" json:"nome,omitempty" yaml:"nome,omitempty"`
	Logradouro                *string `xml:"logradouro,omitempty" json:"logradouro,omitempty" yaml:"logradouro,omitempty"`
	Numero                    *string `xml:"numero,omitempty" json:"numero,omitempty" yaml:"numero,omitempty"`
	Complemento               *string `xml:"complemento,omitempty" json:"complemento,omitempty" yaml:"complemento,omitempty"`
	Bairro                    *string `xml:"bairro,omitempty" json:"bairro,omitempty" yaml:"bairro,omitempty"`
	Referencia                *string `xml:"referencia,omitempty" json:"referencia,omitempty" yaml:"referencia,omitempty"`
	Cidade                    *string `xml:"cidade,omitempty" json:"cidade,omitempty" yaml:"cidade,omitempty"`
	Uf                        *string `xml:"uf,omitempty" json:"uf,omitempty" yaml:"uf,omitempty"`
	Cep                       *string `xml:"cep,omitempty" json:"cep,omitempty" yaml:"cep,omitempty"`
	Ddd                       *string `xml:"ddd,omitempty" json:"ddd,omitempty" yaml:"ddd,omitempty"`
	Telefone                  *string `xml:"telefone,omitempty" json:"telefone,omitempty" yaml:"telefone,omitempty"`
	Ddd_celular               *string `xml:"ddd_celular,omitempty" json:"ddd_celular,omitempty" yaml:"ddd_celular,omitempty"`
	Celular                   *string `xml:"celular,omitempty" json:"celular,omitempty" yaml:"celular,omitempty"`
	Email                     *string `xml:"email,omitempty" json:"email,omitempty" yaml:"email,omitempty"`
	Identificacao             *string `xml:"identificacao,omitempty" json:"identificacao,omitempty" yaml:"identificacao,omitempty"`
	Ciencia_conteudo_proibido *string `xml:"ciencia_conteudo_proibido,omitempty" json:"ciencia_conteudo_proibido,omitempty" yaml:"ciencia_conteudo_proibido,omitempty"`
	Dnec                      *DnecTO `xml:"dnec,omitempty" json:"dnec,omitempty" yaml:"dnec,omitempty"`
}

// DnecTO was auto-generated from WSDL.
type DnecTO struct {
	Dnec_mensagem  *string  `xml:"dnec_mensagem,omitempty" json:"dnec_mensagem,omitempty" yaml:"dnec_mensagem,omitempty"`
	Dnec_resultset *string  `xml:"dnec_resultset,omitempty" json:"dnec_resultset,omitempty" yaml:"dnec_resultset,omitempty"`
	Dnec_retorno   *int64   `xml:"dnec_retorno,omitempty" json:"dnec_retorno,omitempty" yaml:"dnec_retorno,omitempty"`
	Dnec_total     *float64 `xml:"dnec_total,omitempty" json:"dnec_total,omitempty" yaml:"dnec_total,omitempty"`
}

// EnderecoResumoTO was auto-generated from WSDL.
type EnderecoResumoTO struct {
	Nome        *string `xml:"nome,omitempty" json:"nome,omitempty" yaml:"nome,omitempty"`
	Logradouro  *string `xml:"logradouro,omitempty" json:"logradouro,omitempty" yaml:"logradouro,omitempty"`
	Numero      *string `xml:"numero,omitempty" json:"numero,omitempty" yaml:"numero,omitempty"`
	Complemento *string `xml:"complemento,omitempty" json:"complemento,omitempty" yaml:"complemento,omitempty"`
	Bairro      *string `xml:"bairro,omitempty" json:"bairro,omitempty" yaml:"bairro,omitempty"`
	Localidade  *string `xml:"localidade,omitempty" json:"localidade,omitempty" yaml:"localidade,omitempty"`
	Uf          *string `xml:"uf,omitempty" json:"uf,omitempty" yaml:"uf,omitempty"`
	Cep         *string `xml:"cep,omitempty" json:"cep,omitempty" yaml:"cep,omitempty"`
}

// HistoricoColeta was auto-generated from WSDL.
type HistoricoColeta struct {
	Status           *int    `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
	Descricao_status *string `xml:"descricao_status,omitempty" json:"descricao_status,omitempty" yaml:"descricao_status,omitempty"`
	Data_atualizacao *string `xml:"data_atualizacao,omitempty" json:"data_atualizacao,omitempty" yaml:"data_atualizacao,omitempty"`
	Hora_atualizacao *string `xml:"hora_atualizacao,omitempty" json:"hora_atualizacao,omitempty" yaml:"hora_atualizacao,omitempty"`
	Observacao       *string `xml:"observacao,omitempty" json:"observacao,omitempty" yaml:"observacao,omitempty"`
}

// Objeto was auto-generated from WSDL.
type ObjetoReversa struct {
	Item    *string `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
	Desc    *string `xml:"desc,omitempty" json:"desc,omitempty" yaml:"desc,omitempty"`
	Entrega *string `xml:"entrega,omitempty" json:"entrega,omitempty" yaml:"entrega,omitempty"`
	Num     *string `xml:"num,omitempty" json:"num,omitempty" yaml:"num,omitempty"`
	Id      *string `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
}

// ObjetoPostal was auto-generated from WSDL.
type ObjetoPostal struct {
	Numero_etiqueta         *string `xml:"numero_etiqueta,omitempty" json:"numero_etiqueta,omitempty" yaml:"numero_etiqueta,omitempty"`
	Controle_objeto_cliente *string `xml:"controle_objeto_cliente,omitempty" json:"controle_objeto_cliente,omitempty" yaml:"controle_objeto_cliente,omitempty"`
	Ultimo_status           *string `xml:"ultimo_status,omitempty" json:"ultimo_status,omitempty" yaml:"ultimo_status,omitempty"`
	Descricao_status        *string `xml:"descricao_status,omitempty" json:"descricao_status,omitempty" yaml:"descricao_status,omitempty"`
	Data_ultima_atualizacao *string `xml:"data_ultima_atualizacao,omitempty" json:"data_ultima_atualizacao,omitempty" yaml:"data_ultima_atualizacao,omitempty"`
	Hora_ultima_atualizacao *string `xml:"hora_ultima_atualizacao,omitempty" json:"hora_ultima_atualizacao,omitempty" yaml:"hora_ultima_atualizacao,omitempty"`
	Peso_cubico             *string `xml:"peso_cubico,omitempty" json:"peso_cubico,omitempty" yaml:"peso_cubico,omitempty"`
	Peso_real               *string `xml:"peso_real,omitempty" json:"peso_real,omitempty" yaml:"peso_real,omitempty"`
	Valor_postagem          *string `xml:"valor_postagem,omitempty" json:"valor_postagem,omitempty" yaml:"valor_postagem,omitempty"`
}

// ObjetoSimplificado was auto-generated from WSDL.
type ObjetoSimplificado struct {
	Numero_pedido         *string `xml:"numero_pedido,omitempty" json:"numero_pedido,omitempty" yaml:"numero_pedido,omitempty"`
	Status_pedido         *string `xml:"status_pedido,omitempty" json:"status_pedido,omitempty" yaml:"status_pedido,omitempty"`
	Datahora_cancelamento *string `xml:"datahora_cancelamento,omitempty" json:"datahora_cancelamento,omitempty" yaml:"datahora_cancelamento,omitempty"`
}

// Pessoa was auto-generated from WSDL.
type Pessoa struct {
	Nome                      *string `xml:"nome,omitempty" json:"nome,omitempty" yaml:"nome,omitempty"`
	Logradouro                *string `xml:"logradouro,omitempty" json:"logradouro,omitempty" yaml:"logradouro,omitempty"`
	Numero                    *string `xml:"numero,omitempty" json:"numero,omitempty" yaml:"numero,omitempty"`
	Complemento               *string `xml:"complemento,omitempty" json:"complemento,omitempty" yaml:"complemento,omitempty"`
	Bairro                    *string `xml:"bairro,omitempty" json:"bairro,omitempty" yaml:"bairro,omitempty"`
	Referencia                *string `xml:"referencia,omitempty" json:"referencia,omitempty" yaml:"referencia,omitempty"`
	Cidade                    *string `xml:"cidade,omitempty" json:"cidade,omitempty" yaml:"cidade,omitempty"`
	Uf                        *string `xml:"uf,omitempty" json:"uf,omitempty" yaml:"uf,omitempty"`
	Cep                       *string `xml:"cep,omitempty" json:"cep,omitempty" yaml:"cep,omitempty"`
	Ddd                       *string `xml:"ddd,omitempty" json:"ddd,omitempty" yaml:"ddd,omitempty"`
	Telefone                  *string `xml:"telefone,omitempty" json:"telefone,omitempty" yaml:"telefone,omitempty"`
	Email                     *string `xml:"email,omitempty" json:"email,omitempty" yaml:"email,omitempty"`
	Ciencia_conteudo_proibido *string `xml:"ciencia_conteudo_proibido,omitempty" json:"ciencia_conteudo_proibido,omitempty" yaml:"ciencia_conteudo_proibido,omitempty"`
}

// Produto was auto-generated from WSDL.
type Produto struct {
	Codigo *string `xml:"codigo,omitempty" json:"codigo,omitempty" yaml:"codigo,omitempty"`
	Tipo   *string `xml:"tipo,omitempty" json:"tipo,omitempty" yaml:"tipo,omitempty"`
	Qtd    *string `xml:"qtd,omitempty" json:"qtd,omitempty" yaml:"qtd,omitempty"`
}

// Remetente was auto-generated from WSDL.
type RemetenteReversa struct {
	Nome                      *string `xml:"nome,omitempty" json:"nome,omitempty" yaml:"nome,omitempty"`
	Logradouro                *string `xml:"logradouro,omitempty" json:"logradouro,omitempty" yaml:"logradouro,omitempty"`
	Numero                    *string `xml:"numero,omitempty" json:"numero,omitempty" yaml:"numero,omitempty"`
	Complemento               *string `xml:"complemento,omitempty" json:"complemento,omitempty" yaml:"complemento,omitempty"`
	Bairro                    *string `xml:"bairro,omitempty" json:"bairro,omitempty" yaml:"bairro,omitempty"`
	Referencia                *string `xml:"referencia,omitempty" json:"referencia,omitempty" yaml:"referencia,omitempty"`
	Cidade                    *string `xml:"cidade,omitempty" json:"cidade,omitempty" yaml:"cidade,omitempty"`
	Uf                        *string `xml:"uf,omitempty" json:"uf,omitempty" yaml:"uf,omitempty"`
	Cep                       *string `xml:"cep,omitempty" json:"cep,omitempty" yaml:"cep,omitempty"`
	Ddd                       *string `xml:"ddd,omitempty" json:"ddd,omitempty" yaml:"ddd,omitempty"`
	Telefone                  *string `xml:"telefone,omitempty" json:"telefone,omitempty" yaml:"telefone,omitempty"`
	Email                     *string `xml:"email,omitempty" json:"email,omitempty" yaml:"email,omitempty"`
	Ciencia_conteudo_proibido *string `xml:"ciencia_conteudo_proibido,omitempty" json:"ciencia_conteudo_proibido,omitempty" yaml:"ciencia_conteudo_proibido,omitempty"`
	Identificacao             *string `xml:"identificacao,omitempty" json:"identificacao,omitempty" yaml:"identificacao,omitempty"`
	Ddd_celular               *string `xml:"ddd_celular,omitempty" json:"ddd_celular,omitempty" yaml:"ddd_celular,omitempty"`
	Celular                   *string `xml:"celular,omitempty" json:"celular,omitempty" yaml:"celular,omitempty"`
	Sms                       *string `xml:"sms,omitempty" json:"sms,omitempty" yaml:"sms,omitempty"`
	Restricao_anac            *string `xml:"restricao_anac,omitempty" json:"restricao_anac,omitempty" yaml:"restricao_anac,omitempty"`
	Documento_estrangeiro     *string `xml:"documento_estrangeiro,omitempty" json:"documento_estrangeiro,omitempty" yaml:"documento_estrangeiro,omitempty"`
	TypeAttrXSI               string  `xml:"xsi:type,attr,omitempty"`
	TypeNamespace             string  `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *RemetenteReversa) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:remetente"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "http://service.logisticareversa.correios.com.br/"
	}
}

// ResultadoSolicitacao was auto-generated from WSDL.
type ResultadoSolicitacao struct {
	Tipo             *string `xml:"tipo,omitempty" json:"tipo,omitempty" yaml:"tipo,omitempty"`
	Id_cliente       *string `xml:"id_cliente,omitempty" json:"id_cliente,omitempty" yaml:"id_cliente,omitempty"`
	Numero_coleta    *string `xml:"numero_coleta,omitempty" json:"numero_coleta,omitempty" yaml:"numero_coleta,omitempty"`
	Numero_etiqueta  *string `xml:"numero_etiqueta,omitempty" json:"numero_etiqueta,omitempty" yaml:"numero_etiqueta,omitempty"`
	Id_obj           *string `xml:"id_obj,omitempty" json:"id_obj,omitempty" yaml:"id_obj,omitempty"`
	Status_objeto    *string `xml:"status_objeto,omitempty" json:"status_objeto,omitempty" yaml:"status_objeto,omitempty"`
	Prazo            *string `xml:"prazo,omitempty" json:"prazo,omitempty" yaml:"prazo,omitempty"`
	Data_solicitacao *string `xml:"data_solicitacao,omitempty" json:"data_solicitacao,omitempty" yaml:"data_solicitacao,omitempty"`
	Hora_solicitacao *string `xml:"hora_solicitacao,omitempty" json:"hora_solicitacao,omitempty" yaml:"hora_solicitacao,omitempty"`
	Codigo_erro      *int64  `xml:"codigo_erro,omitempty" json:"codigo_erro,omitempty" yaml:"codigo_erro,omitempty"`
	Descricao_erro   *string `xml:"descricao_erro,omitempty" json:"descricao_erro,omitempty" yaml:"descricao_erro,omitempty"`
}

// RetornoAcompanhamento was auto-generated from WSDL.
type RetornoAcompanhamento struct {
	Codigo_administrativo *string               `xml:"codigo_administrativo,omitempty" json:"codigo_administrativo,omitempty" yaml:"codigo_administrativo,omitempty"`
	Tipo_solicitacao      *string               `xml:"tipo_solicitacao,omitempty" json:"tipo_solicitacao,omitempty" yaml:"tipo_solicitacao,omitempty"`
	Coleta                []*ColetasSolicitadas `xml:"coleta,omitempty" json:"coleta,omitempty" yaml:"coleta,omitempty"`
	Data                  *string               `xml:"data,omitempty" json:"data,omitempty" yaml:"data,omitempty"`
	Hora                  *string               `xml:"hora,omitempty" json:"hora,omitempty" yaml:"hora,omitempty"`
	Cod_erro              *string               `xml:"cod_erro,omitempty" json:"cod_erro,omitempty" yaml:"cod_erro,omitempty"`
	Msg_erro              *string               `xml:"msg_erro,omitempty" json:"msg_erro,omitempty" yaml:"msg_erro,omitempty"`
}

// RetornoCancelamento was auto-generated from WSDL.
type RetornoCancelamento struct {
	Codigo_administrativo *string             `xml:"codigo_administrativo,omitempty" json:"codigo_administrativo,omitempty" yaml:"codigo_administrativo,omitempty"`
	Objeto_postal         *ObjetoSimplificado `xml:"objeto_postal,omitempty" json:"objeto_postal,omitempty" yaml:"objeto_postal,omitempty"`
	Data                  *string             `xml:"data,omitempty" json:"data,omitempty" yaml:"data,omitempty"`
	Hora                  *string             `xml:"hora,omitempty" json:"hora,omitempty" yaml:"hora,omitempty"`
	Cod_erro              *string             `xml:"cod_erro,omitempty" json:"cod_erro,omitempty" yaml:"cod_erro,omitempty"`
	Msg_erro              *string             `xml:"msg_erro,omitempty" json:"msg_erro,omitempty" yaml:"msg_erro,omitempty"`
}

// RetornoDigitoVerificador was auto-generated from WSDL.
type RetornoDigitoVerificador struct {
	Data     *string `xml:"data,omitempty" json:"data,omitempty" yaml:"data,omitempty"`
	Hora     *string `xml:"hora,omitempty" json:"hora,omitempty" yaml:"hora,omitempty"`
	Cod_erro *string `xml:"cod_erro,omitempty" json:"cod_erro,omitempty" yaml:"cod_erro,omitempty"`
	Msg_erro *string `xml:"msg_erro,omitempty" json:"msg_erro,omitempty" yaml:"msg_erro,omitempty"`
	Digito   *int    `xml:"digito,omitempty" json:"digito,omitempty" yaml:"digito,omitempty"`
	Numero   *string `xml:"numero,omitempty" json:"numero,omitempty" yaml:"numero,omitempty"`
}

// RetornoFaixaNumerica was auto-generated from WSDL.
type RetornoFaixaNumerica struct {
	Data          *string `xml:"data,omitempty" json:"data,omitempty" yaml:"data,omitempty"`
	Hora          *string `xml:"hora,omitempty" json:"hora,omitempty" yaml:"hora,omitempty"`
	Cod_erro      *string `xml:"cod_erro,omitempty" json:"cod_erro,omitempty" yaml:"cod_erro,omitempty"`
	Msg_erro      *string `xml:"msg_erro,omitempty" json:"msg_erro,omitempty" yaml:"msg_erro,omitempty"`
	Faixa_inicial *int    `xml:"faixa_inicial,omitempty" json:"faixa_inicial,omitempty" yaml:"faixa_inicial,omitempty"`
	Faixa_final   *int    `xml:"faixa_final,omitempty" json:"faixa_final,omitempty" yaml:"faixa_final,omitempty"`
}

// RetornoPostagem was auto-generated from WSDL.
type RetornoPostagem struct {
	Status_processamento  *string                 `xml:"status_processamento,omitempty" json:"status_processamento,omitempty" yaml:"status_processamento,omitempty"`
	Data_processamento    *string                 `xml:"data_processamento,omitempty" json:"data_processamento,omitempty" yaml:"data_processamento,omitempty"`
	Hora_processamento    *string                 `xml:"hora_processamento,omitempty" json:"hora_processamento,omitempty" yaml:"hora_processamento,omitempty"`
	Cod_erro              *string                 `xml:"cod_erro,omitempty" json:"cod_erro,omitempty" yaml:"cod_erro,omitempty"`
	Msg_erro              *string                 `xml:"msg_erro,omitempty" json:"msg_erro,omitempty" yaml:"msg_erro,omitempty"`
	Resultado_solicitacao []*ResultadoSolicitacao `xml:"resultado_solicitacao,omitempty" json:"resultado_solicitacao,omitempty" yaml:"resultado_solicitacao,omitempty"`
}

// RetornoResumoColeta was auto-generated from WSDL.
type RetornoResumoColeta struct {
	Numero_etiqueta            *string           `xml:"numero_etiqueta,omitempty" json:"numero_etiqueta,omitempty" yaml:"numero_etiqueta,omitempty"`
	Numero_coleta              *string           `xml:"numero_coleta,omitempty" json:"numero_coleta,omitempty" yaml:"numero_coleta,omitempty"`
	EnderecoResumoRemetente    *EnderecoResumoTO `xml:"enderecoResumoRemetente,omitempty" json:"enderecoResumoRemetente,omitempty" yaml:"enderecoResumoRemetente,omitempty"`
	EnderecoResumoDestinatario *EnderecoResumoTO `xml:"enderecoResumoDestinatario,omitempty" json:"enderecoResumoDestinatario,omitempty" yaml:"enderecoResumoDestinatario,omitempty"`
}

// RetornoRevalidarPrazo was auto-generated from WSDL.
type RetornoRevalidarPrazo struct {
	Numero_pedido *string `xml:"numero_pedido,omitempty" json:"numero_pedido,omitempty" yaml:"numero_pedido,omitempty"`
	Prazo         *string `xml:"prazo,omitempty" json:"prazo,omitempty" yaml:"prazo,omitempty"`
	Cod_erro      *string `xml:"cod_erro,omitempty" json:"cod_erro,omitempty" yaml:"cod_erro,omitempty"`
	Msg_erro      *string `xml:"msg_erro,omitempty" json:"msg_erro,omitempty" yaml:"msg_erro,omitempty"`
}

// RetornoSobreWebService was auto-generated from WSDL.
type RetornoSobreWebService struct {
	Data            *string `xml:"data,omitempty" json:"data,omitempty" yaml:"data,omitempty"`
	Hora            *string `xml:"hora,omitempty" json:"hora,omitempty" yaml:"hora,omitempty"`
	Cod_erro        *string `xml:"cod_erro,omitempty" json:"cod_erro,omitempty" yaml:"cod_erro,omitempty"`
	Msg_erro        *string `xml:"msg_erro,omitempty" json:"msg_erro,omitempty" yaml:"msg_erro,omitempty"`
	Versao          *string `xml:"versao,omitempty" json:"versao,omitempty" yaml:"versao,omitempty"`
	DataHomologacao *string `xml:"dataHomologacao,omitempty" json:"dataHomologacao,omitempty" yaml:"dataHomologacao,omitempty"`
	DataProducao    *string `xml:"dataProducao,omitempty" json:"dataProducao,omitempty" yaml:"dataProducao,omitempty"`
	Fase            *string `xml:"fase,omitempty" json:"fase,omitempty" yaml:"fase,omitempty"`
	Resumo          *string `xml:"resumo,omitempty" json:"resumo,omitempty" yaml:"resumo,omitempty"`
}

// RetornoValidacao was auto-generated from WSDL.
type RetornoValidacao struct {
	Cod_erro *int64  `xml:"cod_erro,omitempty" json:"cod_erro,omitempty" yaml:"cod_erro,omitempty"`
	Msg_erro *string `xml:"msg_erro,omitempty" json:"msg_erro,omitempty" yaml:"msg_erro,omitempty"`
}

// RevalidarPrazoAutorizacaoPostagem was auto-generated from WSDL.
type RevalidarPrazoAutorizacaoPostagem struct {
	CodAdministrativo *string `xml:"codAdministrativo,omitempty" json:"codAdministrativo,omitempty" yaml:"codAdministrativo,omitempty"`
	NumeroPedido      *string `xml:"numeroPedido,omitempty" json:"numeroPedido,omitempty" yaml:"numeroPedido,omitempty"`
	QtdeDias          *string `xml:"qtdeDias,omitempty" json:"qtdeDias,omitempty" yaml:"qtdeDias,omitempty"`
}

// RevalidarPrazoAutorizacaoPostagemResponse was auto-generated
// from WSDL.
type RevalidarPrazoAutorizacaoPostagemResponse struct {
	RevalidarPrazoAutorizacaoPostagem *RetornoRevalidarPrazo `xml:"revalidarPrazoAutorizacaoPostagem,omitempty" json:"revalidarPrazoAutorizacaoPostagem,omitempty" yaml:"revalidarPrazoAutorizacaoPostagem,omitempty"`
}

// SobreWebService was auto-generated from WSDL.
type SobreWebService struct {
}

// SobreWebServiceResponse was auto-generated from WSDL.
type SobreWebServiceResponse struct {
	SobreWebService *RetornoSobreWebService `xml:"sobreWebService,omitempty" json:"sobreWebService,omitempty" yaml:"sobreWebService,omitempty"`
}

// SolicitarPostagemReversa was auto-generated from WSDL.
type SolicitarPostagemReversa struct {
	XMLName             xml.Name             `xml:"ser:solicitarPostagemReversa"`
	CodAdministrativo   *string              `xml:"codAdministrativo,omitempty" json:"codAdministrativo,omitempty" yaml:"codAdministrativo,omitempty"`
	Codigo_servico      *string              `xml:"codigo_servico,omitempty" json:"codigo_servico,omitempty" yaml:"codigo_servico,omitempty"`
	Cartao              *string              `xml:"cartao,omitempty" json:"cartao,omitempty" yaml:"cartao,omitempty"`
	Destinatario        *DestinatarioReversa `xml:"destinatario,omitempty" json:"destinatario,omitempty" yaml:"destinatario,omitempty"`
	Coletas_solicitadas []*ColetaReversa     `xml:"coletas_solicitadas,omitempty" json:"coletas_solicitadas,omitempty" yaml:"coletas_solicitadas,omitempty"`
	XMLNsSer            string               `xml:"xmlns:ser,attr" json:"xmlnsSer,omitempty"`
}

// SolicitarPostagemReversaResponse was auto-generated from WSDL.
type SolicitarPostagemReversaResponse struct {
	XMLName                  xml.Name         `xml:"solicitarPostagemReversaResponse"`
	SolicitarPostagemReversa *RetornoPostagem `xml:"solicitarPostagemReversa,omitempty" json:"solicitarPostagemReversa,omitempty" yaml:"solicitarPostagemReversa,omitempty"`
}

// SolicitarPostagemSimultanea was auto-generated from WSDL.
type SolicitarPostagemSimultanea struct {
	CodAdministrativo   *string              `xml:"codAdministrativo,omitempty" json:"codAdministrativo,omitempty" yaml:"codAdministrativo,omitempty"`
	Codigo_servico      *string              `xml:"codigo_servico,omitempty" json:"codigo_servico,omitempty" yaml:"codigo_servico,omitempty"`
	Cartao              *string              `xml:"cartao,omitempty" json:"cartao,omitempty" yaml:"cartao,omitempty"`
	Destinatario        *DestinatarioReversa `xml:"destinatario,omitempty" json:"destinatario,omitempty" yaml:"destinatario,omitempty"`
	Coletas_solicitadas []*ColetaSimultanea  `xml:"coletas_solicitadas,omitempty" json:"coletas_solicitadas,omitempty" yaml:"coletas_solicitadas,omitempty"`
}

// SolicitarPostagemSimultaneaResponse was auto-generated from
// WSDL.
type SolicitarPostagemSimultaneaResponse struct {
	SolicitarPostagemSimultanea *RetornoPostagem `xml:"solicitarPostagemSimultanea,omitempty" json:"solicitarPostagemSimultanea,omitempty" yaml:"solicitarPostagemSimultanea,omitempty"`
}

// SolicitarRange was auto-generated from WSDL.
type SolicitarRange struct {
	CodAdministrativo *string `xml:"codAdministrativo,omitempty" json:"codAdministrativo,omitempty" yaml:"codAdministrativo,omitempty"`
	Tipo              *string `xml:"tipo,omitempty" json:"tipo,omitempty" yaml:"tipo,omitempty"`
	Servico           *string `xml:"servico,omitempty" json:"servico,omitempty" yaml:"servico,omitempty"`
	Quantidade        *string `xml:"quantidade,omitempty" json:"quantidade,omitempty" yaml:"quantidade,omitempty"`
}

// SolicitarRangeResponse was auto-generated from WSDL.
type SolicitarRangeResponse struct {
	SolicitarRange *RetornoFaixaNumerica `xml:"solicitarRange,omitempty" json:"solicitarRange,omitempty" yaml:"solicitarRange,omitempty"`
}

// ValidarPostagemReversa was auto-generated from WSDL.
type ValidarPostagemReversa struct {
	CodAdministrativo *string        `xml:"codAdministrativo,omitempty" json:"codAdministrativo,omitempty" yaml:"codAdministrativo,omitempty"`
	Codigo_servico    *string        `xml:"codigo_servico,omitempty" json:"codigo_servico,omitempty" yaml:"codigo_servico,omitempty"`
	Cartao            *string        `xml:"cartao,omitempty" json:"cartao,omitempty" yaml:"cartao,omitempty"`
	Cep_destinatario  *string        `xml:"cep_destinatario,omitempty" json:"cep_destinatario,omitempty" yaml:"cep_destinatario,omitempty"`
	Coleta            *ColetaReversa `xml:"coleta,omitempty" json:"coleta,omitempty" yaml:"coleta,omitempty"`
}

// ValidarPostagemReversaResponse was auto-generated from WSDL.
type ValidarPostagemReversaResponse struct {
	ValidarPostagemReversa *RetornoValidacao `xml:"validarPostagemReversa,omitempty" json:"validarPostagemReversa,omitempty" yaml:"validarPostagemReversa,omitempty"`
}

// ValidarPostagemSimultanea was auto-generated from WSDL.
type ValidarPostagemSimultanea struct {
	CodAdministrativo *string           `xml:"codAdministrativo,omitempty" json:"codAdministrativo,omitempty" yaml:"codAdministrativo,omitempty"`
	Codigo_servico    *string           `xml:"codigo_servico,omitempty" json:"codigo_servico,omitempty" yaml:"codigo_servico,omitempty"`
	Cartao            *string           `xml:"cartao,omitempty" json:"cartao,omitempty" yaml:"cartao,omitempty"`
	Cep_destinatario  *string           `xml:"cep_destinatario,omitempty" json:"cep_destinatario,omitempty" yaml:"cep_destinatario,omitempty"`
	Coleta            *ColetaSimultanea `xml:"coleta,omitempty" json:"coleta,omitempty" yaml:"coleta,omitempty"`
}

// ValidarPostagemSimultaneaResponse was auto-generated from WSDL.
type ValidarPostagemSimultaneaResponse struct {
	ValidarPostagemSimultanea *RetornoValidacao `xml:"validarPostagemSimultanea,omitempty" json:"validarPostagemSimultanea,omitempty" yaml:"validarPostagemSimultanea,omitempty"`
}

// Operation wrapper for AcompanharPedido.
// OperationAcompanharPedido was auto-generated from WSDL.
type OperationAcompanharPedido struct {
	AcompanharPedido *AcompanharPedido `xml:"acompanharPedido,omitempty" json:"acompanharPedido,omitempty" yaml:"acompanharPedido,omitempty"`
}

// Operation wrapper for AcompanharPedido.
// OperationAcompanharPedidoResponse was auto-generated from WSDL.
type OperationAcompanharPedidoResponse struct {
	AcompanharPedidoResponse *AcompanharPedidoResponse `xml:"acompanharPedidoResponse,omitempty" json:"acompanharPedidoResponse,omitempty" yaml:"acompanharPedidoResponse,omitempty"`
}

// Operation wrapper for AcompanharPedidoPorData.
// OperationAcompanharPedidoPorData was auto-generated from WSDL.
type OperationAcompanharPedidoPorData struct {
	AcompanharPedidoPorData *AcompanharPedidoPorData `xml:"acompanharPedidoPorData,omitempty" json:"acompanharPedidoPorData,omitempty" yaml:"acompanharPedidoPorData,omitempty"`
}

// Operation wrapper for AcompanharPedidoPorData.
// OperationAcompanharPedidoPorDataResponse was auto-generated
// from WSDL.
type OperationAcompanharPedidoPorDataResponse struct {
	AcompanharPedidoPorDataResponse *AcompanharPedidoPorDataResponse `xml:"acompanharPedidoPorDataResponse,omitempty" json:"acompanharPedidoPorDataResponse,omitempty" yaml:"acompanharPedidoPorDataResponse,omitempty"`
}

// Operation wrapper for CalcularDigitoVerificador.
// OperationCalcularDigitoVerificador was auto-generated from WSDL.
type OperationCalcularDigitoVerificador struct {
	CalcularDigitoVerificador *CalcularDigitoVerificador `xml:"calcularDigitoVerificador,omitempty" json:"calcularDigitoVerificador,omitempty" yaml:"calcularDigitoVerificador,omitempty"`
}

// Operation wrapper for CalcularDigitoVerificador.
// OperationCalcularDigitoVerificadorResponse was auto-generated
// from WSDL.
type OperationCalcularDigitoVerificadorResponse struct {
	CalcularDigitoVerificadorResponse *CalcularDigitoVerificadorResponse `xml:"calcularDigitoVerificadorResponse,omitempty" json:"calcularDigitoVerificadorResponse,omitempty" yaml:"calcularDigitoVerificadorResponse,omitempty"`
}

// Operation wrapper for CancelarPedido.
// OperationCancelarPedido was auto-generated from WSDL.
type OperationCancelarPedido struct {
	CancelarPedido *CancelarPedido `xml:"cancelarPedido,omitempty" json:"cancelarPedido,omitempty" yaml:"cancelarPedido,omitempty"`
}

// Operation wrapper for CancelarPedido.
// OperationCancelarPedidoResponse was auto-generated from WSDL.
type OperationCancelarPedidoResponse struct {
	CancelarPedidoResponse *CancelarPedidoResponse `xml:"cancelarPedidoResponse,omitempty" json:"cancelarPedidoResponse,omitempty" yaml:"cancelarPedidoResponse,omitempty"`
}

// Operation wrapper for ConsultarResumoColeta.
// OperationConsultarResumoColeta was auto-generated from WSDL.
type OperationConsultarResumoColeta struct {
	ConsultarResumoColeta *ConsultarResumoColeta `xml:"consultarResumoColeta,omitempty" json:"consultarResumoColeta,omitempty" yaml:"consultarResumoColeta,omitempty"`
}

// Operation wrapper for ConsultarResumoColeta.
// OperationConsultarResumoColetaResponse was auto-generated from
// WSDL.
type OperationConsultarResumoColetaResponse struct {
	ConsultarResumoColetaResponse *ConsultarResumoColetaResponse `xml:"consultarResumoColetaResponse,omitempty" json:"consultarResumoColetaResponse,omitempty" yaml:"consultarResumoColetaResponse,omitempty"`
}

// Operation wrapper for RevalidarPrazoAutorizacaoPostagem.
// OperationRevalidarPrazoAutorizacaoPostagem was auto-generated
// from WSDL.
type OperationRevalidarPrazoAutorizacaoPostagem struct {
	RevalidarPrazoAutorizacaoPostagem *RevalidarPrazoAutorizacaoPostagem `xml:"revalidarPrazoAutorizacaoPostagem,omitempty" json:"revalidarPrazoAutorizacaoPostagem,omitempty" yaml:"revalidarPrazoAutorizacaoPostagem,omitempty"`
}

// Operation wrapper for RevalidarPrazoAutorizacaoPostagem.
// OperationRevalidarPrazoAutorizacaoPostagemResponse was auto-generated
// from WSDL.
type OperationRevalidarPrazoAutorizacaoPostagemResponse struct {
	RevalidarPrazoAutorizacaoPostagemResponse *RevalidarPrazoAutorizacaoPostagemResponse `xml:"revalidarPrazoAutorizacaoPostagemResponse,omitempty" json:"revalidarPrazoAutorizacaoPostagemResponse,omitempty" yaml:"revalidarPrazoAutorizacaoPostagemResponse,omitempty"`
}

// Operation wrapper for SobreWebService.
// OperationSobreWebService was auto-generated from WSDL.
type OperationSobreWebService struct {
	SobreWebService *SobreWebService `xml:"sobreWebService,omitempty" json:"sobreWebService,omitempty" yaml:"sobreWebService,omitempty"`
}

// Operation wrapper for SobreWebService.
// OperationSobreWebServiceResponse was auto-generated from WSDL.
type OperationSobreWebServiceResponse struct {
	SobreWebServiceResponse *SobreWebServiceResponse `xml:"sobreWebServiceResponse,omitempty" json:"sobreWebServiceResponse,omitempty" yaml:"sobreWebServiceResponse,omitempty"`
}

// Operation wrapper for SolicitarPostagemReversa.
// OperationSolicitarPostagemReversa was auto-generated from WSDL.
type OperationSolicitarPostagemReversa struct {
	SolicitarPostagemReversa *SolicitarPostagemReversa `xml:"ser:solicitarPostagemReversa,omitempty" json:"solicitarPostagemReversa,omitempty" yaml:"solicitarPostagemReversa,omitempty"`
}

// Operation wrapper for SolicitarPostagemReversa.
// OperationSolicitarPostagemReversaResponse was auto-generated
// from WSDL.
type OperationSolicitarPostagemReversaResponse struct {
	SolicitarPostagemReversaResponse *SolicitarPostagemReversaResponse `xml:"solicitarPostagemReversaResponse,omitempty" json:"solicitarPostagemReversaResponse,omitempty" yaml:"solicitarPostagemReversaResponse,omitempty"`
}

// Operation wrapper for SolicitarPostagemSimultanea.
// OperationSolicitarPostagemSimultanea was auto-generated from
// WSDL.
type OperationSolicitarPostagemSimultanea struct {
	SolicitarPostagemSimultanea *SolicitarPostagemSimultanea `xml:"solicitarPostagemSimultanea,omitempty" json:"solicitarPostagemSimultanea,omitempty" yaml:"solicitarPostagemSimultanea,omitempty"`
}

// Operation wrapper for SolicitarPostagemSimultanea.
// OperationSolicitarPostagemSimultaneaResponse was auto-generated
// from WSDL.
type OperationSolicitarPostagemSimultaneaResponse struct {
	SolicitarPostagemSimultaneaResponse *SolicitarPostagemSimultaneaResponse `xml:"solicitarPostagemSimultaneaResponse,omitempty" json:"solicitarPostagemSimultaneaResponse,omitempty" yaml:"solicitarPostagemSimultaneaResponse,omitempty"`
}

// Operation wrapper for SolicitarRange.
// OperationSolicitarRange was auto-generated from WSDL.
type OperationSolicitarRange struct {
	SolicitarRange *SolicitarRange `xml:"solicitarRange,omitempty" json:"solicitarRange,omitempty" yaml:"solicitarRange,omitempty"`
}

// Operation wrapper for SolicitarRange.
// OperationSolicitarRangeResponse was auto-generated from WSDL.
type OperationSolicitarRangeResponse struct {
	SolicitarRangeResponse *SolicitarRangeResponse `xml:"solicitarRangeResponse,omitempty" json:"solicitarRangeResponse,omitempty" yaml:"solicitarRangeResponse,omitempty"`
}

// Operation wrapper for ValidarPostagemReversa.
// OperationValidarPostagemReversa was auto-generated from WSDL.
type OperationValidarPostagemReversa struct {
	ValidarPostagemReversa *ValidarPostagemReversa `xml:"validarPostagemReversa,omitempty" json:"validarPostagemReversa,omitempty" yaml:"validarPostagemReversa,omitempty"`
}

// Operation wrapper for ValidarPostagemReversa.
// OperationValidarPostagemReversaResponse was auto-generated from
// WSDL.
type OperationValidarPostagemReversaResponse struct {
	ValidarPostagemReversaResponse *ValidarPostagemReversaResponse `xml:"validarPostagemReversaResponse,omitempty" json:"validarPostagemReversaResponse,omitempty" yaml:"validarPostagemReversaResponse,omitempty"`
}

// Operation wrapper for ValidarPostagemSimultanea.
// OperationValidarPostagemSimultanea was auto-generated from WSDL.
type OperationValidarPostagemSimultanea struct {
	ValidarPostagemSimultanea *ValidarPostagemSimultanea `xml:"validarPostagemSimultanea,omitempty" json:"validarPostagemSimultanea,omitempty" yaml:"validarPostagemSimultanea,omitempty"`
}

// Operation wrapper for ValidarPostagemSimultanea.
// OperationValidarPostagemSimultaneaResponse was auto-generated
// from WSDL.
type OperationValidarPostagemSimultaneaResponse struct {
	ValidarPostagemSimultaneaResponse *ValidarPostagemSimultaneaResponse `xml:"validarPostagemSimultaneaResponse,omitempty" json:"validarPostagemSimultaneaResponse,omitempty" yaml:"validarPostagemSimultaneaResponse,omitempty"`
}

// logisticaReversaWS implements the LogisticaReversaWS interface.
type logisticaReversaWS struct {
	cli *soap.Client
}

// AcompanharPedido was auto-generated from WSDL.
func (p *logisticaReversaWS) AcompanharPedido(AcompanharPedido *AcompanharPedido) (*AcompanharPedidoResponse, error) {
	α := struct {
		OperationAcompanharPedido `xml:"tns:acompanharPedido"`
	}{
		OperationAcompanharPedido{
			AcompanharPedido,
		},
	}

	γ := struct {
		OperationAcompanharPedidoResponse `xml:"acompanharPedidoResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("AcompanharPedido", α, &γ); err != nil {
		return nil, err
	}
	return γ.AcompanharPedidoResponse, nil
}

// AcompanharPedidoPorData was auto-generated from WSDL.
func (p *logisticaReversaWS) AcompanharPedidoPorData(AcompanharPedidoPorData *AcompanharPedidoPorData) (*AcompanharPedidoPorDataResponse, error) {
	α := struct {
		OperationAcompanharPedidoPorData `xml:"tns:acompanharPedidoPorData"`
	}{
		OperationAcompanharPedidoPorData{
			AcompanharPedidoPorData,
		},
	}

	γ := struct {
		OperationAcompanharPedidoPorDataResponse `xml:"acompanharPedidoPorDataResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("AcompanharPedidoPorData", α, &γ); err != nil {
		return nil, err
	}
	return γ.AcompanharPedidoPorDataResponse, nil
}

// CalcularDigitoVerificador was auto-generated from WSDL.
func (p *logisticaReversaWS) CalcularDigitoVerificador(CalcularDigitoVerificador *CalcularDigitoVerificador) (*CalcularDigitoVerificadorResponse, error) {
	α := struct {
		OperationCalcularDigitoVerificador `xml:"tns:calcularDigitoVerificador"`
	}{
		OperationCalcularDigitoVerificador{
			CalcularDigitoVerificador,
		},
	}

	γ := struct {
		OperationCalcularDigitoVerificadorResponse `xml:"calcularDigitoVerificadorResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("CalcularDigitoVerificador", α, &γ); err != nil {
		return nil, err
	}
	return γ.CalcularDigitoVerificadorResponse, nil
}

// CancelarPedido was auto-generated from WSDL.
func (p *logisticaReversaWS) CancelarPedido(CancelarPedido *CancelarPedido) (*CancelarPedidoResponse, error) {
	α := struct {
		OperationCancelarPedido `xml:"tns:cancelarPedido"`
	}{
		OperationCancelarPedido{
			CancelarPedido,
		},
	}

	γ := struct {
		OperationCancelarPedidoResponse `xml:"cancelarPedidoResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("CancelarPedido", α, &γ); err != nil {
		return nil, err
	}
	return γ.CancelarPedidoResponse, nil
}

// ConsultarResumoColeta was auto-generated from WSDL.
func (p *logisticaReversaWS) ConsultarResumoColeta(ConsultarResumoColeta *ConsultarResumoColeta) (*ConsultarResumoColetaResponse, error) {
	α := struct {
		OperationConsultarResumoColeta `xml:"tns:consultarResumoColeta"`
	}{
		OperationConsultarResumoColeta{
			ConsultarResumoColeta,
		},
	}

	γ := struct {
		OperationConsultarResumoColetaResponse `xml:"consultarResumoColetaResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("ConsultarResumoColeta", α, &γ); err != nil {
		return nil, err
	}
	return γ.ConsultarResumoColetaResponse, nil
}

// RevalidarPrazoAutorizacaoPostagem was auto-generated from WSDL.
func (p *logisticaReversaWS) RevalidarPrazoAutorizacaoPostagem(RevalidarPrazoAutorizacaoPostagem *RevalidarPrazoAutorizacaoPostagem) (*RevalidarPrazoAutorizacaoPostagemResponse, error) {
	α := struct {
		OperationRevalidarPrazoAutorizacaoPostagem `xml:"tns:revalidarPrazoAutorizacaoPostagem"`
	}{
		OperationRevalidarPrazoAutorizacaoPostagem{
			RevalidarPrazoAutorizacaoPostagem,
		},
	}

	γ := struct {
		OperationRevalidarPrazoAutorizacaoPostagemResponse `xml:"revalidarPrazoAutorizacaoPostagemResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("RevalidarPrazoAutorizacaoPostagem", α, &γ); err != nil {
		return nil, err
	}
	return γ.RevalidarPrazoAutorizacaoPostagemResponse, nil
}

// SobreWebService was auto-generated from WSDL.
func (p *logisticaReversaWS) SobreWebService(SobreWebService *SobreWebService) (*SobreWebServiceResponse, error) {
	α := struct {
		OperationSobreWebService `xml:"tns:sobreWebService"`
	}{
		OperationSobreWebService{
			SobreWebService,
		},
	}

	γ := struct {
		OperationSobreWebServiceResponse `xml:"sobreWebServiceResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("SobreWebService", α, &γ); err != nil {
		return nil, err
	}
	return γ.SobreWebServiceResponse, nil
}

// SolicitarPostagemReversa was auto-generated from WSDL.
func (p *logisticaReversaWS) SolicitarPostagemReversa(SolicitarPostagemReversa *SolicitarPostagemReversa) (*SolicitarPostagemReversaResponse, error) {
	α := struct {
		OperationSolicitarPostagemReversa `xml:"ser:solicitarPostagemReversa"`
	}{
		OperationSolicitarPostagemReversa{
			SolicitarPostagemReversa,
		},
	}

	γ := struct {
		OperationSolicitarPostagemReversaResponse `xml:"solicitarPostagemReversaResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("SolicitarPostagemReversa", α, &γ); err != nil {
		return nil, err
	}
	return γ.SolicitarPostagemReversaResponse, nil
}

// SolicitarPostagemSimultanea was auto-generated from WSDL.
func (p *logisticaReversaWS) SolicitarPostagemSimultanea(SolicitarPostagemSimultanea *SolicitarPostagemSimultanea) (*SolicitarPostagemSimultaneaResponse, error) {
	α := struct {
		OperationSolicitarPostagemSimultanea `xml:"tns:solicitarPostagemSimultanea"`
	}{
		OperationSolicitarPostagemSimultanea{
			SolicitarPostagemSimultanea,
		},
	}

	γ := struct {
		OperationSolicitarPostagemSimultaneaResponse `xml:"solicitarPostagemSimultaneaResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("SolicitarPostagemSimultanea", α, &γ); err != nil {
		return nil, err
	}
	return γ.SolicitarPostagemSimultaneaResponse, nil
}

// SolicitarRange was auto-generated from WSDL.
func (p *logisticaReversaWS) SolicitarRange(SolicitarRange *SolicitarRange) (*SolicitarRangeResponse, error) {
	α := struct {
		OperationSolicitarRange `xml:"tns:solicitarRange"`
	}{
		OperationSolicitarRange{
			SolicitarRange,
		},
	}

	γ := struct {
		OperationSolicitarRangeResponse `xml:"solicitarRangeResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("SolicitarRange", α, &γ); err != nil {
		return nil, err
	}
	return γ.SolicitarRangeResponse, nil
}

// ValidarPostagemReversa was auto-generated from WSDL.
func (p *logisticaReversaWS) ValidarPostagemReversa(ValidarPostagemReversa *ValidarPostagemReversa) (*ValidarPostagemReversaResponse, error) {
	α := struct {
		OperationValidarPostagemReversa `xml:"tns:validarPostagemReversa"`
	}{
		OperationValidarPostagemReversa{
			ValidarPostagemReversa,
		},
	}

	γ := struct {
		OperationValidarPostagemReversaResponse `xml:"validarPostagemReversaResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("ValidarPostagemReversa", α, &γ); err != nil {
		return nil, err
	}
	return γ.ValidarPostagemReversaResponse, nil
}

// ValidarPostagemSimultanea was auto-generated from WSDL.
func (p *logisticaReversaWS) ValidarPostagemSimultanea(ValidarPostagemSimultanea *ValidarPostagemSimultanea) (*ValidarPostagemSimultaneaResponse, error) {
	α := struct {
		OperationValidarPostagemSimultanea `xml:"tns:validarPostagemSimultanea"`
	}{
		OperationValidarPostagemSimultanea{
			ValidarPostagemSimultanea,
		},
	}

	γ := struct {
		OperationValidarPostagemSimultaneaResponse `xml:"validarPostagemSimultaneaResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("ValidarPostagemSimultanea", α, &γ); err != nil {
		return nil, err
	}
	return γ.ValidarPostagemSimultaneaResponse, nil
}
