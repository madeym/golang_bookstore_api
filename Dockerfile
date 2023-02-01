FROM golang

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o docker-bookstore

EXPOSE 8080

CMD ["go", "run", "main.go"]