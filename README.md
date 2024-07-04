<h1 align="center">br-utils-go</h1>  

br-utils-go é uma biblioteca de utilitários em Go para desenvolvedores brasileiros. Ela oferece diversas funcionalidades úteis para validação de dados, manipulação de strings e muito mais.

## Instalação

Para instalar a biblioteca, você pode utilizar o `go get`:

```bash
go get module github.com/CaioMartinss/br-utils-go

```


## Uso

#### Validação de CPF

A biblioteca inclui uma função para validar CPFs.

#### Inpu

Criando exemplos de validação de CPF usando as funções definidas em `validation/cpf.validation.go`.

```go
package examples

import (
	"fmt"
	"github.com/CaioMartinss/br-utils-go/validation"
)

func Cpf_function() {
	cpf := "123.456.789-09"
	fmt.Printf("CPF %s é válido? %v\n", cpf, validation.ValidaCPF(cpf))
	fmt.Printf("CPF formatado: %s\n", validation.FormataCPF(cpf))

	cpf = "12345678909"
	fmt.Printf("CPF %s é válido? %v\n", cpf, validation.ValidaCPF(cpf))
	fmt.Printf("CPF formatado: %s\n", validation.FormataCPF(cpf))


}
```
No arquivo `main.go`, chamando o exemplo de validação de CPF e exibindo-o:

```go
package main

import (
	"fmt"

	"github.com/CaioMartinss/br-utils-go/examples"
)

func main() {
	fmt.Println("Exemplo de validação de CPF:")
	examples.Cpf_function()

}
```

#### Output

```go
Exemplo de validação de CPF:
CPF 123.456.789-09 é válido? true
CPF formatado: 123.456.789-09
CPF 12345678909 é válido? true
CPF formatado: 123.456.789-09

```

#### Validação de CNPJ

A biblioteca inclui uma função para validar CNPJ.

#### Input

Criando exemplos de validação de CNPJ usando as funções definidas em `validation/cnpj.validation.go`.

```go
package examples

import (
	"fmt"
	"github.com/CaioMartinss/br-utils-go/validation"

	"github.com/CaioMartinss/br-utils-go/external"
)

func Cnpj_function() {
	cnpj := "00000000000191"
	if !validation.ValidaCNPJ(cnpj) {
		fmt.Println("CNPJ inválido")
		return
	}

	empresa, err := external.ConsultaCNPJ(cnpj)
	if err != nil {
		fmt.Println("Erro ao consultar CNPJ:", err)
		return
	}

	fmt.Printf("CNPJ: %s\n", validation.FormataCNPJ(cnpj))
	fmt.Printf("Nome: %s\n", empresa.Nome)
	fmt.Printf("Nome Fantasia: %s\n", empresa.Fantasia)
	if len(empresa.Atividades) > 0 {
		fmt.Printf("Atividade Principal: %s\n", empresa.Atividades[0].Texto)
	}
	fmt.Printf("Município: %s\n", empresa.Municipio)
	fmt.Printf("UF: %s\n", empresa.UF)
	fmt.Printf("Status: %s\n", empresa.Status)
}
```

No arquivo `main.go`, chamando o exemplo de validação e consulta de CNPJ e exibindo-o:

```go
package main

import (
	"fmt"

	"github.com/CaioMartinss/br-utils-go/examples"
)

func main() {
	fmt.Println("\nExemplo de validação de CNPJ:")
	examples.Cnpj_function()
}

```

#### Output

```go
Exemplo de validação de CNPJ:
CNPJ: 00.000.000/0001-91
Nome: BANCO DO BRASIL SA
Nome Fantasia: DIRECAO GERAL
Atividade Principal: Bancos múltiplos, com carteira comercial
Município: BRASILIA
UF: DF
Status: OK

```

### Dependências Externas

<details>
<summary>CEP</summary>

Este projeto utiliza o serviço [ViaCep](https://viacep.com.br/), uma API pública brasileira que permite consultar e validar CEPs. Ela fornece informações atualizadas diretamente dos Correios, incluindo detalhes como logradouro, bairro, cidade e estado associados a um CEP.

</details>

<details>
<summary>CNPJ</summary>

Este projeto utiliza o serviço [receitaWS](https://receitaws.com.br/#section-api), uma API que facilita a consulta de CNPJs no Brasil. Com ele, é possível verificar a existência de um CNPJ, obter informações detalhadas sobre a empresa cadastrada, como razão social, data de abertura, natureza jurídica, situação cadastral e endereço.

</details>



### Lista completa de exemplos
Confira a lista completa de exemplos [aqui](../dev/exemplos.md). Esta lista contém diversos casos de uso e demonstrações que podem ajudar a entender melhor o projeto.



### Branch de Desenvolvimento (dev)

A branch mais atual e ativa para desenvolvimento é a [dev](../dev). Esta branch é onde estão sendo feitas as atualizações mais recentes e onde novos recursos estão sendo desenvolvidos e testados.


### Como Contribuir

Contribuições são bem-vindas! Sinta-se à vontade para abrir uma issue ou enviar um pull request. Para contribuir, leia as diretrizes de contribuição [aqui](CONTRIBUTING.md).



### Licença

Este projeto está licenciado sob a [Licença MIT](LICENSE). Veja o arquivo LICENSE para mais detalhes.




### br-utils-go v1.0.0

[![badge atenção](https://img.shields.io/badge/Atenção-yellow?style=flat&logo=none)](https://example.com)
[![badge versão](https://img.shields.io/badge/Versão-v1.0.0-blue?style=flat&logo=none)](https://example.com)

Biblioteca em Go para validação e formatação de CPFs e CNPJs, na versão inicial com foco em transformação de formatos corretos.



