# Generated from https://github.com/csmith/dockerfiles/blob/master/nitter/Dockerfile.gotpl
# BOM: {"apk:bash":"5.1.4-r0","apk:binutils":"2.35.2-r2","apk:brotli-libs":"1.0.9-r5","apk:busybox":"1.33.1-r3","apk:ca-certificates":"20191127-r5","apk:curl":"7.79.1-r0","apk:expat":"2.4.1-r0","apk:gcc":"10.3.1_git20210424-r2","apk:git":"2.32.0-r0","apk:gmp":"6.2.1-r0","apk:isl22":"0.22-r0","apk:libatomic":"10.3.1_git20210424-r2","apk:libcrypto1.1":"1.1.1l-r0","apk:libcurl":"7.79.1-r0","apk:libgcc":"10.3.1_git20210424-r2","apk:libgomp":"10.3.1_git20210424-r2","apk:libgphobos":"10.3.1_git20210424-r2","apk:libpcre16":"8.44-r0","apk:libpcre32":"8.44-r0","apk:libpcrecpp":"8.44-r0","apk:libsass":"3.6.5-r0","apk:libsass-dev":"3.6.5-r0","apk:libssl1.1":"1.1.1l-r0","apk:libstdc++":"10.3.1_git20210424-r2","apk:mpc1":"1.2.1-r0","apk:mpfr4":"4.1.0-r0","apk:musl":"1.2.2-r3","apk:musl-dev":"1.2.2-r3","apk:ncurses-libs":"6.2_p20210612-r0","apk:ncurses-terminfo-base":"6.2_p20210612-r0","apk:nghttp2-libs":"1.43.0-r0","apk:openssl":"1.1.1l-r0","apk:openssl-dev":"1.1.1l-r0","apk:pcre":"8.44-r0","apk:pcre-dev":"8.44-r0","apk:pcre2":"10.36-r0","apk:pkgconf":"1.7.4-r0","apk:readline":"8.1.0-r0","apk:sqlite-libs":"3.35.5-r0","apk:zlib":"1.2.11-r3","github:nim-lang/nim":"v1.6.0","image:alpine":"sha256:f32b56d2f268e9a4dce68b95adac9f81b8ed8f9bc7fd33ba0d804e7426f603fd","image:base":"sha256:e211f9db3dc7b27fdb9ce4c8155f980370daf4f68b7411a144e82f9d44f16022"}

FROM reg.c5h.io/alpine@sha256:f32b56d2f268e9a4dce68b95adac9f81b8ed8f9bc7fd33ba0d804e7426f603fd as nimbuild

ENV TAG=v1.6.0

ENV PATH=/opt/nim/bin:/opt/nimble/bin:$PATH

RUN set -eux; \
    # Dependencies and source for bootstrapping
    apk add --no-cache --virtual .build-deps \
        brotli-libs=1.0.9-r5 \
        busybox=1.33.1-r3 \
        ca-certificates=20191127-r5 \
        curl=7.79.1-r0 \
        libcrypto1.1=1.1.1l-r0 \
        libcurl=7.79.1-r0 \
        libssl1.1=1.1.1l-r0 \
        musl=1.2.2-r3 \
        nghttp2-libs=1.43.0-r0 \
        zlib=1.2.11-r3 \
        ; \
    curl -sfLo nim.tar.xz https://nim-lang.org/download/nim-${TAG:1}-linux_x64.tar.xz; \
    curl -sfLo nim.tar.xz.sha256 https://nim-lang.org/download/nim-${TAG:1}-linux_x64.tar.xz.sha256; \
    echo "$(awk '{print $1}' nim.tar.xz.sha256) *nim.tar.xz" | sha256sum -wc -; \
    tar -C /tmp -xf nim.tar.xz; \
    mv /tmp/nim-* /tmp/nim; \
    cd /tmp/nim; \
    ./install.sh /usr/bin; \
    cp /tmp/nim/bin/nimble /usr/bin/nimble;

FROM nimbuild as build

#No tags on the project, will just have to update this as needed until some tags are made
ENV TAG=5caf77481f70be5118edf885e9f4417a26074edb

RUN set -eux; \
    apk add --no-cache \
        bash=5.1.4-r0 \
        binutils=2.35.2-r2 \
        brotli-libs=1.0.9-r5 \
        busybox=1.33.1-r3 \
        ca-certificates=20191127-r5 \
        expat=2.4.1-r0 \
        gcc=10.3.1_git20210424-r2 \
        git=2.32.0-r0 \
        gmp=6.2.1-r0 \
        isl22=0.22-r0 \
        libatomic=10.3.1_git20210424-r2 \
        libcrypto1.1=1.1.1l-r0 \
        libcurl=7.79.1-r0 \
        libgcc=10.3.1_git20210424-r2 \
        libgomp=10.3.1_git20210424-r2 \
        libgphobos=10.3.1_git20210424-r2 \
        libpcre16=8.44-r0 \
        libpcre32=8.44-r0 \
        libpcrecpp=8.44-r0 \
        libsass=3.6.5-r0 \
        libsass-dev=3.6.5-r0 \
        libssl1.1=1.1.1l-r0 \
        libstdc++=10.3.1_git20210424-r2 \
        mpc1=1.2.1-r0 \
        mpfr4=4.1.0-r0 \
        musl=1.2.2-r3 \
        musl-dev=1.2.2-r3 \
        ncurses-libs=6.2_p20210612-r0 \
        ncurses-terminfo-base=6.2_p20210612-r0 \
        nghttp2-libs=1.43.0-r0 \
        openssl=1.1.1l-r0 \
        openssl-dev=1.1.1l-r0 \
        pcre=8.44-r0 \
        pcre-dev=8.44-r0 \
        pcre2=10.36-r0 \
        pkgconf=1.7.4-r0 \
        readline=8.1.0-r0 \
        sqlite-libs=3.35.5-r0 \
        zlib=1.2.11-r3 \
        ; \
    git clone -b master --single-branch git://github.com/zedeus/nitter /src/nitter; \
    mkdir -p /src/libs; \
    cp /usr/lib/libsqlite3* /src/libs; \
    cp /usr/lib/libssl* /src/libs; \
    cp /usr/lib/libcrypto* /src/libs; \
    cp /usr/lib/libpcre* /src/libs; \
    cd /src/nitter; \
    git checkout $TAG; \
    nimble build -y -d:release; \
    strip -s nitter; \
    nimble scss

FROM reg.c5h.io/base@sha256:e211f9db3dc7b27fdb9ce4c8155f980370daf4f68b7411a144e82f9d44f16022

COPY --from=build /src/libs/* /usr/lib/
COPY --from=build /src/nitter/nitter /home/nonroot/nitter
COPY --from=build /src/nitter/public /home/nonroot/public

WORKDIR /home/nonroot
ENTRYPOINT ["/home/nonroot/nitter"]
