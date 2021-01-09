FROM golang:1.15

WORKDIR /app
COPY ./src .

RUN go build -o store .

CMD ["/app/store"]