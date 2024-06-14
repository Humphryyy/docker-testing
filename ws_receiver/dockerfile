FROM golang:latest

WORKDIR /build

COPY go.mod .
COPY main.go .

RUN go get

RUN go build -o main .

ENTRYPOINT [ "/build/main" ]