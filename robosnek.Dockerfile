FROM golang:1.14-alpine AS builder
WORKDIR /go/src/github.com/broothie/battlesnake
COPY . .
RUN apk add --update ca-certificates
RUN go build robosnek/main.go

FROM alpine:3.12.0
COPY --from=builder /go/src/github.com/broothie/battlesnake/main main
CMD ./main
