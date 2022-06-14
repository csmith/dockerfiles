# Generated from https://github.com/csmith/dockerfiles/blob/master/haproxy/Dockerfile.gotpl
# BOM: {"apk:binutils":"2.38-r2","apk:brotli-libs":"1.0.9-r6","apk:build-base":"0.5-r2","apk:busybox":"1.35.0-r14","apk:ca-certificates":"20211220-r0","apk:expat":"2.4.8-r0","apk:file":"5.41-r0","apk:fortify-headers":"1.1-r1","apk:g++":"11.2.1_git20220219-r2","apk:gcc":"11.2.1_git20220219-r2","apk:git":"2.36.1-r0","apk:gmp":"6.2.1-r2","apk:isl22":"0.22-r0","apk:libatomic":"11.2.1_git20220219-r2","apk:libc-dev":"0.7.2-r3","apk:libcrypto1.1":"1.1.1o-r0","apk:libcurl":"7.83.1-r1","apk:libgcc":"11.2.1_git20220219-r2","apk:libgomp":"11.2.1_git20220219-r2","apk:libmagic":"5.41-r0","apk:libpcre16":"8.45-r2","apk:libpcre32":"8.45-r2","apk:libpcrecpp":"8.45-r2","apk:libssl1.1":"1.1.1o-r0","apk:libstdc++":"11.2.1_git20220219-r2","apk:linux-headers":"5.16.7-r1","apk:make":"4.3-r0","apk:mpc1":"1.2.1-r0","apk:mpfr4":"4.1.0-r0","apk:musl":"1.2.3-r0","apk:musl-dev":"1.2.3-r0","apk:nghttp2-libs":"1.47.0-r0","apk:openssl-dev":"1.1.1o-r0","apk:openssl-libs-static":"1.1.1o-r0","apk:patch":"2.7.6-r7","apk:pcre":"8.45-r2","apk:pcre-dev":"8.45-r2","apk:pcre2":"10.40-r0","apk:pkgconf":"1.8.0-r0","apk:tzdata":"2022a-r0","apk:zlib":"1.2.12-r1","apk:zlib-dev":"1.2.12-r1","apk:zlib-static":"1.2.12-r1","github:haproxy/haproxy":"v2.6.0","image:alpine":"e1af61331d46ce7a5770c5083e1c81e297fbf3fcc97b63542b0ac62a8f3900a3","image:base":"bffee9632ec79071de34e1a4a4f6c4563a71ad26d5923062bd9d0f9016f72beb"}

FROM reg.c5h.io/alpine@sha256:e1af61331d46ce7a5770c5083e1c81e297fbf3fcc97b63542b0ac62a8f3900a3 AS build

ARG TAG="v2.6.0"

COPY *.patch /src/

RUN set -eux; \
     apk add --no-cache --virtual .build-deps \
        binutils=2.38-r2 \
        brotli-libs=1.0.9-r6 \
        build-base=0.5-r2 \
        busybox=1.35.0-r14 \
        ca-certificates=20211220-r0 \
        expat=2.4.8-r0 \
        file=5.41-r0 \
        fortify-headers=1.1-r1 \
        g++=11.2.1_git20220219-r2 \
        gcc=11.2.1_git20220219-r2 \
        git=2.36.1-r0 \
        gmp=6.2.1-r2 \
        isl22=0.22-r0 \
        libatomic=11.2.1_git20220219-r2 \
        libc-dev=0.7.2-r3 \
        libcrypto1.1=1.1.1o-r0 \
        libcurl=7.83.1-r1 \
        libgcc=11.2.1_git20220219-r2 \
        libgomp=11.2.1_git20220219-r2 \
        libmagic=5.41-r0 \
        libpcre16=8.45-r2 \
        libpcre32=8.45-r2 \
        libpcrecpp=8.45-r2 \
        libssl1.1=1.1.1o-r0 \
        libstdc++=11.2.1_git20220219-r2 \
        linux-headers=5.16.7-r1 \
        make=4.3-r0 \
        mpc1=1.2.1-r0 \
        mpfr4=4.1.0-r0 \
        musl=1.2.3-r0 \
        musl-dev=1.2.3-r0 \
        nghttp2-libs=1.47.0-r0 \
        openssl-dev=1.1.1o-r0 \
        openssl-libs-static=1.1.1o-r0 \
        patch=2.7.6-r7 \
        pcre=8.45-r2 \
        pcre-dev=8.45-r2 \
        pcre2=10.40-r0 \
        pkgconf=1.8.0-r0 \
        tzdata=2022a-r0 \
        zlib=1.2.12-r1 \
        zlib-dev=1.2.12-r1 \
        zlib-static=1.2.12-r1 \
        ; \
     git clone --depth=1 -b $TAG --single-branch https://github.com/haproxy/haproxy /src/haproxy; \
     cd /src/haproxy; \
     patch Makefile < /src/buildstatic.patch; \
     if [ "$TAG" = "v2.5.0" ]; then \
       patch -p1 < /src/header-flags.patch; \
     fi; \
     make TARGET=linux-musl USE_OPENSSL=1 USE_PCRE=1; \
     make install-bin; \
     strip /usr/local/sbin/haproxy; \
     touch --date=@0 /usr/local/sbin/haproxy

FROM reg.c5h.io/base@sha256:bffee9632ec79071de34e1a4a4f6c4563a71ad26d5923062bd9d0f9016f72beb

STOPSIGNAL SIGUSR1

COPY --from=build /usr/local/sbin/haproxy /haproxy

ENTRYPOINT ["/haproxy", "-W", "-db"]

CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
