# Generated from https://github.com/csmith/dockerfiles/blob/master/haproxy/Dockerfile.gotpl
# BOM: {"apk:binutils":"2.37-r3","apk:brotli-libs":"1.0.9-r5","apk:build-base":"0.5-r2","apk:busybox":"1.34.1-r4","apk:ca-certificates":"20211220-r0","apk:expat":"2.4.7-r0","apk:file":"5.41-r0","apk:fortify-headers":"1.1-r1","apk:g++":"10.3.1_git20211027-r0","apk:gcc":"10.3.1_git20211027-r0","apk:git":"2.34.1-r0","apk:gmp":"6.2.1-r1","apk:isl22":"0.22-r0","apk:libatomic":"10.3.1_git20211027-r0","apk:libc-dev":"0.7.2-r3","apk:libcrypto1.1":"1.1.1n-r0","apk:libcurl":"7.80.0-r0","apk:libgcc":"10.3.1_git20211027-r0","apk:libgomp":"10.3.1_git20211027-r0","apk:libgphobos":"10.3.1_git20211027-r0","apk:libmagic":"5.41-r0","apk:libpcre16":"8.45-r1","apk:libpcre32":"8.45-r1","apk:libpcrecpp":"8.45-r1","apk:libssl1.1":"1.1.1n-r0","apk:libstdc++":"10.3.1_git20211027-r0","apk:linux-headers":"5.10.41-r0","apk:make":"4.3-r0","apk:mpc1":"1.2.1-r0","apk:mpfr4":"4.1.0-r0","apk:musl":"1.2.2-r7","apk:musl-dev":"1.2.2-r7","apk:nghttp2-libs":"1.46.0-r0","apk:openssl-dev":"1.1.1n-r0","apk:openssl-libs-static":"1.1.1n-r0","apk:patch":"2.7.6-r7","apk:pcre":"8.45-r1","apk:pcre-dev":"8.45-r1","apk:pcre2":"10.39-r0","apk:pkgconf":"1.8.0-r0","apk:tzdata":"2022a-r0","apk:zlib":"1.2.11-r3","apk:zlib-dev":"1.2.11-r3","apk:zlib-static":"1.2.11-r3","github:haproxy/haproxy":"v2.5.0","image:alpine":"a3b31a0fa91c1fd8833d62e92b57d2c85a1d6181c74b7a680624f193329dd3a1","image:base":"f85825cc6e579732f339638d1cd16d87f0673759127df47e8bcd3586412acb5d"}

FROM reg.c5h.io/alpine@sha256:a3b31a0fa91c1fd8833d62e92b57d2c85a1d6181c74b7a680624f193329dd3a1 AS build

ARG TAG="v2.5.0"

COPY *.patch /src/

RUN set -eux; \
     apk add --no-cache --virtual .build-deps \
        binutils=2.37-r3 \
        brotli-libs=1.0.9-r5 \
        build-base=0.5-r2 \
        busybox=1.34.1-r4 \
        ca-certificates=20211220-r0 \
        expat=2.4.7-r0 \
        file=5.41-r0 \
        fortify-headers=1.1-r1 \
        g++=10.3.1_git20211027-r0 \
        gcc=10.3.1_git20211027-r0 \
        git=2.34.1-r0 \
        gmp=6.2.1-r1 \
        isl22=0.22-r0 \
        libatomic=10.3.1_git20211027-r0 \
        libc-dev=0.7.2-r3 \
        libcrypto1.1=1.1.1n-r0 \
        libcurl=7.80.0-r0 \
        libgcc=10.3.1_git20211027-r0 \
        libgomp=10.3.1_git20211027-r0 \
        libgphobos=10.3.1_git20211027-r0 \
        libmagic=5.41-r0 \
        libpcre16=8.45-r1 \
        libpcre32=8.45-r1 \
        libpcrecpp=8.45-r1 \
        libssl1.1=1.1.1n-r0 \
        libstdc++=10.3.1_git20211027-r0 \
        linux-headers=5.10.41-r0 \
        make=4.3-r0 \
        mpc1=1.2.1-r0 \
        mpfr4=4.1.0-r0 \
        musl=1.2.2-r7 \
        musl-dev=1.2.2-r7 \
        nghttp2-libs=1.46.0-r0 \
        openssl-dev=1.1.1n-r0 \
        openssl-libs-static=1.1.1n-r0 \
        patch=2.7.6-r7 \
        pcre=8.45-r1 \
        pcre-dev=8.45-r1 \
        pcre2=10.39-r0 \
        pkgconf=1.8.0-r0 \
        tzdata=2022a-r0 \
        zlib=1.2.11-r3 \
        zlib-dev=1.2.11-r3 \
        zlib-static=1.2.11-r3 \
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

FROM reg.c5h.io/base@sha256:f85825cc6e579732f339638d1cd16d87f0673759127df47e8bcd3586412acb5d

STOPSIGNAL SIGUSR1

COPY --from=build /usr/local/sbin/haproxy /haproxy

ENTRYPOINT ["/haproxy", "-W", "-db"]

CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
