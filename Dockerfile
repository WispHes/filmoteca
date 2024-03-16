FROM golang:latest

WORKDIR /filmoteca

COPY filmoteca .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]