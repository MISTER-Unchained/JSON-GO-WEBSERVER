FROM golang:latest

RUN mkdir /go-project/

WORKDIR /go-project/

COPY . .

EXPOSE 5000

RUN go mod init main

RUN go mod tidy


CMD ["go", "run", "."]