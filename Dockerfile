FROM golang:1.12-alpine

WORKDIR /go/src/github.com/broothie/battlesnake
COPY . .

RUN apk add --update ca-certificates
RUN go build

CMD ./battlesnake
