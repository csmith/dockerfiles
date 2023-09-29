# Generated from https://github.com/csmith/dockerfiles/blob/master/miniflux/Dockerfile.gotpl
# BOM: {"apk:brotli-libs":"1.0.9-r14","apk:busybox":"1.36.1-r2","apk:busybox-binsh":"1.36.1-r2","apk:ca-certificates":"20230506-r0","apk:git":"2.40.1-r0","apk:libcrypto3":"3.1.3-r0","apk:libcurl":"8.3.0-r0","apk:libexpat":"2.5.0-r1","apk:libidn2":"2.3.4-r1","apk:libssl3":"3.1.3-r0","apk:libunistring":"1.1-r1","apk:musl":"1.2.4-r1","apk:nghttp2-libs":"1.55.1-r0","apk:pcre2":"10.42-r1","apk:zlib":"1.2.13-r1","github:miniflux/v2":"2.0.48","image:base":"8e5d45c90d7bcc56a15f1770933c8052a79773c1c631b96f36b5e25ef5606d5f","image:golang":"e678b10613f2308f383c6f251a67348a2e3f3244b2938d54712b5eb10f803ac0"}

FROM reg.c5h.io/golang@sha256:e678b10613f2308f383c6f251a67348a2e3f3244b2938d54712b5eb10f803ac0 AS build

ARG TAG="2.0.48"

RUN set -eux; \
    apk add --no-cache \
        brotli-libs=1.0.9-r14\
        busybox=1.36.1-r2\
        busybox-binsh=1.36.1-r2\
        ca-certificates=20230506-r0\
        git=2.40.1-r0\
        libcrypto3=3.1.3-r0\
        libcurl=8.3.0-r0\
        libexpat=2.5.0-r1\
        libidn2=2.3.4-r1\
        libssl3=3.1.3-r0\
        libunistring=1.1-r1\
        musl=1.2.4-r1\
        nghttp2-libs=1.55.1-r0\
        pcre2=10.42-r1\
        zlib=1.2.13-r1\
        ; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/miniflux/v2 /go/src/github.com/miniflux/v2; \
    cd /go/src/github.com/miniflux/v2; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main .; \
    go-licenses save ./... --save_path=/notices;

FROM reg.c5h.io/base@sha256:8e5d45c90d7bcc56a15f1770933c8052a79773c1c631b96f36b5e25ef5606d5f

COPY --from=build /go/src/github.com/miniflux/v2/main /miniflux
COPY --from=build /notices /notices
ENTRYPOINT ["/miniflux"]
