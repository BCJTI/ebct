package ebct

const soapCorreios = "logisticaReversaWS/logisticaReversaService/logisticaReversaWS"
const envSerNamespaceReversa = "http://service.logisticareversa.correios.com.br/"

/*@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
*/

// AcompanharPedido was auto-generated from WSDL.
func (c *Client) PostAcompanharPedido(AcompanharPedido *AcompanharPedido) (*AcompanharPedidoResponse, error) {
	AcompanharPedido.XMLNsSer = envSerNamespaceReversa

	sendRequest := struct {
		OperationAcompanharPedido `xml:"tns:acompanharPedido"`
	}{
		OperationAcompanharPedido{
			AcompanharPedido,
		},
	}

	returnRequest := struct {
		OperationAcompanharPedidoResponse `xml:"acompanharPedidoResponse"`
	}{}
	c.DsNamespace = envSerNamespaceReversa
	c.FlSoap = true

	err := c.Soap(soapCorreios, sendRequest, nil, &returnRequest)
	return returnRequest.AcompanharPedidoResponse, err

}

// AcompanharPedidoPorData was auto-generated from WSDL.
func (c *Client) PostAcompanharPedidoPorData(AcompanharPedidoPorData *AcompanharPedidoPorData) (*AcompanharPedidoPorDataResponse, error) {
	AcompanharPedidoPorData.XMLNsSer = envSerNamespaceReversa

	sendRequest := struct {
		OperationAcompanharPedidoPorData `xml:"ser:acompanharPedidoPorData"`
	}{
		OperationAcompanharPedidoPorData{
			AcompanharPedidoPorData,
		},
	}

	returnRequest := struct {
		OperationAcompanharPedidoPorDataResponse `xml:"acompanharPedidoPorDataResponse"`
	}{}
	c.DsNamespace = envSerNamespaceReversa
	c.FlSoap = true

	err := c.Soap(soapCorreios, sendRequest, nil, &returnRequest)
	return returnRequest.AcompanharPedidoPorDataResponse, err

}

// CalcularDigitoVerificador was auto-generated from WSDL.
func (c *Client) PostCalcularDigitoVerificador(CalcularDigitoVerificador *CalcularDigitoVerificador) (*CalcularDigitoVerificadorResponse, error) {
	CalcularDigitoVerificador.XMLNsSer = envSerNamespaceReversa

	sendRequest := struct {
		OperationCalcularDigitoVerificador `xml:"tns:calcularDigitoVerificador"`
	}{
		OperationCalcularDigitoVerificador{
			CalcularDigitoVerificador,
		},
	}

	returnRequest := struct {
		OperationCalcularDigitoVerificadorResponse `xml:"calcularDigitoVerificadorResponse"`
	}{}
	c.DsNamespace = envSerNamespaceReversa
	c.FlSoap = true

	err := c.Soap(soapCorreios, sendRequest, nil, &returnRequest)
	return returnRequest.CalcularDigitoVerificadorResponse, err

}

// CancelarPedido was auto-generated from WSDL.
func (c *Client) PostCancelarPedido(CancelarPedido *CancelarPedido) (*CancelarPedidoResponse, error) {
	CancelarPedido.XMLNsSer = envSerNamespaceReversa

	sendRequest := struct {
		OperationCancelarPedido `xml:"ser:cancelarPedido"`
	}{
		OperationCancelarPedido{
			CancelarPedido,
		},
	}

	returnRequest := struct {
		OperationCancelarPedidoResponse `xml:"cancelarPedidoResponse"`
	}{}
	c.DsNamespace = envSerNamespaceReversa
	c.FlSoap = true

	err := c.Soap(soapCorreios, sendRequest, nil, &returnRequest)
	return returnRequest.CancelarPedidoResponse, err

}

// ConsultarResumoColeta was auto-generated from WSDL.
func (c *Client) PostConsultarResumoColeta(ConsultarResumoColeta *ConsultarResumoColeta) (*ConsultarResumoColetaResponse, error) {
	ConsultarResumoColeta.XMLNsSer = envSerNamespaceReversa

	sendRequest := struct {
		OperationConsultarResumoColeta `xml:"ser:consultarResumoColeta"`
	}{
		OperationConsultarResumoColeta{
			ConsultarResumoColeta,
		},
	}

	returnRequest := struct {
		OperationConsultarResumoColetaResponse `xml:"consultarResumoColetaResponse"`
	}{}
	c.DsNamespace = envSerNamespaceReversa
	c.FlSoap = true

	err := c.Soap(soapCorreios, sendRequest, nil, &returnRequest)
	return returnRequest.ConsultarResumoColetaResponse, err

}

// RevalidarPrazoAutorizacaoPostagem was auto-generated from WSDL.
func (c *Client) PostRevalidarPrazoAutorizacaoPostagem(RevalidarPrazoAutorizacaoPostagem *RevalidarPrazoAutorizacaoPostagem) (*RevalidarPrazoAutorizacaoPostagemResponse, error) {
	RevalidarPrazoAutorizacaoPostagem.XMLNsSer = envSerNamespaceReversa

	sendRequest := struct {
		OperationRevalidarPrazoAutorizacaoPostagem `xml:"ser:revalidarPrazoAutorizacaoPostagem"`
	}{
		OperationRevalidarPrazoAutorizacaoPostagem{
			RevalidarPrazoAutorizacaoPostagem,
		},
	}

	returnRequest := struct {
		OperationRevalidarPrazoAutorizacaoPostagemResponse `xml:"revalidarPrazoAutorizacaoPostagemResponse"`
	}{}
	c.DsNamespace = envSerNamespaceReversa
	c.FlSoap = true

	err := c.Soap(soapCorreios, sendRequest, nil, &returnRequest)
	return returnRequest.RevalidarPrazoAutorizacaoPostagemResponse, err

}

