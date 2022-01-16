FROM golang:latest

COPY . .

EXPOSE 80

RUN go mod init main

CMD ["go", "build", "."]