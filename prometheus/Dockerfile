# Generated from https://github.com/csmith/dockerfiles/blob/master/prometheus/Dockerfile.gotpl
# BOM: {"apk:ada-libs":"2.7.4-r0","apk:bash":"5.2.21-r0","apk:binutils":"2.41-r0","apk:brotli-libs":"1.1.0-r1","apk:build-base":"0.5-r3","apk:busybox":"1.36.1-r15","apk:busybox-binsh":"1.36.1-r15","apk:c-ares":"1.22.1-r0","apk:ca-certificates":"20230506-r0","apk:curl":"8.5.0-r0","apk:file":"5.45-r1","apk:fortify-headers":"1.1-r3","apk:g++":"13.2.1_git20231014-r0","apk:gcc":"13.2.1_git20231014-r0","apk:gmp":"6.3.0-r0","apk:icu-data-en":"74.1-r0","apk:icu-libs":"74.1-r0","apk:isl25":"0.25-r2","apk:jansson":"2.14-r4","apk:libacl":"2.3.1-r4","apk:libatomic":"13.2.1_git20231014-r0","apk:libbase64":"0.5.0-r0","apk:libc-dev":"0.7.2-r5","apk:libcrypto3":"3.1.4-r2","apk:libcurl":"8.5.0-r0","apk:libgcc":"13.2.1_git20231014-r0","apk:libgomp":"13.2.1_git20231014-r0","apk:libidn2":"2.3.4-r4","apk:libmagic":"5.45-r1","apk:libncursesw":"6.4_p20231125-r0","apk:libssl3":"3.1.4-r2","apk:libstdc++":"13.2.1_git20231014-r0","apk:libstdc++-dev":"13.2.1_git20231014-r0","apk:libunistring":"1.1-r2","apk:make":"4.4.1-r2","apk:mpc1":"1.3.1-r1","apk:mpfr4":"4.2.1-r0","apk:musl":"1.2.4_git20230717-r4","apk:musl-dev":"1.2.4_git20230717-r4","apk:ncurses-terminfo-base":"6.4_p20231125-r0","apk:nghttp2-libs":"1.58.0-r0","apk:nodejs":"20.10.0-r1","apk:npm":"10.2.5-r0","apk:patch":"2.7.6-r10","apk:readline":"8.2.1-r2","apk:tar":"1.35-r2","apk:zlib":"1.3-r2","apk:zstd-libs":"1.5.5-r8","github:prometheus/prometheus":"v2.48.1","image:base":"cfa1d5f949b495a6e077bdddb1dbbc3d9a9d55036ab1267f4168a6a5c7e7fdce","image:golang":"72e140a52eb30250acb9f89a7ca4c1ce30ffcffb46d5ba1070f5f73afdf7bf78"}

FROM reg.c5h.io/golang@sha256:72e140a52eb30250acb9f89a7ca4c1ce30ffcffb46d5ba1070f5f73afdf7bf78 AS build

ARG TAG="v2.48.1"

RUN set -eux; \
    # Dependencies
    apk add --no-cache \
        ada-libs=2.7.4-r0 \
        bash=5.2.21-r0 \
        binutils=2.41-r0 \
        brotli-libs=1.1.0-r1 \
        build-base=0.5-r3 \
        busybox=1.36.1-r15 \
        busybox-binsh=1.36.1-r15 \
        c-ares=1.22.1-r0 \
        ca-certificates=20230506-r0 \
        curl=8.5.0-r0 \
        file=5.45-r1 \
        fortify-headers=1.1-r3 \
        g++=13.2.1_git20231014-r0 \
        gcc=13.2.1_git20231014-r0 \
        gmp=6.3.0-r0 \
        icu-data-en=74.1-r0 \
        icu-libs=74.1-r0 \
        isl25=0.25-r2 \
        jansson=2.14-r4 \
        libacl=2.3.1-r4 \
        libatomic=13.2.1_git20231014-r0 \
        libbase64=0.5.0-r0 \
        libc-dev=0.7.2-r5 \
        libcrypto3=3.1.4-r2 \
        libcurl=8.5.0-r0 \
        libgcc=13.2.1_git20231014-r0 \
        libgomp=13.2.1_git20231014-r0 \
        libidn2=2.3.4-r4 \
        libmagic=5.45-r1 \
        libncursesw=6.4_p20231125-r0 \
        libssl3=3.1.4-r2 \
        libstdc++=13.2.1_git20231014-r0 \
        libstdc++-dev=13.2.1_git20231014-r0 \
        libunistring=1.1-r2 \
        make=4.4.1-r2 \
        mpc1=1.3.1-r1 \
        mpfr4=4.2.1-r0 \
        musl=1.2.4_git20230717-r4 \
        musl-dev=1.2.4_git20230717-r4 \
        ncurses-terminfo-base=6.4_p20231125-r0 \
        nghttp2-libs=1.58.0-r0 \
        nodejs=20.10.0-r1 \
        npm=10.2.5-r0 \
        patch=2.7.6-r10 \
        readline=8.2.1-r2 \
        tar=1.35-r2 \
        zlib=1.3-r2 \
        zstd-libs=1.5.5-r8 \
        ;\
    # Build
    git clone --depth=1 -b $TAG --single-branch https://github.com/prometheus/prometheus /tmp/prometheus; \
    cd /tmp/prometheus; \
    make build npm_licenses; \
    # Fiddle with folders and prepare licenses
    mkdir -p /tmp/dist/etc/prometheus /tmp/dist/data /tmp/dist/notices; \
    cp -t /tmp/dist/notices/ /tmp/prometheus/LICENSE /tmp/prometheus/NOTICE /tmp/prometheus/npm_licenses.tar.bz2;

FROM reg.c5h.io/base@sha256:cfa1d5f949b495a6e077bdddb1dbbc3d9a9d55036ab1267f4168a6a5c7e7fdce

COPY --from=build --chown=65532:65532 /tmp/dist/ /
COPY --from=build /tmp/prometheus/prometheus /prometheus

VOLUME /etc/prometheus /data

ENTRYPOINT ["/prometheus"]
CMD ["--config.file=/etc/prometheus/prometheus.yml", "--storage.tsdb.path=/data"]
