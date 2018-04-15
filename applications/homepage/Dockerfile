# stage 0
FROM golang:latest as builder
RUN mkdir -p /go/src/app
WORKDIR /go/src/app
COPY . .
RUN go get github.com/gorilla/handlers && \
    go get github.com/gorilla/mux
RUN GOARCH=amd64 GOOS=linux go build -ldflags "-linkmode external -extldflags -static -w"

# stage 1
FROM scratch
WORKDIR /
COPY --from=builder /go/src/app .
COPY public .
ENTRYPOINT ["/app"]
