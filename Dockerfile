FROM golang:latest

RUN mkdir /go-project/

COPY . /go-project/

EXPOSE 80

WORKDIR /go/

RUN go mod init main

RUN go mod tidy

CMD ["go", "build", "."]