FROM golang:1.10
WORKDIR /go/src/github.com/fanminshi/simple-server/server
COPY ./server/main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .
WORKDIR /go/src/github.com/fanminshi/simple-server/client
COPY ./client/main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o client .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/fanminshi/simple-server/server .
COPY --from=0 /go/src/github.com/fanminshi/simple-server/client .
COPY src src
