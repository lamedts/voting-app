# build stage
FROM golang:1.10 as builder
ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep
COPY . $GOPATH/src/voting-app/voting-app-worker/
WORKDIR $GOPATH/src/voting-app/voting-app-worker/
RUN dep ensure --vendor-only
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/voting-app-worker .


# final stage
FROM alpine
COPY --from=builder /go/bin/voting-app-worker /go/bin/voting-app-worker
COPY config.yaml /go/bin/config.yaml
ENTRYPOINT ["/go/bin/voting-app-worker"]
