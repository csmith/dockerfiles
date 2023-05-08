# Generated from https://github.com/csmith/dockerfiles/blob/master/prometheus/Dockerfile.gotpl
# BOM: {"apk:bash":"5.2.15-r0","apk:binutils":"2.39-r2","apk:brotli-libs":"1.0.9-r9","apk:build-base":"0.5-r3","apk:busybox":"1.35.0-r29","apk:busybox-binsh":"1.35.0-r29","apk:c-ares":"1.18.1-r1","apk:ca-certificates":"20230506-r0","apk:curl":"8.0.1-r0","apk:file":"5.43-r0","apk:fortify-headers":"1.1-r1","apk:g++":"12.2.1_git20220924-r4","apk:gcc":"12.2.1_git20220924-r4","apk:gmp":"6.2.1-r2","apk:icu-data-en":"72.1-r1","apk:icu-libs":"72.1-r1","apk:isl24":"0.24-r1","apk:libacl":"2.3.1-r1","apk:libatomic":"12.2.1_git20220924-r4","apk:libc-dev":"0.7.2-r3","apk:libcrypto3":"3.0.8-r4","apk:libcurl":"8.0.1-r0","apk:libgcc":"12.2.1_git20220924-r4","apk:libgomp":"12.2.1_git20220924-r4","apk:libmagic":"5.43-r0","apk:libssl3":"3.0.8-r4","apk:libstdc++":"12.2.1_git20220924-r4","apk:libstdc++-dev":"12.2.1_git20220924-r4","apk:make":"4.3-r1","apk:mpc1":"1.2.1-r1","apk:mpfr4":"4.1.0-r0","apk:musl":"1.2.3-r4","apk:musl-dev":"1.2.3-r4","apk:ncurses-libs":"6.3_p20221119-r0","apk:ncurses-terminfo-base":"6.3_p20221119-r0","apk:nghttp2-libs":"1.51.0-r0","apk:nodejs":"18.14.2-r0","apk:npm":"9.1.2-r0","apk:patch":"2.7.6-r9","apk:readline":"8.2.0-r0","apk:tar":"1.34-r2","apk:zlib":"1.2.13-r0","github:prometheus/prometheus":"v2.43.1+stringlabels","image:base":"06668724f8daa1866a72d07c44d9c0fe5c42bfd277d47fecc8e4ed037a5d0e84","image:golang":"8cca8e81ef263f66855ba7ad711e884e03e7670276e1cd2011d8f0f3f0eec7e0"}

FROM reg.c5h.io/golang@sha256:8cca8e81ef263f66855ba7ad711e884e03e7670276e1cd2011d8f0f3f0eec7e0 AS build

ARG TAG="v2.43.1+stringlabels"

RUN set -eux; \
    # Dependencies
    apk add --no-cache \
        bash=5.2.15-r0 \
        binutils=2.39-r2 \
        brotli-libs=1.0.9-r9 \
        build-base=0.5-r3 \
        busybox=1.35.0-r29 \
        busybox-binsh=1.35.0-r29 \
        c-ares=1.18.1-r1 \
        ca-certificates=20230506-r0 \
        curl=8.0.1-r0 \
        file=5.43-r0 \
        fortify-headers=1.1-r1 \
        g++=12.2.1_git20220924-r4 \
        gcc=12.2.1_git20220924-r4 \
        gmp=6.2.1-r2 \
        icu-data-en=72.1-r1 \
        icu-libs=72.1-r1 \
        isl24=0.24-r1 \
        libacl=2.3.1-r1 \
        libatomic=12.2.1_git20220924-r4 \
        libc-dev=0.7.2-r3 \
        libcrypto3=3.0.8-r4 \
        libcurl=8.0.1-r0 \
        libgcc=12.2.1_git20220924-r4 \
        libgomp=12.2.1_git20220924-r4 \
        libmagic=5.43-r0 \
        libssl3=3.0.8-r4 \
        libstdc++=12.2.1_git20220924-r4 \
        libstdc++-dev=12.2.1_git20220924-r4 \
        make=4.3-r1 \
        mpc1=1.2.1-r1 \
        mpfr4=4.1.0-r0 \
        musl=1.2.3-r4 \
        musl-dev=1.2.3-r4 \
        ncurses-libs=6.3_p20221119-r0 \
        ncurses-terminfo-base=6.3_p20221119-r0 \
        nghttp2-libs=1.51.0-r0 \
        nodejs=18.14.2-r0 \
        npm=9.1.2-r0 \
        patch=2.7.6-r9 \
        readline=8.2.0-r0 \
        tar=1.34-r2 \
        zlib=1.2.13-r0 \
        ;\
    # Build
    git clone --depth=1 -b $TAG --single-branch https://github.com/prometheus/prometheus /tmp/prometheus; \
    cd /tmp/prometheus; \
    make build npm_licenses; \
    # Fiddle with folders and prepare licenses
    mkdir -p /tmp/dist/etc/prometheus /tmp/dist/data /tmp/dist/notices; \
    cp -t /tmp/dist/notices/ /tmp/prometheus/LICENSE /tmp/prometheus/NOTICE /tmp/prometheus/npm_licenses.tar.bz2;

FROM reg.c5h.io/base@sha256:06668724f8daa1866a72d07c44d9c0fe5c42bfd277d47fecc8e4ed037a5d0e84

COPY --from=build --chown=65532:65532 /tmp/dist/ /
COPY --from=build /tmp/prometheus/prometheus /prometheus

VOLUME /etc/prometheus /data

ENTRYPOINT ["/prometheus"]
CMD ["--config.file=/etc/prometheus/prometheus.yml", "--storage.tsdb.path=/data"]
