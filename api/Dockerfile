FROM golang:1.20 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o ./server

FROM alpine
WORKDIR /app

COPY --from=builder /app/server ./
EXPOSE 8080
CMD ["./server"]