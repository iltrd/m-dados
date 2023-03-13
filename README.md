<p align="center">
<img src="http://img.shields.io/static/v1?label=STATUS&message=EM%20DESENVOLVIMENTO&color=GREEN&style=for-the-badge"/>


# Manipular dados 

O serviço irá ler o arquivo de entrada, separar as linhas em colunas, realizar a higienização dos dados (remover acentos e converter para maiúsculo), validar os CPFs/CNPJs e persistir os dados no banco de dados PostgreSQL.
Além disso, o serviço utiliza o padrão Clean Code ao nomear variáveis, funções e estruturas de dados de forma clara e concisa. Também são incluídos comentários no código para explicar o funcionamento de cada parte do serviço.
Por fim, é importante destacar que o serviço foi otimizado para ter uma performance de 1 minuto, utilizando contextos para controlar o tempo máximo de execução e transações para reduzir a quantidade de operações de I/O no banco de dados.

## Tecnologias utilizadas 
  <div style="display: inline_block"><br>
  <img align="center" alt="let-Go" height="40" width="50" src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original-wordmark.svg">
  <img align="center" alt="let-postgres" height="40" width="50" src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/postgresql/postgresql-original-wordmark.svg">
    <img align="center" alt="let-js" height="40" width="50" src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/docker/docker-original.svg">
</div>


## Pacotes utilizados 

- `database/sql` Fornece uma interface de banco de dados SQL genérica
- `enconding/csv` lê e grava arquivos CSV
- `fmt` fornece funções de formatação 
- `io` fornece primitivas básicas de E/S.
- `log` fornece um pacote de registro simples.
- `os` fornece uma interface independente de uma plataforma para o sistema operacional.
- `regexp` fornece funcionalidade de expressão regilar 
- `strconv` fornece funções para converter strings em tipos numéricos
- `strings` fornece funções para manipular strings
- `github.com/go-playground/validator` fornece um pacote de validação 


O código define uma estrutura chamada Record que apresente um único registro do arquivo csv/txt. Cada campo na struct possui uma marca de validação que especifica as regras de validaçã para o campo. As regras são dfinidas através da função  `validator`. 

A função `CleanData` remove espaços em branco iniciais e finais de cada campo nos dados. 

A função `InsertData` insere os dados em um banco de dados PostgreSQL.

A função `ValidateRecords` valida cada registro nos dados utilizando a função `validator`

A função `main` lê os dados  do arquivo, inseri no banco de dados  e valida os campos. Se qualquer uma dessas etapas falhar, o programa registrará uma mensagem de erro e será encerrado. 

## Para executar o programa 

``` 
docker-compose up 
```


## Status atual do serviço 

Atualmente venho me deparado com o erro abaixo, mesmo tenho feito todas as importações e seguindo os comando de 'go.mod init' e 'go.mod tidy'

```
 > [7/7] RUN go build -o main .:
#0 1.741 go: github.com/go-playground/locales@v0.14.1 requires
#0 1.741        golang.org/x/text@v0.3.8: missing go.sum entry; to add it:
#0 1.741        go mod download golang.org/x/text
```