// SobreWebService was auto-generated from WSDL.
func (c *Client) PostSobreWebService(SobreWebService *SobreWebService) (*SobreWebServiceResponse, error) {
	sendRequest := struct {
		OperationSobreWebService `xml:"tns:sobreWebService"`
	}{
		OperationSobreWebService{
			SobreWebService,
		},
	}

	returnRequest := struct {
		OperationSobreWebServiceResponse `xml:"sobreWebServiceResponse"`
	}{}
	c.DsNamespace = envSerNamespaceReversa
	c.FlSoap = true

	err := c.Soap(soapCorreios, sendRequest, nil, &returnRequest)
	return returnRequest.SobreWebServiceResponse, err

}

func (c *Client) PostSolicitarPostagemReversa(SolicitarPostagemReversa *SolicitarPostagemReversa) (*SolicitarPostagemReversaResponse, error) {
	SolicitarPostagemReversa.XMLNsSer = envSerNamespaceReversa

	sendRequest := struct {
		OperationSolicitarPostagemReversa `xml:"ser:solicitarPostagemReversa"`
	}{
		OperationSolicitarPostagemReversa{
			SolicitarPostagemReversa,
		},
	}

	returnRequest := struct {
		OperationSolicitarPostagemReversaResponse `xml:"solicitarPostagemReversaResponse"`
	}{}

	c.DsNamespace = envSerNamespaceReversa
	c.FlSoap = true

	err := c.Soap(soapCorreios, sendRequest, nil, &returnRequest)
	return returnRequest.SolicitarPostagemReversaResponse, err
}

// SolicitarPostagemSimultanea was auto-generated from WSDL.
func (c *Client) PostSolicitarPostagemSimultanea(SolicitarPostagemSimultanea *SolicitarPostagemSimultanea) (*SolicitarPostagemSimultaneaResponse, error) {
	SolicitarPostagemSimultanea.XMLNsSer = envSerNamespaceReversa

	sendRequest := struct {
		OperationSolicitarPostagemSimultanea `xml:"ser:solicitarPostagemSimultanea"`
	}{
		OperationSolicitarPostagemSimultanea{
			SolicitarPostagemSimultanea,
		},
	}

	returnRequest := struct {
		OperationSolicitarPostagemSimultaneaResponse `xml:"solicitarPostagemSimultaneaResponse"`
	}{}
	c.DsNamespace = envSerNamespaceReversa
	c.FlSoap = true

	err := c.Soap(soapCorreios, sendRequest, nil, &returnRequest)
	return returnRequest.SolicitarPostagemSimultaneaResponse, err

}

// SolicitarRange was auto-generated from WSDL.
func (c *Client) PostSolicitarRange(SolicitarRange *SolicitarRange) (*SolicitarRangeResponse, error) {
	SolicitarRange.XMLNsSer = envSerNamespaceReversa

	sendRequest := struct {
		OperationSolicitarRange `xml:"ser:solicitarRange"`
	}{
		OperationSolicitarRange{
			SolicitarRange,
		},
	}

	returnRequest := struct {
		OperationSolicitarRangeResponse `xml:"solicitarRangeResponse"`
	}{}
	c.DsNamespace = envSerNamespaceReversa
	c.FlSoap = true

	err := c.Soap(soapCorreios, sendRequest, nil, &returnRequest)
	return returnRequest.SolicitarRangeResponse, err

}

// ValidarPostagemReversa was auto-generated from WSDL.
func (c *Client) PostValidarPostagemReversa(ValidarPostagemReversa *ValidarPostagemReversa) (*ValidarPostagemReversaResponse, error) {
	ValidarPostagemReversa.XMLNsSer = envSerNamespaceReversa

	sendRequest := struct {
		OperationValidarPostagemReversa `xml:"ser:validarPostagemReversa"`
	}{
		OperationValidarPostagemReversa{
			ValidarPostagemReversa,
		},
	}

	returnRequest := struct {
		OperationValidarPostagemReversaResponse `xml:"validarPostagemReversaResponse"`
	}{}
	c.DsNamespace = envSerNamespaceReversa
	c.FlSoap = true

	err := c.Soap(soapCorreios, sendRequest, nil, &returnRequest)
	return returnRequest.ValidarPostagemReversaResponse, err

}

// ValidarPostagemSimultanea was auto-generated from WSDL.
func (c *Client) PostValidarPostagemSimultanea(ValidarPostagemSimultanea *ValidarPostagemSimultanea) (*ValidarPostagemSimultaneaResponse, error) {
	ValidarPostagemSimultanea.XMLNsSer = envSerNamespaceReversa

	sendRequest := struct {
		OperationValidarPostagemSimultanea `xml:"ser:validarPostagemSimultanea"`
	}{
		OperationValidarPostagemSimultanea{
			ValidarPostagemSimultanea,
		},
	}

	returnRequest := struct {
		OperationValidarPostagemSimultaneaResponse `xml:"validarPostagemSimultaneaResponse"`
	}{}
	c.DsNamespace = envSerNamespaceReversa
	c.FlSoap = true

	err := c.Soap(soapCorreios, sendRequest, nil, &returnRequest)
	return returnRequest.ValidarPostagemSimultaneaResponse, err

}
