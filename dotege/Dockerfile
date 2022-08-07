# Generated from https://github.com/csmith/dockerfiles/blob/master/dotege/Dockerfile.gotpl
# BOM: {"github:csmith/dotege":"v1.3.1","image:base":"4061e85f242c026f31ae923214310d31cf4a3c2b112953fd9307fe649b3b1d49","image:golang":"7f7f0260925d34f1c0734ccc9eb6106f220ff3f285e2d5bdf07e5fd4231d94e1"}

FROM reg.c5h.io/golang@sha256:7f7f0260925d34f1c0734ccc9eb6106f220ff3f285e2d5bdf07e5fd4231d94e1 AS build

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
