# Generated from https://github.com/csmith/dockerfiles/blob/master/nitter/Dockerfile.gotpl
# BOM: {"apk:bash":"5.1.8-r0","apk:binutils":"2.37-r3","apk:brotli-libs":"1.0.9-r5","apk:busybox":"1.34.1-r4","apk:ca-certificates":"20211220-r0","apk:curl":"7.80.0-r0","apk:expat":"2.4.4-r0","apk:gcc":"10.3.1_git20211027-r0","apk:git":"2.34.1-r0","apk:gmp":"6.2.1-r1","apk:isl22":"0.22-r0","apk:libatomic":"10.3.1_git20211027-r0","apk:libcrypto1.1":"1.1.1l-r8","apk:libcurl":"7.80.0-r0","apk:libgcc":"10.3.1_git20211027-r0","apk:libgomp":"10.3.1_git20211027-r0","apk:libgphobos":"10.3.1_git20211027-r0","apk:libpcre16":"8.45-r1","apk:libpcre32":"8.45-r1","apk:libpcrecpp":"8.45-r1","apk:libsass":"3.6.5-r0","apk:libsass-dev":"3.6.5-r0","apk:libssl1.1":"1.1.1l-r8","apk:libstdc++":"10.3.1_git20211027-r0","apk:mpc1":"1.2.1-r0","apk:mpfr4":"4.1.0-r0","apk:musl":"1.2.2-r7","apk:musl-dev":"1.2.2-r7","apk:ncurses-libs":"6.3_p20211120-r0","apk:ncurses-terminfo-base":"6.3_p20211120-r0","apk:nghttp2-libs":"1.46.0-r0","apk:openssl":"1.1.1l-r8","apk:openssl-dev":"1.1.1l-r8","apk:pcre":"8.45-r1","apk:pcre-dev":"8.45-r1","apk:pcre2":"10.39-r0","apk:pkgconf":"1.8.0-r0","apk:readline":"8.1.1-r0","apk:sqlite-libs":"3.36.0-r0","apk:zlib":"1.2.11-r3","github:nim-lang/nim":"v1.6.2","image:alpine":"ef95559f475cdf56e55db71c168f1616d497e2215bbd9f6bc63398036ae546a1","image:base":"f47ed6eddc8491fd9683c0bb66751d16b0dd6f9e920d96c9d028ec0001ae5968"}

FROM reg.c5h.io/alpine@sha256:ef95559f475cdf56e55db71c168f1616d497e2215bbd9f6bc63398036ae546a1 as nimbuild

ENV TAG=v1.6.2

ENV PATH=/opt/nim/bin:/opt/nimble/bin:$PATH

RUN set -eux; \
    # Dependencies and source for bootstrapping
    apk add --no-cache --virtual .build-deps \
        brotli-libs=1.0.9-r5 \
        busybox=1.34.1-r4 \
        ca-certificates=20211220-r0 \
        curl=7.80.0-r0 \
        libcrypto1.1=1.1.1l-r8 \
        libcurl=7.80.0-r0 \
        libssl1.1=1.1.1l-r8 \
        musl=1.2.2-r7 \
        nghttp2-libs=1.46.0-r0 \
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
        bash=5.1.8-r0 \
        binutils=2.37-r3 \
        brotli-libs=1.0.9-r5 \
        busybox=1.34.1-r4 \
        ca-certificates=20211220-r0 \
        expat=2.4.4-r0 \
        gcc=10.3.1_git20211027-r0 \
        git=2.34.1-r0 \
        gmp=6.2.1-r1 \
        isl22=0.22-r0 \
        libatomic=10.3.1_git20211027-r0 \
        libcrypto1.1=1.1.1l-r8 \
        libcurl=7.80.0-r0 \
        libgcc=10.3.1_git20211027-r0 \
        libgomp=10.3.1_git20211027-r0 \
        libgphobos=10.3.1_git20211027-r0 \
        libpcre16=8.45-r1 \
        libpcre32=8.45-r1 \
        libpcrecpp=8.45-r1 \
        libsass=3.6.5-r0 \
        libsass-dev=3.6.5-r0 \
        libssl1.1=1.1.1l-r8 \
        libstdc++=10.3.1_git20211027-r0 \
        mpc1=1.2.1-r0 \
        mpfr4=4.1.0-r0 \
        musl=1.2.2-r7 \
        musl-dev=1.2.2-r7 \
        ncurses-libs=6.3_p20211120-r0 \
        ncurses-terminfo-base=6.3_p20211120-r0 \
        nghttp2-libs=1.46.0-r0 \
        openssl=1.1.1l-r8 \
        openssl-dev=1.1.1l-r8 \
        pcre=8.45-r1 \
        pcre-dev=8.45-r1 \
        pcre2=10.39-r0 \
        pkgconf=1.8.0-r0 \
        readline=8.1.1-r0 \
        sqlite-libs=3.36.0-r0 \
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

FROM reg.c5h.io/base@sha256:f47ed6eddc8491fd9683c0bb66751d16b0dd6f9e920d96c9d028ec0001ae5968

COPY --from=build /src/libs/* /usr/lib/
COPY --from=build /src/nitter/nitter /home/nonroot/nitter
COPY --from=build /src/nitter/public /home/nonroot/public

WORKDIR /home/nonroot
ENTRYPOINT ["/home/nonroot/nitter"]
