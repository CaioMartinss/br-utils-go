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
    "github.com/CaioMartinss/br-utils-go/brutils"
)

func main() {
    cpf := "12345678909"
    if pacotes.IsValidCPF(cpf) {
        fmt.Println("CPF válido")
    } else {
        fmt.Println("CPF inválido")
    }
}
```

### Como Contribuir

Contribuições são bem-vindas! Sinta-se à vontade para abrir uma issue ou enviar um pull request. Para contribuir, leia as diretrizes de contribuição [aqui](CONTRIBUTING.md).



### Licença

Este projeto está licenciado sob a [Licença MIT](LICENSE). Veja o arquivo LICENSE para mais detalhes.




### br-utils-go v1.0.0

![badge atenção](https://img.shields.io/badge/Atenção-yellow?style=for-the-badge&logo=none)
![badge versão](https://img.shields.io/badge/Versão-v1.0.0-blue?style=for-the-badge)

Biblioteca em Go para validação e formatação de CPFs e CNPJs, na versão inicial com foco em transformação de formatos corretos.



