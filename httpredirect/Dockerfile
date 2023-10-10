# Generated from https://github.com/csmith/dockerfiles/blob/master/httpredirect/Dockerfile.gotpl
# BOM: {"apk:brotli-libs":"1.0.9-r14","apk:busybox":"1.36.1-r4","apk:busybox-binsh":"1.36.1-r4","apk:ca-certificates":"20230506-r0","apk:git":"2.40.1-r0","apk:libcrypto3":"3.1.3-r0","apk:libcurl":"8.3.0-r0","apk:libexpat":"2.5.0-r1","apk:libidn2":"2.3.4-r1","apk:libssl3":"3.1.3-r0","apk:libunistring":"1.1-r1","apk:musl":"1.2.4-r2","apk:nghttp2-libs":"1.55.1-r0","apk:pcre2":"10.42-r1","apk:zlib":"1.2.13-r1","github:greboid/httpredirect":"v1.0.2","image:base":"7442771fc9e4d70fbf8f9d592e28b9ec7b9b81af30ddb87cbaaba90f2cc0064f","image:golang":"5dd959be9e0d267512c13b608c0695e3b857b3c9e2c1d76e8e7a5dc2c6a3fa8e"}

FROM reg.c5h.io/golang@sha256:5dd959be9e0d267512c13b608c0695e3b857b3c9e2c1d76e8e7a5dc2c6a3fa8e AS build

ARG TAG="v1.0.2"

RUN set -eux; \
    apk add --no-cache \
        brotli-libs=1.0.9-r14\
        busybox=1.36.1-r4\
        busybox-binsh=1.36.1-r4\
        ca-certificates=20230506-r0\
        git=2.40.1-r0\
        libcrypto3=3.1.3-r0\
        libcurl=8.3.0-r0\
        libexpat=2.5.0-r1\
        libidn2=2.3.4-r1\
        libssl3=3.1.3-r0\
        libunistring=1.1-r1\
        musl=1.2.4-r2\
        nghttp2-libs=1.55.1-r0\
        pcre2=10.42-r1\
        zlib=1.2.13-r1\
        ; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/greboid/httpredirect /go/src/github.com/greboid/httpredirect; \
    cd /go/src/github.com/greboid/httpredirect; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main .; \
    go-licenses save ./... --save_path=/notices;

FROM reg.c5h.io/base@sha256:7442771fc9e4d70fbf8f9d592e28b9ec7b9b81af30ddb87cbaaba90f2cc0064f

COPY --from=build /go/src/github.com/greboid/httpredirect/main /httpredirect
COPY --from=build /notices /notices
EXPOSE 8080
ENTRYPOINT ["/httpredirect"]
