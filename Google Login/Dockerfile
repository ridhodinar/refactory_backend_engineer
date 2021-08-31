FROM golang:1.14-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /docker-golang
CMD [ "/docker-golang" ]