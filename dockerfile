FROM golang:1.16-alpine

# Define o diretório de trabalho
WORKDIR docker build -t meu-app "https://github.com/iltrd/manipular-dados"


# Copia o arquivo go.mod e go.sum
COPY go.mod go.sum ./
# Baixa as dependências
RUN go mod download
RUN go mod download golang.org/x/text@v0.3.8


# Copia o restante do código
COPY . .

# Compila o aplicativo
RUN go build -o main .

# Define a porta de exposição
EXPOSE 8080

# Define o comando padrão a ser executado quando o contêiner é iniciado
CMD ["./main"]