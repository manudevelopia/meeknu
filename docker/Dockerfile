FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./src /app/src

RUN go build -o meeknu-api src/cmd/meeknu-api/main.go

EXPOSE 1323

CMD [ "/app/meeknu-api" ]