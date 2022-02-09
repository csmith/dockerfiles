# Generated from https://github.com/csmith/dockerfiles/blob/master/grafana/Dockerfile.gotpl
# BOM: {"apk:bash":"5.1.8-r0","apk:binutils":"2.37-r3","apk:brotli-libs":"1.0.9-r5","apk:build-base":"0.5-r2","apk:busybox":"1.34.1-r4","apk:c-ares":"1.18.1-r0","apk:ca-certificates":"20211220-r0","apk:curl":"7.80.0-r0","apk:file":"5.41-r0","apk:fortify-headers":"1.1-r1","apk:g++":"10.3.1_git20211027-r0","apk:gcc":"10.3.1_git20211027-r0","apk:gmp":"6.2.1-r1","apk:isl22":"0.22-r0","apk:libacl":"2.2.53-r0","apk:libatomic":"10.3.1_git20211027-r0","apk:libc-dev":"0.7.2-r3","apk:libcrypto1.1":"1.1.1l-r8","apk:libcurl":"7.80.0-r0","apk:libgcc":"10.3.1_git20211027-r0","apk:libgomp":"10.3.1_git20211027-r0","apk:libgphobos":"10.3.1_git20211027-r0","apk:libmagic":"5.41-r0","apk:libssl1.1":"1.1.1l-r8","apk:libstdc++":"10.3.1_git20211027-r0","apk:make":"4.3-r0","apk:mpc1":"1.2.1-r0","apk:mpfr4":"4.1.0-r0","apk:musl":"1.2.2-r7","apk:musl-dev":"1.2.2-r7","apk:ncurses-libs":"6.3_p20211120-r0","apk:ncurses-terminfo-base":"6.3_p20211120-r0","apk:nghttp2-libs":"1.46.0-r0","apk:nodejs":"16.13.2-r0","apk:patch":"2.7.6-r7","apk:readline":"8.1.1-r0","apk:tar":"1.34-r0","apk:yarn":"1.22.17-r0","apk:zlib":"1.2.11-r3","github:grafana/grafana":"v8.3.5","image:base":"f47ed6eddc8491fd9683c0bb66751d16b0dd6f9e920d96c9d028ec0001ae5968","image:golang":"ac8fa5f4078b0a697796b5d74158ab4e73b5f743433f621d89abb7c5faf94f1f"}

FROM reg.c5h.io/golang@sha256:ac8fa5f4078b0a697796b5d74158ab4e73b5f743433f621d89abb7c5faf94f1f AS build

ARG TAG="v8.3.5"

RUN set -eux; \
    # Dependencies
    apk add --no-cache \
        bash=5.1.8-r0 \
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
        libcrypto1.1=1.1.1l-r8 \
        libcurl=7.80.0-r0 \
        libgcc=10.3.1_git20211027-r0 \
        libgomp=10.3.1_git20211027-r0 \
        libgphobos=10.3.1_git20211027-r0 \
        libmagic=5.41-r0 \
        libssl1.1=1.1.1l-r8 \
        libstdc++=10.3.1_git20211027-r0 \
        make=4.3-r0 \
        mpc1=1.2.1-r0 \
        mpfr4=4.1.0-r0 \
        musl=1.2.2-r7 \
        musl-dev=1.2.2-r7 \
        ncurses-libs=6.3_p20211120-r0 \
        ncurses-terminfo-base=6.3_p20211120-r0 \
        nghttp2-libs=1.46.0-r0 \
        nodejs=16.13.2-r0 \
        patch=2.7.6-r7 \
        readline=8.1.1-r0 \
        tar=1.34-r0 \
        yarn=1.22.17-r0 \
        zlib=1.2.11-r3 \
        ;\
    # Build
    git clone --depth=1 -b $TAG --single-branch https://github.com/grafana/grafana /tmp/grafana; \
    cd /tmp/grafana; \
    yarn install; \
    NODE_ENV=production NODE_OPTIONS=--max_old_space_size=4096 yarn build; \
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

FROM reg.c5h.io/base@sha256:f47ed6eddc8491fd9683c0bb66751d16b0dd6f9e920d96c9d028ec0001ae5968

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
