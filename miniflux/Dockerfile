# Generated from https://github.com/csmith/dockerfiles/blob/master/miniflux/Dockerfile.gotpl
# BOM: {"apk:brotli-libs":"1.1.0-r1","apk:busybox":"1.36.1-r15","apk:busybox-binsh":"1.36.1-r15","apk:c-ares":"1.27.0-r0","apk:ca-certificates":"20240226-r0","apk:git":"2.43.0-r0","apk:libcrypto3":"3.1.4-r6","apk:libcurl":"8.5.0-r0","apk:libexpat":"2.6.2-r0","apk:libidn2":"2.3.4-r4","apk:libssl3":"3.1.4-r6","apk:libunistring":"1.1-r2","apk:musl":"1.2.4_git20230717-r4","apk:nghttp2-libs":"1.58.0-r0","apk:pcre2":"10.42-r2","apk:zlib":"1.3.1-r0","github:miniflux/v2":"2.1.2","image:base":"c4936e066851061b7fd136c869de1dd56413038f3fc44545e606b637a074fd8d","image:golang":"00405f110fbb03f3a820fb90077f26c41b754718998a9f643ccadc778cc3522c"}

FROM reg.c5h.io/golang@sha256:00405f110fbb03f3a820fb90077f26c41b754718998a9f643ccadc778cc3522c AS build

ARG TAG="2.1.2"

RUN set -eux; \
    apk add --no-cache \
        brotli-libs=1.1.0-r1\
        busybox=1.36.1-r15\
        busybox-binsh=1.36.1-r15\
        c-ares=1.27.0-r0\
        ca-certificates=20240226-r0\
        git=2.43.0-r0\
        libcrypto3=3.1.4-r6\
        libcurl=8.5.0-r0\
        libexpat=2.6.2-r0\
        libidn2=2.3.4-r4\
        libssl3=3.1.4-r6\
        libunistring=1.1-r2\
        musl=1.2.4_git20230717-r4\
        nghttp2-libs=1.58.0-r0\
        pcre2=10.42-r2\
        zlib=1.3.1-r0\
        ; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/miniflux/v2 /go/src/github.com/miniflux/v2; \
    cd /go/src/github.com/miniflux/v2; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main .; \
    go-licenses save ./... --save_path=/notices;

FROM reg.c5h.io/base@sha256:c4936e066851061b7fd136c869de1dd56413038f3fc44545e606b637a074fd8d

COPY --from=build /go/src/github.com/miniflux/v2/main /miniflux
COPY --from=build /notices /notices
ENTRYPOINT ["/miniflux"]
