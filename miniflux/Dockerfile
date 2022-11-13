# Generated from https://github.com/csmith/dockerfiles/blob/master/miniflux/Dockerfile.gotpl
# BOM: {"apk:brotli-libs":"1.0.9-r6","apk:busybox":"1.35.0-r17","apk:ca-certificates":"20220614-r0","apk:expat":"2.5.0-r0","apk:git":"2.36.3-r0","apk:libcrypto1.1":"1.1.1s-r0","apk:libcurl":"7.83.1-r4","apk:libssl1.1":"1.1.1s-r0","apk:musl":"1.2.3-r2","apk:nghttp2-libs":"1.47.0-r0","apk:pcre2":"10.40-r0","apk:zlib":"1.2.12-r3","github:miniflux/v2":"2.0.39","image:base":"5d2c54fa31ef9b63b0499aabf0dd0091eee215ae556dd89f68189e7894cef381","image:golang":"d8fd3cc3b73ff5eb50c1e125e6fbf69775e75c4b881a89c035bc0a60372f3b3f"}

FROM reg.c5h.io/golang@sha256:d8fd3cc3b73ff5eb50c1e125e6fbf69775e75c4b881a89c035bc0a60372f3b3f AS build

ARG TAG="2.0.39"

RUN set -eux; \
    apk add --no-cache \
        brotli-libs=1.0.9-r6\
        busybox=1.35.0-r17\
        ca-certificates=20220614-r0\
        expat=2.5.0-r0\
        git=2.36.3-r0\
        libcrypto1.1=1.1.1s-r0\
        libcurl=7.83.1-r4\
        libssl1.1=1.1.1s-r0\
        musl=1.2.3-r2\
        nghttp2-libs=1.47.0-r0\
        pcre2=10.40-r0\
        zlib=1.2.12-r3\
        ; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/miniflux/v2 /go/src/github.com/miniflux/v2; \
    cd /go/src/github.com/miniflux/v2; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main .; \
    go-licenses save ./... --save_path=/notices;

FROM reg.c5h.io/base@sha256:5d2c54fa31ef9b63b0499aabf0dd0091eee215ae556dd89f68189e7894cef381

COPY --from=build /go/src/github.com/miniflux/v2/main /miniflux
COPY --from=build /notices /notices
ENTRYPOINT ["/miniflux"]
