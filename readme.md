# Manipular dados 



# Pacotes utilizados 

database/sql: Fornece uma interface de banco de dados SQL genérica
enconding/csv: lê e grava arquivos CSV
fmt: fornece funções de formatação 
io: fornece primitivas básicas de E/S.
log: fornece um pacore de registro simples.
os: fornece uma interface independente de uma platafora para o sistema operacional.
regexp: fornece funcionalidade de expressão regilar 
strconv: fornece funções para converter strings em tipos numéricos
strings: fornece funções para manipular strings
github.com/go-playground/validator: fornece um pacote de validação 


O código define uma estrutura chamada Record que apresente um único registro do arquivo csv/txt. Cada campo na struct possui uma marca de validação que especifica as regras de validação para esse campo. As regras de validação são defnidas usando a função validator. 

A função CleanData remove espaços em branco iniciais e finais de cada campo nos dados. 

A função InsertData insere os dados em um banco de dados PostgreSQL.