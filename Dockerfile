FROM golang:latest

COPY . .

EXPOSE 80

CMD ["go", "build", "."]