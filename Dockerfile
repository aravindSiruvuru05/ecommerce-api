# Use an official Golang runtime as a parent image
FROM golang:1.20-alpine

# RUN go get -u github.com/kyleconroy/sqlc/cmd/sqlc@latest

WORKDIR /go/src/haste

COPY ./src/haste/go.mod ./

RUN go mod download
RUN go mod tidy

# Install Beego and its dependencies - this is done after coping source code as go installs can only done in a module from 1.18. not out side
RUN apk add git && \
    go install github.com/beego/bee@v1.12.3 && \
    go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY ./src/haste .

EXPOSE 8080

CMD ["bee", "run"]