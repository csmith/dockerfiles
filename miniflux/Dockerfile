# Generated from https://github.com/csmith/dockerfiles/blob/master/miniflux/Dockerfile.gotpl
# BOM: {"apk:brotli-libs":"1.1.0-r2","apk:busybox":"1.36.1-r28","apk:busybox-binsh":"1.36.1-r28","apk:c-ares":"1.28.1-r0","apk:ca-certificates":"20240226-r0","apk:git":"2.45.2-r0","apk:libcrypto3":"3.3.0-r2","apk:libcurl":"8.7.1-r0","apk:libexpat":"2.6.2-r0","apk:libidn2":"2.3.7-r0","apk:libpsl":"0.21.5-r1","apk:libssl3":"3.3.0-r2","apk:libunistring":"1.2-r0","apk:musl":"1.2.5-r0","apk:nghttp2-libs":"1.62.0-r0","apk:pcre2":"10.43-r0","apk:zlib":"1.3.1-r1","apk:zstd-libs":"1.5.6-r0","github:miniflux/v2":"2.1.3","image:base":"b53f7bd15a51e9e9e1c9d1ddea46a2f62223ed7cd8d8fa320c0d6b323f9f3f99","image:golang":"c2a6061e5c0fc3f10832b672d84d043674dd30f2a0cf16852959f6aa97a9b477"}

FROM reg.c5h.io/golang@sha256:c2a6061e5c0fc3f10832b672d84d043674dd30f2a0cf16852959f6aa97a9b477 AS build

ARG TAG="2.1.3"

RUN set -eux; \
    apk add --no-cache \
        brotli-libs=1.1.0-r2\
        busybox=1.36.1-r28\
        busybox-binsh=1.36.1-r28\
        c-ares=1.28.1-r0\
        ca-certificates=20240226-r0\
        git=2.45.2-r0\
        libcrypto3=3.3.0-r2\
        libcurl=8.7.1-r0\
        libexpat=2.6.2-r0\
        libidn2=2.3.7-r0\
        libpsl=0.21.5-r1\
        libssl3=3.3.0-r2\
        libunistring=1.2-r0\
        musl=1.2.5-r0\
        nghttp2-libs=1.62.0-r0\
        pcre2=10.43-r0\
        zlib=1.3.1-r1\
        zstd-libs=1.5.6-r0\
        ; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/miniflux/v2 /go/src/github.com/miniflux/v2; \
    cd /go/src/github.com/miniflux/v2; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main .; \
    go-licenses save ./... --save_path=/notices;

FROM reg.c5h.io/base@sha256:b53f7bd15a51e9e9e1c9d1ddea46a2f62223ed7cd8d8fa320c0d6b323f9f3f99

COPY --from=build /go/src/github.com/miniflux/v2/main /miniflux
COPY --from=build /notices /notices
ENTRYPOINT ["/miniflux"]
