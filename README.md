<h1 align="center">br-utils-go</h1>  

br-utils-go é uma biblioteca de utilitários em Go para desenvolvedores brasileiros. Ela oferece diversas funcionalidades úteis para validação de dados, manipulação de strings e muito mais.

## Instalação

Para instalar a biblioteca, você pode utilizar o `go get`:

```bash
go get github.com/CaioMartinss/br-utils-go

```

## Uso

#### Validação de CPF

A biblioteca inclui uma função para validar CPFs.

```go
package main

import (
    "fmt"
    "github.com/CaioMartinss/br-utils-go/cmd"
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
    "github.com/CaioMartinss/br-utils-go/cmd"
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

### Como Contribuir

Contribuições são bem-vindas! Sinta-se à vontade para abrir uma issue ou enviar um pull request. Para contribuir, leia as diretrizes de contribuição [aqui](CONTRIBUTING.md).



### Licença

Este projeto está licenciado sob a [Licença MIT](LICENSE). Veja o arquivo LICENSE para mais detalhes.




### br-utils-go v1.0.0

![badge atenção](https://img.shields.io/badge/Atenção-yellow?style=for-the-badge&logo=none)
![badge versão](https://img.shields.io/badge/Versão-v1.0.0-blue?style=for-the-badge)

Biblioteca em Go para validação e formatação de CPFs e CNPJs, na versão inicial com foco em transformação de formatos corretos.



