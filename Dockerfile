FROM golang:1.22-alpine

WORKDIR /application

COPY . .

RUN go mod download && \
    go build -o main ./cmd/eulabs

EXPOSE 8080

CMD ["./main"]