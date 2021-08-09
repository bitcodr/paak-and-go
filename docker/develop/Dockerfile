FROM golang:1.16-alpine

RUN apk add --no-cache git make curl

WORKDIR /app

ADD . .

RUN make mod

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="make build" --command=./bin/paak