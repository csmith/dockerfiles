# Generated from https://github.com/csmith/dockerfiles/blob/master/httpredirect/Dockerfile.gotpl
# BOM: {"apk:brotli-libs":"1.0.9-r9","apk:busybox":"1.35.0-r29","apk:busybox-binsh":"1.35.0-r29","apk:ca-certificates":"20230506-r0","apk:git":"2.38.5-r0","apk:libcrypto3":"3.0.8-r4","apk:libcurl":"8.0.1-r0","apk:libexpat":"2.5.0-r0","apk:libssl3":"3.0.8-r4","apk:musl":"1.2.3-r4","apk:nghttp2-libs":"1.51.0-r0","apk:pcre2":"10.42-r0","apk:zlib":"1.2.13-r0","github:greboid/httpredirect":"v1.0.1","image:base":"06668724f8daa1866a72d07c44d9c0fe5c42bfd277d47fecc8e4ed037a5d0e84","image:golang":"8cca8e81ef263f66855ba7ad711e884e03e7670276e1cd2011d8f0f3f0eec7e0"}

FROM reg.c5h.io/golang@sha256:8cca8e81ef263f66855ba7ad711e884e03e7670276e1cd2011d8f0f3f0eec7e0 AS build

ARG TAG="v1.0.1"

RUN set -eux; \
    apk add --no-cache \
        brotli-libs=1.0.9-r9\
        busybox=1.35.0-r29\
        busybox-binsh=1.35.0-r29\
        ca-certificates=20230506-r0\
        git=2.38.5-r0\
        libcrypto3=3.0.8-r4\
        libcurl=8.0.1-r0\
        libexpat=2.5.0-r0\
        libssl3=3.0.8-r4\
        musl=1.2.3-r4\
        nghttp2-libs=1.51.0-r0\
        pcre2=10.42-r0\
        zlib=1.2.13-r0\
        ; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/greboid/httpredirect /go/src/github.com/greboid/httpredirect; \
    cd /go/src/github.com/greboid/httpredirect; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main .; \
    go-licenses save ./... --save_path=/notices;

FROM reg.c5h.io/base@sha256:06668724f8daa1866a72d07c44d9c0fe5c42bfd277d47fecc8e4ed037a5d0e84

COPY --from=build /go/src/github.com/greboid/httpredirect/main /httpredirect
COPY --from=build /notices /notices
EXPOSE 8080
ENTRYPOINT ["/httpredirect"]
