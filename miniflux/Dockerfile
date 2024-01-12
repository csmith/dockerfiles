# Generated from https://github.com/csmith/dockerfiles/blob/master/miniflux/Dockerfile.gotpl
# BOM: {"apk:brotli-libs":"1.1.0-r1","apk:busybox":"1.36.1-r15","apk:busybox-binsh":"1.36.1-r15","apk:c-ares":"1.24.0-r0","apk:ca-certificates":"20230506-r0","apk:git":"2.43.0-r0","apk:libcrypto3":"3.1.4-r3","apk:libcurl":"8.5.0-r0","apk:libexpat":"2.5.0-r2","apk:libidn2":"2.3.4-r4","apk:libssl3":"3.1.4-r3","apk:libunistring":"1.1-r2","apk:musl":"1.2.4_git20230717-r4","apk:nghttp2-libs":"1.58.0-r0","apk:pcre2":"10.42-r2","apk:zlib":"1.3-r2","github:miniflux/v2":"2.0.51","image:base":"cfa1d5f949b495a6e077bdddb1dbbc3d9a9d55036ab1267f4168a6a5c7e7fdce","image:golang":"0336d9c467f9e0a89228f5b2aea2c186437ea4eb0aca68e0b4e4f8aa61ca88bb"}

FROM reg.c5h.io/golang@sha256:0336d9c467f9e0a89228f5b2aea2c186437ea4eb0aca68e0b4e4f8aa61ca88bb AS build

ARG TAG="2.0.51"

RUN set -eux; \
    apk add --no-cache \
        brotli-libs=1.1.0-r1\
        busybox=1.36.1-r15\
        busybox-binsh=1.36.1-r15\
        c-ares=1.24.0-r0\
        ca-certificates=20230506-r0\
        git=2.43.0-r0\
        libcrypto3=3.1.4-r3\
        libcurl=8.5.0-r0\
        libexpat=2.5.0-r2\
        libidn2=2.3.4-r4\
        libssl3=3.1.4-r3\
        libunistring=1.1-r2\
        musl=1.2.4_git20230717-r4\
        nghttp2-libs=1.58.0-r0\
        pcre2=10.42-r2\
        zlib=1.3-r2\
        ; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/miniflux/v2 /go/src/github.com/miniflux/v2; \
    cd /go/src/github.com/miniflux/v2; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main .; \
    go-licenses save ./... --save_path=/notices;

FROM reg.c5h.io/base@sha256:cfa1d5f949b495a6e077bdddb1dbbc3d9a9d55036ab1267f4168a6a5c7e7fdce

COPY --from=build /go/src/github.com/miniflux/v2/main /miniflux
COPY --from=build /notices /notices
ENTRYPOINT ["/miniflux"]
