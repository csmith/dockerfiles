# Generated from https://github.com/csmith/dockerfiles/blob/master/grafana/Dockerfile.gotpl
# BOM: {"apk:bash":"5.1.16-r2","apk:binutils":"2.38-r3","apk:brotli-libs":"1.0.9-r6","apk:build-base":"0.5-r3","apk:busybox":"1.35.0-r14","apk:c-ares":"1.18.1-r0","apk:ca-certificates":"20211220-r0","apk:curl":"7.83.1-r2","apk:file":"5.41-r0","apk:fortify-headers":"1.1-r1","apk:g++":"11.2.1_git20220219-r2","apk:gcc":"11.2.1_git20220219-r2","apk:gmp":"6.2.1-r2","apk:icu-data-en":"71.1-r2","apk:icu-libs":"71.1-r2","apk:isl22":"0.22-r0","apk:libacl":"2.3.1-r0","apk:libatomic":"11.2.1_git20220219-r2","apk:libc-dev":"0.7.2-r3","apk:libcrypto1.1":"1.1.1q-r0","apk:libcurl":"7.83.1-r2","apk:libgcc":"11.2.1_git20220219-r2","apk:libgomp":"11.2.1_git20220219-r2","apk:libmagic":"5.41-r0","apk:libssl1.1":"1.1.1q-r0","apk:libstdc++":"11.2.1_git20220219-r2","apk:make":"4.3-r0","apk:mpc1":"1.2.1-r0","apk:mpfr4":"4.1.0-r0","apk:musl":"1.2.3-r0","apk:musl-dev":"1.2.3-r0","apk:ncurses-libs":"6.3_p20220521-r0","apk:ncurses-terminfo-base":"6.3_p20220521-r0","apk:nghttp2-libs":"1.47.0-r0","apk:nodejs":"16.15.0-r1","apk:patch":"2.7.6-r7","apk:readline":"8.1.2-r0","apk:tar":"1.34-r0","apk:yarn":"1.22.19-r0","apk:zlib":"1.2.12-r1","github:grafana/grafana":"v9.0.2","image:base":"bffee9632ec79071de34e1a4a4f6c4563a71ad26d5923062bd9d0f9016f72beb","image:golang":"27fb5384caaaee607d36986fcb3bbd357693bc221a9e68022399aa1fa1d4f923"}

FROM reg.c5h.io/golang@sha256:27fb5384caaaee607d36986fcb3bbd357693bc221a9e68022399aa1fa1d4f923 AS build

ARG TAG="v9.0.2"

RUN set -eux; \
    # Dependencies
    apk add --no-cache \
        bash=5.1.16-r2 \
        binutils=2.38-r3 \
        brotli-libs=1.0.9-r6 \
        build-base=0.5-r3 \
        busybox=1.35.0-r14 \
        c-ares=1.18.1-r0 \
        ca-certificates=20211220-r0 \
        curl=7.83.1-r2 \
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
        libcrypto1.1=1.1.1q-r0 \
        libcurl=7.83.1-r2 \
        libgcc=11.2.1_git20220219-r2 \
        libgomp=11.2.1_git20220219-r2 \
        libmagic=5.41-r0 \
        libssl1.1=1.1.1q-r0 \
        libstdc++=11.2.1_git20220219-r2 \
        make=4.3-r0 \
        mpc1=1.2.1-r0 \
        mpfr4=4.1.0-r0 \
        musl=1.2.3-r0 \
        musl-dev=1.2.3-r0 \
        ncurses-libs=6.3_p20220521-r0 \
        ncurses-terminfo-base=6.3_p20220521-r0 \
        nghttp2-libs=1.47.0-r0 \
        nodejs=16.15.0-r1 \
        patch=2.7.6-r7 \
        readline=8.1.2-r0 \
        tar=1.34-r0 \
        yarn=1.22.19-r0 \
        zlib=1.2.12-r1 \
        ;\
    # Build
    git clone --depth=1 -b $TAG --single-branch https://github.com/grafana/grafana /tmp/grafana; \
    cd /tmp/grafana; \
    yarn install; \
    NODE_ENV=production NODE_OPTIONS=--max_old_space_size=6144 yarn build; \
    make build-go; \
    # Fiddle with folders and prepare licenses
    mkdir -p \
        /tmp/dist/etc/grafana/provisioning/dashboards \
        /tmp/dist/etc/grafana/provisioning/datasources \
        /tmp/dist/etc/grafana/provisioning/notifiers \
        /tmp/dist/etc/grafana/provisioning/plugins \
        /tmp/dist/etc/grafana/provisioning/access-control \
        /tmp/dist/var/log/grafana \
        /tmp/dist/var/lib/grafana/plugins \
        /tmp/dist/usr/share/grafana; \
    mv -t /tmp/dist/usr/share/grafana /tmp/grafana/public/ /tmp/grafana/tools/ /tmp/grafana/conf/; \
    mv /tmp/dist/usr/share/grafana/conf/sample.ini /tmp/dist/etc/grafana/grafana.ini;

FROM reg.c5h.io/base@sha256:bffee9632ec79071de34e1a4a4f6c4563a71ad26d5923062bd9d0f9016f72beb

COPY --from=build --chown=65532:65532 /tmp/dist/ /
COPY --from=build /tmp/grafana/bin/*/grafana-server /grafana-server

VOLUME /var/lib/grafana /etc/grafana

ENTRYPOINT ["/grafana-server"]
CMD [ \
    "--homepath=/usr/share/grafana/", \
    "--config=/etc/grafana/grafana.ini", \
    "--packaging=docker", \
    "cfg:default.log.mode=console", \
    "cfg:default.paths.data=/var/lib/grafana", \
    "cfg:default.paths.logs=/var/log/grafana", \
    "cfg:default.paths.plugins=/var/lib/grafana/plugins", \
    "cfg:default.paths.provisioning=/etc/grafana/provisioning" \
]
