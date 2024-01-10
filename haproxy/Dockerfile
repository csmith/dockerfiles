# Generated from https://github.com/csmith/dockerfiles/blob/master/haproxy/Dockerfile.gotpl
# BOM: {"apk:binutils":"2.41-r0","apk:brotli-libs":"1.1.0-r1","apk:build-base":"0.5-r3","apk:busybox":"1.36.1-r15","apk:busybox-binsh":"1.36.1-r15","apk:c-ares":"1.22.1-r0","apk:ca-certificates":"20230506-r0","apk:file":"5.45-r1","apk:fortify-headers":"1.1-r3","apk:g++":"13.2.1_git20231014-r0","apk:gcc":"13.2.1_git20231014-r0","apk:git":"2.43.0-r0","apk:gmp":"6.3.0-r0","apk:isl25":"0.25-r2","apk:jansson":"2.14-r4","apk:libatomic":"13.2.1_git20231014-r0","apk:libc-dev":"0.7.2-r5","apk:libcrypto3":"3.1.4-r3","apk:libcurl":"8.5.0-r0","apk:libexpat":"2.5.0-r2","apk:libgcc":"13.2.1_git20231014-r0","apk:libgomp":"13.2.1_git20231014-r0","apk:libidn2":"2.3.4-r4","apk:libmagic":"5.45-r1","apk:libpcre16":"8.45-r3","apk:libpcre32":"8.45-r3","apk:libpcrecpp":"8.45-r3","apk:libssl3":"3.1.4-r3","apk:libstdc++":"13.2.1_git20231014-r0","apk:libstdc++-dev":"13.2.1_git20231014-r0","apk:libunistring":"1.1-r2","apk:linux-headers":"6.5-r0","apk:make":"4.4.1-r2","apk:mpc1":"1.3.1-r1","apk:mpfr4":"4.2.1-r0","apk:musl":"1.2.4_git20230717-r4","apk:musl-dev":"1.2.4_git20230717-r4","apk:nghttp2-libs":"1.58.0-r0","apk:openssl-dev":"3.1.4-r3","apk:openssl-libs-static":"3.1.4-r3","apk:patch":"2.7.6-r10","apk:pcre":"8.45-r3","apk:pcre-dev":"8.45-r3","apk:pcre2":"10.42-r2","apk:pkgconf":"2.1.0-r0","apk:tzdata":"2023d-r0","apk:zlib":"1.3-r2","apk:zlib-dev":"1.3-r2","apk:zlib-static":"1.3-r2","apk:zstd-libs":"1.5.5-r8","github:haproxy/haproxy":"v2.9.0","image:alpine":"66b5ada484622010e4c04dda462a69c927abeed683648b7019f602b83f3474aa","image:base":"cfa1d5f949b495a6e077bdddb1dbbc3d9a9d55036ab1267f4168a6a5c7e7fdce"}

FROM reg.c5h.io/alpine@sha256:66b5ada484622010e4c04dda462a69c927abeed683648b7019f602b83f3474aa AS build

ARG TAG="v2.9.0"

COPY *.patch /src/

RUN set -eux; \
     apk add --no-cache --virtual .build-deps \
        binutils=2.41-r0 \
        brotli-libs=1.1.0-r1 \
        build-base=0.5-r3 \
        busybox=1.36.1-r15 \
        busybox-binsh=1.36.1-r15 \
        c-ares=1.22.1-r0 \
        ca-certificates=20230506-r0 \
        file=5.45-r1 \
        fortify-headers=1.1-r3 \
        g++=13.2.1_git20231014-r0 \
        gcc=13.2.1_git20231014-r0 \
        git=2.43.0-r0 \
        gmp=6.3.0-r0 \
        isl25=0.25-r2 \
        jansson=2.14-r4 \
        libatomic=13.2.1_git20231014-r0 \
        libc-dev=0.7.2-r5 \
        libcrypto3=3.1.4-r3 \
        libcurl=8.5.0-r0 \
        libexpat=2.5.0-r2 \
        libgcc=13.2.1_git20231014-r0 \
        libgomp=13.2.1_git20231014-r0 \
        libidn2=2.3.4-r4 \
        libmagic=5.45-r1 \
        libpcre16=8.45-r3 \
        libpcre32=8.45-r3 \
        libpcrecpp=8.45-r3 \
        libssl3=3.1.4-r3 \
        libstdc++=13.2.1_git20231014-r0 \
        libstdc++-dev=13.2.1_git20231014-r0 \
        libunistring=1.1-r2 \
        linux-headers=6.5-r0 \
        make=4.4.1-r2 \
        mpc1=1.3.1-r1 \
        mpfr4=4.2.1-r0 \
        musl=1.2.4_git20230717-r4 \
        musl-dev=1.2.4_git20230717-r4 \
        nghttp2-libs=1.58.0-r0 \
        openssl-dev=3.1.4-r3 \
        openssl-libs-static=3.1.4-r3 \
        patch=2.7.6-r10 \
        pcre=8.45-r3 \
        pcre-dev=8.45-r3 \
        pcre2=10.42-r2 \
        pkgconf=2.1.0-r0 \
        tzdata=2023d-r0 \
        zlib=1.3-r2 \
        zlib-dev=1.3-r2 \
        zlib-static=1.3-r2 \
        zstd-libs=1.5.5-r8 \
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

FROM reg.c5h.io/base@sha256:cfa1d5f949b495a6e077bdddb1dbbc3d9a9d55036ab1267f4168a6a5c7e7fdce

STOPSIGNAL SIGUSR1

COPY --from=build /usr/local/sbin/haproxy /haproxy

ENTRYPOINT ["/haproxy", "-W", "-db"]

CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
