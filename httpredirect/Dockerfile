# Generated from https://github.com/csmith/dockerfiles/blob/master/httpredirect/Dockerfile.gotpl
# BOM: {"apk:brotli-libs":"1.0.9-r14","apk:busybox":"1.36.1-r0","apk:busybox-binsh":"1.36.1-r0","apk:ca-certificates":"20230506-r0","apk:git":"2.40.1-r0","apk:libcrypto3":"3.1.1-r1","apk:libcurl":"8.1.2-r0","apk:libexpat":"2.5.0-r1","apk:libidn2":"2.3.4-r1","apk:libssl3":"3.1.1-r1","apk:libunistring":"1.1-r1","apk:musl":"1.2.4-r0","apk:nghttp2-libs":"1.53.0-r0","apk:pcre2":"10.42-r1","apk:zlib":"1.2.13-r1","github:greboid/httpredirect":"v1.0.1","image:base":"33b47501f2e31be55c1af6b21082de3c3efaae9ec0062beb1d3255a9fd739874","image:golang":"36bd910612256f7ba9752fd4c69b3705e330ec58f879333871ce6edc82b85e22"}

FROM reg.c5h.io/golang@sha256:36bd910612256f7ba9752fd4c69b3705e330ec58f879333871ce6edc82b85e22 AS build

ARG TAG="v1.0.1"

RUN set -eux; \
    apk add --no-cache \
        brotli-libs=1.0.9-r14\
        busybox=1.36.1-r0\
        busybox-binsh=1.36.1-r0\
        ca-certificates=20230506-r0\
        git=2.40.1-r0\
        libcrypto3=3.1.1-r1\
        libcurl=8.1.2-r0\
        libexpat=2.5.0-r1\
        libidn2=2.3.4-r1\
        libssl3=3.1.1-r1\
        libunistring=1.1-r1\
        musl=1.2.4-r0\
        nghttp2-libs=1.53.0-r0\
        pcre2=10.42-r1\
        zlib=1.2.13-r1\
        ; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/greboid/httpredirect /go/src/github.com/greboid/httpredirect; \
    cd /go/src/github.com/greboid/httpredirect; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main .; \
    go-licenses save ./... --save_path=/notices;

FROM reg.c5h.io/base@sha256:33b47501f2e31be55c1af6b21082de3c3efaae9ec0062beb1d3255a9fd739874

COPY --from=build /go/src/github.com/greboid/httpredirect/main /httpredirect
COPY --from=build /notices /notices
EXPOSE 8080
ENTRYPOINT ["/httpredirect"]
