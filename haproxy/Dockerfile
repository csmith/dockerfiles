# Generated from https://github.com/csmith/dockerfiles/blob/master/haproxy/Dockerfile.gotpl
# BOM: {"apk:binutils":"2.42-r0","apk:brotli-libs":"1.1.0-r2","apk:build-base":"0.5-r3","apk:busybox":"1.36.1-r28","apk:busybox-binsh":"1.36.1-r28","apk:c-ares":"1.28.1-r0","apk:ca-certificates":"20240226-r0","apk:file":"5.45-r1","apk:fortify-headers":"1.1-r3","apk:g++":"13.2.1_git20240309-r0","apk:gcc":"13.2.1_git20240309-r0","apk:git":"2.45.2-r0","apk:gmp":"6.3.0-r1","apk:isl25":"0.25-r2","apk:jansson":"2.14-r4","apk:libatomic":"13.2.1_git20240309-r0","apk:libcrypto3":"3.3.1-r0","apk:libcurl":"8.7.1-r0","apk:libexpat":"2.6.2-r0","apk:libgcc":"13.2.1_git20240309-r0","apk:libgomp":"13.2.1_git20240309-r0","apk:libidn2":"2.3.7-r0","apk:libmagic":"5.45-r1","apk:libpcre16":"8.45-r3","apk:libpcre32":"8.45-r3","apk:libpcrecpp":"8.45-r3","apk:libpsl":"0.21.5-r1","apk:libssl3":"3.3.1-r0","apk:libstdc++":"13.2.1_git20240309-r0","apk:libstdc++-dev":"13.2.1_git20240309-r0","apk:libunistring":"1.2-r0","apk:linux-headers":"6.6-r0","apk:make":"4.4.1-r2","apk:mpc1":"1.3.1-r1","apk:mpfr4":"4.2.1-r0","apk:musl":"1.2.5-r0","apk:musl-dev":"1.2.5-r0","apk:nghttp2-libs":"1.62.0-r0","apk:openssl-dev":"3.3.1-r0","apk:openssl-libs-static":"3.3.1-r0","apk:patch":"2.7.6-r10","apk:pcre":"8.45-r3","apk:pcre-dev":"8.45-r3","apk:pcre2":"10.43-r0","apk:pkgconf":"2.2.0-r0","apk:tzdata":"2024a-r1","apk:zlib":"1.3.1-r1","apk:zlib-dev":"1.3.1-r1","apk:zlib-static":"1.3.1-r1","apk:zstd-libs":"1.5.6-r0","github:haproxy/haproxy":"v3.0.0","image:alpine":"156a9f5dfd9defd08b6f5205b8d263b1810044b015f8cbb33fd469dd34e045fc","image:base":"b53f7bd15a51e9e9e1c9d1ddea46a2f62223ed7cd8d8fa320c0d6b323f9f3f99"}

FROM reg.c5h.io/alpine@sha256:156a9f5dfd9defd08b6f5205b8d263b1810044b015f8cbb33fd469dd34e045fc AS build

ARG TAG="v3.0.0"

COPY *.patch /src/

RUN set -eux; \
     apk add --no-cache --virtual .build-deps \
        binutils=2.42-r0 \
        brotli-libs=1.1.0-r2 \
        build-base=0.5-r3 \
        busybox=1.36.1-r28 \
        busybox-binsh=1.36.1-r28 \
        c-ares=1.28.1-r0 \
        ca-certificates=20240226-r0 \
        file=5.45-r1 \
        fortify-headers=1.1-r3 \
        g++=13.2.1_git20240309-r0 \
        gcc=13.2.1_git20240309-r0 \
        git=2.45.2-r0 \
        gmp=6.3.0-r1 \
        isl25=0.25-r2 \
        jansson=2.14-r4 \
        libatomic=13.2.1_git20240309-r0 \
        libcrypto3=3.3.1-r0 \
        libcurl=8.7.1-r0 \
        libexpat=2.6.2-r0 \
        libgcc=13.2.1_git20240309-r0 \
        libgomp=13.2.1_git20240309-r0 \
        libidn2=2.3.7-r0 \
        libmagic=5.45-r1 \
        libpcre16=8.45-r3 \
        libpcre32=8.45-r3 \
        libpcrecpp=8.45-r3 \
        libpsl=0.21.5-r1 \
        libssl3=3.3.1-r0 \
        libstdc++=13.2.1_git20240309-r0 \
        libstdc++-dev=13.2.1_git20240309-r0 \
        libunistring=1.2-r0 \
        linux-headers=6.6-r0 \
        make=4.4.1-r2 \
        mpc1=1.3.1-r1 \
        mpfr4=4.2.1-r0 \
        musl=1.2.5-r0 \
        musl-dev=1.2.5-r0 \
        nghttp2-libs=1.62.0-r0 \
        openssl-dev=3.3.1-r0 \
        openssl-libs-static=3.3.1-r0 \
        patch=2.7.6-r10 \
        pcre=8.45-r3 \
        pcre-dev=8.45-r3 \
        pcre2=10.43-r0 \
        pkgconf=2.2.0-r0 \
        tzdata=2024a-r1 \
        zlib=1.3.1-r1 \
        zlib-dev=1.3.1-r1 \
        zlib-static=1.3.1-r1 \
        zstd-libs=1.5.6-r0 \
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

FROM reg.c5h.io/base@sha256:b53f7bd15a51e9e9e1c9d1ddea46a2f62223ed7cd8d8fa320c0d6b323f9f3f99

STOPSIGNAL SIGUSR1

COPY --from=build /usr/local/sbin/haproxy /haproxy

ENTRYPOINT ["/haproxy", "-W", "-db"]

CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
