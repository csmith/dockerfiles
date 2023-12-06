# Generated from https://github.com/csmith/dockerfiles/blob/master/prometheus/Dockerfile.gotpl
# BOM: {"apk:bash":"5.2.15-r5","apk:binutils":"2.40-r7","apk:brotli-libs":"1.0.9-r14","apk:build-base":"0.5-r3","apk:busybox":"1.36.1-r5","apk:busybox-binsh":"1.36.1-r5","apk:c-ares":"1.19.1-r0","apk:ca-certificates":"20230506-r0","apk:curl":"8.4.0-r0","apk:file":"5.45-r0","apk:fortify-headers":"1.1-r3","apk:g++":"12.2.1_git20220924-r10","apk:gcc":"12.2.1_git20220924-r10","apk:gmp":"6.2.1-r3","apk:icu-data-en":"73.2-r2","apk:icu-libs":"73.2-r2","apk:isl25":"0.25-r2","apk:libacl":"2.3.1-r3","apk:libatomic":"12.2.1_git20220924-r10","apk:libc-dev":"0.7.2-r5","apk:libcrypto3":"3.1.4-r1","apk:libcurl":"8.4.0-r0","apk:libgcc":"12.2.1_git20220924-r10","apk:libgomp":"12.2.1_git20220924-r10","apk:libidn2":"2.3.4-r1","apk:libmagic":"5.45-r0","apk:libncursesw":"6.4_p20230506-r0","apk:libssl3":"3.1.4-r1","apk:libstdc++":"12.2.1_git20220924-r10","apk:libstdc++-dev":"12.2.1_git20220924-r10","apk:libunistring":"1.1-r1","apk:make":"4.4.1-r1","apk:mpc1":"1.3.1-r1","apk:mpfr4":"4.2.0_p12-r0","apk:musl":"1.2.4-r2","apk:musl-dev":"1.2.4-r2","apk:ncurses-terminfo-base":"6.4_p20230506-r0","apk:nghttp2-libs":"1.57.0-r0","apk:nodejs":"18.18.2-r0","apk:npm":"9.6.6-r0","apk:patch":"2.7.6-r10","apk:readline":"8.2.1-r1","apk:tar":"1.34-r3","apk:zlib":"1.2.13-r1","apk:zstd-libs":"1.5.5-r4","github:prometheus/prometheus":"v2.48.0","image:base":"7442771fc9e4d70fbf8f9d592e28b9ec7b9b81af30ddb87cbaaba90f2cc0064f","image:golang":"5e230c03f6aaf29b89f73ae91102d1b2740706fc131fff0dedc3f72bf189adba"}

FROM reg.c5h.io/golang@sha256:5e230c03f6aaf29b89f73ae91102d1b2740706fc131fff0dedc3f72bf189adba AS build

ARG TAG="v2.48.0"

RUN set -eux; \
    # Dependencies
    apk add --no-cache \
        bash=5.2.15-r5 \
        binutils=2.40-r7 \
        brotli-libs=1.0.9-r14 \
        build-base=0.5-r3 \
        busybox=1.36.1-r5 \
        busybox-binsh=1.36.1-r5 \
        c-ares=1.19.1-r0 \
        ca-certificates=20230506-r0 \
        curl=8.4.0-r0 \
        file=5.45-r0 \
        fortify-headers=1.1-r3 \
        g++=12.2.1_git20220924-r10 \
        gcc=12.2.1_git20220924-r10 \
        gmp=6.2.1-r3 \
        icu-data-en=73.2-r2 \
        icu-libs=73.2-r2 \
        isl25=0.25-r2 \
        libacl=2.3.1-r3 \
        libatomic=12.2.1_git20220924-r10 \
        libc-dev=0.7.2-r5 \
        libcrypto3=3.1.4-r1 \
        libcurl=8.4.0-r0 \
        libgcc=12.2.1_git20220924-r10 \
        libgomp=12.2.1_git20220924-r10 \
        libidn2=2.3.4-r1 \
        libmagic=5.45-r0 \
        libncursesw=6.4_p20230506-r0 \
        libssl3=3.1.4-r1 \
        libstdc++=12.2.1_git20220924-r10 \
        libstdc++-dev=12.2.1_git20220924-r10 \
        libunistring=1.1-r1 \
        make=4.4.1-r1 \
        mpc1=1.3.1-r1 \
        mpfr4=4.2.0_p12-r0 \
        musl=1.2.4-r2 \
        musl-dev=1.2.4-r2 \
        ncurses-terminfo-base=6.4_p20230506-r0 \
        nghttp2-libs=1.57.0-r0 \
        nodejs=18.18.2-r0 \
        npm=9.6.6-r0 \
        patch=2.7.6-r10 \
        readline=8.2.1-r1 \
        tar=1.34-r3 \
        zlib=1.2.13-r1 \
        zstd-libs=1.5.5-r4 \
        ;\
    # Build
    git clone --depth=1 -b $TAG --single-branch https://github.com/prometheus/prometheus /tmp/prometheus; \
    cd /tmp/prometheus; \
    make build npm_licenses; \
    # Fiddle with folders and prepare licenses
    mkdir -p /tmp/dist/etc/prometheus /tmp/dist/data /tmp/dist/notices; \
    cp -t /tmp/dist/notices/ /tmp/prometheus/LICENSE /tmp/prometheus/NOTICE /tmp/prometheus/npm_licenses.tar.bz2;

FROM reg.c5h.io/base@sha256:7442771fc9e4d70fbf8f9d592e28b9ec7b9b81af30ddb87cbaaba90f2cc0064f

COPY --from=build --chown=65532:65532 /tmp/dist/ /
COPY --from=build /tmp/prometheus/prometheus /prometheus

VOLUME /etc/prometheus /data

ENTRYPOINT ["/prometheus"]
CMD ["--config.file=/etc/prometheus/prometheus.yml", "--storage.tsdb.path=/data"]
