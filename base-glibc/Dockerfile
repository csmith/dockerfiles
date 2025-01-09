# Generated from https://github.com/csmith/dockerfiles/blob/master/base-glibc/Dockerfile.gotpl
# BOM: {"apk:acl-libs":"2.3.2-r1","apk:busybox":"1.37.0-r9","apk:busybox-binsh":"1.37.0-r9","apk:ca-certificates":"20241121-r1","apk:gcompat":"1.1.0-r4","apk:libcrypto3":"3.3.2-r4","apk:libgcc":"14.2.0-r4","apk:libstdc++":"14.2.0-r4","apk:libucontext":"1.3.2-r0","apk:libxxhash":"0.8.2-r2","apk:lz4-libs":"1.10.0-r0","apk:musl":"1.2.5-r8","apk:musl-obstack":"1.2.3-r2","apk:popt":"1.19-r4","apk:rsync":"3.3.0-r1","apk:tzdata":"2024b-r1","apk:zlib":"1.3.1-r2","apk:zstd-libs":"1.5.6-r2","image:alpine":"0be8812dc5bc37867ac650dde5331475f4fabd765e7f2f37f3166338e36cf6de"}

FROM reg.c5h.io/alpine@sha256:0be8812dc5bc37867ac650dde5331475f4fabd765e7f2f37f3166338e36cf6de AS build

RUN set -eux; \
    apk add --no-cache \
        busybox=1.37.0-r9 \
        busybox-binsh=1.37.0-r9 \
        ca-certificates=20241121-r1 \
        libcrypto3=3.3.2-r4 \
        musl=1.2.5-r8 \
        acl-libs=2.3.2-r1 \
        libxxhash=0.8.2-r2 \
        lz4-libs=1.10.0-r0 \
        musl=1.2.5-r8 \
        popt=1.19-r4 \
        rsync=3.3.0-r1 \
        zlib=1.3.1-r2 \
        zstd-libs=1.5.6-r2 \
    ;

RUN set -eux; \
    # Set up directories and password file
    mkdir -p /rootfs/etc/ssl/certs /rootfs/tmp /rootfs/home/nonroot /rootfs/root; \
    echo "root:x:0:0:root:/root:/sbin/nologin" >> /rootfs/etc/passwd; \
    echo "nobody:x:65534:65534:nobody:/nonexistent:/sbin/nologin" >> /rootfs/etc/passwd; \
    echo "nonroot:x:65532:65532:nonroot:/home/nonroot:/sbin/nologin" >> /rootfs/etc/passwd; \
    chown 65532:65532 /rootfs/home/nonroot; \
    chmod 01777 /tmp; \
    # Copy our data
    apk add --no-cache gcompat=1.1.0-r4; \
    apk info -qL gcompat | rsync -aq --files-from=- / /rootfs/; \
    apk add --no-cache libgcc=14.2.0-r4; \
    apk info -qL libgcc | rsync -aq --files-from=- / /rootfs/; \
    apk add --no-cache libstdc++=14.2.0-r4; \
    apk info -qL libstdc++ | rsync -aq --files-from=- / /rootfs/; \
    apk add --no-cache libucontext=1.3.2-r0; \
    apk info -qL libucontext | rsync -aq --files-from=- / /rootfs/; \
    apk add --no-cache musl=1.2.5-r8; \
    apk info -qL musl | rsync -aq --files-from=- / /rootfs/; \
    apk add --no-cache musl-obstack=1.2.3-r2; \
    apk info -qL musl-obstack | rsync -aq --files-from=- / /rootfs/; \
    apk add --no-cache tzdata=2024b-r1; \
    apk info -qL tzdata | rsync -aq --files-from=- / /rootfs/; \
    # Generate the CA certificates file and copy the result
    /usr/sbin/update-ca-certificates; \
    cp /etc/ssl/certs/ca-certificates.crt /rootfs/etc/ssl/certs/ca-certificates.crt; \
    # Create mount points for various special things
    mkdir /rootfs/proc /rootfs/dev /rootfs/sys;

FROM scratch

COPY --from=build /rootfs/ /

ENV SSL_CERT_FILE /etc/ssl/certs/ca-certificates.crt
