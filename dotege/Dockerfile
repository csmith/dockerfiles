# Generated from https://github.com/csmith/dockerfiles/blob/master/dotege/Dockerfile.gotpl
# BOM: {"github:csmith/dotege":"v1.3.1","image:base":"d97e687affbcacbd2173a2a58381ecf270adc8d37d409258cfbc04b0d0a8b092","image:golang":"e4a0cd1d6a1ae7a9d0e83ea82a8e434616ce8ef488ef4a20e4876091ae701e21"}

FROM reg.c5h.io/golang@sha256:e4a0cd1d6a1ae7a9d0e83ea82a8e434616ce8ef488ef4a20e4876091ae701e21 AS build

ARG TAG="v1.3.1"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/csmith/dotege /go/src/github.com/csmith/dotege; \
    cd /go/src/github.com/csmith/dotege; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' -X main.GitSHA=$(git rev-parse --short HEAD)" -trimpath -tags lego_httpreq -o main ./cmd/dotege; \
    go-licenses save ./... --save_path=/notices; \
    # Make volume mount points
    mkdir -p /data/config /data/certs /data/output; \
    chown -R 65532:65532 /data;

FROM reg.c5h.io/base@sha256:d97e687affbcacbd2173a2a58381ecf270adc8d37d409258cfbc04b0d0a8b092

COPY --from=build /go/src/github.com/csmith/dotege/main /dotege
COPY --from=build /notices /data /
COPY --from=build /go/src/github.com/csmith/dotege/templates /templates

VOLUME /data/config
VOLUME /data/certs
VOLUME /data/output

ENTRYPOINT ["/dotege"]
