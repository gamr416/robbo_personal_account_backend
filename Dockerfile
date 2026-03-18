FROM golang:latest

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
RUN go build -o robbo_server

EXPOSE 8080

CMD [ "/app/robbo_server" ]