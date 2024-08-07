package ebct

const soapCorreios = "logisticaReversaWS/logisticaReversaService/logisticaReversaWS"
const envNamespaceReversa = "http://service.logisticareversa.correios.com.br/"

/*@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
*/

// AcompanharPedido was auto-generated from WSDL.
func (c *Client) PostAcompanharPedido(AcompanharPedido *AcompanharPedido) (*AcompanharPedidoResponse, error) {
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
	c.DsNamespace = envNamespaceReversa
	c.FlSoap = true

	err := c.Post(soapCorreios, sendRequest, nil, &returnRequest, "XML", "AcompanharPedido")
	return returnRequest.AcompanharPedidoResponse, err

}

// AcompanharPedidoPorData was auto-generated from WSDL.
func (c *Client) PostAcompanharPedidoPorData(AcompanharPedidoPorData *AcompanharPedidoPorData) (*AcompanharPedidoPorDataResponse, error) {
	sendRequest := struct {
		OperationAcompanharPedidoPorData `xml:"tns:acompanharPedidoPorData"`
	}{
		OperationAcompanharPedidoPorData{
			AcompanharPedidoPorData,
		},
	}

	returnRequest := struct {
		OperationAcompanharPedidoPorDataResponse `xml:"acompanharPedidoPorDataResponse"`
	}{}
	c.DsNamespace = envNamespaceReversa
	c.FlSoap = true

	err := c.Post(soapCorreios, sendRequest, nil, &returnRequest, "XML", "AcompanharPedidoPorData")
	return returnRequest.AcompanharPedidoPorDataResponse, err

}

// CalcularDigitoVerificador was auto-generated from WSDL.
func (c *Client) PostCalcularDigitoVerificador(CalcularDigitoVerificador *CalcularDigitoVerificador) (*CalcularDigitoVerificadorResponse, error) {
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
	c.DsNamespace = envNamespaceReversa
	c.FlSoap = true

	err := c.Post(soapCorreios, sendRequest, nil, &returnRequest, "XML", "CalcularDigitoVerificador")
	return returnRequest.CalcularDigitoVerificadorResponse, err

}

// CancelarPedido was auto-generated from WSDL.
func (c *Client) PostCancelarPedido(CancelarPedido *CancelarPedido) (*CancelarPedidoResponse, error) {
	sendRequest := struct {
		OperationCancelarPedido `xml:"tns:cancelarPedido"`
	}{
		OperationCancelarPedido{
			CancelarPedido,
		},
	}

	returnRequest := struct {
		OperationCancelarPedidoResponse `xml:"cancelarPedidoResponse"`
	}{}
	c.DsNamespace = envNamespaceReversa
	c.FlSoap = true

	err := c.Post(soapCorreios, sendRequest, nil, &returnRequest, "XML", "CancelarPedido")
	return returnRequest.CancelarPedidoResponse, err

}

// ConsultarResumoColeta was auto-generated from WSDL.
func (c *Client) PostConsultarResumoColeta(ConsultarResumoColeta *ConsultarResumoColeta) (*ConsultarResumoColetaResponse, error) {
	sendRequest := struct {
		OperationConsultarResumoColeta `xml:"tns:consultarResumoColeta"`
	}{
		OperationConsultarResumoColeta{
			ConsultarResumoColeta,
		},
	}

	returnRequest := struct {
		OperationConsultarResumoColetaResponse `xml:"consultarResumoColetaResponse"`
	}{}
	c.DsNamespace = envNamespaceReversa
	c.FlSoap = true

	err := c.Post(soapCorreios, sendRequest, nil, &returnRequest, "XML", "ConsultarResumoColeta")
	return returnRequest.ConsultarResumoColetaResponse, err

}

// RevalidarPrazoAutorizacaoPostagem was auto-generated from WSDL.
func (c *Client) PostRevalidarPrazoAutorizacaoPostagem(RevalidarPrazoAutorizacaoPostagem *RevalidarPrazoAutorizacaoPostagem) (*RevalidarPrazoAutorizacaoPostagemResponse, error) {
	sendRequest := struct {
		OperationRevalidarPrazoAutorizacaoPostagem `xml:"tns:revalidarPrazoAutorizacaoPostagem"`
	}{
		OperationRevalidarPrazoAutorizacaoPostagem{
			RevalidarPrazoAutorizacaoPostagem,
		},
	}

	returnRequest := struct {
		OperationRevalidarPrazoAutorizacaoPostagemResponse `xml:"revalidarPrazoAutorizacaoPostagemResponse"`
	}{}
	c.DsNamespace = envNamespaceReversa
	c.FlSoap = true

	err := c.Post(soapCorreios, sendRequest, nil, &returnRequest, "XML", "RevalidarPrazoAutorizacaoPostagem")
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
	c.DsNamespace = envNamespaceReversa
	c.FlSoap = true

	err := c.Post(soapCorreios, sendRequest, nil, &returnRequest, "XML", "SobreWebService")
	return returnRequest.SobreWebServiceResponse, err

}

