FROM golang:1.20.1

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

RUN export GO111MODULE=on

#RUN go mod init

#RUN go build -o /main.go

EXPOSE 9090

ENTRYPOINT [ "app/main.go" ]