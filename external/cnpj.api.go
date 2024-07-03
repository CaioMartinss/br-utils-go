package external

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/CaioMartinss/br-utils-go/validation"
)

// Atividade define a estrutura da atividade principal da empresa
type Atividade struct {
	Texto string `json:"text"`
}

// Empresa define a estrutura da resposta da API ReceitaWS
type Empresa struct {
	Nome       string      `json:"nome"`
	Fantasia   string      `json:"fantasia"`
	Atividades []Atividade `json:"atividade_principal"`
	Municipio  string      `json:"municipio"`
	UF         string      `json:"uf"`
	Status     string      `json:"status"`
}

// ConsultaCNPJ consulta a API ReceitaWS para obter dados da empresa
func ConsultaCNPJ(cnpj string) (*Empresa, error) {
	url := fmt.Sprintf("https://www.receitaws.com.br/v1/cnpj/%s", validation.Limpa(cnpj))
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer a requisição: %v", err)
	}
	defer resp.Body.Close()

	var empresa Empresa
	if err := json.NewDecoder(resp.Body).Decode(&empresa); err != nil {
		return nil, fmt.Errorf("erro ao decodificar a resposta: %v", err)
	}

	return &empresa, nil
}
