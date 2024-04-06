FROM golang:latest

WORKDIR /app

COPY . .
COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go build -o main .
RUN chmod a+x main

CMD ["./main"]