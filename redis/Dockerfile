# Generated from https://github.com/csmith/dockerfiles/blob/master/redis/Dockerfile.gotpl
# BOM: {"apk:binutils":"2.43.1-r1","apk:brotli-libs":"1.1.0-r2","apk:build-base":"0.5-r3","apk:busybox":"1.37.0-r12","apk:busybox-binsh":"1.37.0-r12","apk:c-ares":"1.34.3-r0","apk:ca-certificates":"20241121-r1","apk:ca-certificates-bundle":"20241121-r1","apk:file":"5.46-r2","apk:fortify-headers":"1.1-r5","apk:g++":"14.2.0-r4","apk:gcc":"14.2.0-r4","apk:git":"2.47.2-r0","apk:gmp":"6.3.0-r2","apk:isl25":"0.25-r2","apk:jansson":"2.14-r4","apk:libatomic":"14.2.0-r4","apk:libcrypto3":"3.3.2-r6","apk:libcurl":"8.11.1-r1","apk:libexpat":"2.6.4-r0","apk:libgcc":"14.2.0-r4","apk:libgomp":"14.2.0-r4","apk:libidn2":"2.3.7-r0","apk:libmagic":"5.46-r2","apk:libpcre16":"8.45-r3","apk:libpcre32":"8.45-r3","apk:libpcrecpp":"8.45-r3","apk:libpsl":"0.21.5-r3","apk:libssl3":"3.3.2-r6","apk:libstdc++":"14.2.0-r4","apk:libstdc++-dev":"14.2.0-r4","apk:libunistring":"1.2-r0","apk:linux-headers":"6.6-r1","apk:make":"4.4.1-r2","apk:mpc1":"1.3.1-r1","apk:mpfr4":"4.2.1-r0","apk:musl":"1.2.5-r8","apk:musl-dev":"1.2.5-r8","apk:nghttp2-libs":"1.64.0-r0","apk:openssl-dev":"3.3.2-r6","apk:openssl-libs-static":"3.3.2-r6","apk:patch":"2.7.6-r10","apk:pcre":"8.45-r3","apk:pcre-dev":"8.45-r3","apk:pcre2":"10.43-r0","apk:pkgconf":"2.3.0-r0","apk:tzdata":"2024b-r1","apk:zlib":"1.3.1-r2","apk:zlib-dev":"1.3.1-r2","apk:zlib-static":"1.3.1-r2","apk:zstd-libs":"1.5.6-r2","github:redis/redis":"7.4.2","image:alpine":"0be8812dc5bc37867ac650dde5331475f4fabd765e7f2f37f3166338e36cf6de","image:base":"27b64203bd464e43e2314afe482fa80c5c8a3258cada24e4c97b40ab9be1bbd4"}

FROM reg.c5h.io/alpine@sha256:0be8812dc5bc37867ac650dde5331475f4fabd765e7f2f37f3166338e36cf6de AS build

RUN set -eux; \
    apk add --no-cache \
        brotli-libs=1.1.0-r2 \
        c-ares=1.34.3-r0 \
        ca-certificates-bundle=20241121-r1 \
        git=2.47.2-r0 \
        libcrypto3=3.3.2-r6 \
        libcurl=8.11.1-r1 \
        libexpat=2.6.4-r0 \
        libidn2=2.3.7-r0 \
        libpsl=0.21.5-r3 \
        libssl3=3.3.2-r6 \
        libunistring=1.2-r0 \
        musl=1.2.5-r8 \
        nghttp2-libs=1.64.0-r0 \
        pcre2=10.43-r0 \
        zlib=1.3.1-r2 \
        zstd-libs=1.5.6-r2 \
        binutils=2.43.1-r1 \
        build-base=0.5-r3 \
        file=5.46-r2 \
        fortify-headers=1.1-r5 \
        g++=14.2.0-r4 \
        gcc=14.2.0-r4 \
        gmp=6.3.0-r2 \
        isl25=0.25-r2 \
        jansson=2.14-r4 \
        libatomic=14.2.0-r4 \
        libgcc=14.2.0-r4 \
        libgomp=14.2.0-r4 \
        libmagic=5.46-r2 \
        libstdc++=14.2.0-r4 \
        libstdc++-dev=14.2.0-r4 \
        make=4.4.1-r2 \
        mpc1=1.3.1-r1 \
        mpfr4=4.2.1-r0 \
        musl=1.2.5-r8 \
        musl-dev=1.2.5-r8 \
        patch=2.7.6-r10 \
        zlib=1.3.1-r2 \
        zstd-libs=1.5.6-r2 \
        tzdata=2024b-r1 \
        busybox=1.37.0-r12 \
        busybox-binsh=1.37.0-r12 \
        ca-certificates=20241121-r1 \
        libcrypto3=3.3.2-r6 \
        musl=1.2.5-r8 \
        linux-headers=6.6-r1 \
        libgcc=14.2.0-r4 \
        libpcre16=8.45-r3 \
        libpcre32=8.45-r3 \
        libpcrecpp=8.45-r3 \
        libstdc++=14.2.0-r4 \
        musl=1.2.5-r8 \
        pcre=8.45-r3 \
        pcre-dev=8.45-r3 \
        pkgconf=2.3.0-r0 \
        musl=1.2.5-r8 \
        musl-dev=1.2.5-r8 \
        libcrypto3=3.3.2-r6 \
        libssl3=3.3.2-r6 \
        musl=1.2.5-r8 \
        openssl-dev=3.3.2-r6 \
        pkgconf=2.3.0-r0 \
        openssl-libs-static=3.3.2-r6 \
        musl=1.2.5-r8 \
        pkgconf=2.3.0-r0 \
        zlib=1.3.1-r2 \
        zlib-dev=1.3.1-r2 \
        zlib-static=1.3.1-r2 \
    ;

RUN set -eux; \
    git clone --depth=1 -b 7.4.2 --single-branch https://github.com/redis/redis /src/redis

WORKDIR /src/redis


ENV LDFLAGS="-static"

COPY redis.conf /redis.conf

RUN set -eux; \
     make distclean; \
     make; \
     mkdir /data; \
     strip /src/redis/src/redis-server; \
     touch --date=@0 /src/redis/src/redis-server /data /redis.conf

FROM reg.c5h.io/base@sha256:27b64203bd464e43e2314afe482fa80c5c8a3258cada24e4c97b40ab9be1bbd4

COPY --from=build --chown=65532:65532 /redis.conf /home/nonroot/redis.conf
COPY --from=build --chown=65532:65532 /data /home/nonroot/database
COPY --from=build --chown=65532:65532 /src/redis/src/redis-server /home/nonroot/redis

ENTRYPOINT ["/home/nonroot/redis"]

EXPOSE 6379
VOLUME "/home/nonroot/database"
CMD ["/home/nonroot/redis.conf"]
