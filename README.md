<p align="center">
<img src="http://img.shields.io/static/v1?label=STATUS&message=EM%20DESENVOLVIMENTO&color=GREEN&style=for-the-badge"/>


# Manipular dados 
Serviço criação para ler arquivo csv/txt, valida os dados e  os inseri no banco de dados PostgreSQL. O código executa as seguintes tarefas. 

1 - lê os dados do arquivo. 
2 - Limpa os dados cortando os espaços em branco de cada campo.
3 - Inseri os dados limpos em um banco de dados 
4 - Valida cada registro usando uma struct e um pacote validador. 

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
