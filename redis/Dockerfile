# Generated from https://github.com/csmith/dockerfiles/blob/master/redis/Dockerfile.gotpl
# BOM: {"apk:binutils":"2.42-r0","apk:brotli-libs":"1.1.0-r2","apk:build-base":"0.5-r3","apk:busybox":"1.36.1-r29","apk:busybox-binsh":"1.36.1-r29","apk:c-ares":"1.33.1-r0","apk:ca-certificates":"20240705-r0","apk:file":"5.45-r1","apk:fortify-headers":"1.1-r3","apk:g++":"13.2.1_git20240309-r0","apk:gcc":"13.2.1_git20240309-r0","apk:git":"2.45.2-r0","apk:gmp":"6.3.0-r1","apk:isl25":"0.25-r2","apk:jansson":"2.14-r4","apk:libatomic":"13.2.1_git20240309-r0","apk:libcrypto3":"3.3.2-r0","apk:libcurl":"8.9.1-r2","apk:libexpat":"2.6.3-r0","apk:libgcc":"13.2.1_git20240309-r0","apk:libgomp":"13.2.1_git20240309-r0","apk:libidn2":"2.3.7-r0","apk:libmagic":"5.45-r1","apk:libpcre16":"8.45-r3","apk:libpcre32":"8.45-r3","apk:libpcrecpp":"8.45-r3","apk:libpsl":"0.21.5-r1","apk:libssl3":"3.3.2-r0","apk:libstdc++":"13.2.1_git20240309-r0","apk:libstdc++-dev":"13.2.1_git20240309-r0","apk:libunistring":"1.2-r0","apk:linux-headers":"6.6-r0","apk:make":"4.4.1-r2","apk:mpc1":"1.3.1-r1","apk:mpfr4":"4.2.1-r0","apk:musl":"1.2.5-r0","apk:musl-dev":"1.2.5-r0","apk:nghttp2-libs":"1.62.1-r0","apk:openssl-dev":"3.3.2-r0","apk:openssl-libs-static":"3.3.2-r0","apk:patch":"2.7.6-r10","apk:pcre":"8.45-r3","apk:pcre-dev":"8.45-r3","apk:pcre2":"10.43-r0","apk:pkgconf":"2.2.0-r0","apk:tzdata":"2024b-r0","apk:zlib":"1.3.1-r1","apk:zlib-dev":"1.3.1-r1","apk:zlib-static":"1.3.1-r1","apk:zstd-libs":"1.5.6-r0","github:redis/redis":"7.4.0","image:alpine":"2e8d81e98838573646bf82865b487cee723829f133bc3e599d5ac3e08f682e5a","image:base":"fa9cdd83581affde03cb2234170235ff287588e4da3786133d72846838210621"}

FROM reg.c5h.io/alpine@sha256:2e8d81e98838573646bf82865b487cee723829f133bc3e599d5ac3e08f682e5a AS build

ARG TAG="7.4.0"

ENV LDFLAGS="-static"

COPY redis.conf /redis.conf

RUN set -eux; \
     apk add --no-cache --virtual .build-deps \
        binutils=2.42-r0 \
        brotli-libs=1.1.0-r2 \
        build-base=0.5-r3 \
        busybox=1.36.1-r29 \
        busybox-binsh=1.36.1-r29 \
        c-ares=1.33.1-r0 \
        ca-certificates=20240705-r0 \
        file=5.45-r1 \
        fortify-headers=1.1-r3 \
        g++=13.2.1_git20240309-r0 \
        gcc=13.2.1_git20240309-r0 \
        git=2.45.2-r0 \
        gmp=6.3.0-r1 \
        isl25=0.25-r2 \
        jansson=2.14-r4 \
        libatomic=13.2.1_git20240309-r0 \
        libcrypto3=3.3.2-r0 \
        libcurl=8.9.1-r2 \
        libexpat=2.6.3-r0 \
        libgcc=13.2.1_git20240309-r0 \
        libgomp=13.2.1_git20240309-r0 \
        libidn2=2.3.7-r0 \
        libmagic=5.45-r1 \
        libpcre16=8.45-r3 \
        libpcre32=8.45-r3 \
        libpcrecpp=8.45-r3 \
        libpsl=0.21.5-r1 \
        libssl3=3.3.2-r0 \
        libstdc++=13.2.1_git20240309-r0 \
        libstdc++-dev=13.2.1_git20240309-r0 \
        libunistring=1.2-r0 \
        linux-headers=6.6-r0 \
        make=4.4.1-r2 \
        mpc1=1.3.1-r1 \
        mpfr4=4.2.1-r0 \
        musl=1.2.5-r0 \
        musl-dev=1.2.5-r0 \
        nghttp2-libs=1.62.1-r0 \
        openssl-dev=3.3.2-r0 \
        openssl-libs-static=3.3.2-r0 \
        patch=2.7.6-r10 \
        pcre=8.45-r3 \
        pcre-dev=8.45-r3 \
        pcre2=10.43-r0 \
        pkgconf=2.2.0-r0 \
        tzdata=2024b-r0 \
        zlib=1.3.1-r1 \
        zlib-dev=1.3.1-r1 \
        zlib-static=1.3.1-r1 \
        zstd-libs=1.5.6-r0 \
        ; \
     git clone --depth=1 -b $TAG --single-branch https://github.com/redis/redis /src/redis; \
     cd /src/redis; \
     make distclean; \
     make; \
     mkdir /data; \
     strip /src/redis/src/redis-server; \
     touch --date=@0 /src/redis/src/redis-server /data /redis.conf

FROM reg.c5h.io/base@sha256:fa9cdd83581affde03cb2234170235ff287588e4da3786133d72846838210621

COPY --from=build --chown=65532:65532 /redis.conf /home/nonroot/redis.conf
COPY --from=build --chown=65532:65532 /data /home/nonroot/database
COPY --from=build --chown=65532:65532 /src/redis/src/redis-server /home/nonroot/redis

ENTRYPOINT ["/home/nonroot/redis"]

EXPOSE 6379
VOLUME "/home/nonroot/database"
CMD ["/home/nonroot/redis.conf"]
