FROM golang:1.15

COPY server.go server.go

RUN go build -o server .

CMD ["./server"]