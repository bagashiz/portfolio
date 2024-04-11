FROM golang:1.22-alpine as builder

WORKDIR /app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main /app/cmd/main.go

FROM scratch as final

WORKDIR /app
COPY --from=builder /app/main /app/

ENTRYPOINT [ "/app/main" ]
