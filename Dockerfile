FROM golang:1.10-alpine as builder
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

ARG VERSION=0.0.1

# build
WORKDIR /go/src/k8s-scheduler-extender-example
COPY . .
RUN go install -ldflags "-s -w -X main.version=$VERSION" k8s-scheduler-extender-example

# runtime image
FROM gcr.io/google_containers/ubuntu-slim:0.14
COPY --from=builder /go/bin/k8s-scheduler-extender-example /usr/bin/k8s-scheduler-extender-example
ENTRYPOINT ["k8s-scheduler-extender-example"]
