# Generated from https://github.com/csmith/dockerfiles/blob/master/miniflux/Dockerfile.gotpl
# BOM: {"apk:brotli-libs":"1.0.9-r6","apk:busybox":"1.35.0-r14","apk:ca-certificates":"20211220-r0","apk:expat":"2.4.8-r0","apk:git":"2.36.1-r0","apk:libcrypto1.1":"1.1.1o-r0","apk:libcurl":"7.83.1-r1","apk:libssl1.1":"1.1.1o-r0","apk:musl":"1.2.3-r0","apk:nghttp2-libs":"1.47.0-r0","apk:pcre2":"10.39-r0","apk:zlib":"1.2.12-r1","github:miniflux/v2":"2.0.37","image:base":"bffee9632ec79071de34e1a4a4f6c4563a71ad26d5923062bd9d0f9016f72beb","image:golang":"c9df3e62448ba6e02f7597fa8e8ca8e6df40ad1dd5b9e5d1cb6f59984261e632"}

FROM reg.c5h.io/golang@sha256:c9df3e62448ba6e02f7597fa8e8ca8e6df40ad1dd5b9e5d1cb6f59984261e632 AS build

ARG TAG="2.0.37"

RUN set -eux; \
    apk add --no-cache \
        brotli-libs=1.0.9-r6\
        busybox=1.35.0-r14\
        ca-certificates=20211220-r0\
        expat=2.4.8-r0\
        git=2.36.1-r0\
        libcrypto1.1=1.1.1o-r0\
        libcurl=7.83.1-r1\
        libssl1.1=1.1.1o-r0\
        musl=1.2.3-r0\
        nghttp2-libs=1.47.0-r0\
        pcre2=10.39-r0\
        zlib=1.2.12-r1\
        ; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/miniflux/v2 /go/src/github.com/miniflux/v2; \
    cd /go/src/github.com/miniflux/v2; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main .; \
    go-licenses save ./... --save_path=/notices;

FROM reg.c5h.io/base@sha256:bffee9632ec79071de34e1a4a4f6c4563a71ad26d5923062bd9d0f9016f72beb

COPY --from=build /go/src/github.com/miniflux/v2/main /miniflux
COPY --from=build /notices /notices
ENTRYPOINT ["/miniflux"]
