package brutils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// limpa remove caracteres especiais do CNPJ
func Limpa(cnpj string) string {
	return strings.NewReplacer(".", "", "-", "", "/", "").Replace(cnpj)
}

// todosDigitosIguais verifica se todos os dígitos do CNPJ são iguais
func TodosDigitosIguais(cnpj string) bool {
	for i := 1; i < len(cnpj); i++ {
		if cnpj[i] != cnpj[0] {
			return false
		}
	}
	return true
}

// calculaDigito calcula o dígito verificador do CNPJ
func CalculaDigito(cnpj string, peso int) int {
	soma := 0
	for i := 0; i < len(cnpj); i++ {
		num, _ := strconv.Atoi(string(cnpj[i]))
		soma += num * peso
		peso--
		if peso < 2 {
			peso = 9
		}
	}
	resto := soma % 11
	if resto < 2 {
		return 0
	}
	return 11 - resto
}

// ValidaCNPJ verifica se um CNPJ é válido
func ValidaCNPJ(cnpj string) bool {
	cnpj = limpa(cnpj)
	if len(cnpj) != 14 {
		return false
	}
	if todosDigitosIguais(cnpj) {
		return false
	}
	d1 := calculaDigito(cnpj[:12], 5)
	d2 := calculaDigito(cnpj[:12]+strconv.Itoa(d1), 6)
	return cnpj == cnpj[:12]+strconv.Itoa(d1)+strconv.Itoa(d2)
}

// FormataCNPJ formata um CNPJ com pontos, barras e hífen
func FormataCNPJ(cnpj string) string {
	cnpj = limpa(cnpj)
	if len(cnpj) != 14 {
		return ""
	}
	return fmt.Sprintf("%s.%s.%s/%s-%s", cnpj[:2], cnpj[2:5], cnpj[5:8], cnpj[8:12], cnpj[12:])
}

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
	url := fmt.Sprintf("https://www.receitaws.com.br/v1/cnpj/%s", limpa(cnpj))
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
