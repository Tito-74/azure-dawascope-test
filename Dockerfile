FROM golang:1.18-alpine3.17
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]
RUN apk add --no-cache bash