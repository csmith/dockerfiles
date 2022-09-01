# Generated from https://github.com/csmith/dockerfiles/blob/master/haproxy/Dockerfile.gotpl
# BOM: {"apk:binutils":"2.38-r3","apk:brotli-libs":"1.0.9-r6","apk:build-base":"0.5-r3","apk:busybox":"1.35.0-r17","apk:ca-certificates":"20220614-r0","apk:expat":"2.4.8-r0","apk:file":"5.41-r0","apk:fortify-headers":"1.1-r1","apk:g++":"11.2.1_git20220219-r2","apk:gcc":"11.2.1_git20220219-r2","apk:git":"2.36.2-r0","apk:gmp":"6.2.1-r2","apk:isl22":"0.22-r0","apk:libatomic":"11.2.1_git20220219-r2","apk:libc-dev":"0.7.2-r3","apk:libcrypto1.1":"1.1.1q-r0","apk:libcurl":"7.83.1-r3","apk:libgcc":"11.2.1_git20220219-r2","apk:libgomp":"11.2.1_git20220219-r2","apk:libmagic":"5.41-r0","apk:libpcre16":"8.45-r2","apk:libpcre32":"8.45-r2","apk:libpcrecpp":"8.45-r2","apk:libssl1.1":"1.1.1q-r0","apk:libstdc++":"11.2.1_git20220219-r2","apk:linux-headers":"5.16.7-r1","apk:make":"4.3-r0","apk:mpc1":"1.2.1-r0","apk:mpfr4":"4.1.0-r0","apk:musl":"1.2.3-r0","apk:musl-dev":"1.2.3-r0","apk:nghttp2-libs":"1.47.0-r0","apk:openssl-dev":"1.1.1q-r0","apk:openssl-libs-static":"1.1.1q-r0","apk:patch":"2.7.6-r7","apk:pcre":"8.45-r2","apk:pcre-dev":"8.45-r2","apk:pcre2":"10.40-r0","apk:pkgconf":"1.8.0-r1","apk:tzdata":"2022c-r0","apk:zlib":"1.2.12-r3","apk:zlib-dev":"1.2.12-r3","apk:zlib-static":"1.2.12-r3","github:haproxy/haproxy":"v2.6.0","image:alpine":"a475fe9f28063dd6914342b78877af564ac1144cacba56ddb2aed1bce0c3fc14","image:base":"dad97084335f003c2ec5251b8f9d22e59704f59c8e2ffcc6108470131341b2df"}

FROM reg.c5h.io/alpine@sha256:a475fe9f28063dd6914342b78877af564ac1144cacba56ddb2aed1bce0c3fc14 AS build

ARG TAG="v2.6.0"

COPY *.patch /src/

RUN set -eux; \
     apk add --no-cache --virtual .build-deps \
        binutils=2.38-r3 \
        brotli-libs=1.0.9-r6 \
        build-base=0.5-r3 \
        busybox=1.35.0-r17 \
        ca-certificates=20220614-r0 \
        expat=2.4.8-r0 \
        file=5.41-r0 \
        fortify-headers=1.1-r1 \
        g++=11.2.1_git20220219-r2 \
        gcc=11.2.1_git20220219-r2 \
        git=2.36.2-r0 \
        gmp=6.2.1-r2 \
        isl22=0.22-r0 \
        libatomic=11.2.1_git20220219-r2 \
        libc-dev=0.7.2-r3 \
        libcrypto1.1=1.1.1q-r0 \
        libcurl=7.83.1-r3 \
        libgcc=11.2.1_git20220219-r2 \
        libgomp=11.2.1_git20220219-r2 \
        libmagic=5.41-r0 \
        libpcre16=8.45-r2 \
        libpcre32=8.45-r2 \
        libpcrecpp=8.45-r2 \
        libssl1.1=1.1.1q-r0 \
        libstdc++=11.2.1_git20220219-r2 \
        linux-headers=5.16.7-r1 \
        make=4.3-r0 \
        mpc1=1.2.1-r0 \
        mpfr4=4.1.0-r0 \
        musl=1.2.3-r0 \
        musl-dev=1.2.3-r0 \
        nghttp2-libs=1.47.0-r0 \
        openssl-dev=1.1.1q-r0 \
        openssl-libs-static=1.1.1q-r0 \
        patch=2.7.6-r7 \
        pcre=8.45-r2 \
        pcre-dev=8.45-r2 \
        pcre2=10.40-r0 \
        pkgconf=1.8.0-r1 \
        tzdata=2022c-r0 \
        zlib=1.2.12-r3 \
        zlib-dev=1.2.12-r3 \
        zlib-static=1.2.12-r3 \
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

FROM reg.c5h.io/base@sha256:dad97084335f003c2ec5251b8f9d22e59704f59c8e2ffcc6108470131341b2df

STOPSIGNAL SIGUSR1

COPY --from=build /usr/local/sbin/haproxy /haproxy

ENTRYPOINT ["/haproxy", "-W", "-db"]

CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
