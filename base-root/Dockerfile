# Generated from https://github.com/csmith/dockerfiles/blob/master/base-root/Dockerfile.gotpl
# BOM: {"apk:acl-libs":"2.3.2-r1","apk:busybox":"1.37.0-r31","apk:busybox-binsh":"1.37.0-r31","apk:ca-certificates":"20261020-r0","apk:libcrypto3":"3.5.9-r0","apk:libxxhash":"0.8.3-r1","apk:lz4-libs":"1.10.0-r1","apk:musl":"1.2.6-r3","apk:popt":"1.19-r4","apk:rsync":"3.5.0-r0","apk:tzdata":"2026e-r0","apk:zlib":"1.3.2-r1","apk:zstd-libs":"1.5.7-r2","image:alpine":"930c2b55771daf688e754e943c002fa02e169f69219c5005f666a7f1e312ce4c"}

FROM reg.c5h.io/alpine@sha256:930c2b55771daf688e754e943c002fa02e169f69219c5005f666a7f1e312ce4c AS build

RUN set -eux; \
    apk add --no-cache \
        busybox=1.37.0-r31 \
        busybox-binsh=1.37.0-r31 \
        ca-certificates=20261020-r0 \
        libcrypto3=3.5.9-r0 \
        musl=1.2.6-r3 \
        acl-libs=2.3.2-r1 \
        libxxhash=0.8.3-r1 \
        lz4-libs=1.10.0-r1 \
        musl=1.2.6-r3 \
        popt=1.19-r4 \
        rsync=3.5.0-r0 \
        zlib=1.3.2-r1 \
        zstd-libs=1.5.7-r2 \
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
    apk add --no-cache musl=1.2.6-r3; \
    apk info -qL musl | rsync -aq --files-from=- / /rootfs/; \
    apk add --no-cache tzdata=2026e-r0; \
    apk info -qL tzdata | rsync -aq --files-from=- / /rootfs/; \
    # Generate the CA certificates file and copy the result
    /usr/sbin/update-ca-certificates; \
    cp /etc/ssl/certs/ca-certificates.crt /rootfs/etc/ssl/certs/ca-certificates.crt; \
    # Create mount points for various special things
    mkdir /rootfs/proc /rootfs/dev /rootfs/sys;

FROM scratch

COPY --from=build /rootfs/ /

ENV SSL_CERT_FILE /etc/ssl/certs/ca-certificates.crt
