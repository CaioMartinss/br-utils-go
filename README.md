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

```go
package main

import (
    "fmt"
    "github.com/CaioMartinss/br-utils-go/brutils"
)

func main() {
	cpf := "123.456.789-09"
	fmt.Printf("CPF %s é válido? %v\n", cpf, brutils.ValidaCPF(cpf))
	fmt.Printf("CPF formatado: %s\n", brutils.FormataCPF(cpf))

}
```

#### Output

```go
CPF 123.456.789-09 é válido? true
CPF formatado: 123.456.789-09

```

#### Validação de CNPJ

A biblioteca inclui uma função para validar CNPJ.

```go
package main

import (
    "fmt"
    "github.com/CaioMartinss/br-utils-go"
)

func main() {

	cnpj := "12.345.678/0001-95"
	fmt.Printf("CNPJ %s é válido? %v\n", cnpj, brutils.ValidaCNPJ(cnpj))
	fmt.Printf("CNPJ formatado: %s\n", brutils.FormataCNPJ(cnpj))
}
```
#### Output

```go
CNPJ 12.345.678/0001-95 é válido? true
CNPJ formatado: 12.345.678/0001-95

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



