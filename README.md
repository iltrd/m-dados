
# Manipular dados 
Este é um aplicativo Go que lê um arquivo csv/txt e insere os dados em um banco de dados PostgreSQL. O aplicativo usa a biblioteca github.com/lib/pq para se conectar ao banco de dados.

O aplicativo consiste em dois arquivos principais: main.go e database.go. O arquivo main.go contém o código principal do aplicativo e o arquivo database.go contém as funções de manipulação de banco de dados.
Configuração do Banco de Dados
Antes de executar o aplicativo, você precisará configurar um banco de dados PostgreSQL. O aplicativo espera que o banco de dados tenha os seguintes detalhes de conexão:

host: localhost (para conexão local)
porta: 5432
usuário: postgres
senha: 1404
nome do banco de dados: go_postgresql
Você pode usar o Docker para configurar um banco de dados PostgreSQL com as configurações acima. Para fazer isso, você precisará ter o Docker instalado em sua máquina.

Para iniciar um contêiner PostgreSQL com Docker, execute o seguinte comando:

```
docker run --name postgres -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres
```
Este comando iniciará um contêiner PostgreSQL com a senha 1404 e exporá a porta 5432. O nome do banco de dados padrão é postgres.

# Executando o Aplicativo

Antes de executar o aplicativo, você precisará ter o Go instalado em sua máquina. Você pode baixar o Go em https://golang.org/dl/.

Para executar o aplicativo, execute o seguinte comando no diretório do arquivo main.go:
```
go run main.go
```
Este comando compilará e executará o aplicativo.

O aplicativo lerá o arquivo base_teste.txt no diretório raiz do projeto e inserirá os dados em um banco de dados PostgreSQL com a tabela clientes. A tabela terá os seguintes campos:

- `cpf: inteiro`
- `private: inteiro`
- `incompleto: inteiro`
- `ultima_compra: data`
- `ticket_medio: decimal`
- `ticket_ult_comp: decimal`
- `loja_frequente: inteiro`
- `loja_ult_comp: inteiro`

Os dados no arquivo base_teste.txt devem estar separados por vírgula e devem ser do seguinte formato:

```
12345678900,1,0,2022-01-01,100.00,200.00,1,2
```
O aplicativo irá fazer o split dos dados em colunas no banco de dados.

# Executando o Aplicativo com Docker Compose

Você também pode executar o aplicativo usando Docker Compose. O Docker Compose é uma ferramenta que permite definir e executar aplicativos Docker com vários contêineres.

O arquivo docker-compose.yml define dois serviços: db e app. O serviço db usa a imagem oficial do PostgreSQL e define as configurações do banco de dados. O serviço app usa o Dockerfile no diretório atual para construir uma imagem Docker do aplicativo e define as variáveis de ambiente para conectar ao banco de dados.

Para executar o aplicativo com Docker Compose, execute o seguinte comando no diretório do arquivo docker-compose.yml:

```
docker-compose up --build
```

