# Generated from https://github.com/csmith/dockerfiles/blob/master/redis/Dockerfile.gotpl
# BOM: {"apk:binutils":"2.44-r3","apk:brotli-libs":"1.1.0-r2","apk:build-base":"0.5-r3","apk:busybox":"1.37.0-r19","apk:busybox-binsh":"1.37.0-r19","apk:c-ares":"1.34.5-r0","apk:ca-certificates":"20250911-r0","apk:ca-certificates-bundle":"20250911-r0","apk:file":"5.46-r2","apk:fortify-headers":"1.1-r5","apk:g++":"14.2.0-r6","apk:gcc":"14.2.0-r6","apk:git":"2.49.1-r0","apk:gmp":"6.3.0-r3","apk:isl25":"0.25-r2","apk:jansson":"2.14.1-r0","apk:libatomic":"14.2.0-r6","apk:libcrypto3":"3.5.4-r0","apk:libcurl":"8.14.1-r2","apk:libexpat":"2.7.3-r0","apk:libgcc":"14.2.0-r6","apk:libgomp":"14.2.0-r6","apk:libidn2":"2.3.7-r0","apk:libmagic":"5.46-r2","apk:libpcre16":"8.45-r4","apk:libpcre32":"8.45-r4","apk:libpcrecpp":"8.45-r4","apk:libpsl":"0.21.5-r3","apk:libssl3":"3.5.4-r0","apk:libstdc++":"14.2.0-r6","apk:libstdc++-dev":"14.2.0-r6","apk:libunistring":"1.3-r0","apk:linux-headers":"6.14.2-r0","apk:make":"4.4.1-r3","apk:mpc1":"1.3.1-r1","apk:mpfr4":"4.2.1_p1-r0","apk:musl":"1.2.5-r10","apk:musl-dev":"1.2.5-r10","apk:nghttp2-libs":"1.65.0-r0","apk:openssl-dev":"3.5.4-r0","apk:openssl-libs-static":"3.5.4-r0","apk:patch":"2.8-r0","apk:pcre":"8.45-r4","apk:pcre-dev":"8.45-r4","apk:pcre2":"10.46-r0","apk:pkgconf":"2.4.3-r0","apk:tzdata":"2025b-r0","apk:zlib":"1.3.1-r2","apk:zlib-dev":"1.3.1-r2","apk:zlib-static":"1.3.1-r2","apk:zstd-libs":"1.5.7-r0","github:redis/redis":"8.2.2","image:alpine":"7870f9b13b86c48d20a50151dd813a32dff5cca8c3f71a968dec5ec2cd4880a2","image:base":"a09ba1c5a9b2ec41293a8fa0cef935f6ab81978fcbb62271f490ebb07bbf6024"}

FROM reg.c5h.io/alpine@sha256:7870f9b13b86c48d20a50151dd813a32dff5cca8c3f71a968dec5ec2cd4880a2 AS build

RUN set -eux; \
    apk add --no-cache \
        brotli-libs=1.1.0-r2 \
        c-ares=1.34.5-r0 \
        ca-certificates-bundle=20250911-r0 \
        git=2.49.1-r0 \
        libcrypto3=3.5.4-r0 \
        libcurl=8.14.1-r2 \
        libexpat=2.7.3-r0 \
        libidn2=2.3.7-r0 \
        libpsl=0.21.5-r3 \
        libssl3=3.5.4-r0 \
        libunistring=1.3-r0 \
        musl=1.2.5-r10 \
        nghttp2-libs=1.65.0-r0 \
        pcre2=10.46-r0 \
        zlib=1.3.1-r2 \
        zstd-libs=1.5.7-r0 \
        binutils=2.44-r3 \
        build-base=0.5-r3 \
        file=5.46-r2 \
        fortify-headers=1.1-r5 \
        g++=14.2.0-r6 \
        gcc=14.2.0-r6 \
        gmp=6.3.0-r3 \
        isl25=0.25-r2 \
        jansson=2.14.1-r0 \
        libatomic=14.2.0-r6 \
        libgcc=14.2.0-r6 \
        libgomp=14.2.0-r6 \
        libmagic=5.46-r2 \
        libstdc++=14.2.0-r6 \
        libstdc++-dev=14.2.0-r6 \
        make=4.4.1-r3 \
        mpc1=1.3.1-r1 \
        mpfr4=4.2.1_p1-r0 \
        musl=1.2.5-r10 \
        musl-dev=1.2.5-r10 \
        patch=2.8-r0 \
        zlib=1.3.1-r2 \
        zstd-libs=1.5.7-r0 \
        tzdata=2025b-r0 \
        busybox=1.37.0-r19 \
        busybox-binsh=1.37.0-r19 \
        ca-certificates=20250911-r0 \
        libcrypto3=3.5.4-r0 \
        musl=1.2.5-r10 \
        linux-headers=6.14.2-r0 \
        libgcc=14.2.0-r6 \
        libpcre16=8.45-r4 \
        libpcre32=8.45-r4 \
        libpcrecpp=8.45-r4 \
        libstdc++=14.2.0-r6 \
        musl=1.2.5-r10 \
        pcre=8.45-r4 \
        pcre-dev=8.45-r4 \
        pkgconf=2.4.3-r0 \
        musl=1.2.5-r10 \
        musl-dev=1.2.5-r10 \
        libcrypto3=3.5.4-r0 \
        libssl3=3.5.4-r0 \
        musl=1.2.5-r10 \
        openssl-dev=3.5.4-r0 \
        pkgconf=2.4.3-r0 \
        openssl-libs-static=3.5.4-r0 \
        musl=1.2.5-r10 \
        pkgconf=2.4.3-r0 \
        zlib=1.3.1-r2 \
        zlib-dev=1.3.1-r2 \
        zlib-static=1.3.1-r2 \
    ;

RUN set -eux; \
    git clone --depth=1 -b 8.2.2 --single-branch https://github.com/redis/redis /src/redis

WORKDIR /src/redis


ENV LDFLAGS="-static"

COPY redis.conf /redis.conf

RUN set -eux; \
     make distclean; \
     make; \
     mkdir /data; \
     strip /src/redis/src/redis-server; \
     touch --date=@0 /src/redis/src/redis-server /data /redis.conf

FROM reg.c5h.io/base@sha256:a09ba1c5a9b2ec41293a8fa0cef935f6ab81978fcbb62271f490ebb07bbf6024

COPY --from=build --chown=65532:65532 /redis.conf /home/nonroot/redis.conf
COPY --from=build --chown=65532:65532 /data /home/nonroot/database
COPY --from=build --chown=65532:65532 /src/redis/src/redis-server /home/nonroot/redis

ENTRYPOINT ["/home/nonroot/redis"]

EXPOSE 6379
VOLUME "/home/nonroot/database"
CMD ["/home/nonroot/redis.conf"]
