# Generated from https://github.com/csmith/dockerfiles/blob/master/redis/Dockerfile.gotpl
# BOM: {"apk:binutils":"2.39-r2","apk:brotli-libs":"1.0.9-r9","apk:build-base":"0.5-r3","apk:busybox":"1.35.0-r29","apk:busybox-binsh":"1.35.0-r29","apk:ca-certificates":"20220614-r4","apk:file":"5.43-r0","apk:fortify-headers":"1.1-r1","apk:g++":"12.2.1_git20220924-r4","apk:gcc":"12.2.1_git20220924-r4","apk:git":"2.38.2-r0","apk:gmp":"6.2.1-r2","apk:isl24":"0.24-r1","apk:libatomic":"12.2.1_git20220924-r4","apk:libc-dev":"0.7.2-r3","apk:libcrypto3":"3.0.7-r2","apk:libcurl":"7.87.0-r1","apk:libexpat":"2.5.0-r0","apk:libgcc":"12.2.1_git20220924-r4","apk:libgomp":"12.2.1_git20220924-r4","apk:libmagic":"5.43-r0","apk:libpcre16":"8.45-r2","apk:libpcre32":"8.45-r2","apk:libpcrecpp":"8.45-r2","apk:libssl3":"3.0.7-r2","apk:libstdc++":"12.2.1_git20220924-r4","apk:libstdc++-dev":"12.2.1_git20220924-r4","apk:linux-headers":"5.19.5-r0","apk:make":"4.3-r1","apk:mpc1":"1.2.1-r1","apk:mpfr4":"4.1.0-r0","apk:musl":"1.2.3-r4","apk:musl-dev":"1.2.3-r4","apk:nghttp2-libs":"1.51.0-r0","apk:openssl-dev":"3.0.7-r2","apk:openssl-libs-static":"3.0.7-r2","apk:patch":"2.7.6-r8","apk:pcre":"8.45-r2","apk:pcre-dev":"8.45-r2","apk:pcre2":"10.42-r0","apk:pkgconf":"1.9.3-r0","apk:tzdata":"2022f-r1","apk:zlib":"1.2.13-r0","apk:zlib-dev":"1.2.13-r0","apk:zlib-static":"1.2.13-r0","github:redis/redis":"7.0.8","image:alpine":"c3b87feb2ca33e6430ed010a3b799eb6739e137aead5a02bdf5977cc6ae0a882","image:base":"45ece244ca0ace29f890ad27114ff83a36c00e74e6cbec3d1e226b7e62d0de7a"}

FROM reg.c5h.io/alpine@sha256:c3b87feb2ca33e6430ed010a3b799eb6739e137aead5a02bdf5977cc6ae0a882 AS build

ARG TAG="7.0.8"

ENV LDFLAGS="-static"

COPY redis.conf /redis.conf

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
        git=2.38.2-r0 \
        gmp=6.2.1-r2 \
        isl24=0.24-r1 \
        libatomic=12.2.1_git20220924-r4 \
        libc-dev=0.7.2-r3 \
        libcrypto3=3.0.7-r2 \
        libcurl=7.87.0-r1 \
        libexpat=2.5.0-r0 \
        libgcc=12.2.1_git20220924-r4 \
        libgomp=12.2.1_git20220924-r4 \
        libmagic=5.43-r0 \
        libpcre16=8.45-r2 \
        libpcre32=8.45-r2 \
        libpcrecpp=8.45-r2 \
        libssl3=3.0.7-r2 \
        libstdc++=12.2.1_git20220924-r4 \
        libstdc++-dev=12.2.1_git20220924-r4 \
        linux-headers=5.19.5-r0 \
        make=4.3-r1 \
        mpc1=1.2.1-r1 \
        mpfr4=4.1.0-r0 \
        musl=1.2.3-r4 \
        musl-dev=1.2.3-r4 \
        nghttp2-libs=1.51.0-r0 \
        openssl-dev=3.0.7-r2 \
        openssl-libs-static=3.0.7-r2 \
        patch=2.7.6-r8 \
        pcre=8.45-r2 \
        pcre-dev=8.45-r2 \
        pcre2=10.42-r0 \
        pkgconf=1.9.3-r0 \
        tzdata=2022f-r1 \
        zlib=1.2.13-r0 \
        zlib-dev=1.2.13-r0 \
        zlib-static=1.2.13-r0 \
        ; \
     git clone --depth=1 -b $TAG --single-branch https://github.com/redis/redis /src/redis; \
     cd /src/redis; \
     make distclean; \
     make; \
     mkdir /data; \
     strip /src/redis/src/redis-server; \
     touch --date=@0 /src/redis/src/redis-server /data /redis.conf

FROM reg.c5h.io/base@sha256:45ece244ca0ace29f890ad27114ff83a36c00e74e6cbec3d1e226b7e62d0de7a

COPY --from=build --chown=65532:65532 /redis.conf /home/nonroot/redis.conf
COPY --from=build --chown=65532:65532 /data /home/nonroot/database
COPY --from=build --chown=65532:65532 /src/redis/src/redis-server /home/nonroot/redis

ENTRYPOINT ["/home/nonroot/redis"]

EXPOSE 6379
VOLUME "/home/nonroot/database"
CMD ["/home/nonroot/redis.conf"]
