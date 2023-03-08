FROM golang:1.16-alpine

# Define o diretório de trabalho
WORKDIR /app

# Copia o arquivo go.mod e go.sum
COPY go.mod .
COPY go.sum .

# Baixa as dependências
RUN go mod download

# Copia o restante do código
COPY . .

# Compila o aplicativo
RUN go build -o main .

# Define a porta de exposição
EXPOSE 8080

# Define o comando padrão a ser executado quando o contêiner é iniciado
CMD ["./main"]