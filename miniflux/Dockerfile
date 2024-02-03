# Generated from https://github.com/csmith/dockerfiles/blob/master/miniflux/Dockerfile.gotpl
# BOM: {"apk:brotli-libs":"1.1.0-r1","apk:busybox":"1.36.1-r15","apk:busybox-binsh":"1.36.1-r15","apk:c-ares":"1.24.0-r1","apk:ca-certificates":"20230506-r0","apk:git":"2.43.0-r0","apk:libcrypto3":"3.1.4-r5","apk:libcurl":"8.5.0-r0","apk:libexpat":"2.5.0-r2","apk:libidn2":"2.3.4-r4","apk:libssl3":"3.1.4-r5","apk:libunistring":"1.1-r2","apk:musl":"1.2.4_git20230717-r4","apk:nghttp2-libs":"1.58.0-r0","apk:pcre2":"10.42-r2","apk:zlib":"1.3.1-r0","github:miniflux/v2":"2.0.51","image:base":"fd06c9da248d637dfa37579addc089c993016cf3e21ee1142b5aea23095ad4f9","image:golang":"45a1591f8121fd2361f10b3b4960542e36ca8ab1d3787d7b5adc0fcbafc9b7fc"}

FROM reg.c5h.io/golang@sha256:45a1591f8121fd2361f10b3b4960542e36ca8ab1d3787d7b5adc0fcbafc9b7fc AS build

ARG TAG="2.0.51"

RUN set -eux; \
    apk add --no-cache \
        brotli-libs=1.1.0-r1\
        busybox=1.36.1-r15\
        busybox-binsh=1.36.1-r15\
        c-ares=1.24.0-r1\
        ca-certificates=20230506-r0\
        git=2.43.0-r0\
        libcrypto3=3.1.4-r5\
        libcurl=8.5.0-r0\
        libexpat=2.5.0-r2\
        libidn2=2.3.4-r4\
        libssl3=3.1.4-r5\
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

FROM reg.c5h.io/base@sha256:fd06c9da248d637dfa37579addc089c993016cf3e21ee1142b5aea23095ad4f9

COPY --from=build /go/src/github.com/miniflux/v2/main /miniflux
COPY --from=build /notices /notices
ENTRYPOINT ["/miniflux"]
