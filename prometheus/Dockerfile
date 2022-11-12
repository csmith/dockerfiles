# Generated from https://github.com/csmith/dockerfiles/blob/master/prometheus/Dockerfile.gotpl
# BOM: {"apk:bash":"5.1.16-r2","apk:binutils":"2.38-r3","apk:brotli-libs":"1.0.9-r6","apk:build-base":"0.5-r3","apk:busybox":"1.35.0-r17","apk:c-ares":"1.18.1-r0","apk:ca-certificates":"20220614-r0","apk:curl":"7.83.1-r4","apk:file":"5.41-r0","apk:fortify-headers":"1.1-r1","apk:g++":"11.2.1_git20220219-r2","apk:gcc":"11.2.1_git20220219-r2","apk:gmp":"6.2.1-r2","apk:icu-data-en":"71.1-r2","apk:icu-libs":"71.1-r2","apk:isl22":"0.22-r0","apk:libacl":"2.3.1-r0","apk:libatomic":"11.2.1_git20220219-r2","apk:libc-dev":"0.7.2-r3","apk:libcrypto1.1":"1.1.1s-r0","apk:libcurl":"7.83.1-r4","apk:libgcc":"11.2.1_git20220219-r2","apk:libgomp":"11.2.1_git20220219-r2","apk:libmagic":"5.41-r0","apk:libssl1.1":"1.1.1s-r0","apk:libstdc++":"11.2.1_git20220219-r2","apk:make":"4.3-r0","apk:mpc1":"1.2.1-r0","apk:mpfr4":"4.1.0-r0","apk:musl":"1.2.3-r1","apk:musl-dev":"1.2.3-r1","apk:ncurses-libs":"6.3_p20220521-r0","apk:ncurses-terminfo-base":"6.3_p20220521-r0","apk:nghttp2-libs":"1.47.0-r0","apk:nodejs":"16.17.1-r0","apk:npm":"8.10.0-r0","apk:patch":"2.7.6-r7","apk:readline":"8.1.2-r0","apk:tar":"1.34-r0","apk:zlib":"1.2.12-r3","github:prometheus/prometheus":"v2.40.1","image:base":"2ad8fb5e47298cb81072a36ee939a2024bc3acb6221b105cb3928c0e6a97c665","image:golang":"dc4eac488c6c21ead1b8d08cf44dfa9e8d33ccec71eb063fe8b3e2ede109d082"}

FROM reg.c5h.io/golang@sha256:dc4eac488c6c21ead1b8d08cf44dfa9e8d33ccec71eb063fe8b3e2ede109d082 AS build

ARG TAG="v2.40.1"

RUN set -eux; \
    # Dependencies
    apk add --no-cache \
        bash=5.1.16-r2 \
        binutils=2.38-r3 \
        brotli-libs=1.0.9-r6 \
        build-base=0.5-r3 \
        busybox=1.35.0-r17 \
        c-ares=1.18.1-r0 \
        ca-certificates=20220614-r0 \
        curl=7.83.1-r4 \
        file=5.41-r0 \
        fortify-headers=1.1-r1 \
        g++=11.2.1_git20220219-r2 \
        gcc=11.2.1_git20220219-r2 \
        gmp=6.2.1-r2 \
        icu-data-en=71.1-r2 \
        icu-libs=71.1-r2 \
        isl22=0.22-r0 \
        libacl=2.3.1-r0 \
        libatomic=11.2.1_git20220219-r2 \
        libc-dev=0.7.2-r3 \
        libcrypto1.1=1.1.1s-r0 \
        libcurl=7.83.1-r4 \
        libgcc=11.2.1_git20220219-r2 \
        libgomp=11.2.1_git20220219-r2 \
        libmagic=5.41-r0 \
        libssl1.1=1.1.1s-r0 \
        libstdc++=11.2.1_git20220219-r2 \
        make=4.3-r0 \
        mpc1=1.2.1-r0 \
        mpfr4=4.1.0-r0 \
        musl=1.2.3-r1 \
        musl-dev=1.2.3-r1 \
        ncurses-libs=6.3_p20220521-r0 \
        ncurses-terminfo-base=6.3_p20220521-r0 \
        nghttp2-libs=1.47.0-r0 \
        nodejs=16.17.1-r0 \
        npm=8.10.0-r0 \
        patch=2.7.6-r7 \
        readline=8.1.2-r0 \
        tar=1.34-r0 \
        zlib=1.2.12-r3 \
        ;\
    # Build
    git clone --depth=1 -b $TAG --single-branch https://github.com/prometheus/prometheus /tmp/prometheus; \
    cd /tmp/prometheus; \
    make build npm_licenses; \
    # Fiddle with folders and prepare licenses
    mkdir -p /tmp/dist/etc/prometheus /tmp/dist/data /tmp/dist/notices; \
    cp -t /tmp/dist/notices/ /tmp/prometheus/LICENSE /tmp/prometheus/NOTICE /tmp/prometheus/npm_licenses.tar.bz2;

FROM reg.c5h.io/base@sha256:2ad8fb5e47298cb81072a36ee939a2024bc3acb6221b105cb3928c0e6a97c665

COPY --from=build --chown=65532:65532 /tmp/dist/ /
COPY --from=build /tmp/prometheus/prometheus /prometheus

VOLUME /etc/prometheus /data

ENTRYPOINT ["/prometheus"]
CMD ["--config.file=/etc/prometheus/prometheus.yml", "--storage.tsdb.path=/data"]
