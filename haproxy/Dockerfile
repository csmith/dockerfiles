# Generated from https://github.com/csmith/dockerfiles/blob/master/haproxy/Dockerfile.gotpl
# BOM: {"apk:binutils":"2.35.2-r2","apk:brotli-libs":"1.0.9-r5","apk:build-base":"0.5-r2","apk:busybox":"1.33.1-r5","apk:ca-certificates":"20191127-r5","apk:expat":"2.4.1-r0","apk:file":"5.40-r1","apk:fortify-headers":"1.1-r1","apk:g++":"10.3.1_git20210424-r2","apk:gcc":"10.3.1_git20210424-r2","apk:git":"2.32.0-r0","apk:gmp":"6.2.1-r0","apk:isl22":"0.22-r0","apk:libatomic":"10.3.1_git20210424-r2","apk:libc-dev":"0.7.2-r3","apk:libcrypto1.1":"1.1.1l-r0","apk:libcurl":"7.79.1-r0","apk:libgcc":"10.3.1_git20210424-r2","apk:libgomp":"10.3.1_git20210424-r2","apk:libgphobos":"10.3.1_git20210424-r2","apk:libmagic":"5.40-r1","apk:libpcre16":"8.44-r0","apk:libpcre32":"8.44-r0","apk:libpcrecpp":"8.44-r0","apk:libssl1.1":"1.1.1l-r0","apk:libstdc++":"10.3.1_git20210424-r2","apk:linux-headers":"5.10.41-r0","apk:make":"4.3-r0","apk:mpc1":"1.2.1-r0","apk:mpfr4":"4.1.0-r0","apk:musl":"1.2.2-r3","apk:musl-dev":"1.2.2-r3","apk:nghttp2-libs":"1.43.0-r0","apk:openssl-dev":"1.1.1l-r0","apk:openssl-libs-static":"1.1.1l-r0","apk:patch":"2.7.6-r7","apk:pcre":"8.44-r0","apk:pcre-dev":"8.44-r0","apk:pcre2":"10.36-r0","apk:pkgconf":"1.7.4-r0","apk:tzdata":"2021e-r0","apk:zlib":"1.2.11-r3","apk:zlib-dev":"1.2.11-r3","apk:zlib-static":"1.2.11-r3","github:haproxy/haproxy":"v2.4.0","image:alpine":"f32b56d2f268e9a4dce68b95adac9f81b8ed8f9bc7fd33ba0d804e7426f603fd","image:base":"b2e73bfd5ae1ce12341dd36f24573d1316bf5172809c65dd669ff2720582967f"}

FROM reg.c5h.io/alpine@sha256:f32b56d2f268e9a4dce68b95adac9f81b8ed8f9bc7fd33ba0d804e7426f603fd AS build

ARG TAG="v2.4.0"

COPY buildstatic.patch /src/buildstatic.patch

RUN set -eux; \
     apk add --no-cache --virtual .build-deps \
        binutils=2.35.2-r2 \
        brotli-libs=1.0.9-r5 \
        build-base=0.5-r2 \
        busybox=1.33.1-r5 \
        ca-certificates=20191127-r5 \
        expat=2.4.1-r0 \
        file=5.40-r1 \
        fortify-headers=1.1-r1 \
        g++=10.3.1_git20210424-r2 \
        gcc=10.3.1_git20210424-r2 \
        git=2.32.0-r0 \
        gmp=6.2.1-r0 \
        isl22=0.22-r0 \
        libatomic=10.3.1_git20210424-r2 \
        libc-dev=0.7.2-r3 \
        libcrypto1.1=1.1.1l-r0 \
        libcurl=7.79.1-r0 \
        libgcc=10.3.1_git20210424-r2 \
        libgomp=10.3.1_git20210424-r2 \
        libgphobos=10.3.1_git20210424-r2 \
        libmagic=5.40-r1 \
        libpcre16=8.44-r0 \
        libpcre32=8.44-r0 \
        libpcrecpp=8.44-r0 \
        libssl1.1=1.1.1l-r0 \
        libstdc++=10.3.1_git20210424-r2 \
        linux-headers=5.10.41-r0 \
        make=4.3-r0 \
        mpc1=1.2.1-r0 \
        mpfr4=4.1.0-r0 \
        musl=1.2.2-r3 \
        musl-dev=1.2.2-r3 \
        nghttp2-libs=1.43.0-r0 \
        openssl-dev=1.1.1l-r0 \
        openssl-libs-static=1.1.1l-r0 \
        patch=2.7.6-r7 \
        pcre=8.44-r0 \
        pcre-dev=8.44-r0 \
        pcre2=10.36-r0 \
        pkgconf=1.7.4-r0 \
        tzdata=2021e-r0 \
        zlib=1.2.11-r3 \
        zlib-dev=1.2.11-r3 \
        zlib-static=1.2.11-r3 \
        ; \
     git clone --depth=1 -b $TAG --single-branch git://github.com/haproxy/haproxy /src/haproxy; \
     cd /src/haproxy; \
     patch Makefile < /src/buildstatic.patch; \
     make TARGET=linux-musl USE_OPENSSL=1 USE_PCRE=1; \
     make install-bin; \
     strip /usr/local/sbin/haproxy; \
     touch --date=@0 /usr/local/sbin/haproxy

FROM reg.c5h.io/base@sha256:b2e73bfd5ae1ce12341dd36f24573d1316bf5172809c65dd669ff2720582967f

STOPSIGNAL SIGUSR1

COPY --from=build /usr/local/sbin/haproxy /haproxy

ENTRYPOINT ["/haproxy", "-W", "-db"]

CMD ["-f", "/usr/local/etc/haproxy/haproxy.cfg"]
