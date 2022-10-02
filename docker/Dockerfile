FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ../cmd/meeknu-api/*.go ./

RUN go build -o /meeknu-api

EXPOSE 1323

CMD [ "/meeknu-api" ]