func (c *Client) PostSolicitarPostagemReversa(SolicitarPostagemReversa *SolicitarPostagemReversa) (*SolicitarPostagemReversaResponse, error) {
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

	c.DsNamespace = envNamespaceReversa
	c.FlSoap = true

	err := c.Post(soapCorreios, sendRequest, nil, &returnRequest, "XML", "SolicitarPostagemReversa")
	return returnRequest.SolicitarPostagemReversaResponse, err
}

// SolicitarPostagemSimultanea was auto-generated from WSDL.
func (c *Client) PostSolicitarPostagemSimultanea(SolicitarPostagemSimultanea *SolicitarPostagemSimultanea) (*SolicitarPostagemSimultaneaResponse, error) {
	sendRequest := struct {
		OperationSolicitarPostagemSimultanea `xml:"tns:solicitarPostagemSimultanea"`
	}{
		OperationSolicitarPostagemSimultanea{
			SolicitarPostagemSimultanea,
		},
	}

	returnRequest := struct {
		OperationSolicitarPostagemSimultaneaResponse `xml:"solicitarPostagemSimultaneaResponse"`
	}{}
	c.DsNamespace = envNamespaceReversa
	c.FlSoap = true

	err := c.Post(soapCorreios, sendRequest, nil, &returnRequest, "XML", "SolicitarPostagemSimultanea")
	return returnRequest.SolicitarPostagemSimultaneaResponse, err

}

// SolicitarRange was auto-generated from WSDL.
func (c *Client) PostSolicitarRange(SolicitarRange *SolicitarRange) (*SolicitarRangeResponse, error) {
	sendRequest := struct {
		OperationSolicitarRange `xml:"tns:solicitarRange"`
	}{
		OperationSolicitarRange{
			SolicitarRange,
		},
	}

	returnRequest := struct {
		OperationSolicitarRangeResponse `xml:"solicitarRangeResponse"`
	}{}
	c.DsNamespace = envNamespaceReversa
	c.FlSoap = true

	err := c.Post(soapCorreios, sendRequest, nil, &returnRequest, "XML", "SolicitarRange")
	return returnRequest.SolicitarRangeResponse, err

}

// ValidarPostagemReversa was auto-generated from WSDL.
func (c *Client) PostValidarPostagemReversa(ValidarPostagemReversa *ValidarPostagemReversa) (*ValidarPostagemReversaResponse, error) {
	sendRequest := struct {
		OperationValidarPostagemReversa `xml:"tns:validarPostagemReversa"`
	}{
		OperationValidarPostagemReversa{
			ValidarPostagemReversa,
		},
	}

	returnRequest := struct {
		OperationValidarPostagemReversaResponse `xml:"validarPostagemReversaResponse"`
	}{}
	c.DsNamespace = envNamespaceReversa
	c.FlSoap = true

	err := c.Post(soapCorreios, sendRequest, nil, &returnRequest, "XML", "ValidarPostagemReversa")
	return returnRequest.ValidarPostagemReversaResponse, err

}

// ValidarPostagemSimultanea was auto-generated from WSDL.
func (c *Client) PostValidarPostagemSimultanea(ValidarPostagemSimultanea *ValidarPostagemSimultanea) (*ValidarPostagemSimultaneaResponse, error) {
	sendRequest := struct {
		OperationValidarPostagemSimultanea `xml:"tns:validarPostagemSimultanea"`
	}{
		OperationValidarPostagemSimultanea{
			ValidarPostagemSimultanea,
		},
	}

	returnRequest := struct {
		OperationValidarPostagemSimultaneaResponse `xml:"validarPostagemSimultaneaResponse"`
	}{}
	c.DsNamespace = envNamespaceReversa
	c.FlSoap = true

	err := c.Post(soapCorreios, sendRequest, nil, &returnRequest, "XML", "ValidarPostagemSimultanea")
	return returnRequest.ValidarPostagemSimultaneaResponse, err

}
