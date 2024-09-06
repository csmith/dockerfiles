# Generated from https://github.com/csmith/dockerfiles/blob/master/httpredirect/Dockerfile.gotpl
# BOM: {"apk:brotli-libs":"1.1.0-r2","apk:busybox":"1.36.1-r29","apk:busybox-binsh":"1.36.1-r29","apk:c-ares":"1.28.1-r0","apk:ca-certificates":"20240705-r0","apk:git":"2.45.2-r0","apk:libcrypto3":"3.3.2-r0","apk:libcurl":"8.9.1-r1","apk:libexpat":"2.6.3-r0","apk:libidn2":"2.3.7-r0","apk:libpsl":"0.21.5-r1","apk:libssl3":"3.3.2-r0","apk:libunistring":"1.2-r0","apk:musl":"1.2.5-r0","apk:nghttp2-libs":"1.62.1-r0","apk:pcre2":"10.43-r0","apk:zlib":"1.3.1-r1","apk:zstd-libs":"1.5.6-r0","github:greboid/httpredirect":"v1.0.3","image:base":"b92ffcd70a3cea70972a600a917a409451afe2732844d09e3afe212ad8e041ff","image:golang":"9f947b59fef44d1d38f356e3006b0f22580b969828f660d1c04890965cc3962f"}

FROM reg.c5h.io/golang@sha256:9f947b59fef44d1d38f356e3006b0f22580b969828f660d1c04890965cc3962f AS build

ARG TAG="v1.0.3"

RUN set -eux; \
    apk add --no-cache \
        brotli-libs=1.1.0-r2\
        busybox=1.36.1-r29\
        busybox-binsh=1.36.1-r29\
        c-ares=1.28.1-r0\
        ca-certificates=20240705-r0\
        git=2.45.2-r0\
        libcrypto3=3.3.2-r0\
        libcurl=8.9.1-r1\
        libexpat=2.6.3-r0\
        libidn2=2.3.7-r0\
        libpsl=0.21.5-r1\
        libssl3=3.3.2-r0\
        libunistring=1.2-r0\
        musl=1.2.5-r0\
        nghttp2-libs=1.62.1-r0\
        pcre2=10.43-r0\
        zlib=1.3.1-r1\
        zstd-libs=1.5.6-r0\
        ; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/greboid/httpredirect /go/src/github.com/greboid/httpredirect; \
    cd /go/src/github.com/greboid/httpredirect; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main .; \
    go-licenses save ./... --save_path=/notices;

FROM reg.c5h.io/base@sha256:b92ffcd70a3cea70972a600a917a409451afe2732844d09e3afe212ad8e041ff

COPY --from=build /go/src/github.com/greboid/httpredirect/main /httpredirect
COPY --from=build /notices /notices
EXPOSE 8080
ENTRYPOINT ["/httpredirect"]
