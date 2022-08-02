# Generated from https://github.com/csmith/dockerfiles/blob/master/dotege/Dockerfile.gotpl
# BOM: {"github:csmith/dotege":"v1.3.1","image:base":"4061e85f242c026f31ae923214310d31cf4a3c2b112953fd9307fe649b3b1d49","image:golang":"a35928efa4dba33cbb0af9b0f5f9931290bb96d4df37bc3948fa40ec4ba4bcec"}

FROM reg.c5h.io/golang@sha256:a35928efa4dba33cbb0af9b0f5f9931290bb96d4df37bc3948fa40ec4ba4bcec AS build

ARG TAG="v1.3.1"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/csmith/dotege /go/src/github.com/csmith/dotege; \
    cd /go/src/github.com/csmith/dotege; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' -X main.GitSHA=$(git rev-parse --short HEAD)" -trimpath -tags lego_httpreq -o main ./cmd/dotege; \
    go-licenses save ./... --save_path=/notices; \
    # Make volume mount points
    mkdir -p /data/config /data/certs /data/output; \
    chown -R 65532:65532 /data;

FROM reg.c5h.io/base@sha256:4061e85f242c026f31ae923214310d31cf4a3c2b112953fd9307fe649b3b1d49

COPY --from=build /go/src/github.com/csmith/dotege/main /dotege
COPY --from=build /notices /data /
COPY --from=build /go/src/github.com/csmith/dotege/templates /templates

VOLUME /data/config
VOLUME /data/certs
VOLUME /data/output

ENTRYPOINT ["/dotege"]
