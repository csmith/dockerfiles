# Generated from https://github.com/csmith/dockerfiles/blob/master/haproxy/Dockerfile.gotpl
# BOM: {"apk:binutils":"2.43.1-r2","apk:brotli-libs":"1.1.0-r2","apk:build-base":"0.5-r3","apk:busybox":"1.37.0-r12","apk:busybox-binsh":"1.37.0-r12","apk:c-ares":"1.34.5-r0","apk:ca-certificates":"20241121-r1","apk:ca-certificates-bundle":"20241121-r1","apk:file":"5.46-r2","apk:fortify-headers":"1.1-r5","apk:g++":"14.2.0-r4","apk:gcc":"14.2.0-r4","apk:git":"2.47.2-r0","apk:gmp":"6.3.0-r2","apk:isl25":"0.25-r2","apk:jansson":"2.14-r4","apk:libatomic":"14.2.0-r4","apk:libcrypto3":"3.3.3-r0","apk:libcurl":"8.12.1-r1","apk:libexpat":"2.7.0-r0","apk:libgcc":"14.2.0-r4","apk:libgomp":"14.2.0-r4","apk:libidn2":"2.3.7-r0","apk:libmagic":"5.46-r2","apk:libpcre16":"8.45-r3","apk:libpcre32":"8.45-r3","apk:libpcrecpp":"8.45-r3","apk:libpsl":"0.21.5-r3","apk:libssl3":"3.3.3-r0","apk:libstdc++":"14.2.0-r4","apk:libstdc++-dev":"14.2.0-r4","apk:libunistring":"1.2-r0","apk:linux-headers":"6.6-r1","apk:make":"4.4.1-r2","apk:mpc1":"1.3.1-r1","apk:mpfr4":"4.2.1-r0","apk:musl":"1.2.5-r9","apk:musl-dev":"1.2.5-r9","apk:nghttp2-libs":"1.64.0-r0","apk:openssl-dev":"3.3.3-r0","apk:openssl-libs-static":"3.3.3-r0","apk:patch":"2.7.6-r10","apk:pcre":"8.45-r3","apk:pcre-dev":"8.45-r3","apk:pcre2":"10.43-r0","apk:pkgconf":"2.3.0-r0","apk:tzdata":"2025b-r0","apk:zlib":"1.3.1-r2","apk:zlib-dev":"1.3.1-r2","apk:zlib-static":"1.3.1-r2","apk:zstd-libs":"1.5.6-r2","github:haproxy/haproxy":"v3.1.0","image:alpine":"9dd02454e2bd77b983bf4db2e9ca788ce76b28aaafa607a1dab7097f616223a9","image:base":"4fea2559a3c315f01c6241c689da902344f4694dbdbdab4f5c0a1d24a8a78363"}

FROM reg.c5h.io/alpine@sha256:9dd02454e2bd77b983bf4db2e9ca788ce76b28aaafa607a1dab7097f616223a9 AS build

RUN set -eux; \
    apk add --no-cache \
        brotli-libs=1.1.0-r2 \
        c-ares=1.34.5-r0 \
        ca-certificates-bundle=20241121-r1 \
        git=2.47.2-r0 \
        libcrypto3=3.3.3-r0 \
        libcurl=8.12.1-r1 \
        libexpat=2.7.0-r0 \
        libidn2=2.3.7-r0 \
        libpsl=0.21.5-r3 \
        libssl3=3.3.3-r0 \
        libunistring=1.2-r0 \
        musl=1.2.5-r9 \
        nghttp2-libs=1.64.0-r0 \
        pcre2=10.43-r0 \
        zlib=1.3.1-r2 \
        zstd-libs=1.5.6-r2 \
        binutils=2.43.1-r2 \
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
        musl=1.2.5-r9 \
        musl-dev=1.2.5-r9 \
        patch=2.7.6-r10 \
        zlib=1.3.1-r2 \
        zstd-libs=1.5.6-r2 \
        tzdata=2025b-r0 \
        busybox=1.37.0-r12 \
        busybox-binsh=1.37.0-r12 \
        ca-certificates=20241121-r1 \
        libcrypto3=3.3.3-r0 \
        musl=1.2.5-r9 \
        linux-headers=6.6-r1 \
        libgcc=14.2.0-r4 \
        libpcre16=8.45-r3 \
        libpcre32=8.45-r3 \
        libpcrecpp=8.45-r3 \
        libstdc++=14.2.0-r4 \
        musl=1.2.5-r9 \
        pcre=8.45-r3 \
        pcre-dev=8.45-r3 \
        pkgconf=2.3.0-r0 \
        musl=1.2.5-r9 \
        musl-dev=1.2.5-r9 \
        libcrypto3=3.3.3-r0 \
        libssl3=3.3.3-r0 \
        musl=1.2.5-r9 \
        openssl-dev=3.3.3-r0 \
        pkgconf=2.3.0-r0 \
        openssl-libs-static=3.3.3-r0 \
        musl=1.2.5-r9 \
        pkgconf=2.3.0-r0 \
        zlib=1.3.1-r2 \
        zlib-dev=1.3.1-r2 \
        zlib-static=1.3.1-r2 \
    ;



RUN set -eux; \
    git clone --depth=1 -b v3.1.0 --single-branch https://github.com/haproxy/haproxy /src/haproxy

WORKDIR /src/haproxy


COPY *.patch /src/

RUN set -eux; \
     patch Makefile < /src/buildstatic.patch; \
     if [ "v3.1.0" = "v2.5.0" ]; then \
       patch -p1 < /src/header-flags.patch; \
     fi; \
     make TARGET=linux-musl USE_OPENSSL=1 USE_PCRE=1; \
     make install-bin; \
     strip /usr/local/sbin/haproxy; \
     touch --date=@0 /usr/local/sbin/haproxy

FROM reg.c5h.io/base@sha256:4fea2559a3c315f01c6241c689da902344f4694dbdbdab4f5c0a1d24a8a78363

STOPSIGNAL SIGUSR1

COPY --from=build /usr/local/sbin/haproxy /haproxy

ENTRYPOINT ["/haproxy", "-W", "-db"]

CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
