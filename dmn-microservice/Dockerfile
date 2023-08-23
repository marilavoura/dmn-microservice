FROM golang:1.19

RUN mkdir ../home/dmn-microservice
WORKDIR /home/dmn-microservice

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -o .

CMD ["./dmn-microservice"]