// Package external contém funções para interagir com APIs externas.
package external

import (
	"encoding/json" // Pacote json para codificação e decodificação JSON
	"fmt"           // Pacote fmt para formatação de saída
	"net/http"      // Pacote http para realizar requisições HTTP

	"github.com/CaioMartinss/br-utils-go/validation" // Pacote validation para validação de CEP
)

// DadosCEP representa a estrutura dos dados retornados pela API do ViaCEP
type DadosCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

// ConsultaCEP realiza uma consulta ao ViaCEP e retorna os dados do CEP
func ConsultaCEP(cep string) (*DadosCEP, error) {
	// Verifica se o CEP está no formato válido
	if !validation.ValidaCEP(cep) {
		return nil, fmt.Errorf("formato de CEP inválido")
	}

	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)

	// Realiza a requisição GET para a API do ViaCEP
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Verifica o status da resposta
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erro ao consultar o ViaCEP. Status: %d", resp.StatusCode)
	}

	// Decodifica a resposta JSON para a struct DadosCEP
	var dados DadosCEP
	err = json.NewDecoder(resp.Body).Decode(&dados)
	if err != nil {
		return nil, err
	}

	return &dados, nil
}
