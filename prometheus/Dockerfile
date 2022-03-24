# Generated from https://github.com/csmith/dockerfiles/blob/master/prometheus/Dockerfile.gotpl
# BOM: {"apk:bash":"5.1.16-r0","apk:binutils":"2.37-r3","apk:brotli-libs":"1.0.9-r5","apk:build-base":"0.5-r2","apk:busybox":"1.34.1-r4","apk:c-ares":"1.18.1-r0","apk:ca-certificates":"20211220-r0","apk:curl":"7.80.0-r0","apk:file":"5.41-r0","apk:fortify-headers":"1.1-r1","apk:g++":"10.3.1_git20211027-r0","apk:gcc":"10.3.1_git20211027-r0","apk:gmp":"6.2.1-r1","apk:isl22":"0.22-r0","apk:libacl":"2.2.53-r0","apk:libatomic":"10.3.1_git20211027-r0","apk:libc-dev":"0.7.2-r3","apk:libcrypto1.1":"1.1.1n-r0","apk:libcurl":"7.80.0-r0","apk:libgcc":"10.3.1_git20211027-r0","apk:libgomp":"10.3.1_git20211027-r0","apk:libgphobos":"10.3.1_git20211027-r0","apk:libmagic":"5.41-r0","apk:libssl1.1":"1.1.1n-r0","apk:libstdc++":"10.3.1_git20211027-r0","apk:make":"4.3-r0","apk:mpc1":"1.2.1-r0","apk:mpfr4":"4.1.0-r0","apk:musl":"1.2.2-r7","apk:musl-dev":"1.2.2-r7","apk:ncurses-libs":"6.3_p20211120-r0","apk:ncurses-terminfo-base":"6.3_p20211120-r0","apk:nghttp2-libs":"1.46.0-r0","apk:nodejs":"16.14.0-r0","apk:npm":"8.1.3-r0","apk:patch":"2.7.6-r7","apk:readline":"8.1.1-r0","apk:tar":"1.34-r0","apk:zlib":"1.2.11-r3","github:prometheus/prometheus":"v2.34.0","image:base":"f85825cc6e579732f339638d1cd16d87f0673759127df47e8bcd3586412acb5d","image:golang":"aaf0cd2c7afeee9750b1b891ec05210afa4a508ab9525cd102e5df363b7bcb2a"}

FROM reg.c5h.io/golang@sha256:aaf0cd2c7afeee9750b1b891ec05210afa4a508ab9525cd102e5df363b7bcb2a AS build

ARG TAG="v2.34.0"

RUN set -eux; \
    # Dependencies
    apk add --no-cache \
        bash=5.1.16-r0 \
        binutils=2.37-r3 \
        brotli-libs=1.0.9-r5 \
        build-base=0.5-r2 \
        busybox=1.34.1-r4 \
        c-ares=1.18.1-r0 \
        ca-certificates=20211220-r0 \
        curl=7.80.0-r0 \
        file=5.41-r0 \
        fortify-headers=1.1-r1 \
        g++=10.3.1_git20211027-r0 \
        gcc=10.3.1_git20211027-r0 \
        gmp=6.2.1-r1 \
        isl22=0.22-r0 \
        libacl=2.2.53-r0 \
        libatomic=10.3.1_git20211027-r0 \
        libc-dev=0.7.2-r3 \
        libcrypto1.1=1.1.1n-r0 \
        libcurl=7.80.0-r0 \
        libgcc=10.3.1_git20211027-r0 \
        libgomp=10.3.1_git20211027-r0 \
        libgphobos=10.3.1_git20211027-r0 \
        libmagic=5.41-r0 \
        libssl1.1=1.1.1n-r0 \
        libstdc++=10.3.1_git20211027-r0 \
        make=4.3-r0 \
        mpc1=1.2.1-r0 \
        mpfr4=4.1.0-r0 \
        musl=1.2.2-r7 \
        musl-dev=1.2.2-r7 \
        ncurses-libs=6.3_p20211120-r0 \
        ncurses-terminfo-base=6.3_p20211120-r0 \
        nghttp2-libs=1.46.0-r0 \
        nodejs=16.14.0-r0 \
        npm=8.1.3-r0 \
        patch=2.7.6-r7 \
        readline=8.1.1-r0 \
        tar=1.34-r0 \
        zlib=1.2.11-r3 \
        ;\
    # Build
    git clone --depth=1 -b $TAG --single-branch https://github.com/prometheus/prometheus /tmp/prometheus; \
    cd /tmp/prometheus; \
    make assets npm_licenses common-build; \
    # Fiddle with folders and prepare licenses
    mkdir -p /tmp/dist/etc/prometheus /tmp/dist/data /tmp/dist/notices; \
    cp -t /tmp/dist/notices/ /tmp/prometheus/LICENSE /tmp/prometheus/NOTICE /tmp/prometheus/npm_licenses.tar.bz2;

FROM reg.c5h.io/base@sha256:f85825cc6e579732f339638d1cd16d87f0673759127df47e8bcd3586412acb5d

COPY --from=build --chown=65532:65532 /tmp/dist/ /
COPY --from=build /tmp/prometheus/prometheus /prometheus

VOLUME /etc/prometheus /data

ENTRYPOINT ["/prometheus"]
CMD ["--config.file=/etc/prometheus/prometheus.yml", "--storage.tsdb.path=/data"]
