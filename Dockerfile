FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN go install github.com/githubnemo/CompileDaemon

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.8.0/wait /wait
RUN chmod +x /wait

COPY . .

ENTRYPOINT sh -c "/wait" && CompileDaemon --build="go build -o main ./cmd/main/main.go" --command=./main
