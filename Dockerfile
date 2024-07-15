FROM golang:1.22 as BUILDER


# CRIANDO A PASTA RAIZ ########################################
WORKDIR /app
# COPIAR TODOS OS ARQUIVOS PARA A PASTA DOCKER ################
COPY src src
COPY go.mod go.mod
COPY go.sum go.sum
COPY main.go main.go

# RUN go mod download
# BUILDANDO O PROJETO GO #######################################
RUN CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -v -o meuprimeirocrudgo .

# ==============================================================

# CRIANDO OUTRO FROM PARA RODAR O CONTEXTO CRIADO ACIMA ########
FROM golang:1.22-alpine as RUNNER

# COPIANDO O CONTEXTO A CIMA PARA A NOVA IMAGEM ################
COPY --from=BUILDER /app/meuprimeirocrudgo .

# EXPOR A PORTA ###############################################
EXPOSE 8080

# EXECUTANDO COMANDO NO CMD ##################################
CMD [ "./meuprimeirocrudgo" ]
