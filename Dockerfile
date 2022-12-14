FROM golang:1.19.0-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o ./is-it-open cmd/is-it-open/main.go

EXPOSE 1323
USER 1337

CMD [ "./is-it-open" ]

