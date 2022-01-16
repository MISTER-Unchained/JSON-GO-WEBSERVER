FROM golang:latest

RUN mkdir /go/

COPY . /go/

EXPOSE 80

WORKDIR /go/

RUN go mod init main

RUN go mod tidy

CMD ["go", "build", "."]