# To comply with standard BKN practice, we use alpine.
ARG BUILDER_IMAGE=golang:1.17-buster
ARG BASE_IMAGE=bitnami/minideb:buster

FROM $BUILDER_IMAGE as builder

COPY . /root/go/src/app/

ARG BUILD_VERSION
ARG GOPROXY
ARG GOSUMDB=sum.golang.org

WORKDIR /root/go/src/app

ENV PATH="${PATH}:/usr/local/go/bin"
ENV BUILD_VERSION=$BUILD_VERSION
ENV GOPROXY=$GOPROXY
ENV GOSUMDB=$GOSUMDB

RUN go mod vendor

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -v -ldflags "-X main.version=$(BUILD_VERSION)" -installsuffix cgo -o app ./server

FROM $BASE_IMAGE

# RUN install_packages ca-certificates

WORKDIR /usr/app

COPY --from=builder /root/go/src/app/app /usr/app/app
COPY --from=builder /root/go/src/app/www /usr/app/www
COPY --from=builder /root/go/src/app/grpc_health_probe-linux-amd64 /usr/app/grpc_health_probe-linux-amd64

LABEL authors="Arya Utama Putra <arya@ordent.co>"

# PotatoBeans Co. adheres to OCI image specification.
LABEL org.opencontainers.image.authors="Arya Utama Putra <arya@ordent.co>"
LABEL org.opencontainers.image.title="go-base"
LABEL org.opencontainers.image.description="create go bases server"
LABEL org.opencontainers.image.vendor=""

EXPOSE 9090
EXPOSE 3000

ENTRYPOINT ["/usr/app/app"]
CMD ["grpc-gw-server", "--port1", "9090", "--port2", "3000", "--grpc-endpoint", ":9090"]