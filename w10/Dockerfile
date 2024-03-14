FROM golang:1.22 as builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o torontoTime .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/torontoTime .

EXPOSE 8014

CMD ["./torontoTime"]
