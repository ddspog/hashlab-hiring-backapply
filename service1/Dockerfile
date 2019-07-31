# Development Container
FROM ddsdok/goserve as builder

COPY . /go/src/github.com/ddspog/hashlab-hiring-backapply/service1
WORKDIR /go/src/github.com/ddspog/hashlab-hiring-backapply/service1

# Install dependencies and build the app
ENV CGO_ENABLED 0
ENV GOOS "linux"
ENV GOARCH=amd64

# Show contents and build
RUN ls
RUN go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/service1
RUN setcap CAP_NET_BIND_SERVICE=+ep /go/bin/service1

# Production Container
FROM scratch

# Copy app from dev container.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/service1 /

# Ready to execute binary
WORKDIR /

# Exposer default port
EXPOSE 5001

# Execute binary, in Heroku as well
ENTRYPOINT [ "service1", "-v" ]