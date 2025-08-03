FROM golang:1.24.5

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o main src/cmd/server/main.go

EXPOSE 8080

ENTRYPOINT [ "./main" ]