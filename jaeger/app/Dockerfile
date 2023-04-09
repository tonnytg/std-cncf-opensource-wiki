# Imagem de base que contém as dependências do Go
FROM golang:1.16

# Define o diretório de trabalho
WORKDIR /go/src/app

# Copia o conteúdo da pasta local para o contêiner
COPY . .

# Baixa as dependências e compila a aplicação
RUN go get -d -v ./...
RUN go install -v ./...

# Expõe a porta que a aplicação utiliza
EXPOSE 8080

# Define o comando padrão que será executado quando o contêiner for iniciado
CMD ["app"]
