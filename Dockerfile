# Assuming you are using Go for the server
FROM golang:latest

WORKDIR /app
RUN go mod download
COPY ./server .

RUN go build -o server ./cmd/

CMD ["./server"]