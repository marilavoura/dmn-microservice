FROM golang:1.19

RUN mkdir ../home/producer
WORKDIR /home/producer

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

WORKDIR /home/producer

RUN go build -o .

CMD ["./producer"]