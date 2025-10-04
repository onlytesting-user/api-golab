FROM golang:1.25.1-alpine as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN apk add --no-cache binutils
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o myExec ./cmd/main.go && strip myExec

FROM alpine:3.20

RUN apk --no-cache add ca-certificates
COPY --from=builder /app/myExec /myExec

EXPOSE 8000

ENTRYPOINT ["/myExec"]
