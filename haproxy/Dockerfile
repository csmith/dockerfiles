# Generated from https://github.com/csmith/dockerfiles/blob/master/haproxy/Dockerfile.gotpl
# BOM: {"apk:binutils":"2.39-r2","apk:brotli-libs":"1.0.9-r9","apk:build-base":"0.5-r3","apk:busybox":"1.35.0-r29","apk:busybox-binsh":"1.35.0-r29","apk:ca-certificates":"20220614-r4","apk:file":"5.43-r0","apk:fortify-headers":"1.1-r1","apk:g++":"12.2.1_git20220924-r4","apk:gcc":"12.2.1_git20220924-r4","apk:git":"2.38.4-r1","apk:gmp":"6.2.1-r2","apk:isl24":"0.24-r1","apk:libatomic":"12.2.1_git20220924-r4","apk:libc-dev":"0.7.2-r3","apk:libcrypto3":"3.0.8-r0","apk:libcurl":"7.88.1-r0","apk:libexpat":"2.5.0-r0","apk:libgcc":"12.2.1_git20220924-r4","apk:libgomp":"12.2.1_git20220924-r4","apk:libmagic":"5.43-r0","apk:libpcre16":"8.45-r2","apk:libpcre32":"8.45-r2","apk:libpcrecpp":"8.45-r2","apk:libssl3":"3.0.8-r0","apk:libstdc++":"12.2.1_git20220924-r4","apk:libstdc++-dev":"12.2.1_git20220924-r4","apk:linux-headers":"5.19.5-r0","apk:make":"4.3-r1","apk:mpc1":"1.2.1-r1","apk:mpfr4":"4.1.0-r0","apk:musl":"1.2.3-r4","apk:musl-dev":"1.2.3-r4","apk:nghttp2-libs":"1.51.0-r0","apk:openssl-dev":"3.0.8-r0","apk:openssl-libs-static":"3.0.8-r0","apk:patch":"2.7.6-r8","apk:pcre":"8.45-r2","apk:pcre-dev":"8.45-r2","apk:pcre2":"10.42-r0","apk:pkgconf":"1.9.4-r0","apk:tzdata":"2022f-r1","apk:zlib":"1.2.13-r0","apk:zlib-dev":"1.2.13-r0","apk:zlib-static":"1.2.13-r0","github:haproxy/haproxy":"v2.7.0","image:alpine":"3d271830a936728b41e38fcf89a9703ee4a9983af40618722eefead94279b55c","image:base":"45ece244ca0ace29f890ad27114ff83a36c00e74e6cbec3d1e226b7e62d0de7a"}

FROM reg.c5h.io/alpine@sha256:3d271830a936728b41e38fcf89a9703ee4a9983af40618722eefead94279b55c AS build

ARG TAG="v2.7.0"

COPY *.patch /src/

RUN set -eux; \
     apk add --no-cache --virtual .build-deps \
        binutils=2.39-r2 \
        brotli-libs=1.0.9-r9 \
        build-base=0.5-r3 \
        busybox=1.35.0-r29 \
        busybox-binsh=1.35.0-r29 \
        ca-certificates=20220614-r4 \
        file=5.43-r0 \
        fortify-headers=1.1-r1 \
        g++=12.2.1_git20220924-r4 \
        gcc=12.2.1_git20220924-r4 \
        git=2.38.4-r1 \
        gmp=6.2.1-r2 \
        isl24=0.24-r1 \
        libatomic=12.2.1_git20220924-r4 \
        libc-dev=0.7.2-r3 \
        libcrypto3=3.0.8-r0 \
        libcurl=7.88.1-r0 \
        libexpat=2.5.0-r0 \
        libgcc=12.2.1_git20220924-r4 \
        libgomp=12.2.1_git20220924-r4 \
        libmagic=5.43-r0 \
        libpcre16=8.45-r2 \
        libpcre32=8.45-r2 \
        libpcrecpp=8.45-r2 \
        libssl3=3.0.8-r0 \
        libstdc++=12.2.1_git20220924-r4 \
        libstdc++-dev=12.2.1_git20220924-r4 \
        linux-headers=5.19.5-r0 \
        make=4.3-r1 \
        mpc1=1.2.1-r1 \
        mpfr4=4.1.0-r0 \
        musl=1.2.3-r4 \
        musl-dev=1.2.3-r4 \
        nghttp2-libs=1.51.0-r0 \
        openssl-dev=3.0.8-r0 \
        openssl-libs-static=3.0.8-r0 \
        patch=2.7.6-r8 \
        pcre=8.45-r2 \
        pcre-dev=8.45-r2 \
        pcre2=10.42-r0 \
        pkgconf=1.9.4-r0 \
        tzdata=2022f-r1 \
        zlib=1.2.13-r0 \
        zlib-dev=1.2.13-r0 \
        zlib-static=1.2.13-r0 \
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

FROM reg.c5h.io/base@sha256:45ece244ca0ace29f890ad27114ff83a36c00e74e6cbec3d1e226b7e62d0de7a

STOPSIGNAL SIGUSR1

COPY --from=build /usr/local/sbin/haproxy /haproxy

ENTRYPOINT ["/haproxy", "-W", "-db"]

CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
