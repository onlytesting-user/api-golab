FROM golang:1.25.1

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8000

RUN go build -o main -ldflags="-s -w" cmd/main.go

CMD [ "./main" ]
