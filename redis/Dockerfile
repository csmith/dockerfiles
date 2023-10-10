# Generated from https://github.com/csmith/dockerfiles/blob/master/redis/Dockerfile.gotpl
# BOM: {"apk:binutils":"2.40-r7","apk:brotli-libs":"1.0.9-r14","apk:build-base":"0.5-r3","apk:busybox":"1.36.1-r4","apk:busybox-binsh":"1.36.1-r4","apk:ca-certificates":"20230506-r0","apk:file":"5.45-r0","apk:fortify-headers":"1.1-r3","apk:g++":"12.2.1_git20220924-r10","apk:gcc":"12.2.1_git20220924-r10","apk:git":"2.40.1-r0","apk:gmp":"6.2.1-r3","apk:isl25":"0.25-r2","apk:libatomic":"12.2.1_git20220924-r10","apk:libc-dev":"0.7.2-r5","apk:libcrypto3":"3.1.3-r0","apk:libcurl":"8.3.0-r0","apk:libexpat":"2.5.0-r1","apk:libgcc":"12.2.1_git20220924-r10","apk:libgomp":"12.2.1_git20220924-r10","apk:libidn2":"2.3.4-r1","apk:libmagic":"5.45-r0","apk:libpcre16":"8.45-r3","apk:libpcre32":"8.45-r3","apk:libpcrecpp":"8.45-r3","apk:libssl3":"3.1.3-r0","apk:libstdc++":"12.2.1_git20220924-r10","apk:libstdc++-dev":"12.2.1_git20220924-r10","apk:libunistring":"1.1-r1","apk:linux-headers":"6.3-r0","apk:make":"4.4.1-r1","apk:mpc1":"1.3.1-r1","apk:mpfr4":"4.2.0_p12-r0","apk:musl":"1.2.4-r2","apk:musl-dev":"1.2.4-r2","apk:nghttp2-libs":"1.55.1-r0","apk:openssl-dev":"3.1.3-r0","apk:openssl-libs-static":"3.1.3-r0","apk:patch":"2.7.6-r10","apk:pcre":"8.45-r3","apk:pcre-dev":"8.45-r3","apk:pcre2":"10.42-r1","apk:pkgconf":"1.9.5-r0","apk:tzdata":"2023c-r1","apk:zlib":"1.2.13-r1","apk:zlib-dev":"1.2.13-r1","apk:zlib-static":"1.2.13-r1","apk:zstd-libs":"1.5.5-r4","github:redis/redis":"7.2.1","image:alpine":"2fd95023d8b736c8d820fe158e5e388f24260350dd63516eb1f36f34f73fbc15","image:base":"7442771fc9e4d70fbf8f9d592e28b9ec7b9b81af30ddb87cbaaba90f2cc0064f"}

FROM reg.c5h.io/alpine@sha256:2fd95023d8b736c8d820fe158e5e388f24260350dd63516eb1f36f34f73fbc15 AS build

ARG TAG="7.2.1"

ENV LDFLAGS="-static"

COPY redis.conf /redis.conf

RUN set -eux; \
     apk add --no-cache --virtual .build-deps \
        binutils=2.40-r7 \
        brotli-libs=1.0.9-r14 \
        build-base=0.5-r3 \
        busybox=1.36.1-r4 \
        busybox-binsh=1.36.1-r4 \
        ca-certificates=20230506-r0 \
        file=5.45-r0 \
        fortify-headers=1.1-r3 \
        g++=12.2.1_git20220924-r10 \
        gcc=12.2.1_git20220924-r10 \
        git=2.40.1-r0 \
        gmp=6.2.1-r3 \
        isl25=0.25-r2 \
        libatomic=12.2.1_git20220924-r10 \
        libc-dev=0.7.2-r5 \
        libcrypto3=3.1.3-r0 \
        libcurl=8.3.0-r0 \
        libexpat=2.5.0-r1 \
        libgcc=12.2.1_git20220924-r10 \
        libgomp=12.2.1_git20220924-r10 \
        libidn2=2.3.4-r1 \
        libmagic=5.45-r0 \
        libpcre16=8.45-r3 \
        libpcre32=8.45-r3 \
        libpcrecpp=8.45-r3 \
        libssl3=3.1.3-r0 \
        libstdc++=12.2.1_git20220924-r10 \
        libstdc++-dev=12.2.1_git20220924-r10 \
        libunistring=1.1-r1 \
        linux-headers=6.3-r0 \
        make=4.4.1-r1 \
        mpc1=1.3.1-r1 \
        mpfr4=4.2.0_p12-r0 \
        musl=1.2.4-r2 \
        musl-dev=1.2.4-r2 \
        nghttp2-libs=1.55.1-r0 \
        openssl-dev=3.1.3-r0 \
        openssl-libs-static=3.1.3-r0 \
        patch=2.7.6-r10 \
        pcre=8.45-r3 \
        pcre-dev=8.45-r3 \
        pcre2=10.42-r1 \
        pkgconf=1.9.5-r0 \
        tzdata=2023c-r1 \
        zlib=1.2.13-r1 \
        zlib-dev=1.2.13-r1 \
        zlib-static=1.2.13-r1 \
        zstd-libs=1.5.5-r4 \
        ; \
     git clone --depth=1 -b $TAG --single-branch https://github.com/redis/redis /src/redis; \
     cd /src/redis; \
     make distclean; \
     make; \
     mkdir /data; \
     strip /src/redis/src/redis-server; \
     touch --date=@0 /src/redis/src/redis-server /data /redis.conf

FROM reg.c5h.io/base@sha256:7442771fc9e4d70fbf8f9d592e28b9ec7b9b81af30ddb87cbaaba90f2cc0064f

COPY --from=build --chown=65532:65532 /redis.conf /home/nonroot/redis.conf
COPY --from=build --chown=65532:65532 /data /home/nonroot/database
COPY --from=build --chown=65532:65532 /src/redis/src/redis-server /home/nonroot/redis

ENTRYPOINT ["/home/nonroot/redis"]

EXPOSE 6379
VOLUME "/home/nonroot/database"
CMD ["/home/nonroot/redis.conf"]
