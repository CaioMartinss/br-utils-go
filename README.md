<h3 align="center">br-utils-go</h3>

br-utils-go é uma biblioteca de utilitários em Go para desenvolvedores brasileiros. Ela oferece diversas funcionalidades úteis para validação de dados, manipulação de strings e muito mais.

### Instalação

Para instalar a biblioteca, você pode utilizar o `go get`:

```bash
go get github.com/CaioMartinss/br-utils-go

```

### Uso

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

#### Como Contribuir

Contribuições são bem-vindas! Sinta-se à vontade para abrir uma issue ou enviar um pull request.

1. Faça um fork do repositório
2. Crie uma branch para sua feature (`git checkout -b minha-feature`)
3. Comite suas mudanças (`git commit -am 'Adiciona minha feature'`)
4. Faça push para a branch (`git push origin minha-feature`)
5. Crie um novo Pull Request



### Licença

Este projeto está licenciado sob a [Licença MIT](LICENSE). Veja o arquivo LICENSE para mais detalhes.




### br-utils-go v1.0.0

![badge](https://img.shields.io/badge/Atenção-yellow?style=for-the-badge&logo=none)

Obs: Esta é a versão v1.0.0 que inclui, por enquanto, verificação e transformação de formatos corretos para CPFs e CNPJs.



