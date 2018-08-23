FROM golang:1.10-alpine as build

ARG VERSION=0.0.1

RUN apk add --no-cache git bash curl
RUN curl https://glide.sh/get | sh

WORKDIR /go/src/k8s-scheduler-extender-example

# constructing vender layer
COPY glide.yaml glide.lock /go/src/k8s-scheduler-extender-example/
RUN glide install -v

COPY . .
RUN go install -ldflags "-s -w -X main.version=$VERSION" k8s-scheduler-extender-example


FROM gcr.io/google_containers/ubuntu-slim:0.14
COPY --from=build /go/bin/k8s-scheduler-extender-example /usr/bin/k8s-scheduler-extender-example
ENTRYPOINT ["k8s-scheduler-extender-example"]
