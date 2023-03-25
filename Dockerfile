# Use an official Golang runtime as a parent image
FROM golang:1.16-alpine

# # Install Beego and its dependencies
RUN apk add --no-cache git && \
    go get github.com/beego/bee@v1.12.3

WORKDIR /go/src/haste

COPY ./src/haste .

# RUN go build -o main .

RUN go mod download
RUN go mod tidy

EXPOSE 8080

CMD ["bee", "run"]
