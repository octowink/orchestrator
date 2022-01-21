FROM golang:1.17-alpine as build
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux

RUN apk add --no-cache make git

WORKDIR /go/src/github.com/octowink/orchestrator

# Pulling dependencies
COPY ./Makefile ./go.* ./
RUN make deps

# Building stuff
COPY . /go/src/github.com/octowink/orchestrator
RUN make build


FROM alpine:3.7
RUN adduser -D -u 1000 octowink

RUN apk add --no-cache ca-certificates
COPY --from=build /go/src/github.com/octowink/orchestrator/orchestrator /usr/local/bin/orchestrator

USER octowink
CMD ["orchestrator"]
