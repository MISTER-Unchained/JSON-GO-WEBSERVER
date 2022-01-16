FROM golang:latest

RUN mkdir /go-project/

WORKDIR /go-project/

COPY . .

EXPOSE 80


RUN go mod init main

RUN go mod tidy

RUN go build -o out .

CMD ["sudo", "./out"]