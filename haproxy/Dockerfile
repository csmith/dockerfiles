# Generated from https://github.com/csmith/dockerfiles/blob/master/haproxy/Dockerfile.gotpl
# BOM: {"apk:binutils":"2.40-r6","apk:brotli-libs":"1.0.9-r14","apk:build-base":"0.5-r3","apk:busybox":"1.36.0-r9","apk:busybox-binsh":"1.36.0-r9","apk:ca-certificates":"20230506-r0","apk:file":"5.44-r3","apk:fortify-headers":"1.1-r3","apk:g++":"12.2.1_git20220924-r10","apk:gcc":"12.2.1_git20220924-r10","apk:git":"2.40.1-r0","apk:gmp":"6.2.1-r3","apk:isl25":"0.25-r2","apk:libatomic":"12.2.1_git20220924-r10","apk:libc-dev":"0.7.2-r5","apk:libcrypto3":"3.1.0-r4","apk:libcurl":"8.0.1-r2","apk:libexpat":"2.5.0-r1","apk:libgcc":"12.2.1_git20220924-r10","apk:libgomp":"12.2.1_git20220924-r10","apk:libidn2":"2.3.4-r1","apk:libmagic":"5.44-r3","apk:libpcre16":"8.45-r3","apk:libpcre32":"8.45-r3","apk:libpcrecpp":"8.45-r3","apk:libssl3":"3.1.0-r4","apk:libstdc++":"12.2.1_git20220924-r10","apk:libstdc++-dev":"12.2.1_git20220924-r10","apk:libunistring":"1.1-r1","apk:linux-headers":"6.3-r0","apk:make":"4.4.1-r1","apk:mpc1":"1.3.1-r1","apk:mpfr4":"4.2.0-r2","apk:musl":"1.2.4-r0","apk:musl-dev":"1.2.4-r0","apk:nghttp2-libs":"1.53.0-r0","apk:openssl-dev":"3.1.0-r4","apk:openssl-libs-static":"3.1.0-r4","apk:patch":"2.7.6-r10","apk:pcre":"8.45-r3","apk:pcre-dev":"8.45-r3","apk:pcre2":"10.42-r1","apk:pkgconf":"1.9.5-r0","apk:tzdata":"2023c-r1","apk:zlib":"1.2.13-r1","apk:zlib-dev":"1.2.13-r1","apk:zlib-static":"1.2.13-r1","apk:zstd-libs":"1.5.5-r4","github:haproxy/haproxy":"v2.7.0","image:alpine":"f182b51cb5007336ee5e05fe5389f7c5a1d26428229ab27dbf502a7b4b3b6455","image:base":"33b47501f2e31be55c1af6b21082de3c3efaae9ec0062beb1d3255a9fd739874"}

FROM reg.c5h.io/alpine@sha256:f182b51cb5007336ee5e05fe5389f7c5a1d26428229ab27dbf502a7b4b3b6455 AS build

ARG TAG="v2.7.0"

COPY *.patch /src/

RUN set -eux; \
     apk add --no-cache --virtual .build-deps \
        binutils=2.40-r6 \
        brotli-libs=1.0.9-r14 \
        build-base=0.5-r3 \
        busybox=1.36.0-r9 \
        busybox-binsh=1.36.0-r9 \
        ca-certificates=20230506-r0 \
        file=5.44-r3 \
        fortify-headers=1.1-r3 \
        g++=12.2.1_git20220924-r10 \
        gcc=12.2.1_git20220924-r10 \
        git=2.40.1-r0 \
        gmp=6.2.1-r3 \
        isl25=0.25-r2 \
        libatomic=12.2.1_git20220924-r10 \
        libc-dev=0.7.2-r5 \
        libcrypto3=3.1.0-r4 \
        libcurl=8.0.1-r2 \
        libexpat=2.5.0-r1 \
        libgcc=12.2.1_git20220924-r10 \
        libgomp=12.2.1_git20220924-r10 \
        libidn2=2.3.4-r1 \
        libmagic=5.44-r3 \
        libpcre16=8.45-r3 \
        libpcre32=8.45-r3 \
        libpcrecpp=8.45-r3 \
        libssl3=3.1.0-r4 \
        libstdc++=12.2.1_git20220924-r10 \
        libstdc++-dev=12.2.1_git20220924-r10 \
        libunistring=1.1-r1 \
        linux-headers=6.3-r0 \
        make=4.4.1-r1 \
        mpc1=1.3.1-r1 \
        mpfr4=4.2.0-r2 \
        musl=1.2.4-r0 \
        musl-dev=1.2.4-r0 \
        nghttp2-libs=1.53.0-r0 \
        openssl-dev=3.1.0-r4 \
        openssl-libs-static=3.1.0-r4 \
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

FROM reg.c5h.io/base@sha256:33b47501f2e31be55c1af6b21082de3c3efaae9ec0062beb1d3255a9fd739874

STOPSIGNAL SIGUSR1

COPY --from=build /usr/local/sbin/haproxy /haproxy

ENTRYPOINT ["/haproxy", "-W", "-db"]

CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
