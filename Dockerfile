FROM golang:alpine

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o testjwtauth .

EXPOSE 8080

CMD ["./testjwtauth"]