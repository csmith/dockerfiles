# Generated from https://github.com/csmith/dockerfiles/blob/master/miniflux/Dockerfile.gotpl
# BOM: {"apk:brotli-libs":"1.0.9-r14","apk:busybox":"1.36.0-r9","apk:busybox-binsh":"1.36.0-r9","apk:ca-certificates":"20230506-r0","apk:git":"2.40.1-r0","apk:libcrypto3":"3.1.0-r4","apk:libcurl":"8.0.1-r2","apk:libexpat":"2.5.0-r1","apk:libidn2":"2.3.4-r1","apk:libssl3":"3.1.0-r4","apk:libunistring":"1.1-r1","apk:musl":"1.2.4-r0","apk:nghttp2-libs":"1.53.0-r0","apk:pcre2":"10.42-r1","apk:zlib":"1.2.13-r1","github:miniflux/v2":"2.0.44","image:base":"33b47501f2e31be55c1af6b21082de3c3efaae9ec0062beb1d3255a9fd739874","image:golang":"4c46b35f960d889b4b388a6501c507357c5f16948eb3dc73411ac85dbd72520e"}

FROM reg.c5h.io/golang@sha256:4c46b35f960d889b4b388a6501c507357c5f16948eb3dc73411ac85dbd72520e AS build

ARG TAG="2.0.44"

RUN set -eux; \
    apk add --no-cache \
        brotli-libs=1.0.9-r14\
        busybox=1.36.0-r9\
        busybox-binsh=1.36.0-r9\
        ca-certificates=20230506-r0\
        git=2.40.1-r0\
        libcrypto3=3.1.0-r4\
        libcurl=8.0.1-r2\
        libexpat=2.5.0-r1\
        libidn2=2.3.4-r1\
        libssl3=3.1.0-r4\
        libunistring=1.1-r1\
        musl=1.2.4-r0\
        nghttp2-libs=1.53.0-r0\
        pcre2=10.42-r1\
        zlib=1.2.13-r1\
        ; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/miniflux/v2 /go/src/github.com/miniflux/v2; \
    cd /go/src/github.com/miniflux/v2; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main .; \
    go-licenses save ./... --save_path=/notices;

FROM reg.c5h.io/base@sha256:33b47501f2e31be55c1af6b21082de3c3efaae9ec0062beb1d3255a9fd739874

COPY --from=build /go/src/github.com/miniflux/v2/main /miniflux
COPY --from=build /notices /notices
ENTRYPOINT ["/miniflux"]
