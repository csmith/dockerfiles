# Generated from https://github.com/csmith/dockerfiles/blob/master/miniflux/Dockerfile.gotpl
# BOM: {"apk:brotli-libs":"1.1.0-r2","apk:c-ares":"1.34.3-r0","apk:ca-certificates-bundle":"20241010-r0","apk:git":"2.47.1-r0","apk:libcrypto3":"3.3.2-r4","apk:libcurl":"8.11.1-r0","apk:libexpat":"2.6.4-r0","apk:libidn2":"2.3.7-r0","apk:libpsl":"0.21.5-r3","apk:libssl3":"3.3.2-r4","apk:libunistring":"1.2-r0","apk:musl":"1.2.5-r8","apk:nghttp2-libs":"1.64.0-r0","apk:pcre2":"10.43-r0","apk:zlib":"1.3.1-r2","apk:zstd-libs":"1.5.6-r1","github:miniflux/v2":"2.2.4","image:base":"3d9bc81e06db6829d63423bb674f142586ed9bd48613aab2849ada8f7b431a99","image:golang":"fb31d040911eec2e5068e1b3283d39c60b916549a52671b1aabe091488fb66d0"}

FROM reg.c5h.io/golang@sha256:fb31d040911eec2e5068e1b3283d39c60b916549a52671b1aabe091488fb66d0 AS build

ARG TAG="2.2.4"

RUN set -eux; \
    apk add --no-cache \
        brotli-libs=1.1.0-r2\
        c-ares=1.34.3-r0\
        ca-certificates-bundle=20241010-r0\
        git=2.47.1-r0\
        libcrypto3=3.3.2-r4\
        libcurl=8.11.1-r0\
        libexpat=2.6.4-r0\
        libidn2=2.3.7-r0\
        libpsl=0.21.5-r3\
        libssl3=3.3.2-r4\
        libunistring=1.2-r0\
        musl=1.2.5-r8\
        nghttp2-libs=1.64.0-r0\
        pcre2=10.43-r0\
        zlib=1.3.1-r2\
        zstd-libs=1.5.6-r1\
        ; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/miniflux/v2 /go/src/github.com/miniflux/v2; \
    cd /go/src/github.com/miniflux/v2; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main .; \
    go-licenses save ./... --save_path=/notices;

FROM reg.c5h.io/base@sha256:3d9bc81e06db6829d63423bb674f142586ed9bd48613aab2849ada8f7b431a99

COPY --from=build /go/src/github.com/miniflux/v2/main /miniflux
COPY --from=build /notices /notices
ENTRYPOINT ["/miniflux"]
