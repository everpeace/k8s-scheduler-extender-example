FROM golang:1.9 as build

ARG VERSION

WORKDIR /go/src/k8s-scheduler-extender-example
COPY . .
RUN go install -ldflags "-s -w -X main.version=$VERSION" k8s-scheduler-extender-example

FROM gcr.io/google_containers/ubuntu-slim:0.14
COPY --from=build /go/bin/k8s-scheduler-extender-example /usr/bin/k8s-scheduler-extender-example

ENTRYPOINT ["k8s-scheduler-extender-example"]
