FROM golang:1.13-stretch AS builder
RUN mkdir /myapp

# Adding first only dependencies files,
# by doing this, we help Docker's caching mechanism.
# If we change the code but not the dependencies,
# fetching the dependencies will be in a cached state.
COPY go.mod go.sum /myapp/
ADD Makefile /myapp
WORKDIR /myapp
RUN make fetch-deps

# Once dependencies were fetched, we proceed to build
ADD . /myapp
RUN make build-on-docker

FROM alpine:latest
RUN mkdir -p /go/bin
WORKDIR /go/bin
COPY --from=builder /myapp/myapp .
EXPOSE 8080
ENTRYPOINT /go/bin/myapp