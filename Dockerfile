FROM golang:1.10
WORKDIR /go/src/github.com/fanminshi/simple-server/
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/fanminshi/simple-server/server .
COPY server.key server.key
COPY server.crt server.crt
COPY src src
CMD ["./server"] 
